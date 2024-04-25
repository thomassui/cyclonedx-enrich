package api

import (
	"log/slog"
	"net/http"
	"reflect"

	"cyclonedx-enrich/enrichers/hashes"
	"cyclonedx-enrich/enrichers/licenses"
	"cyclonedx-enrich/enrichers/properties"
	"cyclonedx-enrich/enrichers/references"
	"cyclonedx-enrich/models"

	"github.com/CycloneDX/cyclonedx-go"
	"github.com/gin-gonic/gin"
)

var items []models.Enricher = loadEnrichers()

func loadEnrichers() []models.Enricher {
	return []models.Enricher{
		//licenses
		&licenses.DatabaseEnricher{},
		&licenses.RegexpEnricher{},
		&licenses.ManagerEnricher{},

		//hashes
		&hashes.DatabaseEnricher{},
		&hashes.RegexpEnricher{},

		//properties
		&properties.DatabaseEnricher{},
		&properties.RegexpEnricher{},

		//references
		&references.DatabaseEnricher{},
		&references.RegexpEnricher{},
	}
}

func ParseSBOM(c *gin.Context) {

	var request = &cyclonedx.BOM{}

	decoder := cyclonedx.NewBOMDecoder(c.Request.Body, cyclonedx.BOMFileFormatJSON)
	err := decoder.Decode(request)

	if err != nil {
		log.Error("Unable to bind sbom",
			slog.String("error", err.Error()))
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	processSBOM(request)

	c.JSON(http.StatusOK, request)
}

func processSBOM(bom *cyclonedx.BOM) {
	if bom.Metadata.Component != nil {
		component := bom.Metadata.Component
		parseComponent(component)

		parseComponents(component.Components)
	}

	parseComponents(bom.Components)
}

func parseComponents(components *[]cyclonedx.Component) {
	if components != nil {
		for i := 0; i < len(*components); i++ {
			component := &(*components)[i]
			parseComponent(component)
		}
	}
}

func parseComponent(component *cyclonedx.Component) {
	log.Debug("parsing component",
		slog.String("component", component.PackageURL))

	for _, enricher := range items {
		if !enricher.Skip(component) {
			err := enricher.Enrich(component)
			if err != nil {
				log.Debug("Unable to enrich",
					slog.String("component", component.PackageURL),
					slog.String("enricher", getType(enricher)),
					slog.String("error", err.Error()))
			} else {
				log.Debug("Enriched",
					slog.String("component", component.PackageURL),
					slog.String("enricher", getType(enricher)))
			}
		}
	}
}

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
