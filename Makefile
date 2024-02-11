BINARY_NAME := "wisee"

ARCH := $(or $(GOARCH),$(shell uname -m))
OS := $(or $(GOOS),$(shell uname))

ifneq (,$(filter $(OS),Darwin darwin MacOS macos))
	OS := darwin
else ifneq (,$(filter $(OS),Linux linux))
	OS := linux
else
	OS := windows
endif

ifeq ($(ARCH),x86_64)
	ARCH := amd64
else ifeq ($(ARCH),i386)
	ARCH := 386
endif

build: $(BUILDDEPS)
	@echo "Building $(OS) $(ARCH) binary..."
	@GOOS=$(OS) GOARCH=$(ARCH) go build $(ARGS) -o "bin/$(BINARY_NAME)" src/main.go

test:
	@echo "Running tests..."
	@@GOOS=$(OS) GOARCH=$(ARCH) ENV=test go test -race -covermode=atomic -v -coverpkg=./src/...  ./tests/... ./src/...

clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf bin/*

run: build
	@echo "Running..."
	@
