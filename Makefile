.PHONY: all build vendor test lint

all:
	go test -v ./...
