.DEFAULT_GOAL := help
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
	sudo docker build -t ubuntu20.04:$(VERSION) .

