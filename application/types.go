package nsb

import (
	system_merkle_proof "github.com/HyperService-Consortium/NSB/contract/system/merkle-proof"
	nsbrpc "github.com/HyperService-Consortium/NSB/grpc/nsbrpc"
	log "github.com/HyperService-Consortium/NSB/log"
	"github.com/HyperService-Consortium/NSB/merkmap"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tm-db"
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
	logger                     log.TendermintLogger

	system struct {
		merkleProof *system_merkle_proof.Contract
	}
}

type NSBState struct {
	db        dbm.DB
	StateRoot []byte `json:"action_root"`
	Height    int64  `json:"height"`
}

type FAPair = nsbrpc.FAPair
