package sbom

import "testing"

func Test_validateFiles(t *testing.T) {
	type args struct {
		expression string
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
			if err := validateFiles(tt.args.expression); (err != nil) != tt.wantErr {
				t.Errorf("validateFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateFile(t *testing.T) {
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
			if err := validateFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("validateFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
