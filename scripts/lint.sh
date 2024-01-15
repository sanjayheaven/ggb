#!/bin/sh

# Run go vet against code
go vet ./...

# set a variable to the path of golangci-lint
GOLANGCI_LINT=$(go env GOPATH)/bin/golangci-lint

# Check if golangci-lint is installed, if not, install it
if ! command -v $GOLANGCI_LINT >/dev/null; then
    echo "golangci-lint not found, downloading..."
    # binary will be $(go env GOPATH)/bin/golangci-lint
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
    # Congratulations! golangci-lint is installed!
    echo "golangci-lint installed!"
    # print the version of golangci-lint
    echo "$($GOLANGCI_LINT --version)"
fi

$GOLANGCI_LINT run
