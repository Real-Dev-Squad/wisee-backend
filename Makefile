BINARY_NAME := "wisee"
DEV_DATABASE_URL := "postgres://postgres:postgres@localhost:5432/wisee_core?sslmode=disable"

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
else ifeq ($(ARCH),aarch64)
    ARCH := arm64
endif

build:
	@echo "Building $(OS) $(ARCH) binary..."
	@GOOS=$(OS) GOARCH=$(ARCH) CGO_ENABLED=0 go build $(ARGS) -o "bin/$(BINARY_NAME)" ./src

test_unit:
	@echo "Running tests..."
	@@GOOS=$(OS) GOARCH=$(ARCH) ENV=test go test -race -covermode=atomic -v -coverpkg=./src/...  ./tests/unit/... ./src/...

test_integration:
	@echo "Running tests..."
	@@GOOS=$(OS) GOARCH=$(ARCH) ENV=test go test -race -covermode=atomic -v -coverpkg=./src/...  ./tests/integration/... ./src/...

test: test_unit test_integration

clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf bin/*

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

# Up all migrations
migrate-all-up:
	@if command -v migrate > /dev/null; then \
	    migrate -database $(DEV_DATABASE_URL) -path ./database/migrations up; \
	else \
		echo "Golang Migrate cli is not installed on your machine. Exiting..."; \
		exit 1; \
	fi

# Drop all migrations when in development
migrate-all-down:
	@if command -v migrate > /dev/null; then \
	    migrate -database $(DEV_DATABASE_URL) -path ./database/migrations down; \
	else \
		echo "Golang Migrate cli is not installed on your machine. Exiting..."; \
		exit 1; \
	fi

# Setup Databse and PGAdmin
docker-run:
	@if command -v docker > /dev/null; then \
	    docker-compose -f docker/dev-docker-compose.yaml up -d; \
	else \
		echo "Docker is not installed on your machine. Exiting..."; \
		exit 1; \
	fi

# Down Database and PGAdmin
docker-down:
	@if command -v docker > /dev/null; then \
	    docker-compose -f docker/dev-docker-compose.yaml down; \
	else \
		echo "Docker is not installed on your machine. Exiting..."; \
		exit 1; \
	fi

# Setup development environment 
setup:
	@echo "--- Copying .env files ---"
	@cp environments/dev.env .env

	@echo "--- Setting up docker ---"
	@make docker-run

	@echo "--- Waiting for database to setup ---"
	@sleep 3

	@echo "--- Running all migrations ---"
	@make migrate-all-up
	@echo "\n"
	@echo "Setup complete. To start the server, run 'make watch'"
