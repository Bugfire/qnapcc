#

.PHONY: $(shell egrep -o ^[a-zA-Z_-]+: $(MAKEFILE_LIST) | sed 's/://')

NAME := qnapcc
VERSION := 0.1.0
REVISION := $(shell git rev-parse --short HEAD)
PKG := github.com/bugfire/qnapcc
LDFLAGS := "-X ${PKG}/cmd.Version=${VERSION} -X ${PKG}/cmd.Revision=${REVISION}"

all: help

tools:
	go generate -tags tools

gen: tools
	go generte ./...

build: ## Build 
	go build -ldflags=${LDFLAGS} -o=./bin/qnapcc

install: ## Install
	go install -ldflags=${LDFLAGS} 

lint: tools ## Lint
	reviewdog -diff="git diff master"

test: ## Test
	go test -race -v ./...

cover: ## Cover
	go test -race -v coverprofile coverage.txt -covermode atomic ./...

help: ## This help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
