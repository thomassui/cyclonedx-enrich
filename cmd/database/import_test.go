package database

import (
	"cyclonedx-enrich/models"
	"cyclonedx-enrich/utils"
	"reflect"
	"testing"

	"github.com/CycloneDX/cyclonedx-go"
)

func Test_importFile(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := importFile(tt.filename); (err != nil) != tt.wantErr {
				t.Errorf("importFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_importComponent(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := importComponent(tt.data); (err != nil) != tt.wantErr {
				t.Errorf("importComponent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getLicenses(t *testing.T) {

	tests := []struct {
		name      string
		component *cyclonedx.Component
		want      []models.License
	}{
		{name: "Test with empty component", component: utils.ComponentEmpty, want: []models.License{}},
		{name: "Test with component without data", component: utils.ComponentWithoutData, want: []models.License{}},
		{name: "Test with component with data", component: utils.ComponentWithData, want: []models.License{
			{License: "Apache 2.0"},
			{License: "EDL 1.0"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLicenses(tt.component); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLicenses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getProperties(t *testing.T) {

	tests := []struct {
		name      string
		component *cyclonedx.Component
		want      []models.Property
	}{
		{name: "Test with empty component", component: utils.ComponentEmpty, want: []models.Property{}},
		{name: "Test with component without data", component: utils.ComponentWithoutData, want: []models.Property{}},
		{name: "Test with component with data", component: utils.ComponentWithData, want: []models.Property{
			{Name: "cdx:npm:package:path", Value: "node_modules/@angular/cdk/node_modules/parse5"},   //TODO: GET BETTER DATA
			{Name: "cdx:npm:package:path2", Value: "node_modules/@angular/cdk/node_modules/parse52"}, //TODO: GET BETTER DATA
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getProperties(tt.component); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getReferences(t *testing.T) {

	tests := []struct {
		name      string
		component *cyclonedx.Component
		want      []models.Reference
	}{
		{name: "Test with empty component", component: utils.ComponentEmpty, want: []models.Reference{}},
		{name: "Test with component without data", component: utils.ComponentWithoutData, want: []models.Reference{}},
		{name: "Test with component with data", component: utils.ComponentWithData, want: []models.Reference{
			{URL: "https://github.com/OpenAPITools/jackson-databind-nullable", Type: "website"},
			{URL: "https://oss.sonatype.org/service/local/staging/deploy/maven2/", Type: "distribution"},
			{URL: "https://github.com/OpenAPITools/jackson-databind-nullable", Type: "vcs"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getReferences(tt.component); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getReferences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHashes(t *testing.T) {

	tests := []struct {
		name      string
		component *cyclonedx.Component
		want      []models.Hash
	}{
		{name: "Test with empty component", component: utils.ComponentEmpty, want: []models.Hash{}},
		{name: "Test with component without data", component: utils.ComponentWithoutData, want: []models.Hash{}},
		{name: "Test with component with data", component: utils.ComponentWithData, want: []models.Hash{
			{Name: "MD5", Value: "a0cbf1399c88a0435be995c0f68a3baa"},
			{Name: "SHA-1", Value: "15ee0d893c452db7c4865eacb37cf5355c6c5cd7"},
			{Name: "SHA-256", Value: "b3f164c170d7a281a5b869b0ead0fbca3f3c9e06a01b1be521460816eef861e5"},
			{Name: "SHA-512", Value: "3cc8334f21ca594e6b62d5d18e6219bb781d2a51edd790bc981b80bcd701310ad4cb83541e19dd17214ff61d8c2fda09c83b2c6473da3a2360c770e8fdad53ce"},
			{Name: "SHA-384", Value: "52b81b209ec0b5f4920420203fff8911f7d61defcbcef588debdfa85d2dc22f8d1126d02905092e7131a23698afec639"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHashes(tt.component); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHashes() = %v, want %v", got, tt.want)
			}
		})
	}
}
