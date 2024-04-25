package models

import (
	"github.com/CycloneDX/cyclonedx-go"
)

type EnricherCategory string //TODO: ENUM

type Enricher interface {
	Category() EnricherCategory
	Skip(*cyclonedx.Component) bool
	Enrich(*cyclonedx.Component) error
}
