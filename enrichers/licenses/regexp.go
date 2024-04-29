package licenses

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
	return "license"
}

func (e *RegexpEnricher) Skip(component *cyclonedx.Component) bool {
	// TODO: VALIDATE IF IT HAS EXPRESSION OR LICENSE OBJECT
	if len(component.PackageURL) == 0 || component.Licenses != nil {
		//SKIP
		return true
	}
	return false
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
			utils.SetLicense(component, item.Licenses)
			return nil
		}
	}

	return fmt.Errorf("component doesn't met criteria")
}
