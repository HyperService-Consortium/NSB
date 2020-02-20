# NSB

Tendermint implementation of the NetworkStatusBlockchain.

#### Start Client: 

Under /root/work/go/src/github.com/HyperService-Consortium/NSB 

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
tendermint node --rpc.laddr tcp://0.0.0.0:26657 --home ./nsb --proxy_app tcp://0.0.0.0:27667
```

# NSB-cli
#### build the execuable
Under path/to/NSB/bin/nsb-cli
```
go build
mv nsb-cli.exe ../
```

#### Create New Wallet, which can contain a group of Accounts. 
Under path/to/NSB/bin
```
nsbcli.exe wallet create --db ./kvstore --wn Alice
```


#### Create a new Account to the wallet. 
Under path/to/NSB/bin
```
nsbcli.exe account create --db ./kvstore --wn Alice
```

#### Show Wallet
Under path/to/NSB/bin
```
nsbcli.exe wallet show --db ./kvstore --wn Alice
```

# py-nsbcli
#### Start
Under path/to/NSB/bin
```
py -3
exec(open("./main.py").read())
```

#### Load Wallet to python
```python
alice = kvdb.load_wallet("Alice")
```
#### Show Wallet Address
```python
alice.address(0).hex()
'5699c73fb5b13dcb860c147dbfe57dd34d5758807f9abe355b38499ba4c93a85'
```

#### create signature
```python
alice.sign(b"signature").hex()
'fcb106038f05d03e688ce852323ebc73adf998864206b10f5d3d2beabe4005c3d49aff40620d8f7e08a1cb896d5c77c9f4f0175853b01dbf4355ebc1799aeb0c'
```

#### Set RPC host
```python
admin.set_rpc_host("http://127.0.0.1:27667")
"http://127.0.0.1:27667"
```

#### test
```python
cli.abci_info()
{
    "id": "",
    "jsonrpc": "2.0",
    "result": {
        "response": {
            "app_version": "1",
            "data": {
                "height": 129047,
                "state_root": "283f38f544854a188d297987931316be82d971db8b30cd6fe746122cef4391c7"
            },
            "version": "0.16.0"
        }
    }
}
```
