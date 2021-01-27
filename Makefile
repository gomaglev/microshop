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
	wire ./internal/app/injector ./internal/app/injector/mock

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

PROTO_FILES ?= $(shell find $(PWD) -type f -path '*.proto' | grep -v "vendor")
#PROTO_FILES ?= $(PWD)/api/application.proto
PROTO_PB_FILES ?= $(shell find $(PWD) -type f -path '*.pb.go' | grep -v "vendor")

PROTOC := ${GOPATH}/bin/protoc
PROTOC_INJECT_TAG := ${GOPATH}/bin/protoc-go-inject-tag
PROTOS_DEST = $(PWD)
PROTOC_FLAGS ?= \
    -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
    -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
	-I ${GOPATH}/src/github.com/gomaglev/protos \
	--proto_path=${PROTOS_DEST} \
	--grpc-gateway_out=logtostderr=true:$(PROTOS_DEST)/internal/app \
	--go_out=$(PROTOS_DEST)/internal/app \
	--go-grpc_out=require_unimplemented_servers=false:$(PROTOS_DEST)/internal/app \
	--validate_out=lang=go:$(PROTOS_DEST)/internal/app \
	--openapiv2_out ${PROTOS_DEST} --openapiv2_opt logtostderr=true \

protos: # generate protobuf files
	$(PROTOC) $(PROTOC_FLAGS) ${PROTO_FILES}
	$(PROTOC_INJECT_TAG) -input="${PROTO_PB_FILES}" -verbose=false

faker: # add faker tag to proto files
	$(PROTOC_INJECT_TAG) -input="${PROTO_PB_FILES}" -verbose=false