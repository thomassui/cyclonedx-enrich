package references

import (
	"cyclonedx-enrich/utils"
	"os"
	"testing"

	"github.com/CycloneDX/cyclonedx-go"
)

func setup(tb testing.TB, valid bool) func(tb testing.TB) {
	utils.ResetDatabase()
	utils.ResetRules()

	if valid {
		os.Setenv("DATABASE_FILE", "testdata/database.db")
		os.Setenv("REGEXP_FILE", "testdata/regexp.yaml")
	} else {
		os.Setenv("DATABASE_FILE", "testdata/invalid.db")
		os.Setenv("REGEXP_FILE", "testdata/invalid.yaml")
	}

	// Return a function to teardown the test
	return func(tb testing.TB) {}
}

var emptyArray = make([]cyclonedx.ExternalReference, 0)
var filledArray = []cyclonedx.ExternalReference{
	{URL: "url1", Type: "type1", Comment: "comment1"},
	{URL: "url2", Type: "type2", Comment: "comment2"},
	{URL: "url3", Type: "type3", Comment: "comment3"},
}

func Test_hasKey(t *testing.T) {
	type args struct {
		references []cyclonedx.ExternalReference
		url        string
		refType    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test without args", args: args{}, want: false},
		{name: "Test with empty array", args: args{references: emptyArray, url: "url1", refType: "type1"}, want: false},
		{name: "Test with empty invalid url", args: args{references: filledArray, url: "", refType: "type1"}, want: false},
		{name: "Test with empty invalid type", args: args{references: filledArray, url: "url1", refType: ""}, want: false},
		{name: "Test with array without key", args: args{references: filledArray, url: "no_key", refType: ""}, want: false},
		{name: "Test with array invalid combination", args: args{references: filledArray, url: "url2", refType: "type1"}, want: false},
		{name: "Test with array with key 1", args: args{references: filledArray, url: "url1", refType: "type1"}, want: true},
		{name: "Test with array with key 2", args: args{references: filledArray, url: "url2", refType: "type2"}, want: true},
		{name: "Test with array with key 3", args: args{references: filledArray, url: "url3", refType: "type3"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasKey(tt.args.references, tt.args.url, tt.args.refType); got != tt.want {
				t.Errorf("hasKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
