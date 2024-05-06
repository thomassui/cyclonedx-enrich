package managers

import (
	"testing"

	"github.com/CycloneDX/cyclonedx-go"
)

func Test_skip(t *testing.T) {
	type args struct {
		component *cyclonedx.Component
		prefix    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := skip(tt.args.component, tt.args.prefix); got != tt.want {
				t.Errorf("skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_enrich(t *testing.T) {
	type args struct {
		component  *cyclonedx.Component
		licenses   []string
		hashes     map[string]string
		references map[string]string
		properties map[string]string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enrich(tt.args.component, tt.args.licenses, tt.args.hashes, tt.args.references, tt.args.properties)
		})
	}
}
