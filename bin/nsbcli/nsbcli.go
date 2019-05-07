package main

import (
	"fmt"
	urcli "github.com/urfave/cli"
	"io"
	"log"
	"os"
)

const (
	CliName   = "go-cli-client"
	Usage     = "interactive with cli"
	UsageText = "go-tendermint implementation of Network Status Blockchain"
	Version   = "0.2.0"
)

type NSBCli struct {
	handler *urcli.App

	// submodules
	acc *AccCmd
	wlt *WltCmd

	logfiledir string
	logfile    *os.File
	logger     *log.Logger

	port int
}

func (cli *NSBCli) SetLog(rd io.Writer) {
	cli.logger = log.New(rd, "", log.LstdFlags|log.Lshortfile)
	cli.logger.SetFlags(log.LstdFlags | log.Lshortfile)
}

func (cli *NSBCli) Before(c *urcli.Context) (err error) {
	cli.logfile, err = os.OpenFile(cli.logfiledir, os.O_APPEND|os.O_CREATE, 666)
	if err != nil {
		cli.logfile = nil
		return err
	}
	cli.SetLog(cli.logfile)
	fmt.Println("app Before")
	return nil
}

func (cli *NSBCli) After(c *urcli.Context) error {
	fmt.Println("app After")
	return nil
}
func (cli *NSBCli) SetInfo() {
	cli.handler.Name = CliName
	cli.handler.Usage = Usage
	cli.handler.UsageText = UsageText
	cli.handler.Version = Version
}

func (cli *NSBCli) Init() {
	cli.SetInfo()

	cli.handler.Before = cli.Before
	cli.handler.Action = nil
	cli.handler.After = cli.After

	cli.handler.Flags = []urcli.Flag{
		urcli.IntFlag{
			Name:        "port, p",
			Value:       23766,
			Usage:       "listening port",
			Destination: &cli.port,
		},
		urcli.StringFlag{
			Name:        "logdir, ld",
			Value:       "nsbcli.log",
			Usage:       "logger address",
			Destination: &cli.logfiledir,
		},
	}
	urcli.HelpFlag = urcli.BoolFlag{
		Name:  "help, h",
		Usage: "show manual",
	}

	cli.acc = NewAccCmd(cli)
	cli.wlt = NewWltCmd(cli)
	cli.handler.Commands = []urcli.Command{
		*cli.acc.cmd,
		*cli.wlt.cmd,
	}

}

func (cli *NSBCli) CommandNotFound(c *urcli.Context, cmdString string) {
	fmt.Println("command not found,", cmdString)
}

func (cli *NSBCli) Stop() {
	if cli.logfile != nil {
		cli.logfile.Close()
	}
}

func NewNSBCli() *NSBCli {
	return &NSBCli{
		handler: urcli.NewApp(),
	}
}

func (cli *NSBCli) Run() {
	if err := cli.handler.Run(os.Args); err != nil {
		if cli.logger == nil {
			cli.SetLog(os.Stdout)
		}
		cli.logger.Fatal(err)
	}
}

func (cli *NSBCli) CliExit(status int) {
	fmt.Println("nsbcli exit with", status)
	cli.Stop()
	os.Exit(status)
}

func main() {
	var cli = NewNSBCli()
	urcli.OsExiter = cli.CliExit
	cli.Init()
	cli.Run()
	cli.Stop()
}
