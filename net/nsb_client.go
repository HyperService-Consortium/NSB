package nsbnet

import (
	"fmt"
	"flag"
	abcinsb "github.com/HyperServiceOne/NSB/application"
	abcicli "github.com/tendermint/tendermint/abci/client"
	abcisrv "github.com/tendermint/tendermint/abci/server"
	"github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/libs/log"
)

const (
	// nsb_port           = ":27667"
	// nsb_tcp            = "tcp://0.0.0.0:27667"
	nsb_net_type       = "socket"
	nsb_must_connected = false
)

var (
	nsb_port = flag.String("port", ":27667", "port")
	nsb_db_dir = flag.String("db", "./data/", "db")
	nsb_tcp = flag.String("server", "tcp://0.0.0.0:27667", "server address")
)


type NSB struct {
	app    types.Application
	srv    cmn.Service
	cli    abcicli.Client
	logger log.Logger
}

func NewNSBClient() (cli abcicli.Client, err error) {
	cli, err = abcicli.NewClient(*nsb_port, nsb_net_type, nsb_must_connected)
	return
}

func NewNSBServer(app types.Application) (srv cmn.Service, err error) {
	srv, err = abcisrv.NewServer(*nsb_tcp, nsb_net_type, app)
	return
}

func NewNSB() (nsb NSB, err error) {

	nsb.logger = log.NewNopLogger()

	fmt.Println("create app...")
	nsb.app, err = abcinsb.NewNSBApplication(*nsb_db_dir)
	if err != nil {
		return
	}

	fmt.Println("create server... on", *nsb_tcp)
	nsb.srv, err = NewNSBServer(nsb.app)
	if err != nil {
		return
	}
	nsb.srv.SetLogger(log.NewNopLogger())

	fmt.Println("create client... on", *nsb_port)
	nsb.cli, err = NewNSBClient()
	if err != nil {
		return
	}
	nsb.cli.SetLogger(log.NewNopLogger())

	return
}
func (nsb *NSB) Start() (err error) {

	fmt.Println("start server...")
	if err = nsb.srv.Start(); err != nil {
		return
	}

	fmt.Println("start client...")
	if err = nsb.cli.Start(); err != nil {
		nsb.srv.Stop()
		return
	}

	fmt.Printf("the application is listening %v\n", nsb_tcp)
	return
}

func (nsb *NSB) LoopUntilStop() {
	go func() {

	ForeverLoop:
		fmt.Println("looping")
		cmn.TrapSignal(
			nsb.logger, func() {
				// Cleanup
				nsb.app.(*abcinsb.NSBApplication).Stop()
				nsb.srv.Stop()
				nsb.cli.Stop()
				fmt.Println("stopped")
			})
		select {}
		goto ForeverLoop
	}()

	select {}
}


func init() {
	flag.Parse()
}

