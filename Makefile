
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
	./scripts/lint.sh

.PHONY: build
# build the binary
build: 
	./scripts/build.sh

.PHONY: cache
# clean cache
clean:
	./scripts/clean.sh

.PHONY: swagger
# swager docs generation
swagger:
	./scripts/swagger.sh

.PHONY: all
# default target
all: goimports format lint