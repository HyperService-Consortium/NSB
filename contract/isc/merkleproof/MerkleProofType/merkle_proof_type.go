package MerkleProofType

type Type uint8;
const (
	EthereumMerkleProof Type = 0 + iota
	NebulasMerkleProof
	TendermintMerkleProof
)