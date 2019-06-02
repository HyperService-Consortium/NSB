package main

import (
	nsbnet "github.com/HyperServiceOne/NSB/net"
)

func main() {
	nsb, err := nsbnet.NewNSB()
	if err != nil {
		panic(err)
	}
	err = nsb.Start()
	if err != nil {
		panic(err)
	}
	nsb.LoopUntilStop()
}
