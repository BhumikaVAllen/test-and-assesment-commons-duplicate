GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
#GOPRIVATE="github.com/Allen-Career-Institute/*"
GOVERSION := $(shell go version | cut -d " " -f 3 | cut -c 3-)
GOROOT:=$(shell go env GOROOT)
PROJECT_DIR = $(shell pwd)
PROJECT_BIN = $(PROJECT_DIR)/bin
GOLANGCI_LINT = $(PROJECT_BIN)/golangci-lint

INTERFACE_SOURCES := $(shell find pkg -name '*.go' -not -name '*_gen.go' -not -name '*_test.go')

# Check if mockery is installed
ifeq (, $(shell which mockery))
$(warning "mockery not found in $$PATH, installing now")
T:=$(shell cd .. && go install github.com/vektra/mockery/v2@v2.17.0)
endif

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find pkg -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find pkg -name *.proto)
endif

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2


lint:
	golangci-lint run

test: mockery
	go test -v ./... -covermode=count -coverprofile=coverage.out

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./pkg \
 	       --go_out=paths=source_relative:./pkg \
	       $(INTERNAL_PROTO_FILES)


.PHONY: generate
# generate
generate:
	go mod tidy
	go generate ./...

.PHONY: mockery
# mockery
mockery: $(INTERFACE_SOURCES)
	@echo "Generating mocks with mockery"
	mockery --all --with-expecter --quiet

.PHONY: all
# generate all
all:
	make config;
	make generate;
	#make lint;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
