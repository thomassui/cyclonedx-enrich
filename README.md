# cyclonedx-enrich
![Coverage](https://img.shields.io/badge/Coverage-0.0%25-red)

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