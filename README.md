[![Tests](https://github.com/aricodes-oss/gofish/actions/workflows/test.yml/badge.svg)](https://github.com/aricodes-oss/gofish/actions/workflows/test.yml)

# gofish

A simple ><> interpreter in Go

## Requirements

Go >= 1.21.0

## Building

```bash
go generate ./...
go build
```

## Running tests

```bash
go test -v ./... -bench=.
```
