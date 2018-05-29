SHELL:=/bin/bash
VERSION?=$(shell git describe --tags --always)

build: 
	go build -o context

release:
	CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build -ldflags "-s -X main.Version=$(VERSION)" -a -o build/linux/context .
	CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -ldflags "-s -X main.Version=$(VERSION)" -a -o build/darwin/context .
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -X main.Version=$(VERSION)" -a -o build/windows/context.exe .

.PHONY: build release
