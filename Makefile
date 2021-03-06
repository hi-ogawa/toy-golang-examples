.PHONY: $(shell grep --no-filename -E '^[a-zA-Z_-]+:' $(MAKEFILE_LIST) | sed 's/:.*//')
SHELL := /bin/bash

all:

build:
	go build ./cmd/...

test:
	go test ./...

fmt:
	go fmt -x ./...

fmt/check:
	gofmt -l -d .

vet:
	go vet ./...
