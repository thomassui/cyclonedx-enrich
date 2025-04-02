package references

import (
	"github.com/fnxpt/cyclonedx-enrich/models"
	"github.com/fnxpt/cyclonedx-enrich/utils"

	"github.com/CycloneDX/cyclonedx-go"
)

type RegexpEnricher struct {
	models.Enricher
}

func (e *RegexpEnricher) Skip(component *cyclonedx.Component) bool {
	return skip(component) || len(utils.LoadRules()) == 0
}

func (e *RegexpEnricher) Enrich(component *cyclonedx.Component) error {
	return utils.EnrichRules(component, func(item *models.RuleEntry) error {
		return enrich(component, item.References)
	})
}
