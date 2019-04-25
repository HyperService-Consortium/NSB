package action

import (
	"testing"
	"fmt"
	"encoding/json"
	"github.com/Myriad-Dreamin/NSB/application/action/ActionType"
	"github.com/Myriad-Dreamin/NSB/application/action/ActionError"
)


func TestAddEthAction(t *testing.T) {
	// ethereum merkle proof test
	var action = &Action{
		Atype: ActionType.EthereumAction,
		Signature: []byte("\x01\x02\x03\x04\x05\x06\x07\x08"),
		MsgHash: []byte("\x01\x02\x03\x04\x05\x06\x07\x08")}
	var jsonAction, err = json.Marshal(action)
	if err != nil {
		t.Error(err)
		return
	}
	var proof *Action
	proof, err = AddAction(jsonAction)
	fmt.Println(proof)
	if err != nil {
		t.Error(err)
	}
}

func TestInvalidAction(t *testing.T) {
	// invalid merkle proof test
	var action = &Action{
		Atype: 5,
		Signature: []byte(""),
		MsgHash: []byte("")}	
	
	var jsonAction, err = json.Marshal(action)
	if err != nil {
		t.Error(err)
		return
	}
	
	var proof *Action
	proof, err = AddAction(jsonAction)
	fmt.Println(proof)
	if err != nil && err != ActionError.UnrecognizedType {
		t.Error(err)
	}
}