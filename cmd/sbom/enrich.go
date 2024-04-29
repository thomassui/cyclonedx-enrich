package sbom

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"os"
	"reflect"
	"strconv"

	"cyclonedx-enrich/enrichers/hashes"
	"cyclonedx-enrich/enrichers/licenses"
	"cyclonedx-enrich/enrichers/managers/cocoapods"
	"cyclonedx-enrich/enrichers/managers/maven"
	"cyclonedx-enrich/enrichers/managers/npm"
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

	processSBOM(request)

	return request, err
}

var enrichers []models.Enricher = loadEnrichers()

func loadEnrichers() []models.Enricher {
	value, _ := strconv.ParseBool(os.Getenv("ALLOW_EXTRACT"))

	items := []models.Enricher{
		//licenses
		&licenses.DatabaseEnricher{},
		&licenses.RegexpEnricher{},

		//hashes
		&hashes.DatabaseEnricher{},
		&hashes.RegexpEnricher{},

		//properties
		&properties.DatabaseEnricher{},
		&properties.RegexpEnricher{},

		//references
		&references.DatabaseEnricher{},
		&references.RegexpEnricher{},

		//managers
		&maven.MavenEnricher{},
		&npm.NPMEnricher{},
		&cocoapods.CocoapodsEnricher{},
	}

	if value {
		items = append(items, &licenses.ExtractEnricher{})
	}

	return items
}

func processSBOM(bom *cyclonedx.BOM) {
	if bom.Metadata != nil && bom.Metadata.Component != nil {
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

	for _, enricher := range enrichers {
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
