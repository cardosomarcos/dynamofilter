PROJECT_NAME := "dynamofilter"
PKG := "gitlab.com/cardosomarcos/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all dep build clean test coverage coverhtml lint

all: build

race: dep ## Run data race detector
  @go test -race -short ${PKG_LIST}

msan: dep ## Run memory sanitizer
  @go test -msan -short ${PKG_LIST}

dep: ## Get the dependencies
  @go get -v -d ./...

build: dep ## Build the binary file
  @go build -i -v $(PKG)

clean: ## Remove previous build
  @rm -f $(PROJECT_NAME)

help: ## Display this help screen
  @grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
