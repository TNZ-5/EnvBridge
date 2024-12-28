# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=envbridge
BINARY_UNIX=$(BINARY_NAME)_unix
GOTIDY=$(GOCMD) mod tidy

# Clean build files
clean: 
	$(GOCLEAN)
ifeq ($(OS),Windows_NT)
	del $(BINARY_NAME).exe 
else
	rm $(BINARY_UNIX)
endif


# Run the application
run: build
ifeq ($(OS),Windows_NT)
	 $(BINARY_NAME).exe 
else
	 $(BINARY_UNIX)
endif


# Install dependencies
deps:
	$(GOGET) -v ./...

# Cross compilation for Linux and Windows
build-linux:
	$(GOBUILD) -o $(BINARY_UNIX) -v ./env

build-windows:
	$(GOBUILD) -o $(BINARY_NAME).exe -v ./env

# Automatically decide which system is being used
build: deps
	$(GOTIDY)
ifeq ($(OS),Windows_NT)
	$(MAKE) build-windows
else
	$(MAKE) build-linux
endif

.PHONY: all test build clean run deps build-linux build-windows build-auto