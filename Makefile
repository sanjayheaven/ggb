
.PHONY: format
# Run go fmt against code
format:
	go fmt ./...

.PHONY: fmt
# fmt is an alias for format
fmt: format

.PHONY: goimports
# Run goimports against code
goimports:
	goimports -w .

.PHONY: vet
# Run go vet against code
vet:
	go vet ./...

.PHONY: lint
# lint the code
lint:
	golangci-lint run

.PHONY: build
# build the binary
build: 
	go build -v -o build/ggb main.go


.PHONY: cache
# clean cache
clean:
	rm -rf build/ggb

.PHONY: swagger
# swager docs generation
swagger:
	swag init -g main.go -o api/swagger


.PHONY: tidy
# tidy the go modules
tidy: 
	go mod tidy


.PHONY: gt generate-template
# generate template, THIS IS A DEV ONLY COMMAND DO NOT USE IN PRODUCTION
gt: generate-template
generate-template:
	go run scripts/gen-tmpl.go

.PHONY: all
# default target
all: goimports format vet lint tidy

