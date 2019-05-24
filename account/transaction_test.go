package account

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/HyperServiceOne/NSB/math"
	"testing"
)

// createContract
// \x19
// isc
// \x18
// txHeader

type TransactionHeader struct {
	From            []byte        `json:"from"`
	ContractAddress []byte        `json:"to"`
	JsonParas       []byte        `json:"data"`
	Value           *math.Uint256 `json:"value"`
	Nonce           *math.Uint256 `json:"nonce"`
	Signature       []byte        `json:"signature"`
}

type TransactionIntent struct {
	Fr   []byte `json:"from"`
	To   []byte `json:"to"`
	Seq  uint   `json:"seq"`
	Amt  uint   `json:"amt"`
	Meta []byte `json:"meta"`
}

type ArgsCreateNewContract struct {
	IscOwners          [][]byte            `json:"isc_owners"`
	Funds              []uint32            `json:"required_funds"`
	VesSig             []byte              `json:"ves_signature"`
	TransactionIntents []TransactionIntent `json:"transactionIntents"`
}

func TestGenerateTransaction(t *testing.T) {
	acc := NewAccount([]byte("1234"))

	TxIntents := []TransactionIntent{
		TransactionIntent{
			Fr:   []byte(acc.PublicKey),
			To:   []byte(acc.PublicKey),
			Seq:  0,
			Amt:  0,
			Meta: []byte(""),
		},
	}

	bt, err := json.Marshal(TxIntents)
	fmt.Println(string(bt))
	if err != nil {
		t.Error(err)
		return
	}

	paras := &ArgsCreateNewContract{
		IscOwners:          [][]byte{[]byte(acc.PublicKey)},
		Funds:              []uint32{0},
		VesSig:             acc.Sign(bt),
		TransactionIntents: TxIntents,
	}

	bt, err = json.Marshal(paras)
	fmt.Println(string(bt))
	if err != nil {
		t.Error(err)
		return
	}

	txHeader := &TransactionHeader{
		From:      acc.PublicKey,
		JsonParas: bt,
		Value:     math.NewUint256FromBytes([]byte{0}),
		Nonce:     math.NewUint256FromBytes([]byte{233, 233}),
	}
	
	// modified , concat them!
	bt, err = json.Marshal(txHeader)
	fmt.Println(string(bt))
	if err != nil {
		t.Error(err)
		return
	}
	bt = acc.Sign(bt)
	txHeader.Signature = bt
	bt, err = json.Marshal(txHeader)
	fmt.Println(string(bt))
	if err != nil {
		t.Error(err)
		return
	}
	var buff bytes.Buffer
	buff.WriteString("createContract")
	buff.WriteByte(byte(0x19))
	buff.WriteString("isc")
	buff.WriteByte(byte(0x18))
	buff.Write(bt)
	fmt.Println(bt)
	fmt.Println(hex.EncodeToString(buff.Bytes()))
	fmt.Println(buff.String())
}

/*
0x637265617465436f6e747261637419697363187b2266726f6d223a223168586d34595a6e30346234726d7041446552632b7643627177666e424650552b72316d5a2b724c344e493d222c22746f223a6e756c6c2c2264617461223a2265794a7063324e66623364755a584a7a496a7062496a466f57473030575670754d4452694e484a74634546455a564a6a4b335a44596e46335a6d3543526c42564b33497862566f72636b7730546b6b39496c3073496e4a6c63585670636d566b58325a31626d527a496a70624d463073496e5a6c6331397a6157647559585231636d55694f694a5a4e6b393064476455515568436155686a523268345545787151565a3262585a30515446465230356d5156564955474a5154335a5556477076546d647352456c345a315a35516b6f79516d523155574e425a47466b526d467a4e444642553078766456704a4d5459336569745854565a45515430394969776964484a68626e4e68593352706232354a626e526c626e527a496a706265794a6d636d3974496a6f694d5768596254525a576d34774e474930636d31775155526c556d4d72646b4e696358646d626b4a4755465572636a4674576974795444524f535430694c434a3062794936496a466f57473030575670754d4452694e484a74634546455a564a6a4b335a44596e46335a6d3543526c42564b33497862566f72636b7730546b6b394969776963325678496a6f774c434a68625851694f6a4173496d316c644745694f69496966563139222c2276616c7565223a22222c226e6f6e6365223a2236656b3d222c227369676e6174757265223a22452f76586e3458364c48664a6a772b795a7154534d36366a7a796b3755356e4a39733449397163364a6b695a68754463566255566c444269524132685a356c7a4945724d575644416f30374d4e612f416f51485741773d3d227d
637265617465436f6e747261637413697363127b2266726f6d223a223168586d34595a6e30346234726d7041446552632b7643627177666e424650552b72316d5a2b724c344e493d222c22746f223a6e756c6c2c2264617461223a2265794a7063324e66623364755a584a7a496a7062496a466f57473030575670754d4452694e484a74634546455a564a6a4b335a44596e46335a6d3543526c42564b33497862566f72636b7730546b6b39496c3073496e4a6c63585670636d566b58325a31626d527a496a70624d463073496e5a6c6331397a6157647559585231636d55694f694a5a4e6b393064476455515568436155686a523268345545787151565a3262585a30515446465230356d5156564955474a5154335a5556477076546d647352456c345a315a35516b6f79516d523155574e425a47466b526d467a4e444642553078766456704a4d5459336569745854565a45515430394969776964484a68626e4e68593352706232354a626e526c626e527a496a706265794a6d636d3974496a6f694d5768596254525a576d34774e474930636d31775155526c556d4d72646b4e696358646d626b4a4755465572636a4674576974795444524f535430694c434a3062794936496a466f57473030575670754d4452694e484a74634546455a564a6a4b335a44596e46335a6d3543526c42564b33497862566f72636b7730546b6b394969776963325678496a6f774c434a68625851694f6a4173496d316c644745694f69496966563139222c2276616c7565223a22222c226e6f6e6365223a2236656b3d222c227369676e6174757265223a22452f76586e3458364c48664a6a772b795a7154534d36366a7a796b3755356e4a39733449397163364a6b695a68754463566255566c444269524132685a356c7a4945724d575644416f30374d4e612f416f51485741773d3d227d
*/
