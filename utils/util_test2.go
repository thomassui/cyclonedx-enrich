package utils

import (
	"net/http"
	"os"
	"reflect"
	"testing"

	"github.com/CycloneDX/cyclonedx-go"
)

func TestGetRealPurl(t *testing.T) {
	tests := []struct {
		name string
		purl string
		want string
	}{
		{name: "Test without qualifiers", purl: "pkg:cocoapods/AppAuth@1.6.2", want: "pkg:cocoapods/AppAuth@1.6.2"},
		{name: "Test with anchors", purl: "pkg:cocoapods/AppAuth@1.6.2#Core", want: "pkg:cocoapods/AppAuth@1.6.2"},
		{name: "Test with aar type qualifier", purl: "pkg:maven/com.github.bumptech.glide/glide@4.15.0?type=aar", want: "pkg:maven/com.github.bumptech.glide/glide@4.15.0"},
		{name: "Test with pom type qualifier", purl: "pkg:maven/org.javamoney/moneta@1.4.2?type=pom", want: "pkg:maven/org.javamoney/moneta@1.4.2"},
		{name: "Test with jar type qualifier", purl: "pkg:maven/io.swagger/swagger-annotations@1.6.9?type=jar", want: "pkg:maven/io.swagger/swagger-annotations@1.6.9"},
		{name: "Test with multiple classifiers", purl: "pkg:maven/io.netty/netty-transport-native-epoll@4.1.85.Final?classifier=linux-x86_64&type=jar", want: "pkg:maven/io.netty/netty-transport-native-epoll@4.1.85.Final"},
		{name: "Test with url encoding", purl: "pkg:npm/%40angular/core@16.2.12", want: "pkg:npm/@angular/core@16.2.12"},
		{name: "Test without url encoding", purl: "pkg:npm/@angular/core@16.2.12", want: "pkg:npm/@angular/core@16.2.12"},
		{name: "Test without npm group", purl: "pkg:npm/parse5@7.1.2", want: "pkg:npm/parse5@7.1.2"},
		{name: "Test with cocoapods format", purl: "pkg:cocoapods/Sample@2.0.0?repository_url=https%3A%2F%2Fartifacts.example.com%2Fapi%2Fpods%2Fios3#Sources", want: "pkg:cocoapods/Sample@2.0.0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRealPurl(tt.purl); got != tt.want {
				t.Errorf("GetRealPurl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoded(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{name: "Test empty", value: "", want: ""},
		{name: "Test without encoding", value: "pkg:cocoapods/AppAuth@1.6.2", want: "pkg:cocoapods/AppAuth@1.6.2"},
		{name: "Test with encoding", value: "pkg:npm/%40angular/core@16.2.12", want: "pkg:npm/@angular/core@16.2.12"},
		{name: "Test with error", value: "%ya", want: "%ya"}, //TODO: FIND ERROR SCENARIO
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decoded(tt.value); got != tt.want {
				t.Errorf("Decoded() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetLicense(t *testing.T) {
	type args struct {
		component    *cyclonedx.Component
		licenseNames []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLicense(tt.args.component, tt.args.licenseNames)
		})
	}
}

func TestRequest(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		wantResp *http.Response
		wantErr  bool
	}{
		//TODO: CONTINUE
		// {name: "", url: "", wantResp: nil, wantErr: false}, //TODO: FAILING
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := Request(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Request() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestReadFile(t *testing.T) {
	type args struct {
		filename string
		fn       func(*os.File) error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test with invalid file", args: args{filename: "no_file", fn: func(f *os.File) error { return nil }}, wantErr: true},
		// {name: "Test with valid file", args: args{filename: "../regexp.yaml", fn: func(f *os.File) error { return nil }}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReadFile(tt.args.filename, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetenv(t *testing.T) {
	type args struct {
		key          string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test with valid env variable", args: args{key: "TEST_ENV_1", defaultValue: "MY_VALUE"}, want: "VALUE1"},
		{name: "Test with valid env variable 2", args: args{key: "TEST_ENV_2", defaultValue: "MY_VALUE"}, want: "VALUE2"},
		{name: "Test with invalid env variable without default value", args: args{key: "TEST_NO_ENV", defaultValue: ""}, want: ""},
		{name: "Test with invalid env variable with default value", args: args{key: "TEST_NO_ENV", defaultValue: "MY_VALUE"}, want: "MY_VALUE"},
	}

	os.Setenv("TEST_ENV_1", "VALUE1")
	os.Setenv("TEST_ENV_2", "VALUE2")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Getenv(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("Getenv() = %v, want %v", got, tt.want)
			}
		})
	}
}
