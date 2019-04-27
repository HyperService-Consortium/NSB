package nsb

import (
	"encoding/json"
	"github.com/Myriad-Dreamin/NSB/math"
	"github.com/tendermint/tendermint/abci/types"
)


type NSBApplication struct {
	types.BaseApplication
	state       *NSBState
	stateMap    *merkmap.MerkMap
	accMap      *merkmap.MerkMap
	txMap       *merkmap.MerkMap
	statedb     *leveldb.DB
	ValUpdates  []types.ValidatorUpdate
	logger      log.Logger
}


type NSBState struct {
	db dbm.DB
	StateRoot trie.Hash `json:"action_root"`
	Height  int64  `json:"height"`
}


type AccountInfo struct {
	Balance     *math.Uint256 `json:"balance"`
	CodeHash    []byte        `json:"code_hash"`
	StorageRoot []byte        `json:"storage_root"`
}


type TransactionHeader struct {
	From []byte  `json:"from"`
	ContractAddress []byte  `json:"to"`
	JsonParas []byte `json:"data"`
	Value *math.Uint256 `json:"value"`
	Nonce *math.Uint256 `json:"nonce"`
	Signature []byte `json:"signature"`
}


type ContractEnvironment struct {
	Storage *localstorage.LocalStorage
	From []byte
	fromInfo *AccountInfo
	ContractAddress []byte
	toInfo *AccountInfo
	Data []byte
	Value []byte
}

type KVPair interface {
	// must be bytes
	Key() []byte
	// must be bytes
	Value() []byte
}

type ContractCallBackInfo struct {
	// type responceDeliverTx
	CodeResponse uint32
	Log string
	Info string
	Tags []KVPair
}