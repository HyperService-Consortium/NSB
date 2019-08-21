package main

import (
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/NSB/account"
	"github.com/syndtr/goleveldb/leveldb"
	urcli "github.com/urfave/cli"
)

type WltShowCmd struct {
	parentCmd *WltCmd
	datadir   string
	name      string
}

func (wltcmd *WltShowCmd) Action(c *urcli.Context) error {
	if wltcmd.name == "" {
		return LogicError(errors.New("must enter the name of the wallet"))
	}
	if wltcmd.datadir == "" {
		return LogicError(errors.New("must enter the path of database"))
	}

	db, err := leveldb.OpenFile(wltcmd.datadir, nil) //, leveldb.Options{ErrorIfMissing:true})
	if err != nil {
		return IOError(err)
	}
	defer db.Close()
	var wlt *account.Wallet
	wlt, err = account.ReadWallet(db, wltcmd.name)
	if err != nil {
		return IOError(err)
	}

	fmt.Println(wlt)

	return nil
}

func NewWltShowCmd(wltcmd *WltCmd) urcli.Command {
	var wltShow = &WltShowCmd{parentCmd: wltcmd}
	return urcli.Command{
		Name:      "show",
		ShortName: "s",
		Usage:     "wallet api",
		UsageText: "look through the existing wallet in a database",
		Category:  "wallet",
		Flags: []urcli.Flag{
			urcli.StringFlag{
				Name:        "database, db",
				Value:       "",
				Usage:       "Specify the database path to store the wallet",
				Destination: &wltShow.datadir,
			},
			urcli.StringFlag{
				Name:        "walletname, wn",
				Value:       "",
				Usage:       "Enter the name of the wallet",
				Destination: &wltShow.name,
			},
		},
		Action:      wltShow.Action,
		After:       nil,
		Subcommands: nil,
	}
}
