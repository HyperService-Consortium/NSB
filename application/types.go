package nsb

import (
	"fmt"
	"encoding/hex"
	"github.com/Myriad-Dreamin/NSB/math"
	"github.com/Myriad-Dreamin/NSB/merkmap"
	"github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/syndtr/goleveldb/leveldb"
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
	StateRoot []byte `json:"action_root"`
	Height  int64  `json:"height"`
}

type AccountInfo struct {
	Balance     *math.Uint256 `json:"balance"`
	CodeHash    []byte        `json:"code_hash"`
	StorageRoot []byte        `json:"storage_root"`
}

func (accInfo *AccountInfo) String() string {
	return fmt.Sprintf(
		"Balance: %v\nodeHash: %v\nStorageRoot: %v\n",
		accInfo.Balance.String(),
		hex.EncodeToString(accInfo.CodeHash),
		hex.EncodeToString(accInfo.StorageRoot),
	)
}
