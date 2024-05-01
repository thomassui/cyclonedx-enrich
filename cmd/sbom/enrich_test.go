package sbom

import (
	"cyclonedx-enrich/models"
	"io"
	"reflect"
	"testing"

	"github.com/CycloneDX/cyclonedx-go"
)

func Test_enrichFiles(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := enrichFiles(tt.args.expression); (err != nil) != tt.wantErr {
				t.Errorf("enrichFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_enrichFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := enrichFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("enrichFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEnrich(t *testing.T) {
	type args struct {
		data io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *cyclonedx.BOM
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Enrich(tt.args.data)
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
	type args struct {
		data io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *cyclonedx.BOM
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := load(tt.args.data)
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
	type args struct {
		bom *cyclonedx.BOM
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processSBOM(tt.args.bom)
		})
	}
}

func Test_parseComponents(t *testing.T) {
	type args struct {
		components *[]cyclonedx.Component
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseComponents(tt.args.components)
		})
	}
}

func Test_parseComponent(t *testing.T) {
	type args struct {
		component *cyclonedx.Component
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseComponent(tt.args.component)
		})
	}
}

func Test_getType(t *testing.T) {
	type args struct {
		myvar interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getType(tt.args.myvar); got != tt.want {
				t.Errorf("getType() = %v, want %v", got, tt.want)
			}
		})
	}
}
