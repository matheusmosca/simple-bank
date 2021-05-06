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

.PHONY: start
start: 
	@echo "==> Starting application..."
	docker-compose up --build

.PHONY: test
test:
	@echo "==> Running Tests"
	go test -v ./...

.PHONY: test-coverage
test-coverage:
	@echo "Running tests"
	go get github.com/kyoh86/richgo
	@richgo test -failfast -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	rm coverage.out
	rm coverage.html

.PHONY: generate
generate:
	@echo "Go Generating"
	go get github.com/kevinburke/go-bindata/...
	@go generate ./...
	go mod tidy

