#!/bin/sh

echo "Checking Go file formatting..."

files=$(find . -name '*.go')

for file in $files; do
    result=$(gofmt -l $file)
    if [ ! -z "$result" ]; then
        echo "File $file is not formatted correctly. Please run 'gofmt -w $file' to format it."
    fi
done

echo "Format check completed."
