package ethclient

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/HyperService-Consortium/NSB/crypto"
	"github.com/HyperService-Consortium/NSB/lib/prover"
	"github.com/HyperService-Consortium/NSB/util"
	"github.com/HyperService-Consortium/go-rlp"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
	"testing"
)

const (
	testHost = "121.89.200.234:8545"
)

func TestGetEthAccounts(t *testing.T) {
	x, err := NewEthClient(testHost).GetEthAccounts()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(x)
}

func TestUnlock(t *testing.T) {
	ok, err := NewEthClient(testHost).PersonalUnlockAccout("0x4b3a59cd1183ab81b3c31b5a22bce26adf928ac2", "123456", 600)

	if ok == false || err != nil {
		if ok == false {
			if err != nil {

				t.Error(err)
			} else {
				t.Errorf("not ok..")
			}
		} else {

			t.Error(err)
		}
		return
	}
}

const objjj = `{"from":"0x4b3a59cd1183ab81b3c31b5a22bce26adf928ac2", "to": "0x4b3a59cd1183ab81b3c31b5a22bce26adf928ac2", "value": "0x1"}`

func TestSendTransaction(t *testing.T) {
	b, err := NewEthClient(testHost).SendTransaction([]byte(objjj))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(b))
}

func TestGetStorageAt(t *testing.T) {
	var addr = "1234567812345678123456781234567812345678"
	baddr, err := hex.DecodeString(addr)
	if err != nil {
		t.Error(err)
		return
	}
	var pos = []byte{1}
	b, err := NewEthClient(testHost).GetStorageAt(baddr, pos, "latest")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(b))
}

func TestGetTransactionByHash(t *testing.T) {
	txb, err := hex.DecodeString("a41d03fde4e7cf4c58870092c65709db7532956f7d0882156f11f503a6d88d2f")
	if err != nil {
		t.Error(err)
		return
	}
	b, err := NewEthClient(testHost).GetTransactionByHash(txb)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(b))
	b, err = NewEthClient(testHost).GetTransactionByStringHash("0xa41d03fde4e7cf4c58870092c65709db7532956f7d0882156f11f503a6d88d2f")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(b))

}

func TestGetBlockByHash(t *testing.T) {
	txb, err := hex.DecodeString("8a8b9aaa48e0fb024abb7105798ad48057cf4fd14100505addabc319ed3d41c6")
	if err != nil {
		t.Error(err)
		return
	}

	b, err := NewEthClient(testHost).GetBlockByHash(txb, true)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(b))

	b, err = NewEthClient(testHost).GetBlockByHash(txb, false)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(string(b))
}

func TestEthClient_GetBlock(t *testing.T) {
	assert.EqualValues(t, string(sugar.HandlerError(NewEthClient(testHost).
		GetBlockByNumber(255, false)).([]byte)),
		`{"difficulty":"0x100","extraData":"0xd88301090d846765746888676f312e31342e32856c696e7578","gasLimit":"0xc78b2528","gasUsed":"0x0","hash":"0x5249777d3de7a6fce47e2008cc22f0b0b3e73a6a6c259d1f99c56cd4113678a0","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","miner":"0x4b3a59cd1183ab81b3c31b5a22bce26adf928ac2","mixHash":"0xe709dab483a5fdfda21d6bdf1e244d910b5f00e631e687168a59729386b77fce","nonce":"0x673156c28c6a64fc","number":"0xff","parentHash":"0xa9b51a9e35c1f6ce9854c04ed00e7426eed56274bb5d87d67ed2c7e8c99cc61c","receiptsRoot":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","size":"0x219","stateRoot":"0x114982b0f02effa4ce8db9132d9db0260f77b02e980021d681a160507878cec5","timestamp":"0x5e95735e","totalDifficulty":"0x10000","transactions":[],"transactionsRoot":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","uncles":[]}`)
}

func getBlockByNumber(t *testing.T) (uint64, string) {

	b := sugar.HandlerError(NewEthClient(testHost).
		GetBlockByTag(TagLatest, false)).([]byte)
	res := gjson.ParseBytes(b)

	blockNumber := sugar.HandlerError(
		strconv.ParseUint(res.Get("number").String()[2:], 16, 64)).(uint64)

	assert.EqualValues(t, string(sugar.HandlerError(NewEthClient(testHost).
		GetBlockByNumber(blockNumber, false)).([]byte)),
		b)
	stateRoot := res.Get("stateRoot").String()

	return blockNumber, stateRoot
}

func TestEthClient_GetBlockByTag(t *testing.T) {
	getBlockByNumber(t)
}

func decodeHexString(b string) []byte {
	if strings.HasPrefix(b, "0x") {
		b = b[2:]
	}
	if len(b)&1 == 1 {
		b = "0" + b
	}
	return sugar.HandlerError(hex.DecodeString(b)).([]byte)
}

func TestEthClient_GetProofByNumberSR(t *testing.T) {
	bn, sr := getBlockByNumber(t)

	b := sugar.HandlerError(NewEthClient(testHost).
		GetProofByNumberSR("0x4b3a59cd1183ab81b3c31b5a22bce26adf928ac2",
			[]byte("[]"), bn)).([]byte)

	var reply EthereumGetProofReply
	sugar.HandlerError0(json.Unmarshal(b, &reply))
	proof := sugar.HandlerError(
		util.ConvertBytesSlice(reply.AccountProof)).([][]byte)
	pv := sugar.HandlerError(prover.GetMerkleProofValueWithValidateMPTSecure(
		sugar.HandlerError(util.ConvertBytes(sr)).([]byte), proof,
		hex.EncodeToString(crypto.Keccak256(
			(sugar.HandlerError(hex.DecodeString(
				"4b3a59cd1183ab81b3c31b5a22bce26adf928ac2"))).([]byte))))).([]byte)
	var v [][]byte
	sugar.HandlerError0(rlp.DecodeBytes(pv, &v))

	fmt.Println(hex.EncodeToString(pv))
	assert.EqualValues(t, 4, len(v))
	assert.EqualValues(t, v[0], decodeHexString(reply.Nonce))
	assert.EqualValues(t, v[1], decodeHexString(reply.Balance))
	assert.EqualValues(t, v[2], decodeHexString(reply.StorageHash))
	assert.EqualValues(t, v[3], decodeHexString(reply.CodeHash))
}
