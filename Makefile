OAPI_CODEGEN_VERSION=latest
OAPI_GENERATE_CFG=./api/rest/generate.cfg.yaml
OAPI_SPEC=./api/rest/contract.yaml
OAPI_CODEGEN_BIN=$(shell go env GOPATH)/bin/oapi-codegen

BUILD_DIR=./tmp

.PHONY: __install_deps clean generate generate-rest generate-grpc configure-grpc

__install_deps:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@${OAPI_CODEGEN_VERSION}

generate-rest: __install_deps
	$(OAPI_CODEGEN_BIN) -config ${OAPI_GENERATE_CFG} ${OAPI_SPEC}

generate: generate-rest

configure-grpc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	@echo please install protobuf using your package manager, example: pacman -S protobuf

generate-grpc:
	@echo compile proto files
	@protoc \
		--proto_path=./api/grpc \
		--go_opt=paths=source_relative \
		--go_out=./internal/adapter/grpc/pb \
    	--go-grpc_opt=paths=source_relative \
    	--go-grpc_out=./internal/adapter/grpc/pb \
    	./api/grpc/*.proto

clean:
	@echo "cleaning generated file"