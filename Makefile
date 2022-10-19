all: prepare build test
.PHONY: all

.ONESHELL:
prepare:
	export GOPRIVATE="git.your-company-hub.io"
	go mod download -x
	go generate ./...
	go mod tidy
.PHONY: prepare

build:
	go build -o application .
.PHONY: build

test:
	go test -v -coverprofile coverage.out ./
	go tool cover -func coverage.out
.PHONY: test

lint:
	golangci-lint run
.PHONY: lint
