package nsb

import (
	"github.com/tendermint/tendermint/abci/types"
	"github.com/HyperServiceOne/NSB/merkmap"
	// "encoding/hex"
	"encoding/json"
)


type ArgsGetStorageAt struct {
	Address []byte `json:"1"`
	Key []byte `json:"2"`
}


func (nsb *NSBApplication) QueryIndex(req *types.RequestQuery) string {
	switch req.Path {
	case "acc_getAccInfo":
		return nsb.getAccInfo(req.Data, req.Height)

	case "prove_on_state_trie":
		if req.Data == nil {
			return nsb.stateMap.MakeErrorProofFromString("nil data")
		}
		return nsb.stateMap.MakeProof(req.Data)

	case "prove_on_tx_trie":
		if req.Data == nil {
			return nsb.txMap.MakeErrorProofFromString("nil data")
		}
		return nsb.txMap.MakeProof(req.Data)

	case "prove_on_account_trie":
		if req.Data == nil {
			return nsb.accMap.MakeErrorProofFromString("nil data")
		}
		return nsb.accMap.MakeProof(req.Data)

	case "prove_on_action_trie":
		if req.Data == nil {
			return nsb.actionMap.MakeErrorProofFromString("nil data")
		}
		return nsb.actionMap.MakeProof(req.Data)

	case "prove_on_valid_merkle_proof_trie":
		if req.Data == nil {
			return nsb.validMerkleProofMap.MakeErrorProofFromString("nil data")
		}
		return nsb.validMerkleProofMap.MakeProof(req.Data)

	case "prove_on_valid_on_chain_merkle_proof_trie":
		if req.Data == nil {
			return nsb.validOnchainMerkleProofMap.MakeErrorProofFromString("nil data")
		}
		return nsb.validOnchainMerkleProofMap.MakeProof(req.Data)

	case "prove_get_storage_at":
		var args ArgsGetStorageAt
		err := json.Unmarshal(req.Data, &args)

		if err != nil {
			return nsb.accMap.MakeErrorProof(err)
		}

		if args.Address == nil {
			return nsb.accMap.MakeErrorProofFromString("nil account address")
		}
		
		if len(args.Address) != 32 {
			return nsb.accMap.MakeErrorProofFromString("err account address: the length is not 32")
		}

		var bytesAccInfo []byte
		bytesAccInfo, err = nsb.accMap.TryGet(args.Address)
		if err != nil {
			return nsb.accMap.MakeErrorProof(err)
		}

		accInfo, errInfo := nsb.parseAccInfo(bytesAccInfo)
		if errInfo.IsErr() {
			return nsb.accMap.MakeErrorProofFromString(errInfo.Log)
		}
		if accInfo.StorageRoot == nil {
			return nsb.accMap.MakeErrorProofFromString("empty map")
		}
		accStorageTrie, _ := merkmap.NewMerkMapFromDB(nsb.statedb, accInfo.StorageRoot, []byte{})
		return accStorageTrie.MakeProof(req.Data)
	default:
		return "unknown query type"
	}
}

func (nsb *NSBApplication) getAccInfo(paras []byte, height int64) string {
	// assuming height == 0 // latest
	bytesInfo, err := nsb.accMap.TryGet(paras)
	if err != nil || bytesInfo == nil {
		return "the account is not on this AccTrie"
	}
	var accInfo AccountInfo
	err = json.Unmarshal(bytesInfo, &accInfo)
	if err != nil {
		return err.Error()
	}
	return accInfo.String()
}
