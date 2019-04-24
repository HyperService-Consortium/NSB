package MerkMapError

import "errors"

var (
	UnrecognizedType = errors.New("MerkMapError: Unrecognized Merkle Proof Type")
)