package nsb

import (
	_	"bytes"
)

type MerkleProofType uint8;
const (
	EthereumMerkleProof MerkleProofType = 0 + iota
	NebulasMerkleProof
	TendermintMerkleProof
)

type ResponseCode uint32
const (
	CodeOK ResponseCode = 0 + iota
	CodeFail
	CodeUnknown
	CodeMissing
	CodeTODO
)



type MerkleProof struct {
	Mtype       MerkleProofType     `json:"merkle_proof_type"`
	ChainId     string              `json:"chain_id"`
	StorageHash []byte              `json:"storage_hash"`
	key         []byte              `json:"key"`
	value       []byte              `json:"value"`
}

type TransactionIntent struct {
	TxHash      []byte              `json:"transaction_hash"`
	ActionRoot  []byte              `json:"action_root_hash"`
	ProofRoot   []byte              `json:"proof_root_hash"`
}

type TransactionIntents struct {
	ContractId  []byte              `json:"contract_id"`
	txs         []TransactionIntent `json:"transactions"`
}
