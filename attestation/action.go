package action

import (
	// dbm "github.com/tendermint/tendermint/libs/db"
	"encoding/json"
	"fmt"
	"github.com/HyperServiceOne/NSB/contract/isc/action/ActionType"
	"github.com/HyperServiceOne/NSB/contract/isc/action/ActionError"
)


type Action struct {
	Atype       uint8             `json:"action_type"`
	Signature   []byte                      `json:"signatrue"`
	MsgHash     []byte                      `json:"msg_hash"`
}

func checkEthereumAction(action *Action) (retAction *Action, err error) {
	fmt.Println("adding EthereumAction", action)
	return nil, nil
}

func checkNebulasAction(action *Action) (retAction *Action, err error) {
	fmt.Println("adding NebulasAction", action)
	return nil, nil
}

func checkTendermintAction(action *Action) (retAction *Action, err error) {
	fmt.Println("adding TendermintAction", action)
	return nil, nil
}

func AddAction(byteJson []byte) (*Action, error) {
	var action = new(Action)
	
	err := json.Unmarshal(byteJson, action)
	if err != nil {
		return nil, err
	}

	switch action.Atype {

	case ActionType.EthereumAction:
		return checkEthereumAction(action)

	case ActionType.NebulasAction:
		return checkNebulasAction(action)

	case ActionType.TendermintAction:
		return checkTendermintAction(action)

	default:
		return nil, ActionError.UnrecognizedType
	}
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

// func getActionByMsgHash(db dbm.DB, msgHash []byte) Action {
// 	actionBytes := db.Get(msgHash)
// 	var action Action
// 	if len(actionBytes) != 0 {
// 		err := json.Unmarshal(actionBytes, &action)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// 	return action
// }


// func getActionBySignature(db dbm.DB, signature []byte) Action {
// 	msgHash := db.Get(signature)
// 	var action Action
// 	if len(msgHash) != 0 {
// 		action = getActionByMsgHash(db, msgHash)
// 	}
// 	return action
// }
