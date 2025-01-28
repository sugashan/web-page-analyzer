# Define variables
GO = go
BINARY_NAME = web-page-analyzer
BUILD_DIR = build
LINTER = revive

# Default target when running 'make'
all: build

# Build the Go project
build:
	@echo "Building the project..."
	$(GO) build -o $(BINARY_NAME) .

# Run the project
run: build
	@echo "Running the project..."
	./$(BINARY_NAME)

# Clean up the build (remove binary)
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)

# Run tests
test:
	@echo "Running tests..."
	$(GO) test ./...

# Install dependencies (for Go modules)
install:
	@echo "Installing dependencies..."
	$(GO) mod tidy

# Lint the project
lint:
	@echo "Running linter..."
	revive -formatter friendly ./...

# Run modd to watch for changes and rebuild automatically
modd:
	@echo "Running modd for live-reloading..."
	modd


# Help: show available make targets
help:
	@echo "Makefile for web-page-analyzer project"
	@echo "Available targets:"
	@echo "  build      - Build the project"
	@echo "  run        - Run the project"
	@echo "  clean      - Clean up the build artifacts"
	@echo "  test       - Run tests"
	@echo "  install    - Install dependencies"
	@echo "  lint    	- Lint the project"
	@echo "  modd    	- Live reload project for dev"
	@echo "  help       - Show this help message"
