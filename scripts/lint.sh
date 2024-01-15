#!/bin/sh

go vet ./...

golangci-lint run
