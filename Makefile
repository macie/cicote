.POSIX:
.SUFFIXES:

# MAIN TARGETS

all: install-dependencies

clean:
	@echo '# Remove build cache: go clean -cache' >&2
	@go clean -cache

info:
	@printf '# OS info: '
	@uname -rsv;
	@echo '# Development dependencies:'
	@go version || true
	@echo '# Go environment variables:'
	@go env || true

check:
	@echo '# Static analysis: go vet' >&2
	@go vet
	
test:
	@echo '# Unit tests: go test .' >&2
	@go test .

install-dependencies:
	@echo '# Install dependencies:' >&2
	@go get -v -x .
