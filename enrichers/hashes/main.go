package hashes

import (
	"cyclonedx-enrich/models"
	"log/slog"

	"github.com/CycloneDX/cyclonedx-go"
)

var log = slog.Default()

func skip(component *cyclonedx.Component) bool {
	return component == nil || len(component.PackageURL) == 0 || component.Hashes != nil
}

func hasKey(hashes []cyclonedx.Hash, key string) bool {
	for _, hash := range hashes {
		if string(hash.Algorithm) == key {
			return true
		}
	}

	return false
}

func enrich(component *cyclonedx.Component, items map[string]string) error {
	if component.Hashes == nil {
		component.Hashes = &[]cyclonedx.Hash{}
	}

	for key, value := range items {
		if !hasKey(*component.Hashes, key) {
			*component.Hashes = append(*component.Hashes, cyclonedx.Hash{
				Algorithm: cyclonedx.HashAlgorithm(key),
				Value:     value,
			})
		}
	}
	return nil
}

func toMap(items []models.Hash) map[string]string {
	output := make(map[string]string)

	for _, item := range items {
		output[item.Name] = item.Value
	}

	return output
}
