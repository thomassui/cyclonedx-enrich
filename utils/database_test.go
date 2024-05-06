package utils

import (
	"os"
	"testing"
)

func TestConnectDatabase(t *testing.T) {
	tests := []struct {
		name     string
		database string
		want     bool
	}{
		{name: "Test with invalid database", database: "testdata/invalid.db", want: false},
		{name: "Test with valid database", database: "testdata/database.db", want: true},
	}
	for _, tt := range tests {

		ResetDatabase()
		os.Setenv("DATABASE_FILE", tt.database)

		t.Run(tt.name, func(t *testing.T) {
			got := ConnectDatabase()
			loaded := (got != nil)
			if loaded != tt.want {
				t.Errorf("ConnectDatabase() = %v, want %v", loaded, tt.want)
			}
		})
	}
}

//TODO: NOT REALLY SURE HOW TO TEST THIS
// func TestRegister(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := Register(); (err != nil) != tt.wantErr {
// 				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
