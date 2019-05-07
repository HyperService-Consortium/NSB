package MerkMapError

import "errors"

var (
	UnrecognizedType = errors.New("MerkMapError: Unrecognized Merkle Proof Type")
	DecodeOverflow   = errors.New("MerkMapError: overflow when decoding the merk-slot([32]byte)")
)
