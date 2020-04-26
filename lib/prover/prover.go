package prover

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/NSB/crypto"
	trie "github.com/HyperService-Consortium/go-mpt"
	"github.com/HyperService-Consortium/go-rlp"
	"io"
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

func _getMerkleProofValueWithValidateMPTSecure(currentRoot []byte, proof [][]byte, key string, index int) ([]byte, error) {
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
		} else if len(proof) > 0 {
			return _getMerkleProofValueWithValidateMPTSecure(
				node[1], proof[1:], key, index+1)
		} else {
			return nil, fmt.Errorf("proof consumed (path %v)", key[:index])
		}
	} else if len(node) == 17 { // ('0' ~ 'f' index, v) node
		if len(proof) > 0 {
			return _getMerkleProofValueWithValidateMPTSecure(
				node[key[index]-'0'], proof[1:], key, index+1)
		} else {
			return nil, fmt.Errorf("proof consumed (path %v)", key[:index])
		}
	}
	return nil, errors.New("unknown tree type")
}

func GetMerkleProofValueWithValidateMPTSecure(currentRoot []byte, proof [][]byte, key string) ([]byte, error) {
	return _getMerkleProofValueWithValidateMPTSecure(currentRoot, proof, key, 0)
}

func GetMerkleProofValueWithValidateNSBMPT(currentHash []byte, hashChain [][]byte, key []byte) ([]byte, error) {
	var (
		keyBuf     = bytes.NewBuffer(key)
		keyByte    byte
		proofDepth = len(hashChain)
	)
	for {
		if len(hashChain) == 0 {
			return nil, fmt.Errorf(
				"proof comsumed (depth %v)", len(key)-keyBuf.Len())
		}
		if !bytes.Equal(currentHash, crypto.Keccak256(hashChain[0])) {
			return nil, fmt.Errorf(
				"hash not exists (currentHash %v, actual %v, depth %v)",
				currentHash, crypto.Keccak256(hashChain[0]), len(key)-keyBuf.Len())
		}

		curNode, err := trie.DecodeNode(currentHash, hashChain[0])
		if err != nil {
			return nil, fmt.Errorf(
				"decode node error: %v, depth %v", err, len(key)-keyBuf.Len())
		}
		hashChain = hashChain[1:]

		switch n := curNode.(type) {
		case *trie.FullNode:
			keyByte, err = keyBuf.ReadByte()
			if err == io.EOF {
				if len(hashChain) != 0 {
					return nil, fmt.Errorf(
						"key consumed (proof depth %v)", proofDepth-len(hashChain))
				}
				return n.Children[16].(trie.ValueNode), nil
			} else if err != nil {
				return nil, fmt.Errorf(
					"key get bytes error: %v (depth %v, proof depth %v)",
					err, len(key)-keyBuf.Len(), proofDepth-len(hashChain))
			} else if keyByte >= 16 {
				return nil, fmt.Errorf(
					"bad nibble: %v (depth %v, proof depth %v)",
					keyByte, len(key)-keyBuf.Len(), proofDepth-len(hashChain))
			}
			ch := n.Children[keyByte]
			if ch == nil {
				return nil, fmt.Errorf(
					"hash not exists (index: %v, depth %v)",
					keyByte, len(key)-keyBuf.Len())
			}
			currentHash = ch.(trie.HashNode)
		case *trie.ShortNode:
			for idx := 0; idx < len(n.Key); idx++ {
				keyByte, err = keyBuf.ReadByte()
				if err == io.EOF {
					if len(hashChain) != 0 {
						return nil, fmt.Errorf(
							"key consumed (proof depth %v)", proofDepth-len(hashChain))
					}
					return nil, nil
				} else if err != nil {
					return nil, fmt.Errorf(
						"key get bytes error: %v (depth %v, proof depth %v)",
						err, len(key)-keyBuf.Len(), proofDepth-len(hashChain))
				}
				if keyByte != n.Key[idx] {
					if len(hashChain) != 0 {
						return nil, fmt.Errorf(
							"key compare failed (proof depth %v)", proofDepth-len(hashChain))
					}
					return nil, nil
				}
			}

			if keyBuf.Len() == 0 {
				if len(hashChain) != 0 {
					return nil, fmt.Errorf(
						"key consumed (proof depth %v)", proofDepth-len(hashChain))
				}

				return n.Val.(trie.ValueNode), nil
			}

			currentHash = n.Val.(trie.HashNode)
		default:
			return nil, errors.New("unknown tree type")
		}
	}
}

func NSBKeyBytesToHex(str []byte) []byte {
	l := len(str)*2 + 1
	var nibbles = make([]byte, l)
	for i, b := range str {
		nibbles[i*2] = b / 16
		nibbles[i*2+1] = b % 16
	}
	nibbles[l-1] = 16
	return nibbles
}

func NSBGetSlot(accountAddress, slotName []byte) []byte {
	return crypto.Sha256(accountAddress, slotName)
}

func NSBGetPos(slot, key []byte) []byte {
	return crypto.Keccak256(slot, key)
}

func NSBGetKeyBySlot(slot, key []byte) []byte {
	return NSBKeyBytesToHex(NSBGetPos(slot, key))
}

func NSBKeyByPure(accountAddress, slotName, key []byte) []byte {
	return NSBGetKeyBySlot(NSBGetSlot(accountAddress, slotName), key)
}
