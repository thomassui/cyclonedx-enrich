package properties

import (
	"cyclonedx-enrich/models"
	"fmt"
	"log/slog"
	"os"

	"cyclonedx-enrich/utils"

	"github.com/CycloneDX/cyclonedx-go"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseEnricher struct {
	models.Enricher
	isInitialized bool
	database      *gorm.DB
}

func connect() (*gorm.DB, error) {
	filename := os.Getenv("DATABASE_FILE")
	return gorm.Open(sqlite.Open(filename), &gorm.Config{})
}

func (e *DatabaseEnricher) Category() models.EnricherCategory {
	return "properties" //TODO: REFLECT package name
}

func (e *DatabaseEnricher) Skip(component *cyclonedx.Component) bool {
	return false
}

func (e *DatabaseEnricher) Enrich(component *cyclonedx.Component) error {
	if !e.isInitialized {
		db, err := connect()

		if err != nil {
			log.Error("unable to load rules",
				slog.String("error", err.Error()))
		}

		e.database = db
		e.isInitialized = true
	}

	if e.database == nil {
		return nil
	}

	var item *models.Component
	e.database.Where("purl = ?", utils.GetRealPurl(component.PackageURL)).Preload("Properties").First(&item)

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
