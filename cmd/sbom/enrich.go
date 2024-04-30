package sbom

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"reflect"
	"strconv"

	"cyclonedx-enrich/enrichers/hashes"
	"cyclonedx-enrich/enrichers/licenses"
	"cyclonedx-enrich/enrichers/managers/cocoapods"
	"cyclonedx-enrich/enrichers/managers/maven"
	"cyclonedx-enrich/enrichers/managers/npm"
	"cyclonedx-enrich/enrichers/managers/pypi"
	"cyclonedx-enrich/enrichers/properties"
	"cyclonedx-enrich/enrichers/references"
	"cyclonedx-enrich/models"

	"cyclonedx-enrich/utils"

	"github.com/CycloneDX/cyclonedx-go"
)

func enrichFiles(expression string) error {

	paths, err := filepath.Glob(expression)

	if err != nil {
		return err
	}

	errs := make([]error, 0)

	if len(paths) > 0 {
		for _, file := range paths {
			log.Info("Enriching file",
				slog.String("file", file))
			if err := enrichFile(file); err != nil {
				log.Info(err.Error())
				errs = append(errs, err)
			}
		}
	} else {
		return fmt.Errorf("file not found %s", expression)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func enrichFile(filename string) error {
	return utils.ReadFile(filename, func(file *os.File) error {
		bom, err := Enrich(file)

		if err != nil {
			return err
		}

		buf := new(bytes.Buffer)
		encoder := cyclonedx.NewBOMEncoder(buf, cyclonedx.BOMFileFormatJSON)
		encoder.SetPretty(true)
		err = encoder.Encode(bom)

		if err != nil {
			return err
		}

		err = os.WriteFile(filename, buf.Bytes(), 0644)

		if err != nil {
			return err
		}

		return nil
	})
}

func Enrich(data io.Reader) (*cyclonedx.BOM, error) {
	request, err := load(data)

	if err != nil {
		return nil, err
	}

	processSBOM(request)

	return request, err
}

func load(data io.Reader) (*cyclonedx.BOM, error) {
	var request = &cyclonedx.BOM{}

	decoder := cyclonedx.NewBOMDecoder(data, cyclonedx.BOMFileFormatJSON)
	err := decoder.Decode(request)

	return request, err
}

var enrichers []models.Enricher = loadEnrichers()

func loadEnrichers() []models.Enricher {
	items := []models.Enricher{
		//licenses
		&licenses.RegexpEnricher{},
		&licenses.DatabaseEnricher{},

		//hashes
		&hashes.RegexpEnricher{},
		&hashes.DatabaseEnricher{},

		//properties
		&properties.RegexpEnricher{},
		&properties.DatabaseEnricher{},

		//references
		&references.RegexpEnricher{},
		&references.DatabaseEnricher{},

		//managers
		&maven.MavenEnricher{},
		&npm.NPMEnricher{},
		&cocoapods.CocoapodsEnricher{},
		&pypi.PyPiEnricher{},
	}

	if value, _ := strconv.ParseBool(os.Getenv("ALLOW_EXTRACT")); value {
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
