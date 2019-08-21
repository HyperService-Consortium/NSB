package nsb

import (
	"bytes"
	_ "bytes"
	"errors"
	"fmt"

	"github.com/HyperService-Consortium/NSB/application/response"
	cmn "github.com/HyperService-Consortium/NSB/common"
	"github.com/HyperService-Consortium/NSB/crypto"
	"github.com/HyperService-Consortium/NSB/util"
	merkleprooftype "github.com/Myriad-Dreamin/go-uip/const/merkle-proof-type"
	"github.com/tendermint/tendermint/abci/types"
	// "github.com/Myriad-Dreamin/go-mpt"
)

/*
 * storage := validMerkleProofMap
 * storage2 := validOnchainMerkleProofMap
 */

var (
	bytesOne                      = []byte{1}
	unrecognizedMerkleProofType   = errors.New("unknown merkle proof type")
	evenlenSimpleMerkleProofError = errors.New(
		"MerkleProofError: simple merkle proof must have an odd number of hash nodes",
	)
	wrongMerkleTreeHash = errors.New(
		"MerkleProofError: fail to match the given hash value",
	)
	mptNodesConsumed = errors.New(
		"MerkleProofError: the hash chain is too short to match the key",
	)
	keyConsumed = errors.New(
		"MerkleProofError: the key is too short to match the hash chain",
	)
	wrongValue = errors.New(
		"MerkleProofError: the key does not match the value",
	)
	runeDecodeError = errors.New(
		"MerkleProofError: can not decode rune from key buffer",
	)
	unrecognizedHashFuncType = errors.New(
		"unknown hash function type",
	)
	firstPartMerkleProofMissing = errors.New(
		"can't find the proof of key-value existing on the merkle tree",
	)
	secondPartMerkleProofMismatch = errors.New(
		"the root hash is required to prove or wrong",
	)
)

func errithNode(i int) error {
	return errors.New(fmt.Sprintf("Wrong proof on %v-th node", i))
}

type ArgsValidateMerkleProof struct {
	Type     uint16 `json:"1"`
	RootHash []byte `json:"2"`
	Key      []byte `json:"3"`
	Value    []byte `json:"4"`
	Proof    []byte `json:"5"`
}

type SimpleMerkleProof struct {
	HashChain [][]byte `json:"h"`
}

type MPTMerkleProof struct {
	RootHash  []byte   `json:"r"`
	HashChain [][]byte `json:"h"`
}

func (nsb *NSBApplication) MerkleProofRigisteredMethod(
	env *cmn.TransactionHeader,
	frInfo *AccountInfo,
	toInfo *AccountInfo,
	funcName string,
	args []byte,
) *types.ResponseDeliverTx {
	switch funcName {
	case "validateMerkleProof":
		return nsb.validateMerkleProof(args)
	case "getMerkleProof":
		return nsb.getMerkleProof(args)
	case "addBlockCheck":
		return nsb.addBlockCheck(args)
	default:
		return response.InvalidFuncTypeError(MethodMissing)
	}
}

func validateMerkleProofKey(typeId uint16, rootHash, key []byte) []byte {
	return crypto.Sha512([]byte{uint8(typeId & 0xff), uint8(typeId >> 8)}, rootHash, key)
}

func (nsb *NSBApplication) validateMerkleProof(bytesArgs []byte) *types.ResponseDeliverTx {
	var args ArgsValidateMerkleProof
	MustUnmarshal(bytesArgs, &args)
	switch args.Type {
	case merkleprooftype.SimpleMerkleTreeUsingSha256, merkleprooftype.SimpleMerkleTreeUsingSha512:
		return nsb.validateSimpleMerkleTree(args.Proof, args.Key, args.Type)
	case merkleprooftype.MerklePatriciaTrieUsingKeccak256:
		return nsb.validateMerklePatriciaTrie(args.Proof, args.Key, args.Value, args.Type)
	case merkleprooftype.SecureMerklePatriciaTrieUsingKeccak256:
		return nsb.validateMerklePatriciaTrie(args.Proof, args.Key, args.Value, args.Type)
	default:
		return response.ExecContractError(unrecognizedMerkleProofType)
	}
}

