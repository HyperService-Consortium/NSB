package nsbnet

import (
	"fmt"
	cmn "github.com/tendermint/tendermint/libs/common"
	abcicli "github.com/tendermint/tendermint/abci/client"
	abcisrv "github.com/tendermint/tendermint/abci/server"
	abcinsb "github.com/Myriad-Dreamin/NSB/nsb_abci/nsb"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
)

const (
	nsb_port = ":27667"
	nsb_tcp = "tcp://0.0.0.0:27667"
	nsb_net_type = "socket"
	nsb_db_dir = "./data/"
	nsb_must_connected = false
)

type NSB struct {
	app types.Application
	srv cmn.Service
	cli abcicli.Client
	logger log.Logger
}

func NewNSBClient() (cli abcicli.Client, err error) {
	cli, err = abcicli.NewClient(nsb_port, nsb_net_type, nsb_must_connected)
	return
}

func NewNSBServer(app types.Application) (srv cmn.Service, err error) {
	srv, err = abcisrv.NewServer(nsb_tcp, nsb_net_type, app)
	return 
}

func NewNSB() (nsb NSB, err error) {
	nsb.logger = log.NewNopLogger()
	nsb.app, err =  abcinsb.NewNSBApplication(nsb_db_dir)
	if err != nil {
		return 
	}
	nsb.srv, err = NewNSBServer(nsb.app)
	if err != nil {
		return 
	}
	nsb.srv.SetLogger(logger.With("module", "nsbabci-server"))
	
	nsb.cli, err = NewNSBClient()
	if err != nil {
		server.Stop()
	}
	nsb.cli.SetLogger(logger.With("module", "nsbabci-client"))
	return
}
func (nsb *NSB) Start() (err error) {
	if err = nsb.srv.Start(); err != nil {
		fmt.Println(err)
		fmt.Println("start error")
	}
	return
}

func (nsb *NSB) LoopUntilStop() {
	cmn.TrapSignal(
		nsb.logger, func() {
		// Cleanup
		nsb.srv.Stop()
	})
}