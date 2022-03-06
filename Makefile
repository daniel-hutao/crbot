SHELL := /bin/bash
BASEDIR = $(shell pwd)

.PHONY: help
help:
		@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

fmt: ## Run 'go fmt' & goimports against code.
	goimports -local="github.com/daniel-hutao/crbot" -d -w cmd
	goimports -local="github.com/daniel-hutao/crbot" -d -w internal
	go fmt ./...

.PHONY: build
build: fmt ## Build crbot
	go build -o crbot cmd/crbot/main.go
