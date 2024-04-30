package properties

import (
	"cyclonedx-enrich/models"
	"fmt"

	"cyclonedx-enrich/utils"

	"github.com/CycloneDX/cyclonedx-go"
)

type DatabaseEnricher struct {
	models.Enricher
}

func (e *DatabaseEnricher) Skip(component *cyclonedx.Component) bool {
	if utils.ConnectDatabase() == nil {
		return true
	}

	return false
}

func (e *DatabaseEnricher) Enrich(component *cyclonedx.Component) error {
	db := utils.ConnectDatabase()

	var item *models.Component
	db.Where("purl = ?", utils.GetRealPurl(component.PackageURL)).Preload("Properties").First(&item)

	if item != nil {

		for _, property := range item.Properties {
			if !hasKey(*component.Properties, property.Name) {
				*component.Properties = append(*component.Properties, cyclonedx.Property{
					Name:  property.Name,
					Value: property.Value,
				})
			}
		}
	}

	return fmt.Errorf("component doesn't met criteria")
}
