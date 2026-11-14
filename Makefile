SHELL := /bin/sh

GO := go
GOLANGCI_LINT_VERSION ?= v2.14.0
GOLANGCI_LINT := $(if $(shell command -v golangci-lint 2>/dev/null),$(shell command -v golangci-lint 2>/dev/null),$(shell $(GO) env GOPATH)/bin/golangci-lint)

DEV_ENV := $(if $(wildcard build/dev/.env),build/dev/.env,build/dev/.env-example)
PROD_ENV := $(if $(wildcard build/prod/.env),build/prod/.env,build/prod/.env-example)
DEV_COMPOSE := docker compose --project-name fcupcntr-dev --env-file $(DEV_ENV) -f build/dev/compose.yaml
PROD_COMPOSE := docker compose --project-name fcupcntr-prod --env-file $(PROD_ENV) -f build/prod/compose.yaml

.PHONY: help build test test-race fmt vet lint lint-install ci dev dev-down prod prod-down prod-build compose-check

help: ## Show available targets
	@awk 'BEGIN {FS = ":.*## "} /^[a-zA-Z_-]+:.*## / {printf "%-16s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the application binary
	mkdir -p bin
	$(GO) build -buildvcs=false -o bin/fcupcntr ./cmd/fcupcntr

test: ## Run Go tests
	$(GO) test ./... -v

test-race: ## Run Go tests with the race detector
	$(GO) test -race ./... -v

fmt: ## Check gofmt formatting
	@test -z "$$(gofmt -l .)"

vet: ## Run go vet
	$(GO) vet ./...

lint: ## Run golangci-lint
	@test -x "$(GOLANGCI_LINT)" || { echo 'golangci-lint is missing; run make lint-install' >&2; exit 1; }
	$(GOLANGCI_LINT) run ./... -v

lint-install: ## Install the pinned golangci-lint version
	$(GO) install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

ci: fmt vet lint test-race build ## Run all CI checks

compose-check: ## Validate dev and production Compose using committed examples
	docker compose --project-name fcupcntr-dev --env-file build/dev/.env-example -f build/dev/compose.yaml config --quiet
	docker compose --project-name fcupcntr-prod --env-file build/prod/.env-example -f build/prod/compose.yaml config --quiet

dev: ## Start development with Reflex reload and automatic tests
	@test -f build/dev/.env || cp build/dev/.env-example build/dev/.env
	$(DEV_COMPOSE) up --build

dev-down: ## Stop development without deleting SQLite or Go caches
	$(DEV_COMPOSE) down

prod-build: ## Build the local production image
	docker build -t fcupcntr:prod -f build/prod/Containerfile .

prod: ## Build and start production (requires build/prod/.env)
	@test -f build/prod/.env || { echo 'Copy build/prod/.env-example to build/prod/.env and review it first' >&2; exit 1; }
	$(PROD_COMPOSE) up --build

prod-down: ## Stop production without deleting SQLite data
	$(PROD_COMPOSE) down
