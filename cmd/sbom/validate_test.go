package sbom

import "testing"

func Test_validateFiles(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateFiles(tt.expression); (err != nil) != tt.wantErr {
				t.Errorf("validateFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateFile(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateFile(tt.filename); (err != nil) != tt.wantErr {
				t.Errorf("validateFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
