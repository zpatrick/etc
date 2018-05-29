SHELL:=/bin/bash
VERSION?=$(shell git describe --tags --always)

build: 
	go build -o rp

install:
	go install

release:
	CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build -ldflags "-s -X main.Version=$(VERSION)" -a -o build/linux/rp .
	CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -ldflags "-s -X main.Version=$(VERSION)" -a -o build/darwin/rp .
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -X main.Version=$(VERSION)" -a -o build/windows/rp.exe .

.PHONY: build install release
