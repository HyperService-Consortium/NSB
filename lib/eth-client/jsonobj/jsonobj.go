package jsonobj

import (
	"bytes"
	"encoding/hex"
	bytespool "github.com/HyperService-Consortium/NSB/lib/bytes-pool"
	"strconv"
	"strings"
)

const (
	maxBufferSize = 1024
	splitByte     = ','
	ssplitByte    = '"'
	cbx           = `","`
	endParamByte  = ']'
	endJSONByte   = '}'
	hexPrefix     = "0x"

	XTagLatest = "latest"

	maxBytesSize = 1024
)

var (
	reqGetAccount                 = []byte(`{"id":1,"jsonrpc":"2.0","method":"eth_accounts","params":[]}`)
	reqPersonalUnlock             = []byte(`{"id":64,"jsonrpc":"2.0","method":"personal_unlockAccount","params":["`)
	reqSendTransaction            = []byte(`{"id":1,"jsonrpc":"2.0","method":"eth_sendTransaction","params":[`)
	reqGetStorageAt               = []byte(`{"id":1,"jsonrpc":"2.0","method":"eth_getStorageAt","params":[`)
	reqGetTransactionByHash       = []byte(`{"id":1,"jsonrpc":"2.0","method":"eth_getTransactionByHash","params":["0x`)
	reqGetTransactionByStringHash = []byte(`{"id":1,"jsonrpc":"2.0","method":"eth_getTransactionByHash","params":["`)
	reqGetBlockByHash             = []byte(`{"id":1,"jsonrpc":"2.0","method":"eth_getBlockByHash","params":["0x`)
	reqGetBlockByNumber           = []byte(`{"id":1,"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["`)
	reqGetProof                   = []byte(`{"id":1,"jsonrpc":"2.0","method":"eth_getProof","params":["`)
	reqGetTransactionByHashEnd    = []byte(`"]}`)
	reqGetBlockByHashEnd          = []byte(`]}`)
	bp                            = bytespool.NewMultiThreadBytesPool(maxBytesSize)
)

// GetAccount return all accounts on eth
func GetAccount() []byte {
	return reqGetAccount
}

// GetPersonalUnlock return whether unlocked
// do not send too long passphrase
func GetPersonalUnlock(addr string, passphrase string, duration int) []byte {
	var b = bp.Get()
	var buf = bytes.NewBuffer(b)
	buf.Reset()

	buf.Write(reqPersonalUnlock)

	if !strings.HasPrefix(addr, hexPrefix) {
		buf.WriteString(hexPrefix)
	}
	buf.WriteString(addr)

	buf.WriteString(cbx)

	buf.WriteString(passphrase)

	buf.WriteByte(ssplitByte)
	buf.WriteByte(splitByte)

	buf.WriteString(strconv.Itoa(duration))

	buf.WriteByte(endParamByte)
	buf.WriteByte(endJSONByte)

	return buf.Bytes()
}

// GetSendTransaction return whether unlocked
// do not send too long obj
func GetSendTransaction(obj []byte) []byte {
	var b = bp.Get()
	var buf = bytes.NewBuffer(b)
	buf.Reset()

	buf.Write(reqSendTransaction)

	buf.Write(obj)

	buf.WriteByte(endParamByte)
	buf.WriteByte(endJSONByte)

	return buf.Bytes()
}

// GetStorageAt return whether unlocked
// do not send too long obj
func GetStorageAt(address, pos []byte, tag string) []byte {
	var b = bp.Get()
	var buf = bytes.NewBuffer(b)
	buf.Reset()

	buf.Write(reqGetStorageAt)

	buf.WriteByte(ssplitByte)

	buf.WriteString(hexPrefix)
	buf.WriteString(hex.EncodeToString(address))

	buf.WriteString(cbx)

	buf.WriteString(hexPrefix)
	buf.Write(pos)

	buf.WriteString(cbx)

	buf.WriteString(tag)

	buf.WriteByte(ssplitByte)

	buf.WriteByte(endParamByte)
	buf.WriteByte(endJSONByte)

	return buf.Bytes()
}

func GetTransactionByHash(transactionHash []byte) []byte {
	var b = bp.Get()
	var buf = bytes.NewBuffer(b)
	buf.Reset()

	buf.Write(reqGetTransactionByHash)
	buf.WriteString(hex.EncodeToString(transactionHash))
	buf.Write(reqGetTransactionByHashEnd)
	return buf.Bytes()
}

func GetTransactionByStringHash(transactionHash string) []byte {
	var b = bp.Get()
	var buf = bytes.NewBuffer(b)
	buf.Reset()

	buf.Write(reqGetTransactionByStringHash)
	buf.WriteString(transactionHash)
	buf.Write(reqGetTransactionByHashEnd)
	return buf.Bytes()
}

func GetBlockByHash(blockHash []byte, returnFull bool) []byte {
	var b = bp.Get()
	var buf = bytes.NewBuffer(b)
	buf.Reset()

	buf.Write(reqGetBlockByHash)

	buf.WriteString(hex.EncodeToString(blockHash))

	buf.WriteByte(ssplitByte)
	buf.WriteByte(splitByte)

	buf.WriteString(strconv.FormatBool(returnFull))

	buf.Write(reqGetBlockByHashEnd)
	return buf.Bytes()
}

func GetBlockByNumber(blockNumber uint64, returnFull bool) []byte {
	var b = bp.Get()
	var buf = bytes.NewBuffer(b)
	buf.Reset()

	buf.Write(reqGetBlockByNumber)

	buf.WriteString(hexPrefix)
	buf.WriteString(strconv.FormatUint(blockNumber, 16))

	buf.WriteByte(ssplitByte)
	buf.WriteByte(splitByte)

	buf.WriteString(strconv.FormatBool(returnFull))

	buf.WriteByte(endParamByte)
	buf.WriteByte(endJSONByte)
	return buf.Bytes()
}

func GetBlockByTag(str string, returnFull bool) []byte {
	var b = bp.Get()
	var buf = bytes.NewBuffer(b)
	buf.Reset()

	buf.Write(reqGetBlockByNumber)

	buf.WriteString(str)

	buf.WriteByte(ssplitByte)
	buf.WriteByte(splitByte)

	buf.WriteString(strconv.FormatBool(returnFull))

	buf.WriteByte(endParamByte)
	buf.WriteByte(endJSONByte)
	return buf.Bytes()
}

func GetProofByStringAddress(address string, reqProofs []byte, blockNumber uint64) []byte {
	var b = bp.Get()
	var buf = bytes.NewBuffer(b)
	buf.Reset()

	buf.Write(reqGetProof)

	buf.WriteString(address)

	buf.WriteByte(ssplitByte)
	buf.WriteByte(splitByte)

	buf.Write(reqProofs)

	buf.WriteByte(splitByte)
	buf.WriteByte(ssplitByte)
	buf.WriteString(hexPrefix)
	buf.WriteString(strconv.FormatUint(blockNumber, 16))
	buf.WriteByte(ssplitByte)

	buf.WriteByte(endParamByte)
	buf.WriteByte(endJSONByte)
	return buf.Bytes()
}

// ReturnBytes to Pool
func ReturnBytes(b []byte) {
	bp.Put(b)
}