func (nsb *NSBApplication) validateSimpleMerkleTree(
	Proof []byte,
	Key []byte,
	hfType uint16,
) *types.ResponseDeliverTx {
	var jsonProof SimpleMerkleProof
	MustUnmarshal(Proof, &jsonProof)
	if (len(jsonProof.HashChain) & 1) == 0 {
		return response.ExecContractError(evenlenSimpleMerkleProofError)
	}

	// var hf crypto.HashFunc
	// switch hfType {
	// case simpleMerkleTreeUsingSha256:
	// 	hf = crypto.Sha256
	// case simpleMerkleTreeUsingSha512:
	// 	hf = crypto.Sha512
	// default:
	// 	return response.ExecContractError(unrecognizedHashFuncType)
	// }

	// hashChain := append(append(jsonProof.HashChain, Key), []byte{})

	// for idx := len(hashChain) - 2; idx >= 0; idx -= 2 {
	// 	if !bytes.Equal(hf(hashChain[idx], hashChain[idx + 1]), hashChain[idx - 1]) {
	// 		return response.ExecContractError(wrongMerkleTreeHash)
	// 	}
	// }

	// existence
	err := nsb.validMerkleProofMap.TryUpdate(
		validateMerkleProofKey(hfType, jsonProof.HashChain[0], Key),
		bytesOne,
	)
	if err != nil {
		return response.ExecContractError(err)
	}

	return &types.ResponseDeliverTx{
		Code: uint32(response.CodeOK()),
		Info: "nice!",
	}
}

func (nsb *NSBApplication) validateMerklePatriciaTrie(
	Proof []byte,
	Key []byte,
	Value []byte,
	hfType uint16,
) *types.ResponseDeliverTx {
	var jsonProof MPTMerkleProof
	MustUnmarshal(Proof, &jsonProof)

	// var hf crypto.HashFunc
	// switch hfType{
	// case merklePatriciaTrieUsingKeccak256:
	// 	hf = crypto.Keccak256
	// default:
	// 	return response.ExecContractError(unrecognizedHashFuncType)
	// }

	// keybuf := bytes.NewReader(Key)

	// var keyrune rune
	// var keybyte byte
	// var rsize int
	// var err error
	// var hashChain = jsonProof.HashChain
	// var curNode trie.Node
	// var curHash []byte = jsonProof.RootHash
	// // TODO: export node decoder
	// for {

	// 	if len(hashChain) == 0 {
	// 		// TODO: key may be nil here
	// 		return response.ExecContractError(mptNodesConsumed)
	// 	}
	// 	if !bytes.Equal(curHash, hf(hashChain[0])) {
	// 		return response.ExecContractError(wrongMerkleTreeHash)
	// 	}

	// 	curNode, err = trie.DecodeNode(curHash, hashChain[0])
	// 	if err != nil {
	// 		return response.ExecContractError(err)
	// 	}
	// 	hashChain = hashChain[1:]

	// 	switch n := curNode.(type) {
	// 	case *trie.FullNode:
	// 		keyrune, rsize, err = keybuf.ReadRune()
	// 		if err == io.EOF {
	// 			if len(hashChain) != 0 {
	// 				return response.ExecContractError(keyConsumed)
	// 			}
	// 			if !bytes.Equal(n[16], Value) {
	// 				return response.ExecContractError(wrongValue)
	// 			}
	// 			// else:
	// 			goto CheckKeyValueOK;
	// 		} else if err != nil {
	// 			return require.ExecContractError(err)
	// 		}
	// 		if keyrune == utf8.RuneError {
	// 			return response.ExecContractError(runeDecodeError)
	// 		}

	// 		curHash = []byte(curNode[int(keyrune)])
	// 	case *trie.ShortNode:
	// 		for idx := 0; idx < len(n.Key); idx++ {
	// 			keybyte, err = keybuf.ReadByte()
	// 			if err == io.EOF {
	// 				if idx != len(n.Key) - 1 {
	// 					if Value != nil {
	// 						return response.ExecContractError(wrongValue)
	// 					} else {
	// 						goto CheckKeyValueOK;
	// 					}
	// 				} else {
	// 					if len(hashChain) != 0 {
	// 						return response.ExecContractError(keyConsumed)
	// 					}
	// 					if !bytes.Equal([]byte(n.Val), Value) {
	// 						return response.ExecContractError(wrongValue)
	// 					}
	// 					// else:
	// 					goto CheckKeyValueOK;
	// 				}
	// 			} else if err != nil {
	// 				return require.ExecContractError(err)
	// 			}
	// 			if keybyte != n.Key[i] {
	// 				if Value != nil {
	// 					return response.ExecContractError(wrongValue)
	// 				} else {
	// 					goto CheckKeyValueOK;
	// 				}
	// 			}
	// 		}

	// 		curHash = []byte(n.Value)
	// 	}
	// }
	// CheckKeyValueOK:
	// existence
	err := nsb.validMerkleProofMap.TryUpdate(
		validateMerkleProofKey(hfType, jsonProof.RootHash, Key),
		util.ConcatBytes(bytesOne, Value),
	)
	if err != nil {
		return response.ExecContractError(err)
	}

	return &types.ResponseDeliverTx{
		Code: uint32(response.CodeOK()),
		Info: "nice!",
	}
}

