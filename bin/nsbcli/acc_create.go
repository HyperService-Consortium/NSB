package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/NSB/account"
	"github.com/syndtr/goleveldb/leveldb"
	urcli "github.com/urfave/cli"
	"os"
)

type AccCreateCmd struct {
	parentCmd *AccCmd
	seed      string
	outfile   string
	datadir   string
	wltname   string
	show      bool
}

func (acc *AccCreateCmd) Action(c *urcli.Context) error {

	bt, err := hex.DecodeString(acc.seed)
	if err != nil {
		return InternalError(err)
	}

	var a = account.NewAccount(bt)
	var apri, apub = hex.EncodeToString(a.PrivateKey), hex.EncodeToString(a.PublicKey)

	if acc.show {
		fmt.Println("Private Key:", apri, "\nPublic Key:", apub)
	}

	if acc.outfile != "" {
		fptr, err := os.OpenFile(acc.outfile, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 666)

		if err != nil {
			return IOError(err)
		}
		defer fptr.Close()

		_, err = fptr.Write([]byte(fmt.Sprintf("Private Key: %s\nPublic Key:: %s\n", apri, apub)))
		if err != nil {
			return IOError(err)
		}

	}

	if acc.datadir != "" {
		if acc.wltname == "" {
			return LogicError(errors.New("must enter the name of the wallet"))
		}
		db, err := leveldb.OpenFile(acc.datadir, nil) //, leveldb.Options{ErrorIfMissing:true})
		if err != nil {
			return IOError(err)
		}
		defer db.Close()

		wlt, err := account.ReadWallet(db, acc.wltname)
		if err != nil {
			return InternalError(err)
		}
		wlt.AppendAccount(a)
		err = wlt.Save()
		if err != nil {
			return IOError(err)
		}
	}

	return nil
}

func NewAccCreateCmd(acc *AccCmd) urcli.Command {
	var accCreate = &AccCreateCmd{parentCmd: acc}
	return urcli.Command{
		Name:      "create",
		ShortName: "ct",
		Usage:     "account api",
		UsageText: "create new account",
		Category:  "account",
		Flags: []urcli.Flag{
			urcli.StringFlag{
				Name:        "seed, sd",
				Value:       "",
				Usage:       "The ed25519 seed(hex to bytes)",
				Destination: &accCreate.seed,
			},
			urcli.BoolFlag{
				Name:        "show, s",
				Usage:       "Display the key to the screen",
				Destination: &accCreate.show,
			},
			urcli.StringFlag{
				Name:        "out, o",
				Value:       "",
				Usage:       "Specify the file to store the key",
				Destination: &accCreate.outfile,
			},
			urcli.StringFlag{
				Name:        "database, db",
				Value:       "",
				Usage:       "Specify the database path to store the key",
				Destination: &accCreate.datadir,
			},
			urcli.StringFlag{
				Name:        "walletname, wn",
				Value:       "",
				Usage:       "Enter the name of the wallet",
				Destination: &accCreate.wltname,
			},
		},
		Action:      accCreate.Action,
		After:       nil,
		Subcommands: nil,
	}
}
