package sbom

import (
	"io"
	"reflect"
	"testing"

	"github.com/fnxpt/cyclonedx-enrich/models"

	"github.com/CycloneDX/cyclonedx-go"
)

func Test_enrichFiles(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := enrichFiles(tt.expression); (err != nil) != tt.wantErr {
				t.Errorf("enrichFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_enrichFile(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := enrichFile(tt.filename); (err != nil) != tt.wantErr {
				t.Errorf("enrichFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEnrich(t *testing.T) {
	tests := []struct {
		name    string
		data    io.Reader
		want    *cyclonedx.BOM
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Enrich(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Enrich() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enrich() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_load(t *testing.T) {
	tests := []struct {
		name    string
		data    io.Reader
		want    *cyclonedx.BOM
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := load(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("load() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadEnrichers(t *testing.T) {
	tests := []struct {
		name string
		want []models.Enricher
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loadEnrichers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadEnrichers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processSBOM(t *testing.T) {
	tests := []struct {
		name string
		bom  *cyclonedx.BOM
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processSBOM(tt.bom)
		})
	}
}

func Test_parseComponents(t *testing.T) {
	tests := []struct {
		name       string
		components *[]cyclonedx.Component
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseComponents(tt.components)
		})
	}
}

func Test_parseComponent(t *testing.T) {

	tests := []struct {
		name      string
		component *cyclonedx.Component
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseComponent(tt.component)
		})
	}
}

func Test_getType(t *testing.T) {
	tests := []struct {
		name  string
		myvar interface{}
		want  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getType(tt.myvar); got != tt.want {
				t.Errorf("getType() = %v, want %v", got, tt.want)
			}
		})
	}
}
