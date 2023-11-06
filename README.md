# cyclonedx-enrich
![Coverage](https://img.shields.io/badge/Coverage-0%25-red)
Enrich cyclonedx files with a pattern

## Install

```
go install github.com/fnxpt/cyclonedx-enrich@latest
```

## Run with docker

```
docker run -v `pwd`/sbom:/sbom fnxpt/cyclonedx-enrich:latest --file sbom/bom.json --pattern "(pkg\:maven\/com.example.+)|(pkg:npm\/(@|%40)example\/.+)" --license "MIT License" > output.json
```

## Usage
```
Usage:
  -file value
    	file to be processed
  -format value
    	output format - json/xml (default: json)
  -license string
    	sets license to be added
  -output value
    	output file (default: stdout)
  -pattern string
    	sets the pattern to add license
```