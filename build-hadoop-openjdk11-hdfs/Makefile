.DEFAULT_GOAL := help

REPO_NAME=georgelza
IMAGE_NAME=hadoop-base

HADOOP_VERSION=3.3.5-java11
VERSION=1.0.1

define HELP

Based on Ubuntu 20.04 / OpenJDK 11 / Apache Hadoop 3.3.8

Available commands:

- buildbase: Build the base image project

- build: Build this project

- run: Run the project

- stop: Stop the project

- down: Stop and remove the project

- logs: Show the logs

- ps: Show the status

endef

export HELP
help:
	@echo "$$HELP"
.PHONY: help

buildbase:
	cd base; sudo docker build -t $(REPO_NAME)/$(IMAGE_NAME)-$(HADOOP_VERSION):$(VERSION) .

buildwgbase:
	cd base; sudo docker build -t $(REPO_NAME)/$(IMAGE_NAME)-$(HADOOP_VERSION):$(VERSION) --file Dockerfile.wget .

build:  
	sudo docker build -t hadoop-namenode-$(HADOOP_VERSION):$(VERSION) ./namenode
	sudo docker build -t hadoop-datanode-$(HADOOP_VERSION):$(VERSION) ./datanode
	sudo docker build -t hadoop-resourcemanager-$(HADOOP_VERSION):$(VERSION) ./resourcemanager
	sudo docker build -t hadoop-nodemanager-$(HADOOP_VERSION):$(VERSION) ./nodemanager
	sudo docker build -t hadoop-historyserver-$(HADOOP_VERSION):$(VERSION) ./historyserver

run:
	docker compose -p devlab up -d

stop:
	docker compose stop

down:
	docker compose down

logs:
	docker compose logs

ps:
	docker compose ps