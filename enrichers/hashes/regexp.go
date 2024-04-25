package hashes

import (
	"cyclonedx-enrich/models"
	"cyclonedx-enrich/utils"
	"fmt"
	"log/slog"
	"os"
	"regexp"

	"github.com/CycloneDX/cyclonedx-go"
	"gopkg.in/yaml.v3"
)

type RegexpEnricher struct {
	models.Enricher
	isInitialized bool
	rules         []models.RuleEntry
}

func loadRules() ([]models.RuleEntry, error) {
	filename := os.Getenv("REGEXP_FILE")
	data, err := os.ReadFile(filename)

	rules := []models.RuleEntry{}

	if err != nil {
		return rules, err
	}

	err = yaml.Unmarshal(data, &rules)

	return rules, err

}

func (e *RegexpEnricher) Category() models.EnricherCategory {
	return "hashes"
}

func (e *RegexpEnricher) Skip(component *cyclonedx.Component) bool {
	return component.Hashes != nil
}

func (e *RegexpEnricher) Enrich(component *cyclonedx.Component) error {
	if !e.isInitialized {
		rules, err := loadRules()

		if err != nil {
			log.Error("unable to load rules",
				slog.String("error", err.Error()))
		}

		e.rules = rules
		e.isInitialized = true
	}

	for _, item := range e.rules {
		r, err := regexp.Compile(item.Rule)

		if err != nil {
			return err
		}

		if r.MatchString(utils.GetRealPurl(component.PackageURL)) {
			for key, value := range item.Hashes {
				if !hasKey(*component.Hashes, key) {
					*component.Hashes = append(*component.Hashes, cyclonedx.Hash{
						Algorithm: cyclonedx.HashAlgorithm(key),
						Value:     value,
					})
				}
			}
			return nil
		}
	}

	return fmt.Errorf("component doesn't met criteria")
}
