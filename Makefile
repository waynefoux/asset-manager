# I copied this from an old work project and wanted to see if I can mod it for my use case.
.PHONY: all build test clean build-linux build-arm docker-build

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_NAME=asset-manager
VERSION=1.0-SNAPSHOT
LINUX_ARCH=linux/$(BINARY_NAME)
ARM_ARCH=arm/$(BINARY_NAME)
MAC_ARCH=osx/$(BINARY_NAME)
DEFAULT_ARCH=$(LINUX_ARCH)

all: test build
build:
		$(GOBUILD) -ldflags "-X main.Version=${VERSION}" -o build/$(DEFAULT_ARCH) -v
test:
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -rf build/$(MAC_ARCH)
		rm -rf build/$(LINUX_ARCH)
		rm -rf build/$(ARM_ARCH)
vendor:
		dep ensure
# Cross compilation
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags "-X main.Version=${VERSION}" -o build/$(LINUX_ARCH) -v
build-arm:
		CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GOBUILD) -ldflags "-X main.Version=${VERSION}" -o build/$(LINUX_ARCH) -v
docker-build: build-arm
		docker build --no-cache -t "waynefoux/$(BINARY_NAME):$(VERSION)" . && \
		echo "Version $(VERSION)"
nope:
		docker run --rm -it -v "$(GOPATH)":/go -w /opt/app/$(BINARY_NAME) golang:latest go build -o build/"$(LINUX_ARCH)" -v