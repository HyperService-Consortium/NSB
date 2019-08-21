package transaction

import . "github.com/HyperService-Consortium/NSB/common/contract_response"
import "encoding/json"
import "github.com/HyperService-Consortium/NSB/math"

type TransactionIntent struct {
	Fr          []byte              `json:"from"`
	To          []byte              `json:"to"`
	Seq         *math.Uint256                `json:"seq"`
	Amt         *math.Uint256                `json:"amt"`
	Meta        []byte              `json:"meta"`
}

type TransactionState struct {
	TxHash      []byte              `json:"transaction_hash"`
	// [][]byte?
	ActionRoot  []byte              `json:"action_root_hash"`
	// [][]byte?
	ProofRoot   []byte              `json:"proof_root_hash"`
}


type TransactionStates struct {
	ContractId  []byte              `json:"contract_id"`
	txs         []TransactionIntent `json:"transactions"`
}

func (tx *TransactionIntent) Bytes() []byte {
	bt, err := json.Marshal(tx)
	if err != nil {
		panic(DecodeJsonError(err))
	}
	return bt
}



