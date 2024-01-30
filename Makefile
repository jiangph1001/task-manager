# Go parameters
GOCMD=go
GOBUILD=GO111MODULE=off $(GOCMD) build
GOCLEAN=GO111MODULE=auto $(GOCMD) clean
BINARY_NAME=bin/main

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/main.go
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)-linux -v ./cmd/main.go
