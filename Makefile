
# format the code
format:
	./scripts/format.sh

# lint the code
lint:
	./scripts/lint.sh

# build the binary
build: 
	./scripts/build.sh

# clean cache
clean:
	./scripts/clean.sh

# swager docs generation
swagger:
	./scripts/swagger.sh

# default target
all: format lint