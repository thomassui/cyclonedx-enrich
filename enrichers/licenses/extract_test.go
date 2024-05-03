package licenses

import (
	"testing"

	"github.com/CycloneDX/cyclonedx-go"
)

func TestExtractEnricher_Skip(t *testing.T) {
	tests := []struct {
		name      string
		component *cyclonedx.Component
		want      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ExtractEnricher{}
			if got := e.Skip(tt.component); got != tt.want {
				t.Errorf("ExtractEnricher.Skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractEnricher_Enrich(t *testing.T) {
	tests := []struct {
		name      string
		component *cyclonedx.Component
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ExtractEnricher{}
			if err := e.Enrich(tt.component); (err != nil) != tt.wantErr {
				t.Errorf("ExtractEnricher.Enrich() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
