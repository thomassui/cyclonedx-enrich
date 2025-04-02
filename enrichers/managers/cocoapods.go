package managers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/CycloneDX/cyclonedx-go"
	"github.com/fnxpt/cyclonedx-enrich/models"
)

var cocoapodsEndpoint = "https://cdn.cocoapods.org"
var pathLenght = 3 //TODO: GET THIS FROM https://cdn.cocoapods.org/CocoaPods-version.yml"

type CocoapodsEnricher struct {
	models.Enricher
}

func (e *CocoapodsEnricher) Skip(component *cyclonedx.Component) bool {
	return skip(component, "pkg:cocoapods/")
}

func (e *CocoapodsEnricher) Enrich(component *cyclonedx.Component) error {
	url := fmt.Sprintf("%s/Specs/%s/%s/%s/%s.podspec.json?aaaa", cocoapodsEndpoint, path(component.Name), component.Name, component.Version, component.Name)

	item, err := request[CocoaPodsPackage](url, parseJSON)

	if err != nil {
		return err
	}

	if item != nil {
		licenses := make([]string, 0)
		hashes := make(map[string]string)
		references := make(map[string]string)
		properties := make(map[string]string)

		if item.License != nil {
			licenses = append(licenses, *item.License)
		}

		enrich(component, licenses, hashes, references, properties)
	}

	//TODO: USE MORE DATA

	return nil
}

func path(name string) string {
	hash := md5.Sum([]byte(name))
	hashString := hex.EncodeToString(hash[:])
	value := make([]string, pathLenght)

	for idx := 0; idx < pathLenght; idx++ {
		value[idx] = string(hashString[idx])
	}

	return strings.Join(value, "/")
}
