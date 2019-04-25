package nsb

import (
	"github.com/Myriad-Dreamin/go-mpt"
	dbm "github.com/tendermint/tendermint/libs/db"
	"encoding/json"
)

type NSBState struct {
	db dbm.DB
	StateRoot trie.MerkleHash `json:"action_root"`
	Height  int64  `json:"height"`
}

func (st *NSBState) String() string {
	return string(
		"StateRoot: " + string(st.ActionRoot) + "\n" + 
		"Height: "      + string(st.Height)     + "\n")
}

func NewNSBState() *NSBState {
	return &NSBState{
		db: nil,
		StateRoot: nil,
		Height: 0}
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
