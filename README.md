# cyclonedx-enrich
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/8b6f3ae91fc54703a427c08bb4002cb3)](https://app.codacy.com/gh/fnxpt/cyclonedx-enrich/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)![Coverage](https://img.shields.io/badge/Coverage-0%25-red)

Enrich cyclonedx files with a pattern

## Install

```
go install github.com/fnxpt/cyclonedx-enrich@latest
```

## Run with docker with pattern

```
docker run -v `pwd`/sbom:/sbom fnxpt/cyclonedx-enrich:latest --file sbom/bom.json --pattern "(pkg\:maven\/com.example.+)|(pkg:npm\/(@|%40)example\/.+)" --license "MIT License" > output.json
```

## Run with docker with pattern file
```
docker run -v `pwd`/sbom:/sbom fnxpt/cyclonedx-enrich:latest --file sbom/bom.json --pattern-file sbom/pattern.json > output.json
```

## Usage
```
Usage:
  -file value
    	file to be processed
  -force
    	sets the license even if its already filled
  -format value
    	output format - json/xml (default: json)
  -license string
    	sets license to be added
  -output value
    	output file (default: stdout)
  -pattern string
    	sets the pattern to add license
  -pattern-file string
    	sets file with the patterns to add licenses
```