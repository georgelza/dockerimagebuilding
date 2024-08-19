
.DEFAULT_GOAL := help

define HELP

Available commands:

- pullall: Pull all source images.
- buildall: Build all images.

- os: Docker Pull all source OS images.

- buildos: Build all OS images.
- buildopenjdk: Build all OpenJDK images.
- buildhadoop: Build all Hadoop images.

endef

export HELP
help:
	@echo "$$HELP"
.PHONY: help

pullall: os

buildall: buildos buildopenjdk buildhadoop

# Base images used along the way, used to build 2nd level images.
# remove the arm64v8 to make it work on amd64 compatible.
os: 
	docker pull arm64v8/ubuntu:20.04

buildos:
	cd ./build-ubuntu-os-20.04; make build

buildopenjdk:
	cd build-ubuntu-os-openjdk11; pwd; make build

buildhadoop:
	cd build-hadoop-openjdk11-hdfs; make buildbase; make build



