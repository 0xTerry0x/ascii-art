# ASCII-Art Makefile

.PHONY: all build test clean install run run-shadow run-thinkertoy help

# Default target
all: build

# Build the binary
build:
	go build -o ascii-art ./cmd/ascii-art

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -f ascii-art ascii-art.exe

# Install to GOPATH/bin
install:
	go install ./cmd/ascii-art

# Run examples
run:
	go run ./cmd/ascii-art "Hello World"

run-shadow:
	go run ./cmd/ascii-art "Hello" --shadow

run-thinkertoy:
	go run ./cmd/ascii-art "Hello" --thinkertoy

# Show help
help:
	@echo "Available targets:"
	@echo "  make build          - Build the binary"
	@echo "  make test           - Run all tests"
	@echo "  make clean          - Remove build artifacts"
	@echo "  make install        - Install to GOPATH/bin"
	@echo "  make run            - Run example with standard style"
	@echo "  make run-shadow     - Run example with shadow style"
	@echo "  make run-thinkertoy - Run example with thinkertoy style"
	@echo "  make help           - Show this help message"
