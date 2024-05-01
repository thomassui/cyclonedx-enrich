package models

import (
	"github.com/CycloneDX/cyclonedx-go"
)

type Enricher interface {
	Skip(*cyclonedx.Component) bool
	Enrich(*cyclonedx.Component) error
}
