package utils

import (
	"fmt"
	"log/slog"
	"os"
	"regexp"

	"github.com/fnxpt/cyclonedx-enrich/models"

	"github.com/CycloneDX/cyclonedx-go"
	"gopkg.in/yaml.v3"
)

var isRegexpInitialized bool
var rules = make([]models.RuleEntry, 0)

func ResetRules() {
	isRegexpInitialized = false
	rules = make([]models.RuleEntry, 0)
}

func LoadRules() []models.RuleEntry {
	if !isRegexpInitialized {
		items, err := loadRules()

		if err != nil {
			log.Warn("unable to load rules",
				slog.String("error", err.Error()))
		}

		rules = items
		isRegexpInitialized = true
	}

	return rules
}

func loadRules() ([]models.RuleEntry, error) {
	filename := Getenv("REGEXP_FILE", "regexp.yaml")
	data, err := os.ReadFile(filename)

	rules := []models.RuleEntry{}

	if err != nil {
		return rules, err
	}

	err = yaml.Unmarshal(data, &rules)

	return rules, err
}

func EnrichRules(component *cyclonedx.Component, fn func(item *models.RuleEntry) error) error {
	rules := LoadRules()

	if len(rules) == 0 {
		return fmt.Errorf("unable to access rules")
	}

	for _, item := range rules {
		r, err := regexp.Compile(item.Rule)

		if err != nil {
			return err
		}

		if r.MatchString(GetRealPurl(component.PackageURL)) {
			fn(&item)
		}
	}

	return nil
}
