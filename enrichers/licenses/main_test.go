package licenses

import (
	"os"
	"testing"
)

func setup(tb testing.TB) func(tb testing.TB) {
	os.Setenv("DATABASE_FILE", "testdata/database.db")
	os.Setenv("REGEXP_FILE", "testdata/regexp.yaml")
	// Return a function to teardown the test
	return func(tb testing.TB) {}
}
