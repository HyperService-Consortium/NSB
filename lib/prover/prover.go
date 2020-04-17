package prover

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/NSB/crypto"
	"github.com/HyperService-Consortium/go-rlp"
)

func Validate(expected []byte, data []byte) bool {
	return bytes.Equal(expected, crypto.Keccak256(data))
}

type ValidateError struct {
	Depth int `json:"depth"`
}

func (v ValidateError) Error() string {
	return fmt.Sprintf("validate error: compare hash failed at depth %v", v.Depth)
}

func _getMerkleProofValueWithValidate(currentRoot []byte, proof [][]byte, key string, index int) ([]byte, error) {
	if len(proof) == 0 {
		return nil, errors.New("consumed")
	}
	if !Validate(currentRoot, proof[0]) {
		return nil, ValidateError{Depth: index}
	}

	var node [][]byte

	if err := rlp.DecodeBytes(proof[0], &node); err != nil {
		return nil, err
	}

	if len(node) == 2 { // (k, v)
		var k = node[0]
		if len(k) == 0 {
			return nil, fmt.Errorf("bad proof, key with zero length (path %v)", key[:index])
		}

		var strK = hex.EncodeToString(k)[1:]
		var l = len(strK)
		if l > len(key) {
			l = len(key)
		}
		for i := 0; i+index < l; i++ {
			if strK[i] != key[i+index] {
				return nil, fmt.Errorf("compared key failed, key %v, actually %v (path %v)",
					strK, key[index:], key[:index])
			}
		}

		if l == len(strK) {
			return node[1], nil
		} else {
			return nil, errors.New("todo")
		}
	} else if len(node) == 17 { // ('0' ~ 'f' index, v) node
		return _getMerkleProofValueWithValidate(
			node[key[index]-'0'], proof[1:], key, index+1)
	}
	return nil, errors.New("unknown tree type")
}

func GetMerkleProofValueWithValidate(currentRoot []byte, proof [][]byte, key string) ([]byte, error) {
	return _getMerkleProofValueWithValidate(currentRoot, proof, key, 0)
}
