# Makefile for the tree utility

# Variables
BINARY_NAME := tree
GO := go
GOFLAGS := -trimpath -ldflags="-s -w"

# Default target
.PHONY: all
all: build

# Check if Go is installed and dependencies are satisfied
.PHONY: check-deps
check-deps:
	@which $(GO) > /dev/null || (echo "Error: Go is not installed"; exit 1)
	@$(GO) version
	@echo "Checking dependencies..."
	@$(GO) mod tidy

# Build the binary with optimizations
.PHONY: build
build: check-deps
	@echo "Building $(BINARY_NAME)..."
	@$(GO) build $(GOFLAGS) -o $(BINARY_NAME)

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	@rm -f $(BINARY_NAME)

# Install the binary to system
.PHONY: install
install: build
	@echo "Installing $(BINARY_NAME)..."
	@install -m 755 $(BINARY_NAME) /usr/local/bin/

# Run tests
.PHONY: test
test: check-deps
	@echo "Running tests..."
	@$(GO) test -v ./...

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	@$(GO) fmt ./...

# Show help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  all        : Default target, builds the binary"
	@echo "  check-deps : Check if Go is installed and dependencies are satisfied"
	@echo "  build      : Build the binary with optimizations"
	@echo "  clean      : Remove build artifacts"
	@echo "  install    : Install the binary to /usr/local/bin"
	@echo "  test       : Run tests"
	@echo "  fmt        : Format code"
	@echo "  help       : Show this help message"