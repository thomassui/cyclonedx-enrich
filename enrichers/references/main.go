package references

import (
	"log/slog"

	"github.com/CycloneDX/cyclonedx-go"
)

var log = slog.Default()

func hasKey(references []cyclonedx.ExternalReference, url string, refType string) bool {
	for _, ref := range references {
		if ref.URL == url && ref.Type == cyclonedx.ExternalReferenceType(refType) {
			return true
		}
	}

	return false
}
