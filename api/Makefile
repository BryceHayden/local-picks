# Simple Makefile for a Go project

# Build the application
all: clean build-prod run test clean

# Note if using HTML templates then:
# @templ generate
build-dev:
	@echo "Building binary..."
	@go build -o ./build/main cmd/api/main.go

build-prod: clean
	@echo "Building production binary..."
	@go build -ldflags "-s -w" -o ./build/main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./tests -v

# Clean the binary
clean:
	@if [ -f build/main ]; then \
		rm build/main; \
		echo "Deleted main"; \
	fi
	
# Live Reload
watch: clean build-dev
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


