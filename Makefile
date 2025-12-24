# ASCII-ART Makefile
# Go parameters
BINARY_NAME=ascii-art
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_WINDOWS=$(BINARY_NAME).exe
GO_FILES=$(shell find . -name '*.go' -type f)
CMD_PATH=./cmd/ascii-art
INTERNAL_PATH=./internal/ascii

.PHONY: all build clean test test-verbose test-coverage run install help fmt vet lint

# Default target
all: clean build test

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) $(CMD_PATH)
	@echo "Build complete: $(BINARY_NAME)"

# Build for Linux
build-linux:
	@echo "Building for Linux..."
	@GOOS=linux GOARCH=amd64 go build -o $(BINARY_UNIX) $(CMD_PATH)
	@echo "Linux build complete: $(BINARY_UNIX)"

# Build for Windows
build-windows:
	@echo "Building for Windows..."
	@GOOS=windows GOARCH=amd64 go build -o $(BINARY_WINDOWS) $(CMD_PATH)
	@echo "Windows build complete: $(BINARY_WINDOWS)"

# Build for all platforms
build-all: build-linux build-windows
	@echo "All platform builds complete"

# Run tests
test:
	@echo "Running tests..."
	@go test ./...
	@echo "Tests complete"

# Run tests with verbose output
test-verbose:
	@echo "Running tests (verbose)..."
	@go test ./... -v

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run unit tests only
test-unit:
	@echo "Running unit tests..."
	@go test $(INTERNAL_PATH) -v

# Run edge case tests only
test-edge:
	@echo "Running edge case tests..."
	@go test -v -run TestEdge

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@go clean
	@rm -f $(BINARY_NAME)
	@rm -f $(BINARY_UNIX)
	@rm -f $(BINARY_WINDOWS)
	@rm -f coverage.out coverage.html
	@echo "Clean complete"

# Run with example input
run: build
	@echo "Running example: Hello World"
	@./$(BINARY_NAME) "Hello"

# Run with shadow style
run-shadow: build
	@echo "Running example with shadow style"
	@./$(BINARY_NAME) "Hello" --shadow

# Run with thinkertoy style
run-thinkertoy: build
	@echo "Running example with thinkertoy style"
	@./$(BINARY_NAME) "Hello" --thinkertoy

# Install to $GOPATH/bin
install:
	@echo "Installing $(BINARY_NAME)..."
	@go install $(CMD_PATH)
	@echo "Install complete"

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Format complete"

# Run go vet
vet:
	@echo "Running go vet..."
	@go vet ./...
	@echo "Vet complete"

# Check for common issues
lint: fmt vet
	@echo "Linting complete"

# Show help
help:
	@echo "ASCII-ART Makefile Commands:"
	@echo ""
	@echo "Build Commands:"
	@echo "  make build          - Build the binary"
	@echo "  make build-linux    - Build for Linux (amd64)"
	@echo "  make build-windows  - Build for Windows (amd64)"
	@echo "  make build-all      - Build for all platforms"
	@echo ""
	@echo "Test Commands:"
	@echo "  make test           - Run all tests"
	@echo "  make test-verbose   - Run tests with verbose output"
	@echo "  make test-coverage  - Run tests with coverage report"
	@echo "  make test-unit      - Run unit tests only"
	@echo "  make test-edge      - Run edge case tests only"
	@echo ""
	@echo "Run Commands:"
	@echo "  make run            - Build and run with example (standard style)"
	@echo "  make run-shadow     - Build and run with shadow style"
	@echo "  make run-thinkertoy - Build and run with thinkertoy style"
	@echo ""
	@echo "Utility Commands:"
	@echo "  make clean          - Remove build artifacts"
	@echo "  make install        - Install to $$GOPATH/bin"
	@echo "  make fmt            - Format code with go fmt"
	@echo "  make vet            - Run go vet"
	@echo "  make lint           - Run fmt and vet"
	@echo "  make help           - Show this help message"
	@echo ""
	@echo "Common Workflows:"
	@echo "  make all            - Clean, build, and test"
	@echo "  make && make run    - Build and run quick demo"
