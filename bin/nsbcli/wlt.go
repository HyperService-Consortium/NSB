package main

import (
	"fmt"
	urcli "github.com/urfave/cli"
)

type WltCmd struct {
	cli *NSBCli
	cmd *urcli.Command
}

func (wlt *WltCmd) Before(c *urcli.Context) (err error) {
	fmt.Println("wlt Before")
	return nil
}

func (wlt *WltCmd) After(c *urcli.Context) (err error) {
	fmt.Println("wlt After")
	return nil
}

func (wlt *WltCmd) MakeCommands() urcli.Commands {
	return []urcli.Command{
		NewWltCreateCmd(wlt),
		NewWltShowCmd(wlt),
	}
}

func NewWltCmd(nsbcli *NSBCli) *WltCmd {
	var wlt = &WltCmd{cli: nsbcli}
	wlt.cmd = &urcli.Command{
		Name:        "wallet",
		ShortName:   "wlt",
		Usage:       "wallet api",
		UsageText:   "create new wallet, or read wallet from db",
		ArgsUsage:   "dbdir: the path of leveldb where stores wallets' info",
		Category:    "wallet",
		Before:      wlt.Before,
		Action:      nil,
		Subcommands: wlt.MakeCommands(),
	}
	return wlt
}
