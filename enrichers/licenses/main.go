package licenses

import (
	"log/slog"

	"github.com/CycloneDX/cyclonedx-go"
	"github.com/fnxpt/cyclonedx-enrich/models"
	"github.com/fnxpt/cyclonedx-enrich/utils"
)

var log = slog.Default()

func skip(component *cyclonedx.Component) bool {
	if component == nil {
		return true
	}
	// Make it possible to enrich components from typ "file" there the purl is typically empty
	// but name instead is an idenetifier to match it
	if len(component.PackageURL) == 0 || (component.Licenses != nil && len(*component.Licenses) > 0) && component.Type == cyclonedx.ComponentTypeFile && len(component.Name) > 0 {
		return false
	}

	return len(component.PackageURL) == 0 || (component.Licenses != nil && len(*component.Licenses) > 0)
}

func enrich(component *cyclonedx.Component, items []string) error {

	if len(items) > 0 {
		utils.SetLicense(component, items)
	}
	return nil
}

func toMap(items []models.License) []string {
	output := make([]string, 0)

	for _, item := range items {
		output = append(output, item.License)
	}

	return output
}
