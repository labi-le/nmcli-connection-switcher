.PHONY: build
.DEFAULT_GOAL := build

build-release:
	go build -ldflags "-s" -a -v -o build/package/ellie main.go

build:
	go build -v -o build/package/ellie-debug main.go
