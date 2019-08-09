# Makefile for the "localnode" docker image.

CURPATH=$(shell pwd)
BUILD=${CURPATH}/build
NODENAME=tendermint-nsb/node
SILENT:= 

all:
	mkdir ${BUILD}
	cp ${CURPATH}/node/tendermint ${BUILD}	
	docker build --tag ${NODENAME} node

network:
	docker network create --driver bridge --subnet 192.167.232.0/22  nsb_net

.ONESHELL:
build:
	if ! [ -f ${BUILD}/node0/config/genesis.json ]
	then
		docker run --rm -v ${BUILD}:/tendermint:Z ${NODENAME} testnet --v 4 --o . --populate-persistent-peers --starting-ip-address 192.167.233.2
	fi
	docker-compose -f ${CURPATH}/testnode.yml up

down:
	docker-compose -f ${CURPATH}/testnode.yml down

clean:
	rm -rf -r ${BUILD}/node*
	rm -rf -r ${BUILD}/data*

start:
	docker-compose -f ${CURPATH}/testnode.yml start ${SILENT}

restart:
	docker-compose -f ${CURPATH}/testnode.yml restart ${SILENT}

stop:
	docker-compose -f ${CURPATH}/testnode.yml stop ${SILENT}

.PHONY: all

.PHONY: network

.PHONY: build

.PHONY: down

.PHONY: clean

.PHONY: start

.PHONY: restart

.PHONY: stop

