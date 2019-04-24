package MerkleProofError

import "errors"

var (
	UnrecognizedType = errors.New("MerkleProofError: Unrecognized Merkle Proof Type")
)