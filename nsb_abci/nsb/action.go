package nsb

import (
	dbm "github.com/tendermint/tendermint/libs/db"
	"encoding/json"
)


type ActionType uint8;
const (
	EthereumAction ActionType = 0 + iota
	NebulasAction
	TendermintAction
)


type Action struct {
	Atype       ActionType          `json:"action_type"`
	Signature   []byte              `json:"signatrue"`
	MsgHash     []byte              `json:"msg_hash"`
}


func getActionByMsgHash(db dbm.DB, msgHash []byte) Action {
	actionBytes := db.Get(msgHash)
	var action Action
	if len(actionBytes) != 0 {
		err := json.Unmarshal(actionBytes, &action)
		if err != nil {
			panic(err)
		}
	}
	return action
}


func getActionBySignature(db dbm.DB, signature []byte) Action {
	msgHash := db.Get(signature)
	var action Action
	if len(msgHash) != 0 {
		action = getActionByMsgHash(db, msgHash)
	}
	return action
}
