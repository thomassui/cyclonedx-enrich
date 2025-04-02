package hashes

import (
	"github.com/CycloneDX/cyclonedx-go"
	"github.com/fnxpt/cyclonedx-enrich/models"
	"github.com/fnxpt/cyclonedx-enrich/utils"
)

type RegexpEnricher struct {
	models.Enricher
}

func (e *RegexpEnricher) Skip(component *cyclonedx.Component) bool {
	return skip(component) || len(utils.LoadRules()) == 0
}

func (e *RegexpEnricher) Enrich(component *cyclonedx.Component) error {
	return utils.EnrichRules(component, func(item *models.RuleEntry) error {
		return enrich(component, item.Hashes)
	})
}
