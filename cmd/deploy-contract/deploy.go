package main

import (
	"encoding/json"
	"flag"
	"fmt"
	ethclient "github.com/HyperService-Consortium/NSB/lib/eth-client"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var (
	codePath = flag.String("i", "", "the path of contract to deploy")
	host     = flag.String("host", "", "the ethereum rpc address")
	address  = flag.String("address", "", "transaction.from")
	password = flag.String("password", "123456", "address password")
)

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	if len(*codePath) == 0 {
		log.Fatal("code path empty")
	}
	if len(*host) == 0 {
		log.Fatal("host path empty")
	}

	content, err := ioutil.ReadFile(*codePath)
	if err != nil {
		log.Fatal(err)
	}
	eth := ethclient.NewEthClient(*host)

	ok := sugar.HandlerError(eth.PersonalUnlockAccout(*address, *password, 100)).(bool)
	if !ok {
		log.Fatal(ok)
	}
	//000000000000000000000000dda250dd2646e02ee403da26eb7065dedafb79fd
	//0000000000000000000000000000000000000000000000000000000000000001
	res := sugar.HandlerError(eth.SendTransaction(
		sugar.HandlerError(json.Marshal(struct {
			From  string `json:"from"`
			Data  string `json:"data"`
			Gas   string `json:"gas"`
			Value string `json:"value"`
		}{
			From: *address,
			Data: string(content) +
				formatAddress(*address) +
				"0000000000000000000000000000000000000000000000000000000000000001",
			Gas: "0x" +
				strconv.FormatInt(8000000, 16),
			Value: "0x10",
		})).([]byte)))

	fmt.Println(res)
}

func formatAddress(s string) string {
	if strings.HasPrefix(s, "0x") {
		s = s[2:]
	}
	if len(s) != 40 {
		log.Fatal("bad address")
	}
	return "000000000000000000000000" + s
}
