package models

type Component struct {
	Purl       string      `gorm:"primaryKey"`
	Licenses   []License   `gorm:"many2many:component_licenses;"`
	Properties []Property  `gorm:"many2many:component_properties;"`
	Hashes     []Hash      `gorm:"many2many:component_hashes;"`
	References []Reference `gorm:"many2many:component_references;"`
}

type License struct {
	License string `gorm:"primaryKey"`
}

type Property struct {
	Name  string `gorm:"primaryKey"`
	Value string `gorm:"primaryKey"`
}

type Hash struct {
	Name  string `gorm:"primaryKey"`
	Value string `gorm:"primaryKey"`
}

type Reference struct {
	URL     string `gorm:"primaryKey"`
	Type    string `gorm:"primaryKey"`
	Comment string
}
