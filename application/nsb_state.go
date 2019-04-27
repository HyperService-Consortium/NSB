package nsb

import (
	"errors"
	dbm "github.com/tendermint/tendermint/libs/db"
	"encoding/json"
	"encoding/hex"
)
func (st *NSBState) String() string {
	return "StateRoot: " + hex.EncodeToString(st.StateRoot) + "\nHeight: " + string(st.Height) + "\n"
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

func (st *NSBState) Close() error {
	if st.db == nil {
		return errors.New("the state db is not opened now")
	}
	st.db.Close()
	st.db = nil
	return nil
}