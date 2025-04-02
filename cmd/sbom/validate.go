package sbom

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/fnxpt/cyclonedx-enrich/utils"
)

func validateFiles(expression string) error {

	paths, err := filepath.Glob(expression)

	if err != nil {
		return err
	}

	errs := make([]error, 0)

	if len(paths) > 0 {
		for _, file := range paths {
			log.Info("Validating file",
				slog.String("file", file))
			if err := validateFile(file); err != nil {
				errs = append(errs, err)
			}
		}
	} else {
		return fmt.Errorf("file not found %s", expression)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func validateFile(filename string) error {
	return utils.ReadFile(filename, func(file *os.File) error {
		_, err := load(file)

		return err
	})
}
