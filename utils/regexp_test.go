package utils

import (
	"cyclonedx-enrich/models"
	"os"
	"reflect"
	"testing"
)

func TestLoadRules(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     []models.RuleEntry
	}{
		{name: "Test with invalid regexp", filename: "testdata/invalid.yaml", want: RulesEmpty},
		{name: "Test with valid regexp", filename: "testdata/regexp.yaml", want: Rules},
	}
	for _, tt := range tests {

		ResetRules()
		os.Setenv("REGEXP_FILE", tt.filename)

		t.Run(tt.name, func(t *testing.T) {
			if got := LoadRules(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadRules() = %v, want %v", got, tt.want)
			}
		})
	}
}
