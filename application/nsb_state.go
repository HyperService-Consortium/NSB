package nsb

import (
	"github.com/Myriad-Dreamin/go-mpt"
	dbm "github.com/tendermint/tendermint/libs/db"
	"encoding/json"
)

type NSBState struct {
	db dbm.DB
	ActionRoot trie.MerkleHash `json:"action_root"`
	MerkleProofRoot trie.MerkleHash `json:"merkle_proof_root"`
	ActiveISCRoot trie.MerkleHash `json:"active_isc_root"`
	Height  int64  `json:"height"`
	AppHash []byte `json:"app_hash"`
}

func (st *NSBState) String() string {
	return string(
		"ActionRoot: "      + string(st.ActionRoot)      + "\n" + 
		"MerkleProofRoot: " + string(st.MerkleProofRoot) + "\n" +
		"ActiveISCRoot: "   + string(st.ActiveISCRoot)   + "\n" +
		"Height: "          + string(st.Height)          + "\n" + 
		"AppHash: "         + string(st.AppHash)         + "\n")
}

func NewNSBState() *NSBState {
	return &NSBState{
		db: nil,
		ActionRoot: nil,
		MerkleProofRoot: nil,
		ActiveISCRoot: nil,
		Height: 0,
		AppHash: nil}
}

func loadState(db dbm.DB) *NSBState {
	stateBytes := db.Get(stateKey)
	var state NSBState
	if len(stateBytes) != 0 {
		err := json.Unmarshal(stateBytes, &state)
		if err != nil {
			panic(err)
		}
	}
	state.db = db
	return &state
}

func saveState(state *NSBState) {
	stateBytes, err := json.Marshal(state)
	if err != nil {
		panic(err)
	}
	state.db.Set(stateKey, stateBytes)
}
