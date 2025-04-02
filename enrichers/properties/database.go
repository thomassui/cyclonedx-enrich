package properties

import (
	"github.com/fnxpt/cyclonedx-enrich/models"
	"github.com/fnxpt/cyclonedx-enrich/utils"

	"github.com/CycloneDX/cyclonedx-go"
)

type DatabaseEnricher struct {
	models.Enricher
}

func (e *DatabaseEnricher) Skip(component *cyclonedx.Component) bool {
	return skip(component) || utils.ConnectDatabase() == nil
}

func (e *DatabaseEnricher) Enrich(component *cyclonedx.Component) error {
	return utils.EnrichDB(component, "Properties", func(item *models.Component) error {
		return enrich(component, toMap(item.Properties))
	})
}
