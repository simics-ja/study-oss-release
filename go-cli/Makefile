# Build variables
BINARY_NAME=oss-release-demo-cli
VERSION=1.0.0
BUILD_DIR=dist
LDFLAGS=-ldflags "-X main.version=${VERSION}"

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Platforms
PLATFORMS=darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64

.PHONY: all build clean test help

all: clean build-all

help:
	@echo "Available targets:"
	@echo "  build        - Build for current platform"
	@echo "  build-all    - Build for all platforms"
	@echo "  clean        - Remove build artifacts"
	@echo "  test         - Run tests"
	@echo "  run          - Run the application"

build:
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) -v

build-all:
	@mkdir -p $(BUILD_DIR)
	@for platform in $(PLATFORMS); do \
		GOOS=$$(echo $$platform | cut -d/ -f1) \
		GOARCH=$$(echo $$platform | cut -d/ -f2) \
		OUTPUT_NAME=$(BINARY_NAME)-$$(echo $$platform | sed 's/\//-/g'); \
		if [ "$$(echo $$platform | cut -d/ -f1)" = "windows" ]; then \
			OUTPUT_NAME="$$OUTPUT_NAME.exe"; \
		fi; \
		echo "Building $$OUTPUT_NAME..."; \
		GOOS=$$(echo $$platform | cut -d/ -f1) \
		GOARCH=$$(echo $$platform | cut -d/ -f2) \
		$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$$OUTPUT_NAME -v; \
	done
	@echo "Build complete. Binaries in $(BUILD_DIR)/"
	@ls -la $(BUILD_DIR)/

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

run:
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) -v
	./$(BUILD_DIR)/$(BINARY_NAME)