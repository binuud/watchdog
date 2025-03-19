# Maintainer Binu Udayakumar

# import config.
# You can change the default config with `make cnf="config_special.env" build`
# use build_staging.env for staging server (local too)

# Optional environmental variables
cnf ?= local.env
REPO=binuud
APP_NAME ?= watchdog


BUILD_VER ?= v0.0.1

DOCKER_HUB_TAG ?= v1.0.1

TAGGED_NAME = $(REPO)/$(APP_NAME)

include $(cnf)
export $(shell sed 's/=.*//' $(cnf))

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

run: ## Run code once, for auto run on code change, use make run-watch
	go run services/watchdog/cmd/main.go


vendor: ## go vendor and tidy
	go mod tidy
	go mod vendor	


test: ## Run tests
	go test -v ./pkg/...

testCoverage: ## Run test coverage
	go test ./pkg/... -coverprofile=coverage.out