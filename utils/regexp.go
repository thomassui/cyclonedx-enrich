package utils

import (
	"cyclonedx-enrich/models"
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

var isRegexpInitialized bool
var rules = make([]models.RuleEntry, 0)

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
