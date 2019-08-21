package main

import (
	"errors"
	"github.com/HyperService-Consortium/NSB/account"
	"github.com/syndtr/goleveldb/leveldb"
	urcli "github.com/urfave/cli"
)

type WltCreateCmd struct {
	parentCmd *WltCmd
	datadir   string
	name      string
}

func (wltcmd *WltCreateCmd) Action(c *urcli.Context) error {
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
	var checkErr bool
	checkErr, err = account.WalletExist(db, wltcmd.name)
	if err != nil {
		return IOError(err)
	}
	if checkErr == true {
		return ConflictError(errors.New("the wallet has already existed in the database"))
	}
	wlt = account.NewWallet(db, wltcmd.name)

	err = wlt.Save()
	if err != nil {
		return IOError(err)
	}

	return nil
}

func NewWltCreateCmd(wltcmd *WltCmd) urcli.Command {
	var wltCreate = &WltCreateCmd{parentCmd: wltcmd}
	return urcli.Command{
		Name:      "create",
		ShortName: "ct",
		Usage:     "wallet api",
		UsageText: "create new wallet",
		Category:  "wallet",
		Flags: []urcli.Flag{
			urcli.StringFlag{
				Name:        "database, db",
				Value:       "",
				Usage:       "Specify the database path to store the wallet",
				Destination: &wltCreate.datadir,
			},
			urcli.StringFlag{
				Name:        "walletname, wn",
				Value:       "",
				Usage:       "Enter the name of the wallet",
				Destination: &wltCreate.name,
			},
		},
		Action:      wltCreate.Action,
		After:       nil,
		Subcommands: nil,
	}
}
