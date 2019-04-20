package main

import (
	nsbnet "github.com/Myriad-Dreamin/NetworkStatusBlockChain/nsb_abci/net"
)

func main() {
	nsb := nsbnet.NewNSB()
	nsb.Start()
	nsb.Loop()
}