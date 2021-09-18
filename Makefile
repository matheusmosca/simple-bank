NAME=simple-bank
NAME_COMMAND_HANDLER=server
VERSION=dev
OS ?= linux
PROJECT_PATH ?= github.com/matheusmosca/simple-bank
TERM=xterm-256color
CLICOLOR_FORCE=true
RICHGO_FORCE_COLOR=1
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_BUILD_TIME=$(shell date '+%Y-%m-%d__%I:%M:%S%p')

.PHONY: dev-docker
dev-docker:
	@echo "==> Starting application..."
	docker-compose up --build

.PHONY: dev-local
dev-local:
	@echo "==> Starting application..."
	go mod download
	docker-compose up --d simple_bank_db
	go run cmd/api/main.go

.PHONY: test
test:
	@echo "==> Running Tests..."
	go test -v ./...

.PHONY: test-coverage
test-coverage:
	@echo "==> Checking test coverage..."
	go install github.com/kyoh86/richgo@latest
	@richgo test -failfast -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

.PHONY: generate
generate:
	@echo "==> Go Generating..."
	go get github.com/kevinburke/go-bindata/...
	@go generate ./...
	go mod tidy

.PHONY: lint
lint:
ifeq (, $(shell which $$(go env GOPATH)/bin/golangci-lint))
	@echo "==> golangci-lint not installed, trying to install it..."
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.42.1
endif
	$$(go env GOPATH)/bin/golangci-lint run -c ./.golangci.yml ./...