type ArgsAddBlockCheck struct {
	ChainID  uint64 `json:"1"`
	BlockID  []byte `json:"2"`
	RtType   uint8  `json:"3"`
	RootHash []byte `json:"4"`
}

func merkleProofKey(chainID uint64, blockID []byte, RtType uint8) []byte {
	return crypto.Sha512(util.Uint64ToBytes(chainID), blockID, []byte{RtType})
}

func (nsb *NSBApplication) addBlockCheck(bytesArgs []byte) *types.ResponseDeliverTx {
	var args ArgsAddBlockCheck
	MustUnmarshal(bytesArgs, &args)
	// TODO: check valid isc/tid/blockid
	err := nsb.validOnchainMerkleProofMap.TryUpdate(
		merkleProofKey(args.ChainID, args.BlockID, args.RtType),
		args.RootHash,
	)
	if err != nil {
		return response.ExecContractError(err)
	}

	return &types.ResponseDeliverTx{
		Code: uint32(response.CodeOK()),
		Info: "updateSuccess",
	}
}

type ArgsGetMerkleProof struct {
	Type     uint16 `json:"1"`
	RootHash []byte `json:"2"`
	Key      []byte `json:"3"`
	// Value    []byte `json:"4"`
	// Proof    []byte `json:"5"`
	ChainID uint64 `json:"4"`
	BlockID []byte `json:"5"`
	RtType  uint8  `json:"6"`
	// Value []byte `json:"7"`
}

func (nsb *NSBApplication) getMerkleProof(bytesArgs []byte) *types.ResponseDeliverTx {
	var args ArgsGetMerkleProof
	MustUnmarshal(bytesArgs, &args)
	// TODO: check valid isc/tid/aid
	bt, err := nsb.validOnchainMerkleProofMap.TryGet(
		merkleProofKey(args.ChainID, args.BlockID, args.RtType),
	)
	if err != nil {
		return response.ExecContractError(err)
	}
	if !bytes.Equal(bt, args.RootHash) {
		return response.ExecContractError(secondPartMerkleProofMismatch)
	}

	bt, err = nsb.validMerkleProofMap.TryGet(
		validateMerkleProofKey(args.Type, args.RootHash, args.Key),
	)
	if err != nil {
		return response.ExecContractError(err)
	}

	return &types.ResponseDeliverTx{
		Code: uint32(response.CodeOK()),
		Data: bt,
	}
}
