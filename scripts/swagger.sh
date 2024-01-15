#!/bin/sh

# set a variable to the path of swag
SWAG=$(go env GOPATH)/bin/swag

# Check if swag is installed, if not, install it
if ! command -v $SWAG >/dev/null; then
    echo "swag not found, downloading..."
    # binary will be $(go env GOPATH)/bin/swag
    go install github.com/swaggo/swag/cmd/swag@latest
    # Congratulations! swag is installed!
    echo "swag installed!"
    # print the version of swag
    echo "$($SWAG --version)"
fi

$SWAG init -g main.go -o api/swagger
