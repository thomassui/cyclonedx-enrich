package references

import (
	"cyclonedx-enrich/models"
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

func enrich(component *cyclonedx.Component, items []models.Reference) error {
	if component.ExternalReferences == nil {
		component.ExternalReferences = &[]cyclonedx.ExternalReference{}
	}
	for _, ref := range items {
		if !hasKey(*component.ExternalReferences, ref.URL, ref.Type) {
			*component.ExternalReferences = append(*component.ExternalReferences, cyclonedx.ExternalReference{
				URL:     ref.URL,
				Type:    cyclonedx.ExternalReferenceType(ref.Type),
				Comment: ref.Comment,
			})
		}
	}
	return nil
}
