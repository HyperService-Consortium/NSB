package nsbnet

import (
	"flag"

	abcinsb "github.com/HyperService-Consortium/NSB/application"
	log "github.com/HyperService-Consortium/NSB/log"
	abcicli "github.com/tendermint/tendermint/abci/client"
	abcisrv "github.com/tendermint/tendermint/abci/server"
	"github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
)

const (
	// nsb_port           = ":27667"
	// nsb_tcp            = "tcp://0.0.0.0:27667"
	nsb_net_type       = "socket"
	nsb_must_connected = false
)

var (
	nsb_port   = flag.String("port", ":27667", "port")
	nsb_db_dir = flag.String("db", "./data/", "db")
	nsb_tcp    = flag.String("server", "tcp://0.0.0.0:27667", "server address")
)

type NSB struct {
	app    types.Application
	srv    cmn.Service
	cli    abcicli.Client
	logger log.TendermintLogger
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
	// &splitingIO{os.Stdout, &splitingIO{nil, nil}}
	// if you need to write to multifiles, use command tee

	logger, err := log.NewZapColorfulDevelopmentSugarLogger()
	if err != nil {
		return NSB{}, nil
	}

	nsb.logger = logger

	nsb.logger.Info("create app...")
	nsb.app, err = abcinsb.NewNSBApplication(logger, *nsb_db_dir)
	if err != nil {
		return
	}

	nsb.logger.Info("create server...", "address", *nsb_tcp)
	nsb.srv, err = NewNSBServer(nsb.app)
	if err != nil {
		return
	}
	nsb.srv.SetLogger(nsb.logger)

	nsb.logger.Info("create client...", "port", *nsb_port)
	nsb.cli, err = NewNSBClient()
	if err != nil {
		return
	}
	nsb.cli.SetLogger(nsb.logger)

	return
}
func (nsb *NSB) Start() (err error) {

	nsb.logger.Info("start server...")
	if err = nsb.srv.Start(); err != nil {
		return
	}

	nsb.logger.Info("start client...")
	if err = nsb.cli.Start(); err != nil {
		nsb.srv.Stop()
		return
	}

	nsb.logger.Info("the application is listening", "address", *nsb_tcp)
	return
}

func (nsb *NSB) LoopUntilStop() {
	go func() {
		nsb.logger.Info("looping")
		cmn.TrapSignal(
			nsb.logger, func() {
				// Cleanup
				nsb.app.(*abcinsb.NSBApplication).Stop()
				nsb.srv.Stop()
				nsb.cli.Stop()
				nsb.logger.Info("stopped")
			})
		select {}
	}()
	select {}
}

func init() {
	flag.Parse()
}
