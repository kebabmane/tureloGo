SHELL := /bin/bash

# The name of the executable (default is current directory name)
TARGET := $(shell echo $${PWD\#\#*/})
.DEFAULT_GOAL: $(TARGET)

# The fully qualified import path of the project from within the $GOPATH
PACKAGE := `pwd | sed 's_/.*/go/src/__'`

# These will be provided to the target
VERSION := 1.0.0
COMMIT := `git rev-parse HEAD`
BRANCH := `git rev-parse --abbrev-ref HEAD`
DATE := `date -u +"%Y-%m-%dT%H:%M:%SZ"`

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=$(PACKAGE)/app.Version=$(VERSION) -X=$(PACKAGE)/app.Commit=$(COMMIT) -X=$(PACKAGE)/app.Branch=$(BRANCH) -X=$(PACKAGE)/app.BuildDate=$(DATE)"

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")


.PHONY: all build clean install uninstall fmt simplify check run swagger

all: check install

$(TARGET): $(SRC)
	@go build $(LDFLAGS) -o $(TARGET)

build: $(TARGET)
	@true

clean:
	@rm -f $(TARGET)

install:
	@go install $(LDFLAGS)

uninstall: clean
	@rm -f $$(which ${TARGET})

fmt:
	@gofmt -l -w $(SRC)

simplify:
	@gofmt -s -l -w $(SRC)

check:
	@test -z $(shell gofmt -l server.go | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@for d in $$(go list ./... | grep -v /vendor/); do golint $${d}; done
	@go tool vet ${SRC}

run: install
	@$(TARGET)

# Run all the linters
lint:
	gometalinter --vendor ./...


swagger:
	swagger -apiPackage="github.com/kebabmane/tureloGo" -mainApiFile=github.com/kebabmane/tureloGo/server.go -output=./API.md -format=markdown