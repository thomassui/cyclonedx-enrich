package licenses

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
	return "license"
}

func (e *DatabaseEnricher) Skip(component *cyclonedx.Component) bool {
	if utils.ConnectDatabase() == nil {
		return true
	}
	// TODO: VALIDATE IF IT HAS EXPRESSION OR LICENSE OBJECT
	if len(component.PackageURL) == 0 || component.Licenses != nil {
		//SKIP
		return true
	}
	return false
}

func (e *DatabaseEnricher) Enrich(component *cyclonedx.Component) error {
	db := utils.ConnectDatabase()

	var item *models.Component
	db.Where("purl = ?", utils.GetRealPurl(component.PackageURL)).Preload("Licenses").First(&item)

	if item != nil {
		licenses := make([]string, 0)

		for _, license := range item.Licenses {
			licenses = append(licenses, license.License)
		}

		if len(licenses) > 0 {
			utils.SetLicense(component, licenses)
			return nil
		}
	}

	return fmt.Errorf("component doesn't met criteria")
}
