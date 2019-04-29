package main

import (
	"fmt"
	urcli "github.com/urfave/cli"
	"encoding/hex"
	"github.com/Myriad-Dreamin/NSB/account"
)

type AccCreateCmd struct {
	parentCmd *AccCmd
	pkstring string
}

func (acc *AccCreateCmd) Action(c *urcli.Context) error {

	bt, err := hex.DecodeString(acc.pkstring)
	if err != nil {
		return DecodeError(err)
	}

	var a = account.NewAccount(bt)
	fmt.Println("Private Key:", hex.EncodeToString(a.PrivateKey))
	fmt.Println("Public Key:", hex.EncodeToString(a.PublicKey))

	return nil
}

func NewAccCreateCmd(acc *AccCmd) urcli.Command {
	var accCreate = &AccCreateCmd{parentCmd:acc}
	return urcli.Command {
		Name: "create",
		ShortName: "ct",
		Usage: "account api",
		UsageText: "create new account",
		Category: "account",
		Flags: []urcli.Flag {
			urcli.StringFlag {
				Name: "seed, sd",
				Value: "",
				Usage: "ed25519 seed(hex to bytes)",
				Destination: &accCreate.pkstring,
			},
		},
		Action: accCreate.Action,
		After: nil,
		Subcommands: nil,
	}
}
