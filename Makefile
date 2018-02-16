SOURCE_FILES?=$$(go list ./... | grep -v /vendor/)
TEST_PATTERN?=.
TEST_OPTIONS?=
GO ?= go

# Install all the build and lint dependencies
setup:
	go get -u github.com/alecthomas/gometalinter
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/pierrre/gotestcover
	go get -u golang.org/x/tools/cmd/cover
	gometalinter --install
	dep ensure
.PHONY: setup

# Install from source.
install:
	@echo "==> Installing up ${GOPATH}/bin/lfs-profile-golang"
	@$(GO) install ./...
.PHONY: install

# Run all the tests
test:
	@gotestcover $(TEST_OPTIONS) -covermode=atomic -coverprofile=coverage.txt $(SOURCE_FILES) -run $(TEST_PATTERN) -timeout=2m
.PHONY: test

# Run all the tests and opens the coverage report
cover: test
	@$(GO) tool cover -html=coverage.txt
.PHONY: cover

# Run all the linters
lint:
	gometalinter --vendor ./...
.PHONY: lint

# Run all the tests and code checks
ci: setup test lint
.PHONY: ci

swagger: 
	swagger -apiPackage="github.com/kebabmane/tureloGo" -mainApiFile=github.com/kebabmane/tureloGo/server.go -output=./API.md -format=markdown
.PHONY: swagger