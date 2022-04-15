SHELL := /bin/bash

PROJECT := project
MODULE := github.com/duclos-cavalcanti/go-org

GO ?= go
FMT ?= gofmt
DEBUG ?= dlv debug
DIR := cmd

SRC ?= $(shell find $(DIR) -name "*.go" -type f)
FLAGS ?=
LDFLAGS ?=


all:

.PHONY: init
init:
	$(GO) mod init $(MODULE)

.PHONY: build
build:
	@$(GO) build -v -o $(PROJECT) $(FLAGS) cmd/*.go

.PHONY: fmt
fmt:
	$(FMT) -w $(SRC)

.PHONY: debug
debug:
	$(DEBUG) ${DIR}/main.go

# UNTESTED!
.PHONY: test
test:
	@go generate -v ./...
	@go test -race -v ./...

.PHONY: tidy
tidy:
	$(GO) mod tidy

.PHONY: run
run: build
	@./$(PROJECT) --dir test/numerical

.PHONY: clean
clean:
	@echo "## Cleaning Project ##"
	$(GO) clean -modcache
