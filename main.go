package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/fnxpt/cyclonedx-enrich/enrich"

	"github.com/CycloneDX/cyclonedx-go"
)

var version = "0.0.1"
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

	flag.StringVar(&pattern, "pattern", "", "sets the pattern to add license")
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
	if _, err := os.Stat(value); os.IsNotExist(err) {
		fmt.Printf("file %s doesn't exist\n", value)
		return err
	}

	file, err := os.Open(value)

	if err != nil {
		fmt.Printf("unable to open file %s\n", value)
		return err
	}

	bom, err := parseSBOM(file)

	if err != nil {
		fmt.Printf("unable to parse file %s\n", value)
		return err
	}

	if len(pattern) == 0 || len(license) == 0 {
		panic("unable to enrich without license and pattern")
	}

	enrich.Enrich(bom, pattern, license)

	encoder := cyclonedx.NewBOMEncoder(output, outputFormat)
	encoder.Encode(bom)

	return nil
}

func parseSBOM(input io.Reader) (*cyclonedx.BOM, error) {

	bom := &cyclonedx.BOM{}
	decoder := cyclonedx.NewBOMDecoder(input, cyclonedx.BOMFileFormatJSON)
	err := decoder.Decode(bom)

	if err != nil {
		return nil, err
	}

	return bom, err
}
