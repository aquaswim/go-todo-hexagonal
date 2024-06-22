OAPI_CODEGEN_VERSION=latest
OAPI_GENERATE_CFG=./api/generate.cfg.yaml
OAPI_SPEC=./api/contract.yaml
OAPI_CODEGEN_BIN=$(shell go env GOPATH)/bin/oapi-codegen

BUILD_DIR=./tmp

.PHONY: __install_deps clean generate

__install_deps:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@${OAPI_CODEGEN_VERSION}

generate: __install_deps
	$(OAPI_CODEGEN_BIN) -config ${OAPI_GENERATE_CFG} ${OAPI_SPEC}

clean:
	@echo "cleaning generated file"