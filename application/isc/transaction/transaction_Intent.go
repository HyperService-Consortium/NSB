package transaction


type TransactionIntent struct {
	Fr          []byte              `json:"from"`
	To          []byte              `json:"to"`
	Seq         uint                `json:"seq"`
	Amt         uint                `json:"amt"`
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
