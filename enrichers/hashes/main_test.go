package hashes

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

var emptyArray = make([]cyclonedx.Hash, 0)
var filledArray = []cyclonedx.Hash{
	{Algorithm: "key1", Value: "value1"},
	{Algorithm: "key2", Value: "value2"},
	{Algorithm: `key3`, Value: "value3"},
}

func Test_hasKey(t *testing.T) {
	type args struct {
		hashes []cyclonedx.Hash
		key    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test without arguments", args: args{}, want: false},
		{name: "Test with nil array", args: args{key: "key1"}, want: false},
		{name: "Test with empty array", args: args{hashes: emptyArray, key: "key1"}, want: false},
		{name: "Test with empty invalid key", args: args{hashes: filledArray, key: ""}, want: false},
		{name: "Test with array without key", args: args{hashes: filledArray, key: "no_key"}, want: false},
		{name: "Test with array with key 1", args: args{hashes: filledArray, key: "key1"}, want: true},
		{name: "Test with array with key 2", args: args{hashes: filledArray, key: "key2"}, want: true},
		{name: "Test with array with key 3", args: args{hashes: filledArray, key: "key3"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasKey(tt.args.hashes, tt.args.key); got != tt.want {
				t.Errorf("hasKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
