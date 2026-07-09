# ==============================================================================
# Coela - Development Makefile
# ==============================================================================

.DEFAULT_GOAL := help

GREEN  := \033[0;32m
YELLOW := \033[1;33m
BLUE   := \033[0;34m
RED    := \033[0;31m
RESET  := \033[0m

.PHONY: help run fmt vet test build install clean \
        release-build release-snapshot release \
        doctor version tidy update

help: ## Show available commands
	@echo ""
	@echo "$(BLUE)Coela Development Commands$(RESET)"
	@echo ""
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z0-9_-]+:.*##/ {printf "  \033[32m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""

# ==============================================================================
# Development
# ==============================================================================

run: ## Run Coela
	go run . new

help-cli: ## Display Coela help
	go run . --help

build: ## Build local binary
	@mkdir -p bin
	go build -o bin/coela .
	@echo "$(GREEN)✓ Binary generated in bin/coela$(RESET)"

install: build ## Install Coela locally
	sudo install -m 755 bin/coela /usr/local/bin/coela
	@echo "$(GREEN)✓ Installed in /usr/local/bin/coela$(RESET)"

version: ## Display installed version
	coela --version

# ==============================================================================
# Quality
# ==============================================================================

fmt: ## Format source code
	go fmt ./...

vet: ## Run go vet
	go vet ./...

test: ## Run unit tests
	go test ./...

check: fmt vet test ## Run all quality checks

# ==============================================================================
# Dependencies
# ==============================================================================

tidy: ## Clean go.mod / go.sum
	go mod tidy

update: ## Update dependencies
	go get -u ./...
	go mod tidy

# ==============================================================================
# Release
# ==============================================================================

release-build: ## Build release artifacts (no GitHub release)
	goreleaser build --clean

release-snapshot: ## Simulate a release locally
	goreleaser release --snapshot --clean

release-check: ## Validate GoReleaser configuration
	goreleaser check

# ==============================================================================
# Utilities
# ==============================================================================

clean: ## Remove generated files
	rm -rf bin
	rm -rf dist

doctor: ## Display development environment
	@echo ""
	@echo "$(BLUE)Environment$(RESET)"
	@echo ""
	@go version
	@docker --version
	@docker compose version
	@git --version
	@goreleaser --version