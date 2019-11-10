# Makefile for the "localnode" docker image.

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
makefile_name := $(notdir $(mkfile_path))
CURPATH := $(patsubst %/${makefile_name}, %, $(mkfile_path))
BUILD=${CURPATH}/build/
DOCKER_BUILD=${CURPATH}/docker/build/
SILENT:= 
USER=$(whoami)

all: build

.ONESHELL:
build: image
	# chgrp root .
	# chown root .
	docker run -u root --rm -v ${CURPATH}:/build -it build/nsb
	mkdir -p ${BUILD}
	cp ./NSB ${BUILD}
	mv ./NSB ${DOCKER_BUILD}

image:
	docker build -f ./Dockerfile-Build . -t build/nsb

.PHONY: all image

#.PHONY: network

.PHONY: build

