package properties

import (
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
