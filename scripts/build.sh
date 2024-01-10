#!/bin/bash

SRC="cmd/main.go"
OUTPUT="build/main"

echo "Starting build..."

go build -v -o $OUTPUT $SRC

echo "Build completed."
