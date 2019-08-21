# Local Cluster with Docker Compose

## Requirements

- [Install docker](https://docs.docker.com/engine/installation/)
- [Install docker-compose](https://docs.docker.com/compose/install/)

## Build Node

Build the `nsb` docker image.

Note the binary will be mounted into the container so it can be updated without
rebuilding the image.

```
root@kamiyoru:~/NSB/docker# make
mkdir /root/work/go/src/github.com/HyperService-Consortium/NSB/docker/build
cp /root/work/go/src/github.com/HyperService-Consortium/NSB/docker/node/tendermint /root/work/go/src/github.com/HyperService-Consortium/NSB/docker/build
docker build --tag tendermint-nsb/node node
mkdir: cannot create directory ‘/root/work/go/src/github.com/HyperService-Consortium/NSB/docker/build’: File exists
Sending build context to Docker daemon  46.04MB
Step 1/11 : FROM alpine:3.7
 ---> 6d1ef012b567
Step 2/11 : MAINTAINER Myriad-Dreamin <camiyoru@gmail.com>
 ---> Using cache
 ---> 07c4e6b4bf26
Step 3/11 : RUN apk update &&     apk upgrade &&     apk --no-cache add curl jq file
 ---> Using cache
 ---> bf2d16997cae
Step 4/11 : VOLUME [ /tendermint ]
 ---> Using cache
 ---> 36be42d73e7a
Step 5/11 : WORKDIR /tendermint
 ---> Using cache
 ---> 1e7dbd05e28d
Step 6/11 : EXPOSE 26656 26657
 ---> Using cache
 ---> 41910f0952d5
Step 7/11 : ENTRYPOINT ["/usr/bin/wrapper.sh"]
 ---> Using cache
 ---> 35361bef9bf6
Step 8/11 : STOPSIGNAL SIGTERM
 ---> Running in c0cccfa36177
Removing intermediate container c0cccfa36177
 ---> 84fd7dd78b47
Step 9/11 : COPY tendermint /usr/bin
 ---> 774e6a7bc00c
Step 10/11 : COPY NSB /usr/bin/
 ---> 60dc74841202
Step 11/11 : COPY wrapper.sh /usr/bin/wrapper.sh
 ---> 2e58758eb35c
Successfully built 2e58758eb35c
Successfully tagged tendermint-nsb/node:latest
```

## Set up Network

```bash
root@kamiyoru:~/NSB/docker# make network
docker network create --driver bridge --subnet 192.167.232.0/22  nsb_net
72522d90cf444b57b798a3e77646dc615903660f447fbb3647806865b30fd723
```

## Build and Run Testing Cluster

build and run a 4-node cluster with exposed port ip:26657

```bash
root@kamiyoru:~/NSB/docker# make build
if ! [ -f /root/work/go/src/github.com/HyperService-Consortium/NSB/docker/build/node0/config/genesis.json ]
then
docker run --rm -v /root/work/go/src/github.com/HyperService-Consortium/NSB/docker/build:/tendermint:Z tendermint-nsb/node testnet --v 4 --o . --populate-persistent-peers --starting-ip-address 192.167.233.2
fi
docker-compose -f /root/work/go/src/github.com/HyperService-Consortium/NSB/docker/testnode.yml up
create app...
loading state...
StateRoot:
Height: 0

create server... on
create client... on
start server...
start client...
I[2019-08-09|13:41:16.928] Generated private validator                  module=main keyFile=node0/config/priv_validator_key.json stateFile=node0/data/priv_validator_state.json
I[2019-08-09|13:41:16.928] Generated node key                           module=main path=node0/config/node_key.json
I[2019-08-09|13:41:16.929] Generated genesis file                       module=main path=node0/config/genesis.json
I[2019-08-09|13:41:16.934] Generated private validator                  module=main keyFile=node1/config/priv_validator_key.json stateFile=node1/data/priv_validator_state.json
I[2019-08-09|13:41:16.934] Generated node key                           module=main path=node1/config/node_key.json
I[2019-08-09|13:41:16.935] Generated genesis file                       module=main path=node1/config/genesis.json
I[2019-08-09|13:41:16.940] Generated private validator                  module=main keyFile=node2/config/priv_validator_key.json stateFile=node2/data/priv_validator_state.json
I[2019-08-09|13:41:16.940] Generated node key                           module=main path=node2/config/node_key.json
I[2019-08-09|13:41:16.941] Generated genesis file                       module=main path=node2/config/genesis.json
I[2019-08-09|13:41:16.947] Generated private validator                  module=main keyFile=node3/config/priv_validator_key.json stateFile=node3/data/priv_validator_state.json
I[2019-08-09|13:41:16.947] Generated node key                           module=main path=node3/config/node_key.json
I[2019-08-09|13:41:16.947] Generated genesis file                       module=main path=node3/config/genesis.json
Successfully initialized 4 node directories
Creating node3 ...
Creating node1 ...
Creating node0 ...
Creating node2 ...
Creating node3
Creating node1
Creating node0
Creating node1 ... done
Attaching to node3, node0, node2, node1
...
```


## Start, Restart or Stop Testing Cluster

Start built cluster.

```bash
root@kamiyoru:~/NSB/docker# make start
docker-compose -f /root/work/go/src/github.com/HyperService-Consortium/NSB/docker/testnode.yml start
Starting node1 ... done
Starting node0 ... done
Starting node3 ... done
Starting node2 ... done
```

Restart built cluster.

```bash
root@kamiyoru:~/NSB/docker# make restart
docker-compose -f /root/work/go/src/github.com/HyperService-Consortium/NSB/docker/testnode.yml restart
Restarting node2 ... done
Restarting node0 ... done
Restarting node1 ... done
Restarting node3 ... done
```

Stop built cluster.

```bash
root@kamiyoru:~/NSB/docker# make stop
docker-compose -f /root/work/go/src/github.com/HyperService-Consortium/NSB/docker/testnode.yml stop
Stopping node0 ... done
Stopping node2 ... done
Stopping node1 ... done
Stopping node3 ... done
```

## Push Down and Remove Testing Cluster

```bash
root@kamiyoru:~/NSB/docker# make down
docker-compose -f /root/work/go/src/github.com/HyperService-Consortium/NSB/docker/testnode.yml down
Stopping node2 ... done
Stopping node0 ... done
Stopping node1 ... done
Stopping node3 ... done
Removing node2 ... done
Removing node0 ... done
Removing node1 ... done
Removing node3 ... done
Network nsb_net is external, skipping

root@kamiyoru:~/NSB/docker# make clean
rm -rf -r /root/work/go/src/github.com/HyperService-Consortium/NSB/docker/build/node*
rm -rf -r /root/work/go/src/github.com/HyperService-Consortium/NSB/docker/build/data*
```



