package sbom

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"os"
	"reflect"

	"cyclonedx-enrich/enrichers/hashes"
	"cyclonedx-enrich/enrichers/licenses"
	"cyclonedx-enrich/enrichers/properties"
	"cyclonedx-enrich/enrichers/references"
	"cyclonedx-enrich/models"

	"github.com/CycloneDX/cyclonedx-go"
)

func enrichFile(filename string) error {
	file, err := os.Open(filename)

	if err != nil {
		return err
	}

	bom, err := Enrich(file)

	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	encoder := cyclonedx.NewBOMEncoder(buf, cyclonedx.BOMFileFormatJSON)
	encoder.SetPretty(true)
	encoder.Encode(bom)

	fmt.Println(string(buf.Bytes()))
	return nil
}

func Enrich(data io.Reader) (*cyclonedx.BOM, error) {
	var request = &cyclonedx.BOM{}

	decoder := cyclonedx.NewBOMDecoder(data, cyclonedx.BOMFileFormatJSON)
	err := decoder.Decode(request)

	return request, err
}

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
