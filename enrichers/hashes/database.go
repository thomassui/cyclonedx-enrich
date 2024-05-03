package hashes

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
	return skip(component)
}

func (e *DatabaseEnricher) Enrich(component *cyclonedx.Component) error {
	db := utils.ConnectDatabase()

	if db == nil {
		return fmt.Errorf("Unable to access database")
	}

	var item *models.Component
	db.Where("purl = ?", utils.GetRealPurl(component.PackageURL)).Preload("Hashes").First(&item)

	if item != nil {

		for _, hash := range item.Hashes {
			if !hasKey(*component.Hashes, hash.Name) {
				*component.Hashes = append(*component.Hashes, cyclonedx.Hash{
					Algorithm: cyclonedx.HashAlgorithm(hash.Name),
					Value:     hash.Value,
				})
			}
		}
	}

	return fmt.Errorf("component doesn't met criteria")
}
