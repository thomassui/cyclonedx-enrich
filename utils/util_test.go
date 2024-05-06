package utils

import (
	"testing"
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
