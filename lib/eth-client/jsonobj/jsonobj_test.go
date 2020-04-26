package jsonobj

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBlockByNumber(t *testing.T) {

	assert.EqualValues(t,
		string(GetBlockByNumber(255, false)),
		`{"id":1,"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0xff",false]}`)
}
func TestGetLatestBlock(t *testing.T) {

	assert.EqualValues(t,
		string(GetBlockByTag(XTagLatest, false)),
		`{"id":1,"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["latest",false]}`)
}

func TestGetProofByStringAddress(t *testing.T) {
	fmt.Println(string(GetProofByStringAddress(
		"0x4b3a59cd1183ab81b3c31b5a22bce26adf928ac2",
		[]byte("[]"), 8035)))
}
