package properties

import (
	"os"
	"testing"

	"github.com/CycloneDX/cyclonedx-go"
)

func setup(tb testing.TB) func(tb testing.TB) {
	os.Setenv("DATABASE_FILE", "testdata/database.db")
	os.Setenv("REGEXP_FILE", "testdata/regexp.yaml")
	// Return a function to teardown the test
	return func(tb testing.TB) {}
}

var emptyArray = make([]cyclonedx.Property, 0)
var filledArray = []cyclonedx.Property{
	{Name: "key1", Value: "value1"},
	{Name: "key2", Value: "value2"},
	{Name: `key3`, Value: "value3"},
}

func Test_hasKey(t *testing.T) {
	type args struct {
		properties []cyclonedx.Property
		key        string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test without arguments", args: args{}, want: false},
		{name: "Test with nil array", args: args{key: "key1"}, want: false},
		{name: "Test with empty array", args: args{properties: emptyArray, key: "key1"}, want: false},
		{name: "Test with empty invalid key", args: args{properties: filledArray, key: ""}, want: false},
		{name: "Test with array without key", args: args{properties: filledArray, key: "no_key"}, want: false},
		{name: "Test with array with key 1", args: args{properties: filledArray, key: "key1"}, want: true},
		{name: "Test with array with key 2", args: args{properties: filledArray, key: "key2"}, want: true},
		{name: "Test with array with key 3", args: args{properties: filledArray, key: "key3"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasKey(tt.args.properties, tt.args.key); got != tt.want {
				t.Errorf("hasKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
