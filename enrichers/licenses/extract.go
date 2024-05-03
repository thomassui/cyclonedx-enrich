package licenses

import (
	"bytes"
	"cyclonedx-enrich/utils"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/CycloneDX/cyclonedx-go"
)

type ExtractEnricher struct {
}

func (e *ExtractEnricher) Skip(component *cyclonedx.Component) bool {
	return skip(component)
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
