package properties

import (
	"cyclonedx-enrich/models"
	"log/slog"

	"github.com/CycloneDX/cyclonedx-go"
)

var log = slog.Default()

func hasKey(properties []cyclonedx.Property, key string) bool {
	for _, property := range properties {
		if property.Name == key {
			return true
		}
	}

	return false
}

func enrich(component *cyclonedx.Component, items map[string]string) error {
	if component.Properties == nil {
		component.Properties = &[]cyclonedx.Property{}
	}
	for key, value := range items {
		if !hasKey(*component.Properties, key) {
			*component.Properties = append(*component.Properties, cyclonedx.Property{
				Name:  key,
				Value: value,
			})
		}
	}
	return nil
}

func toMap(items []models.Property) map[string]string {
	output := make(map[string]string)

	for _, item := range items {
		output[item.Name] = item.Value
	}

	return output
}
