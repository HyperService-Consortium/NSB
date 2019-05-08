package nsb

import (
	"encoding/hex"
	"fmt"
	"github.com/Myriad-Dreamin/NSB/math"
	"github.com/Myriad-Dreamin/NSB/merkmap"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
)

type NSBApplication struct {
	types.BaseApplication
	state                      *NSBState
	stateMap                   *merkmap.MerkMap
	accMap                     *merkmap.MerkMap
	txMap                      *merkmap.MerkMap
	actionMap                  *merkmap.MerkMap
	validMerkleProofMap        *merkmap.MerkMap
	validOnchainMerkleProofMap *merkmap.MerkMap
	statedb                    *leveldb.DB
	ValUpdates                 []types.ValidatorUpdate
	logger                     log.Logger
}

type NSBState struct {
	db        dbm.DB
	StateRoot []byte `json:"action_root"`
	Height    int64  `json:"height"`
}

type AccountInfo struct {
	Balance     *math.Uint256 `json:"balance"`
	CodeHash    []byte        `json:"code_hash"`
	StorageRoot []byte        `json:"storage_root"`
	Name        []byte        `json:"name"`
}

type FAPair struct {
	FuncName string `json:"function_name"`
	Args     []byte `json:"args"`
}

func (accInfo *AccountInfo) String() string {
	return fmt.Sprintf(
		"Balance: %v\nodeHash: %v\nStorageRoot: %v, name:%v\n",
		accInfo.Balance.String(),
		hex.EncodeToString(accInfo.CodeHash),
		hex.EncodeToString(accInfo.StorageRoot),
		string(accInfo.Name)
	)
}
