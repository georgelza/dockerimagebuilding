
.ONESHELL:

REPO_NAME=georgelza
BINARY_NAME1=blogapp1
BINARY_NAME2=blogapp2
BINARY_NAME3=blogapp3
BINARY_NAME4=blogapp4
BINARY_NAME5=blogapp5
VERSION=1.0.0

.DEFAULT_GOAL := help

define HELP

Available commands:

- pull: Pull Docker Image.
- build: Build All Docker Image.
- size: Run docker images to see size comparisons.
- push: Push Docker Image.
- run: Run App.
- stop: Stop App.
- down, ps & logs: well what is expected.

endef

export HELP
help:
	@echo "$$HELP"
.PHONY: help

pull:
	docker pull ubuntu:jammy
	docker pull ubuntu:20.04
	docker pull golang:1.22

build: pull build_1 build_2 build_3 build_4 build_5 size

build_1: 
	sudo docker build --platform linux/arm64 --no-cache -t $(BINARY_NAME1):$(VERSION) --file Dockerfile.1 .
	sudo docker tag ${BINARY_NAME1}:${VERSION} $(REPO_NAME)/${BINARY_NAME1}:${VERSION}
	sudo docker rmi ${BINARY_NAME1}:${VERSION}

build_2: 
	sudo docker build --platform linux/arm64 --no-cache -t $(BINARY_NAME2):$(VERSION) --file Dockerfile.2 .
	sudo docker tag ${BINARY_NAME2}:${VERSION} $(REPO_NAME)/${BINARY_NAME2}:${VERSION}
	sudo docker rmi ${BINARY_NAME2}:${VERSION}

build_3: 
	sudo docker build --platform linux/arm64 --no-cache -t $(BINARY_NAME3):$(VERSION) --file Dockerfile.3 .
	sudo docker tag ${BINARY_NAME3}:${VERSION} $(REPO_NAME)/${BINARY_NAME3}:${VERSION}
	sudo docker rmi ${BINARY_NAME3}:${VERSION}

build_4: 
	sudo docker build --platform linux/arm64 --no-cache -t $(BINARY_NAME4):$(VERSION) --file Dockerfile.4 .
	sudo docker tag ${BINARY_NAME4}:${VERSION} $(REPO_NAME)/${BINARY_NAME4}:${VERSION}
	sudo docker rmi ${BINARY_NAME4}:${VERSION}

build_5: 
	sudo docker build --platform linux/arm64 --no-cache -t $(BINARY_NAME5):$(VERSION) --file Dockerfile.5 .
	sudo docker tag ${BINARY_NAME5}:${VERSION} $(REPO_NAME)/${BINARY_NAME5}:${VERSION}
	sudo docker rmi ${BINARY_NAME5}:${VERSION}

size: 
	sudo docker images

push:
	docker push $(REPO_NAME)/${BINARY_NAME1}:${VERSION}
	docker push $(REPO_NAME)/${BINARY_NAME2}:${VERSION}
	docker push $(REPO_NAME)/${BINARY_NAME3}:${VERSION}
	docker push $(REPO_NAME)/${BINARY_NAME4}:${VERSION}
	docker push $(REPO_NAME)/${BINARY_NAME5}:${VERSION}

run: 
	docker compose up -d app
	docker compose logs

down: 
	docker compose down

ps: 
	docker compose ps

logs: 
	docker compose logs
