# Makefile for Go REST Service

BINARY = greet_server
SRC = main.go

.PHONY: all build run clean test fmt

all: build

build:
	go build -o $(BINARY) $(SRC)

run:
	go run $(SRC)

clean:
	rm -f $(BINARY)

fmt:
	go fmt ./...

# Build for multiple platforms
build-all: build-linux build-darwin build-windows

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY)-linux-amd64 $(SRC)

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY)-darwin-amd64 $(SRC)

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BINARY)-windows-amd64.exe $(SRC)

