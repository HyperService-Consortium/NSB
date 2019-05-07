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

# NSB-cli

#### build
Under path/to/NSB/bin/nsb-cli
```
go build
mv nsb-cli ../
```

#### Create New Wallet
Under path/to/NSB/bin
```
nsbcli.exe wallet create --db ./kvstore --wn Alice
```


#### Create New Account
Under path/to/NSB/bin
```
nsbcli.exe account create --db ./kvstore --wn Alice
```

#### Show Wallet
Under path/to/NSB/bin
```
nsbcli.exe wallet show --db ./kvstore --wn Alice
```
