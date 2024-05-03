package hashes

import (
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
