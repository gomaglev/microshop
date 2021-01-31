SHELL := /bin/bash

########################
# Docker configuration #
########################

# Default config args
ENV?=.env

# Default Docker args
DOCKER_HUB?=protomicro
DOCKER_TAG?=make
DOCKER_APP_NAME=microshop

# Config imports
-include $(ENV)

ifndef DOCKER_APP_NAME
  $(error DOCKER_APP_NAME is not set)
endif

DOCKER_IMAGE=$(DOCKER_HUB)/$(DOCKER_APP_NAME)

.PHONY: help

help: ## Show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/##//' 

.DEFAULT_GOAL := up

APP = microshop
SERVER_BIN = ./cmd/${APP}/${APP}

# Docker run
build: ## Build the container
	docker build -t $(DOCKER_IMAGE) .

build-nc: ## Bulid the container without cache
	docker build --no-cache -t $(DOCKER_IMAGE) .

start: 
	go run ./cmd/microshop/main.go start

wire: ## Run the container
	wire ./internal/app/injector ./internal/app/test/mock/model ./internal/app/test/mock/service

run: ## Run the container
	docker run \
	-d \
	--env-file $(ENV) \
	-p $(PORT):$(PORT) \
	--name="$(DOCKER_APP_NAME)" \
	$(DOCKER_IMAGE)

stop: ## Stop the container and remove
	docker stop $(DOCKER_APP_NAME) || true && docker rm $(DOCKER_APP_NAME) || true

up: build stop run ## Build the container and run it

# Docker tagging
tag-release: tag-latest tag-version ## Generate release tags

tag-latest: check-build-existence ## Generate the latest tag
	@echo 'Tagging `latest`'
	docker tag $(DOCKER_IMAGE) $(DOCKER_IMAGE):latest

tag-version: check-build-existence ## Generate a `{DOCKER_TAG}` tag
	@echo 'Tagging `${DOCKER_TAG}`'
	docker tag $(DOCKER_IMAGE) $(DOCKER_IMAGE):$(DOCKER_TAG)

# Docker publishing
publish-release: build publish-latest publish-version ## Publish the release tags

publish-latest: tag-latest ## Publish the latest tag
	@echo 'Publishing `latest` to `$(DOCKER_IMAGE)`'
	docker push $(DOCKER_IMAGE):latest

publish-version: tag-version ## Publish the `{DOCKER_TAG}` tag
	@echo 'Publishing `$(DOCKER_TAG)` to `$(DOCKER_IMAGE)`'
	docker push $(DOCKER_IMAGE):$(DOCKER_TAG)

check-build-existence: ## Check docker existence
	if [ "$(shell docker images -q $(DOCKER_IMAGE) 2> /dev/null)" == "" ]; then\
		make build;\
	fi

########
# gRPC #
########

API_PROTO_FILES ?= $(shell find $(PWD)/api -type f -path '*.proto')
PKG_PROTO_FILES ?= $(shell find $(PWD)/pkg -type f -path '*.proto')
OTR_PROTO_FILES ?= $(shell find $(PWD) -type f -path '*.proto' | grep -v "vendor\|pkg\|api")
PROTO_PB_FILES ?= $(shell find $(PWD) -type f -path '*.pb.go' | grep -v "vendor")

PROTOC_INJECTOR := ${GOPATH}/bin/protoc-go-inject-tag

PROTOC := ${GOPATH}/bin/protoc
PROTO_PATH = $(PWD)

GENSERVICEPROTO=sh -c '$(PROTOC) \
	-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
	-I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
	-I ${PROTO_PATH} \
	--proto_path=$$0 \
	--grpc-gateway_out=logtostderr=true:$$1 \
	--go_out=$$1 \
	--go-grpc_out=require_unimplemented_servers=false:$$1 \
	--validate_out=lang=go:$$1 \
	--openapiv2_out $$0 --openapiv2_opt logtostderr=true $$2'


GENPROTO=sh -c '$(PROTOC) \
	-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
	-I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
	-I ${PROTO_PATH} \
	--proto_path=$$0 \
	--grpc-gateway_out=paths=source_relative,logtostderr=true:$$1 \
	--go_out=paths=source_relative:$$1 \
	--go-grpc_out=paths=source_relative,require_unimplemented_servers=false:$$1 \
	--validate_out=paths=source_relative,lang=go:$$1 \
	--openapiv2_out $$1 --openapiv2_opt logtostderr=true $$2'

protos: # generate protobuf files
	for file in $(shell find ${PKG_PROTO_FILES} -type f -path '*.proto'); \
	do $(GENPROTO) ${PROTO_PATH} ${PROTO_PATH} "$${file}"; done

	for file in $(shell find $(API_PROTO_FILES) -type f -path '*.proto'); \
	do $(GENSERVICEPROTO) ${PROTO_PATH} $(PROTO_PATH)/internal/app "$${file}"; done

	#${GENPROTO} ${PROTO_PATH} ${PROTO_PATH} ${OTR_PROTO_FILES}
	#$(GENPROTO) ${PROTO_PATH} $(PROTO_PATH)/api $(API_PROTO_FILES)
	#${GENPROTO} ${PROTO_PATH} ${PROTO_PATH} ${PKG_PROTO_FILES}
	#$(PROTOC_INJECTOR) -input="${PROTO_PB_FILES}" -verbose=false

faker: # add faker tag to proto files
	$(PROTOC_INJECTOR) -input="${PROTO_PB_FILES}" -verbose=false


