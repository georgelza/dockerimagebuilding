.DEFAULT_GOAL := help

REPO_NAME=georgelza
IMAGE_NAME=ubuntu20.04
VERSION=1.0.1

define HELP

Available commands:

- build: Build this project

endef

export HELP
help:
	@echo "$$HELP"
.PHONY: help

build:  
	sudo docker build -t $(REPO_NAME)/$(IMAGE_NAME):$(VERSION) .

