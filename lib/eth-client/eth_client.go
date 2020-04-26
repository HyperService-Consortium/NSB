package ethclient

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"net/url"

	"github.com/HyperService-Consortium/NSB/lib/eth-client/jsonobj"
	jsonrpc_client "github.com/HyperService-Consortium/NSB/lib/rpc-client"
)

// EthClient provide interface to ethereum rpc service
type EthClient struct {
	*jsonrpc_client.JsonRpcClient
}

// NewEthClient return a pointer of eth-client object
func NewEthClient(host string) *EthClient {
	return &EthClient{
		JsonRpcClient: jsonrpc_client.NewJsonRpcClient(host),
	}
}

//GetEthAccounts return accounts as strings
func (eth *EthClient) GetEthAccounts() ([]string, error) {

	b, err := eth.JsonRpcClient.PostWithBody(jsonobj.GetAccount())
	if err != nil {
		return nil, err
	}

	var x []string
	err = json.Unmarshal(b, &x)
	if err != nil {
		return nil, err
	}

	return x, err
}

// PersonalUnlockAccout return whether the account is unlocked successfully
func (eth *EthClient) PersonalUnlockAccout(addr string, passphrase string, duration int) (bool, error) {
	b := jsonobj.GetPersonalUnlock(addr, passphrase, duration)
	bb, err := eth.JsonRpcClient.PostWithBody(b)
	jsonobj.ReturnBytes(b)
	if err != nil {
		return false, err
	}

	return gjson.ParseBytes(bb).Bool(), err
}

// SendTransaction return the receipt of sending transaction
func (eth *EthClient) SendTransaction(obj []byte) (string, error) {
	b := jsonobj.GetSendTransaction(obj)
	bb, err := eth.JsonRpcClient.PostWithBody(b)
	jsonobj.ReturnBytes(b)
	if err != nil {
		return "", err
	}

	return gjson.ParseBytes(bb).String(), err
}

var ErrInvalidEthPos = errors.New("invalid pos")

// GetStorageAt return the value of position on the address
func (eth *EthClient) GetStorageAt(contractAddress, pos []byte, tag string) (string, error) {
	pos = []byte(hex.EncodeToString(pos))
	for i := range pos {
		if !('0' <= pos[i] && pos[i] <= '9') && !('a' <= pos[i] && pos[i] <= 'f') {
			return "", ErrInvalidEthPos
		}
	}
	b := jsonobj.GetStorageAt(contractAddress, pos, tag)
	bb, err := eth.JsonRpcClient.PostWithBody(b)
	jsonobj.ReturnBytes(b)
	if err != nil {
		return "", err
	}

	return gjson.ParseBytes(bb).String(), err
}

func (eth *EthClient) GetTransactionByHash(transactionHash []byte) ([]byte, error) {
	return eth.postBytes(jsonobj.GetTransactionByHash(transactionHash))
}

func (eth *EthClient) GetTransactionByStringHash(transactionHash string) ([]byte, error) {
	return eth.postBytes(jsonobj.GetTransactionByStringHash(transactionHash))
}

func (eth *EthClient) GetBlockByHash(blockHash []byte, returnFull bool) ([]byte, error) {
	return eth.postBytes(jsonobj.GetBlockByHash(blockHash, returnFull))
}

func (eth *EthClient) GetBlockByNumber(blockNumber uint64, returnFull bool) ([]byte, error) {
	return eth.postBytes(jsonobj.GetBlockByNumber(blockNumber, returnFull))
}

const (
	TagLatest = jsonobj.XTagLatest
)

func (eth *EthClient) GetBlockByTag(tag string, returnFull bool) ([]byte, error) {
	return eth.postBytes(jsonobj.GetBlockByTag(tag, returnFull))
}

func (eth *EthClient) GetProofByNumberSR(address string, reqProofs []byte, blockNumber uint64) ([]byte, error) {
	return eth.postBytes(jsonobj.GetProofByStringAddress(address, reqProofs, blockNumber))
}

func (eth *EthClient) postBytes(b []byte) ([]byte, error) {
	bb, err := eth.JsonRpcClient.PostWithBody(b)

	jsonobj.ReturnBytes(b)
	if err != nil {
		return nil, err
	}
	return bb, nil
}

// Do raw rpc invocation
func Do(url string, jsonBody []byte) ([]byte, error) {
	return jsonrpc_client.Do(url, jsonBody)
}

func HTTPDo(tcpAddr string, jsonBody []byte) ([]byte, error) {
	return Do((&url.URL{Scheme: "http", Host: tcpAddr, Path: "/"}).String(), jsonBody)
}

type StorageProof struct {
	Key   string   `json:"key"`
	Proof []string `json:"proof"`
	Value string   `json:"value"`
}

type EthereumGetProofReply struct {
	AccountProof  []string       `json:"accountProof"`
	Address       string         `json:"address"`
	Balance       string         `json:"balance"`
	CodeHash      string         `json:"codeHash"`
	Nonce         string         `json:"nonce"`
	StorageHash   string         `json:"storageHash"`
	StorageProofs []StorageProof `json:"storageProof"`
}
