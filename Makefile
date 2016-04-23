.PHONY : all
.PHONY : help
.DEFAULT_GOAL := all

STARTD := $(shell pwd)
THIS := $(abspath $(lastword $(MAKEFILE_LIST)))
THISD := $(dir $(THIS))


all: ## Make everything
	go test -v


help: ## This help.
	@echo Targets:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
