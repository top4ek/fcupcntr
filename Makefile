SHELL := /bin/sh

GO := go
GOLANGCI_LINT_VERSION ?= v2.14.0
GOLANGCI_LINT := $(if $(shell command -v golangci-lint 2>/dev/null),$(shell command -v golangci-lint 2>/dev/null),$(shell $(GO) env GOPATH)/bin/golangci-lint)

.PHONY: help build test test-race fmt vet lint lint-install ci dev dev-down compose-check

help: ## Show available targets
	@awk 'BEGIN {FS = ":.*## "} /^[a-zA-Z_-]+:.*## / {printf "%-16s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the application binary
	mkdir -p bin
	$(GO) build -buildvcs=false -o bin/fcupcntr ./cmd/fcupcntr

test: ## Run Go tests
	$(GO) test ./... -v

test-race: ## Run Go tests with the race detector
	$(GO) test -race ./...

fmt: ## Check gofmt formatting
	@test -z "$$(gofmt -l .)"

vet: ## Run go vet
	$(GO) vet ./...

lint: ## Run golangci-lint
	@test -x "$(GOLANGCI_LINT)" || { echo 'golangci-lint is missing; run make lint-install' >&2; exit 1; }
	$(GOLANGCI_LINT) run ./...

lint-install: ## Install the pinned golangci-lint version
	$(GO) install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

ci: fmt vet lint test-race build ## Run all CI checks

compose-check: ## Validate Compose configuration
	docker compose config --quiet

dev: ## Build and start the application container
	docker compose up --build

dev-down: ## Stop the application container
	docker compose down
