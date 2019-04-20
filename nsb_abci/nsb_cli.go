package main

import (
	nsbnet "github.com/Myriad-Dreamin/NSB/nsb_abci/net"
)

func main() {
	nsb, err := nsbnet.NewNSB()
	if err != nil {
		panic(err)
	}
	nsb.Start()
	nsb.LoopUntilStop()
}