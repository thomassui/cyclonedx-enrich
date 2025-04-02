package licenses

import (
	"os"
	"testing"

	"github.com/fnxpt/cyclonedx-enrich/utils"
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
