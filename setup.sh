#!/bin/bash

go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0
golangci-lint --version
