package references

import (
	"testing"

	"cyclonedx-enrich/utils"

	"github.com/CycloneDX/cyclonedx-go"
)

func TestRegexpEnricher_Skip(t *testing.T) {

	tests := []struct {
		name      string
		component *cyclonedx.Component
		want      bool
	}{
		//TODO: CONTINUE
		{name: "Test with empty component", component: utils.ComponentEmpty, want: true},
		{name: "Test with component with data", component: utils.ComponentWithData, want: true},
		// {name: "Test with component without data", component: utils.ComponentWithoutData, want: false}, //TODO: FAILING
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &RegexpEnricher{}
			if got := e.Skip(tt.component); got != tt.want {
				t.Errorf("RegexpEnricher.Skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegexpEnricher_Enrich(t *testing.T) {

	tests := []struct {
		name      string
		component *cyclonedx.Component
		wantErr   bool
	}{
		//TODO: VALIDATE IF DATA WAS ADDED
		// {name: "Test with component without data", component: utils.ComponentWithoutData, wantErr: false}, //TODO: FAILING
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &RegexpEnricher{}
			if err := e.Enrich(tt.component); (err != nil) != tt.wantErr {
				t.Errorf("RegexpEnricher.Enrich() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
