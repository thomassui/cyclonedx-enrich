package database

import (
	"reflect"
	"testing"

	"github.com/fnxpt/cyclonedx-enrich/models"
)

func TestDatabaseCMD_Commands(t *testing.T) {
	tests := []struct {
		name string
		c    DatabaseCMD
		want []models.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Commands(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DatabaseCMD.Commands() = %v, want %v", got, tt.want)
			}
		})
	}
}
