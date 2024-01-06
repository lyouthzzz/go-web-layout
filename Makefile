GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)
BUF_INSTALLED := $(shell command -v buf 2> /dev/null)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: init
# init env
init:
	go install github.com/go-kratos/kratos/cmd/kratos/v2@a7bae93
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.13.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.13.0
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@a7bae93
	go install github.com/lyouthzzz/protoc-gen-go-errors@v0.0.1
	go install github.com/envoyproxy/protoc-gen-validate@v0.9.0
	go install github.com/google/wire/cmd/wire@v0.5.0

.PHONY: api
# generate api proto
api:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
 	       --go-http_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative:. \
 	       --go-errors_out=paths=source_relative:. \
		   --openapiv2_out=. \
	       $(API_PROTO_FILES)

.PHONY: wire
# wire
wire:
	wire ./...

.PHONY: build
# build
build:
	mkdir -p bin/ && GOPROXY="https://goproxy.cn,direct" go build -ldflags '-w -s -extldflags "-static"' -o ./bin/ ./...

buildimage:
	docker build --target api -f Dockerfile -t go-web-layout .