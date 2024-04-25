package hashes

import (
	"log/slog"

	"github.com/CycloneDX/cyclonedx-go"
)

var log = slog.Default()

func hasKey(hashes []cyclonedx.Hash, key string) bool {
	for _, hash := range hashes {
		if hash.Value == key {
			return true
		}
	}

	return false
}
