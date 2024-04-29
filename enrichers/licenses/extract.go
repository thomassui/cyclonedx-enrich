package licenses

import (
	"bytes"
	"cyclonedx-enrich/models"
	"cyclonedx-enrich/utils"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/CycloneDX/cyclonedx-go"
)

type ExtractEnricher struct {
}

func (e *ExtractEnricher) Category() models.EnricherCategory {
	return "license"
}

func (e *ExtractEnricher) Skip(component *cyclonedx.Component) bool {
	// TODO: VALIDATE IF IT HAS EXPRESSION OR LICENSE OBJECT
	if len(component.PackageURL) == 0 || component.Licenses != nil {
		//SKIP
		return true
	}

	if component.Properties != nil && len(*component.Properties) > 0 {
		for _, prop := range *component.Properties {
			if strings.HasPrefix(prop.Name, "aquasecurity") {
				return true
			}
		}
	}

	return false
}

func (e *ExtractEnricher) Enrich(component *cyclonedx.Component) error {
	purl := utils.GetRealPurl(component.PackageURL)
	path := fmt.Sprintf("output/%s.json", purl)

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		bom := cyclonedx.NewBOM()
		bom.Metadata = &cyclonedx.Metadata{
			Component: component,
		}

		buf := new(bytes.Buffer)
		encoder := cyclonedx.NewBOMEncoder(buf, cyclonedx.BOMFileFormatJSON)
		encoder.SetPretty(true)
		encoder.Encode(bom)

		dir := filepath.Dir(path)
		os.MkdirAll(dir, 0700)
		os.Create(path)
		os.WriteFile(path, buf.Bytes(), 0644)
	}

	log.Info("No licenses", slog.String("package", purl))

	return fmt.Errorf("component doesn't met criteria")
}
