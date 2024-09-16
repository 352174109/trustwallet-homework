.PHONY: build test clean

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
BINARY_NAME=bin/trustwallet-homework

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v cmd/*.go

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)