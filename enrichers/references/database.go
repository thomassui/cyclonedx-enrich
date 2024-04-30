package references

import (
	"cyclonedx-enrich/models"
	"fmt"

	"cyclonedx-enrich/utils"

	"github.com/CycloneDX/cyclonedx-go"
)

type DatabaseEnricher struct {
	models.Enricher
}

func (e *DatabaseEnricher) Category() models.EnricherCategory {
	return "references" //TODO: REFLECT package name
}

func (e *DatabaseEnricher) Skip(component *cyclonedx.Component) bool {
	if utils.ConnectDatabase() == nil {
		return true
	}

	return true
}

func (e *DatabaseEnricher) Enrich(component *cyclonedx.Component) error {
	db := utils.ConnectDatabase()

	var item *models.Component
	db.Where("purl = ?", utils.GetRealPurl(component.PackageURL)).Preload("References").First(&item)

	if item != nil {

		for _, ref := range item.References {
			if !hasKey(*component.ExternalReferences, ref.URL, ref.Type) {
				*component.ExternalReferences = append(*component.ExternalReferences, cyclonedx.ExternalReference{
					URL:     ref.URL,
					Type:    cyclonedx.ExternalReferenceType(ref.Type),
					Comment: ref.Comment,
				})
			}
		}
		return nil
	}

	return fmt.Errorf("component doesn't met criteria")
}
