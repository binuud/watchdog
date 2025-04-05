# Maintainer Binu Udayakumar

# import config.
# You can change the default config with `make cnf="config_special.env" build`
# use build_staging.env for staging server (local too)

### WARNING ### NOT-A-PRIVATE-REPO ##########

REPO=dronasys
APP_NAME ?= watchdog

BUILD_VER ?= v0.0.1
DOCKER_HUB_TAG ?= v0.0.1

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

deploy: ## Deploy all images to docker hub
	@echo "====================> Pushing image to dockerhub (PUBLIC) ."
## @docker login --username=$(DOCKER_HUB_USER) --password="$(DOCKER_HUB_PASSWD)"
	docker tag $(REPO)/${APP_NAME} $(REPO)/${APP_NAME}:$(DOCKER_HUB_TAG)
	docker push $(REPO)/${APP_NAME}:$(DOCKER_HUB_TAG)

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
	--oas_out ./gen/web/v1/watchdog/ \
	proto/v1/watchdog/watchdog.proto proto/v1/watchdog/watchdogService.proto  
	yq eval ./gen/web/v1/watchdog/openapi.yaml -o=json -P > ./gen/web/v1/watchdog/openapi.json

run: ## Run code once, for auto run on code change
	go run cmd/watchdog/main.go --file $(PWD)/config.yaml

run-mydomains: ## Run code once, for list of mydomains
	go run cmd/watchdog/main.go --file $(PWD)/local/myDomains.yaml

run-server: ## Start GRPC and HTTP server
	go run cmd/watchdogServer/main.go -grpc_port 10090 -http_port 10080

run-docker: ## run docker image
	docker stop $(APP_NAME); docker rm $(APP_NAME); docker run --name $(APP_NAME) -p 10090:9090 -p 10080:9080 -v  "$(shell pwd)/config.yaml:/configs/config.yaml" $(REPO)/$(APP_NAME)

swagger-ui: ## launch swagger ui
	docker run  -p 10030:8080 -v ./gen/web/v1/watchdog/openapi.json:/tmp/swagger.json -e SWAGGER_FILE=/tmp/swagger.json docker.swagger.io/swaggerapi/swagger-editor

vendor: ## go vendor and tidy
	echo "tidy + vendor"
	go mod tidy
	go mod vendor	

test: ## Run tests
	go test -v ./pkg/...

testCoverage: ## Run test coverage
	go test ./pkg/... -coverprofile=coverage.out

release: ## for releasing the package
## comments for reference
# git tag -a v0.1.0 -m "Release COMMENT"
# git push origin [version]
# export GITHUB_TOKEN="YOUR_GH_TOKEN"
	goreleaser release	

release-check: ## run release build, but do not publish
	goreleaser release --snapshot	