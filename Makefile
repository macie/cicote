.POSIX:
.SUFFIXES:

# MAIN TARGETS

all: test check build

clean:
	@echo '# Remove build cache: go clean -cache' >&2
	@go clean -cache

build:
	@echo '# Build production version: go build -ldflags="-s -w" .' >&2
	@go build -ldflags="-s -w" .

debug:
	@printf '# OS info: '
	@uname -rsv;
	@echo '# Development dependencies:'
	@echo; go version || true

check:
	@printf '# Static analysis: go vet .' >&2
	@go vet .
	
test:
	@echo '# Unit tests: go test .' >&2
	@go test .

