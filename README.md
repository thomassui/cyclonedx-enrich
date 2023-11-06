# cyclonedx-enrich
![Coverage](https://img.shields.io/badge/Coverage-0%25-red)
Enrich cyclonedx files with a pattern

## Install

```
go install github.com/fnxpt/cyclonedx-enrich@latest
```

## Run with docker

```
docker run -v `pwd`:/ fnxpt/cyclonedx-enrich:latest --file bom.json --pattern "com.example" --license "MIT License" > output.json
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