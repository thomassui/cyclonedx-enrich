# cyclonedx-enrich
![Coverage](https://img.shields.io/badge/Coverage-30.7%25-yellow)
[![go report card](https://goreportcard.com/badge/github.com/fnxpt/cyclonedx-enrich "go report card")](https://goreportcard.com/report/github.com/fnxpt/cyclonedx-enrich)
[![status](https://github.com/fnxpt/cyclonedx-enrich/actions/workflows/coverage.yaml/badge.svg?branch=main "status")](https://github.com/fnxpt/cyclonedx-enrich/actions)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

Enrich cyclonedx files

## Install

```
go install github.com/fnxpt/cyclonedx-enrich@latest
```

## Run server with docker

```
docker run --env-file ./.env fnxpt/cyclonedx-enrich:latest --server
```

## Example to call server

```
curl -X POST http://127.0.0.1:8080/sbom/enrich -H "X-Api-Key: DUMMY" -d @bom.json
```

## Data

Database and Data is stored in this [repo](https://github.com/fnxpt/cyclonedx-enrich-data)

## Usage
```
Usage:
    -database-download
    	Downloads database from source
  -database-import value
    	Imports cyclonedx component into database
  -database-register
    	Registers database entities
  -server
    	Starts server
```