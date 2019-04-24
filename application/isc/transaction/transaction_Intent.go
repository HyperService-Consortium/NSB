package transaction


type TransactionIntent struct {
	TxHash      []byte              `json:"transaction_hash"`
	// [][]byte?
	ActionRoot  []byte              `json:"action_root_hash"`
	// [][]byte?
	ProofRoot   []byte              `json:"proof_root_hash"`
}


type TransactionIntents struct {
	ContractId  []byte              `json:"contract_id"`
	txs         []TransactionIntent `json:"transactions"`
}
