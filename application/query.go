package nsb

import (
	"github.com/HyperService-Consortium/NSB/application/nsb_proto"
	"github.com/HyperService-Consortium/NSB/application/response"
	"github.com/HyperService-Consortium/NSB/merkmap"
	"github.com/tendermint/tendermint/abci/types"

	// "encoding/hex"
	"encoding/json"
)

const (
	QueryKeyGetStorageAt = "prove_get_storage_at"
	QueryKeyGetAccInfo   = "acc_getAccInfo"
)

type ArgsGetStorageAt = nsb_proto.ArgsGetStorageAt

func (nsb *NSBApplication) QueryIndex(req *types.RequestQuery) (uint32, string) {
	switch req.Path {
	case QueryKeyGetAccInfo:
		return nsb.getAccInfo(req.Data, req.Height)

	case "prove_on_state_trie":
		if req.Data == nil {
			return uint32(response.CodeProofError), nsb.stateMap.MakeErrorProofFromString("nil data")
		}
		return uint32(response.CodeOK()), nsb.stateMap.MakeProof(req.Data)

	case "prove_on_tx_trie":
		if req.Data == nil {
			return uint32(response.CodeProofError), nsb.txMap.MakeErrorProofFromString("nil data")
		}
		return uint32(response.CodeOK()), nsb.txMap.MakeProof(req.Data)

	case "prove_on_account_trie":
		if req.Data == nil {
			return uint32(response.CodeProofError), nsb.accMap.MakeErrorProofFromString("nil data")
		}
		return uint32(response.CodeOK()), nsb.accMap.MakeProof(req.Data)

	case "prove_on_action_trie":
		if req.Data == nil {
			return uint32(response.CodeProofError), nsb.actionMap.MakeErrorProofFromString("nil data")
		}
		return uint32(response.CodeOK()), nsb.actionMap.MakeProof(req.Data)

	case "prove_on_valid_merkle_proof_trie":
		if req.Data == nil {
			return uint32(response.CodeProofError), nsb.validMerkleProofMap.MakeErrorProofFromString("nil data")
		}
		return uint32(response.CodeOK()), nsb.validMerkleProofMap.MakeProof(req.Data)

	case "prove_on_valid_on_chain_merkle_proof_trie":
		if req.Data == nil {
			return uint32(response.CodeProofError), nsb.validOnchainMerkleProofMap.MakeErrorProofFromString("nil data")
		}
		return uint32(response.CodeOK()), nsb.validOnchainMerkleProofMap.MakeProof(req.Data)

	case QueryKeyGetStorageAt:
		var args ArgsGetStorageAt
		err := json.Unmarshal(req.Data, &args)

		if err != nil {
			return uint32(response.CodeProofError), nsb.accMap.MakeErrorProof(err)
		}

		if args.Address == nil {
			return uint32(response.CodeProofError), nsb.accMap.MakeErrorProofFromString("nil account address")
		}

		if len(args.Address) != 32 {
			return uint32(response.CodeProofError), nsb.accMap.MakeErrorProofFromString("err account address: the length is not 32")
		}

		accInfo, errInfo := nsb.parseAccInfo(args.Address)
		if errInfo != nil && errInfo.IsErr() {
			return errInfo.Code, errInfo.Log
		}
		if accInfo.StorageRoot == nil {
			return uint32(response.CodeProofError), nsb.accMap.MakeErrorProofFromString("empty map")
		}
		accStorageTrie, _ :=
			merkmap.NewMerkMapFromDB(
				nsb.statedb, accInfo.StorageRoot, args.Slot)
		return uint32(response.CodeOK()), accStorageTrie.MakeProof(args.Key)
	default:
		return uint32(response.CodeUnknownQueryType), "unknown query type"
	}
}

func (nsb *NSBApplication) getAccInfo(paras []byte, height int64) (uint32, string) {
	// assuming height == 0 // latest
	bytesInfo, err := nsb.accMap.TryGet(paras)
	if err != nil || bytesInfo == nil {
		return uint32(response.CodeAccountNotFound), "the account is not on this AccTrie"
	}
	return 0, string(bytesInfo)
}
