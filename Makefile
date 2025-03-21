# Maintainer Binu Udayakumar

# import config.
# You can change the default config with `make cnf="config_special.env" build`
# use build_staging.env for staging server (local too)

REPO=binuud
APP_NAME ?= watchdog


BUILD_VER ?= v0.0.1

DOCKER_HUB_TAG ?= v1.0.1

TAGGED_NAME = $(REPO)/$(APP_NAME)

## Needs protoc to be installed

ifndef PROTOC
PROTOC = protoc
endif

ifndef PROTOWEB
PROTOWEB = protoc-gen-grpc-web
endif

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

build: ## Build docker image
	docker build --no-cache  --build-arg BUILD_VER=$(BUILD_VER) -t $(REPO)/$(APP_NAME) -f Dockerfile .

build-binary: ## Build the watchDog project
	go build .

run: ## Run code once, for auto run on code change
	go run services/watchdog/cmd/main.go

run-server: ## Start GRPC and HTTP server
	go run services/watchdogServer/cmd/main.go

protos: ## Buid go and web protos, and swagger openApi json
	$(PROTOC) -I=./proto/.  \
	--go_out=./gen/go/ \
	--go_opt paths=source_relative \
	--go-grpc_out=./gen/go/ \
	--go-grpc_opt paths=source_relative \
	--grpc-gateway_out=./gen/go/ \
	--grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true \
	--grpc-gateway-ts_out=./gen/web/ \
	--grpc-gateway-ts_opt paths=source_relative \
	--grpc-gateway-ts_opt generate_unbound_methods=true \
	proto/v1/watchdog/watchdog.proto proto/v1/watchdog/watchdogService.proto  
##	--oas_out ./gen/web/v1/watchdog/ \
##yq eval ./gen/web/v1/watchdog/openapi.yaml -o=json -P > ./gen/web/v1/watchdog/openapi.json

vendor: ## go vendor and tidy
	go mod tidy && 	go mod vendor	

test: ## Run tests
	go test -v ./pkg/...

testCoverage: ## Run test coverage
	go test ./pkg/... -coverprofile=coverage.out