# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOTOOL=$(GOCMD) tool
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOINST=$(GOCMD) install

# Binary name
BINARY_NAME=app

# Build
build:
	@$(GOBUILD) -o $(BINARY_NAME) .
	@echo "📦 Build Done"

# Clean
clean:
	@$(GOCLEAN)
	@rm -f $(BINARY_NAME)
	@rm -f test.out
	@echo "🧹 Program removed"

# Generate the doc
doc:
	@$(GOINST) github.com/swaggo/swag/cmd/swag@latest
	@swag init --parseDependency=true -g app.go >> output.out
	@rm output.out
	@echo "📓 Docs Generated"

# Run apps from development
dev:
	@$(GOCMD) run .

# Build and run
run: doc build
	@echo "🚀 Running App"
	@./$(BINARY_NAME)