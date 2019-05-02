package main

import (
	"fmt"
	urcli "github.com/urfave/cli"
)

type AccCmd struct {
	cli *NSBCli
	cmd *urcli.Command
}

func (acc *AccCmd) Before(c *urcli.Context) (err error) {
	fmt.Println("acc Before")
	return nil
}

func (acc *AccCmd) After(c *urcli.Context) (err error) {
	fmt.Println("acc After")
	return nil
}

func (acc *AccCmd) MakeCommands() urcli.Commands {
	return []urcli.Command{
		NewAccCreateCmd(acc),
	}
}

func NewAccCmd(nsbcli *NSBCli) *AccCmd {
	var acc = &AccCmd{cli: nsbcli}
	acc.cmd = &urcli.Command{
		Name:        "account",
		ShortName:   "acc",
		Usage:       "account api",
		UsageText:   "create new account, or get accounts from db",
		ArgsUsage:   "dbdir: the path of leveldb where stores accounts' info",
		Category:    "account",
		Before:      acc.Before,
		Action:      nil,
		Subcommands: acc.MakeCommands(),
	}
	return acc
}
