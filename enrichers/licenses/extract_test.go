package licenses

import (
	"testing"

	"github.com/CycloneDX/cyclonedx-go"
	"github.com/fnxpt/cyclonedx-enrich/utils"
)

func TestExtractEnricher_Skip(t *testing.T) {
	tests := []struct {
		name      string
		component cyclonedx.Component
		want      bool
	}{
		{name: "Test with empty component", component: *utils.ComponentEmpty, want: true},
		{name: "Test with component with data", component: *utils.ComponentWithData, want: true},
		{name: "Test with component without data", component: *utils.ComponentWithoutData, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ExtractEnricher{}
			if got := e.Skip(&tt.component); got != tt.want {
				t.Errorf("ExtractEnricher.Skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractEnricher_Enrich(t *testing.T) {
	tests := []struct {
		name      string
		component cyclonedx.Component
		wantErr   bool
	}{
		{name: "Test with component without data", component: *utils.ComponentWithoutData, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ExtractEnricher{}
			if err := e.Enrich(&tt.component); (err != nil) != tt.wantErr {
				t.Errorf("ExtractEnricher.Enrich() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
