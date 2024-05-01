package utils

import (
	"cyclonedx-enrich/models"
	"reflect"
	"testing"
)

func TestLoadRules(t *testing.T) {
	tests := []struct {
		name string
		want []models.RuleEntry
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadRules(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadRules() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadRules(t *testing.T) {
	tests := []struct {
		name    string
		want    []models.RuleEntry
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadRules()
			if (err != nil) != tt.wantErr {
				t.Errorf("loadRules() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadRules() = %v, want %v", got, tt.want)
			}
		})
	}
}
