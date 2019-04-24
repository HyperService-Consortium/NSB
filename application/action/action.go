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

// function addAction(bytes32 msghash, bytes memory signature)
// 	// future will be private
// 	public
// 	// ownerExists(msg.sender)
// 	returns (bytes32 keccakhash)
// {
// 	// require(pa != 0, "invalid pa address");
// 	// require(pz != 0, "invalid pz address");
// 	// require(verify(msg, signature))
// 	Action memory toAdd = Action(msghash, signature);
// 	keccakhash = keccak256(abi.encodePacked(msghash ,signature));
// 	validActionorNot[keccakhash] = true;
// 	ActionTree[keccakhash]= toAdd;
// }

// function getAction(bytes32 keccakhash)
// 	public
// 	view
// 	returns (bytes memory signature)
// {
// 	Action storage toGet = ActionTree[keccakhash];
// 	// pa = toGet.pa;
// 	// pz = toGet.pz;
// 	// msghash = toGet.msghash;
// 	signature = toGet.signature;
// }

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
