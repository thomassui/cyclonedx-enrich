package database

import (
	"cyclonedx-enrich/models"
	"cyclonedx-enrich/utils"
	"reflect"
	"testing"

	"github.com/CycloneDX/cyclonedx-go"
)

func Test_importFile(t *testing.T) {
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
			if err := importFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("importFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_importComponent(t *testing.T) {
	type args struct {
		data []byte
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
			if err := importComponent(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("importComponent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getLicenses(t *testing.T) {
	type args struct {
		component *cyclonedx.Component
	}
	tests := []struct {
		name string
		args args
		want []models.License
	}{
		{name: "Test with empty component", args: args{component: utils.ComponentEmpty}, want: []models.License{}},
		{name: "Test with component without data", args: args{utils.ComponentWithoutData}, want: []models.License{}},
		{name: "Test with component with data", args: args{utils.ComponentWithData}, want: []models.License{
			{License: "Apache-2.0"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLicenses(tt.args.component); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLicenses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getProperties(t *testing.T) {
	type args struct {
		component *cyclonedx.Component
	}
	tests := []struct {
		name string
		args args
		want []models.Property
	}{
		{name: "Test with empty component", args: args{component: utils.ComponentEmpty}, want: []models.Property{}},
		{name: "Test with component without data", args: args{utils.ComponentWithoutData}, want: []models.Property{}},
		{name: "Test with component with data", args: args{utils.ComponentWithData}, want: []models.Property{
			{Name: "cdx:npm:package:path", Value: "node_modules/@angular/cdk/node_modules/parse5"},   //TODO: GET BETTER DATA
			{Name: "cdx:npm:package:path2", Value: "node_modules/@angular/cdk/node_modules/parse52"}, //TODO: GET BETTER DATA
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getProperties(tt.args.component); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getReferences(t *testing.T) {
	type args struct {
		component *cyclonedx.Component
	}
	tests := []struct {
		name string
		args args
		want []models.Reference
	}{
		{name: "Test with empty component", args: args{component: utils.ComponentEmpty}, want: []models.Reference{}},
		{name: "Test with component without data", args: args{utils.ComponentWithoutData}, want: []models.Reference{}},
		{name: "Test with component with data", args: args{utils.ComponentWithData}, want: []models.Reference{
			{URL: "https://github.com/OpenAPITools/jackson-databind-nullable", Type: "website"},
			{URL: "https://oss.sonatype.org/service/local/staging/deploy/maven2/", Type: "distribution"},
			{URL: "https://github.com/OpenAPITools/jackson-databind-nullable", Type: "vcs"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getReferences(tt.args.component); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getReferences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHashes(t *testing.T) {
	type args struct {
		component *cyclonedx.Component
	}
	tests := []struct {
		name string
		args args
		want []models.Hash
	}{
		{name: "Test with empty component", args: args{component: utils.ComponentEmpty}, want: []models.Hash{}},
		{name: "Test with component without data", args: args{utils.ComponentWithoutData}, want: []models.Hash{}},
		{name: "Test with component with data", args: args{utils.ComponentWithData}, want: []models.Hash{
			{Name: "MD5", Value: "479311558bbca63453f8a79e2735aec1"},
			{Name: "SHA-1", Value: "371a38c3d339833edb1b2a0d96c3d249a890bcc4"},
			{Name: "SHA-256", Value: "22c73f6c44eb65cb2ebbd9a0ace61a3951cc259fdc29b89e31a80564cd116ad6"},
			{Name: "SHA-512", Value: "41a4c682635a481f78602087a83a7bbd1f36c0fd8d8fe5daf2ab05907472ca2f345de086fa56bab2d554412f2a1546ec5a2e832e04b1751ba29e6612318b42dc"},
			{Name: "SHA-384", Value: "740ff354152ae7d691590c75d9c0be6decbb18912f56e3aca86243b7e9f5c350c48ca0e97fb3e031aa8aaf82c49e0885"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHashes(tt.args.component); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHashes() = %v, want %v", got, tt.want)
			}
		})
	}
}
