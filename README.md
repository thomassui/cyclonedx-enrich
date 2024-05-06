# cyclonedx-enrich
![Coverage](https://img.shields.io/badge/Coverage-30.7%25-yellow)
[![Go Report Card](https://goreportcard.com/badge/github.com/fnxpt/cyclonedx-enrich)](https://goreportcard.com/report/github.com/fnxpt/cyclonedx-enrich)
[![status](https://github.com/fnxpt/cyclonedx-enrich/actions/workflows/coverage.yaml/badge.svg?branch=main "status")](https://github.com/fnxpt/cyclonedx-enrich/actions)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

Enrich cyclonedx files

## Goal

When generating a sbom file, sometimes its not possible to get all the information, either because the tool we are using is not able to fill that information or because the information is not easily accessible.
With this tool we want to allow increasing the level of quality of generated sboms, by appliyng `enrichers` to the inputed file.
Currently we support the following improvements to the SBOM

* Licenses
* Hashes
* Properties
* References
* Cryptography Information (In development)

## Data

Database and Data are stored in this [repo](https://github.com/fnxpt/cyclonedx-enrich-data)

## Enrichers

An enricher is a small struct that provides 2 methods.
- `Skip` - A method that defines if the enricher should be executed
- `Enrich` - A method which will apply the enrichments

|Type|Enricher|Description|Requires|
|---|---|---|---|
|License|RegexpEnricher|Checks if the components don't have licenses and tries to enrich them based on a regexp file provided|Regexp|
|License|DatabaseEnricher|Checks if the components don't have licenses and tries to enrich them based on a database|DB|
|Hash|RegexpEnricher|Checks if the components don't have hashes and tries to enrich them based on a regexp file provided|Regexp|
|Hash|DatabaseEnricher|Checks if the components don't have hashes and tries to enrich them based on a database|DB|
|Property|RegexpEnricher|Checks if the components don't have properties and tries to enrich them based on a regexp file provided|Regexp|
|Property|DatabaseEnricher|Checks if the components don't have properties and tries to enrich them based on a database|DB|
|Reference|RegexpEnricher|Checks if the components don't have references and tries to enrich them based on a regexp file provided|Regexp|
|Reference|DatabaseEnricher|Checks if the components don't have references and tries to enrich them based on a database|DB|
|Manager|CocoapodsEnricher|Enrich components based on the information available in the [cocoapods](https://cocoapods.org)|-|
|Manager|MavenEnricher|Enrich components based on the information available in the [maven](https://central.sonatype.com)|-|
|Manager|NPMEnricher|Enrich components based on the information available in the [npm](https://npmjs.com)|-|
|Manager|PyPiEnricher|Enrich components based on the information available in the [pypi](https://pypi.org)|-|

## Install

```
go install github.com/fnxpt/cyclonedx-enrich@latest
```

## Usage
```
Usage:
  -database-download
    	Downloads database from source
  -database-import value
    	Imports cyclonedx component into database
  -database-register
    	Registers database entities
  -sbom-enrich value
    	Enrichs sbom
  -sbom-validate value
    	Validates sbom
  -server
    	Starts server
```

### Database Download

Downloads the database file located in [repo](https://github.com/fnxpt/cyclonedx-enrich-data).
Set environment variable `DOWNLOAD_DATABASE_URL` to specify a different path

```
cyclonedx-enrich --download-database
```

### Database Import

Imports a sbom file into the database

```
cyclonedx-enrich -database-import bom.json
```

### SBOM Enrich

Enriches a sbom file, the enriched version overides the input version

```
cyclonedx-enrich -sbom-enrich bom.json
```

### SBOM Validate

Validates a sbom file

```
cyclonedx-enrich -sbom-validate bom.json
```

### Server

Starts a enricher server on port 8080.

```
cyclonedx-enrich -server
```

### Run server with docker

```
docker run --env-file ./.env -p 8080:8080 fnxpt/cyclonedx-enrich:latest --database-download --server
```

### Example to call server

```
curl -X POST http://127.0.0.1:8080/sbom/enrich -H "X-Api-Key: DUMMY" -d @bom.json
```

## Examples

|Input|Output|
|---|---|
|[Bom 1](testdata/sbom/bom1.json)|[Result](testdata/sbom/bom1.json)|


