BIN=timestamp
PKG=github.com/elliottpolk/timestamp
CLI_VERSION=`cat .version`
COMPILED=`date +%s`
GIT_HASH=`git rev-parse --short HEAD`
GOOS?=linux
BUILD_DIR=./build/bin

M = $(shell printf "\033[34;1m◉\033[0m")

default: all ;                                              		@ ## defaulting to clean and build

.PHONY: all
all: clean build

.PHONY: clean
clean: ; $(info $(M) running clean ...)                             @ ## clean up the old build dir
	@rm -vrf build

.PHONY: test
test: unit-test;													@ ## wrapper to run all testing

.PHONY: unit-test
unit-test: ; $(info $(M) running unit tests...)                     @ ## run the unit tests
	@go get -v -u
	@go test -cover ./...

.PHONY: build
build: build-dir; $(info $(M) building ...)                         @ ## build the binary
	@GOOS=$(GOOS) go get -v -u
	@GOOS=$(GOOS) go build \
		-ldflags "-X main.version=$(CLI_VERSION) -X main.compiled=$(COMPILED) -X main.githash=$(GIT_HASH)" \
		-o ./build/bin/$(BIN) ./main.go

.PHONEY: build-dir
build-dir: ;
	@[ ! -d "${BUILD_DIR}" ] && mkdir -vp "${BUILD_DIR}" || true

.PHONEY: install
install: ; $(info $(M) installing locally ...) 						@ ## install binary locally
	@go get -v -u
	@go build \
		-ldflags "-X main.version=$(CLI_VERSION) -X main.compiled=$(COMPILED) -X main.githash=$(GIT_HASH)" \
		-o $(GOPATH)/bin/$(BIN) ./main.go
.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

