# Makefile for the "localnode" docker image.

CURPATH=$(shell pwd)
BUILD=${CURPATH}/build
NODENAME=tendermint-nsb/node
SILENT:= 
COUNT=4
FILE=testnode4.yml

all:
	mkdir ${BUILD}
	cp ${CURPATH}/node/tendermint ${BUILD}	
	docker build --tag ${NODENAME} node

# network:
#	docker-compose -f network.yml up # docker network create --driver bridge --subnet 192.167.232.0/22  nsb_net

.ONESHELL:
build:
	if ! [ -f ${BUILD}/node0/config/genesis.json ]
	then
		docker run --rm -v ${BUILD}:/tendermint:Z ${NODENAME} testnet --v ${COUNT} --o . --populate-persistent-peers --starting-ip-address 192.167.233.2
	fi
	docker-compose -f ${CURPATH}/${FILE} up

down:
	docker-compose -f ${CURPATH}/${FILE} down

clean:
	rm -rf -r ${BUILD}/node*
	rm -rf -r ${BUILD}/data*

start:
	docker-compose -f ${CURPATH}/${FILE} start ${SILENT}

restart:
	docker-compose -f ${CURPATH}/${FILE} restart ${SILENT}

stop:
	docker-compose -f ${CURPATH}/${FILE} stop ${SILENT}

.PHONY: all

#.PHONY: network

.PHONY: build

.PHONY: down

.PHONY: clean

.PHONY: start

.PHONY: restart

.PHONY: stop

