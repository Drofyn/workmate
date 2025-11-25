APP_NAME=workmate
PKG=github.com/drofyn/workmate

VERSION=$(shell git describe --tags --always 2>/dev/null)
COMMIT=$(shell git rev-parse HEAD)
BUILDDATE=$(shell date +'%Y-%m-%dT%H:%M:%SZ')
GOVERSION=$(shell go version | awk '{print $$3}')

LDFLAGS=-ldflags "-s -w \
	-X $(PKG)/cmd/root.Version=$(VERSION) \
	-X $(PKG)/cmd/root.Commit=$(COMMIT) \
	-X $(PKG)/cmd/root.BuildDate=$(BUILDDATE) \
	-X $(PKG)/cmd/root.GoVersion=$(GOVERSION)"

build:
	go build $(LDFLAGS) -o $(APP_NAME)

build-all: build
	GOOS=linux   GOARCH=amd64 go build $(LDFLAGS) -o build/$(APP_NAME)-linux-amd64
	GOOS=darwin  GOARCH=arm64 go build $(LDFLAGS) -o build/$(APP_NAME)-darwin-arm64

clean:
	rm -rf build $(APP_NAME)

.PHONY: build build-all clean
