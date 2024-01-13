#!/bin/bash

SRC="main.go"
OUTPUT="build/ggb"

echo "Starting build..."

go build -v -o $OUTPUT $SRC

echo "Build completed."
