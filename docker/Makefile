# Makefile for the "localnode" docker image.


mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
makefile_name := $(notdir $(mkfile_path))
CURPATH := $(patsubst %/${makefile_name}, %, $(mkfile_path))
BUILD=${CURPATH}/build/
NODE_IMAGE_PATH=${CURPATH}/node/
NODENAME=tendermint-nsb/node
SILENT:= 
COUNT=4
FILE=testnode4.yml


all: start

# network:
#	docker-compose -f network.yml up # docker network create --driver bridge --subnet 192.167.232.0/22  nsb_net

image: nsb_source tendermint
	docker build --tag ${NODENAME} node

nsb_source:
	cp ${BUILD}/NSB ${NODE_IMAGE_PATH}

tendermint:
	mkdir -p ${BUILD}
	if ! [ -f ${BUILD}/tendermint ]
	then
		curl -o ${BUILD}/tendermint -L https://github.com/HyperService-Consortium/NSB/releases/download/v0.7.4/tendermint-linux-v0.32.2-8ba8497a	
	fi


.ONESHELL:
build: image
	mkdir -p ${BUILD}
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

start: build
	docker-compose -f ${CURPATH}/${FILE} start ${SILENT}

restart:
	docker-compose -f ${CURPATH}/${FILE} restart ${SILENT}

stop:
	docker-compose -f ${CURPATH}/${FILE} stop ${SILENT}

.PHONY: all image tendermint nsb_source build down clean start restart stop

#.PHONY: network
