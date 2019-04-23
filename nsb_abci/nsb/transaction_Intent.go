package nsb


type TransactionIntent struct {
	TxHash      []byte              `json:"transaction_hash"`
	ActionRoot  []byte              `json:"action_root_hash"`
	ProofRoot   []byte              `json:"proof_root_hash"`
}


type TransactionIntents struct {
	ContractId  []byte              `json:"contract_id"`
	txs         []TransactionIntent `json:"transactions"`
}
