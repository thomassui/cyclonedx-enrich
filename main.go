package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/fnxpt/cyclonedx-enrich/enrich"

	"github.com/CycloneDX/cyclonedx-go"
)

var version = "0.0.1"
var force bool
var patternFile string = ""
var pattern string = ""
var license string = ""

var outputFormat = cyclonedx.BOMFileFormatJSON
var output = os.Stdout

func main() {
	parseArguments()
}

func showHelpMenu() {
	fmt.Printf("usage: cyclonedx-enrich(%s) [options]\n", version)
	fmt.Println("options:")
	os.Exit(0)
}

func parseArguments() {
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)

	flag.Usage = func() {
		showHelpMenu()
		flag.PrintDefaults()
	}

	flag.BoolVar(&force, "force", false, "sets the license even if its already filled")
	flag.StringVar(&pattern, "pattern", "", "sets the pattern to add license")
	flag.StringVar(&patternFile, "pattern-file", "", "sets file with the patterns to add licenses")
	flag.StringVar(&license, "license", "", "sets license to be added")

	flag.Func("file", "file to be processed", handleFile)

	flag.Func("format", "output format - json/xml (default: json)", func(value string) error {
		switch value {
		case "json":
			outputFormat = cyclonedx.BOMFileFormatJSON
		case "xml":
			outputFormat = cyclonedx.BOMFileFormatXML
		default:
			return fmt.Errorf("invalid output format")
		}
		return nil
	})
	flag.Func("output", "output file (default: stdout)", func(value string) error {
		file, err := os.Create(value)

		if err != nil {
			fmt.Printf("unable to create file %s\n", value)
			return err
		}
		output = file
		return nil
	})

	flag.Parse()
}

func handleFile(value string) error {

	bom, err := parseSBOM(value)

	if err != nil {
		fmt.Printf("unable to parse file %s\n", value)
		return err
	}

	patterns := make(map[string][]string)

	if len(patternFile) > 0 {
		patterns, err = parsePatterns(patternFile)

		if err != nil {
			fmt.Printf("unable to parse patterns file %s\n", patternFile)
			return err
		}

	} else if len(pattern) > 0 && len(license) > 0 {
		patterns[pattern] = []string{license}
	} else {
		panic("unable to enrich without license and pattern")
	}

	enrich.Enrich(bom, patterns, force)

	encoder := cyclonedx.NewBOMEncoder(output, outputFormat)
	encoder.Encode(bom)

	return nil
}

func parseSBOM(value string) (*cyclonedx.BOM, error) {

	file, err := openFile(value)

	if err != nil {
		fmt.Printf("unable to open file %s\n", value)
		return nil, err
	}

	bom := &cyclonedx.BOM{}
	decoder := cyclonedx.NewBOMDecoder(file, cyclonedx.BOMFileFormatJSON)
	err = decoder.Decode(bom)

	if err != nil {
		return nil, err
	}

	return bom, err
}

func parsePatterns(value string) (map[string][]string, error) {
	output := make(map[string][]string)

	file, err := openFile(value)

	if err != nil {
		fmt.Printf("unable to open file %s\n", value)
		return nil, err
	}

	byteValue, err := io.ReadAll(file)

	if err != nil {
		fmt.Printf("unable to open file %s\n", value)
		return nil, err
	}

	if err := json.Unmarshal(byteValue, &output); err != nil {
		fmt.Printf("unable to parse file %s\n", value)
		return nil, err
	}

	return output, nil
}

func openFile(value string) (*os.File, error) {
	if _, err := os.Stat(value); os.IsNotExist(err) {
		fmt.Printf("file %s doesn't exist\n", value)
		return nil, err
	}

	return os.Open(value)
}
