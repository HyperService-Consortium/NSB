# NSB

Tendermint implementation of the NetworkStatusBlockchain.

#### Start Client: 

Under /root/work/go/src/github.com/Myriad-Dreamin/NSB 

```
go run nsb_cli.go
```


#### Initialize Tendermint core:

Under /root

```
tendermint init --home ./nsb
```

#### Start Tendermint core:

Under /root

```
tendermint node --home ./nsb --proxy_app tcp://0.0.0.0:27667
```
