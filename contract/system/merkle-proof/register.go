package system_merkle_proof

import (
	"github.com/HyperService-Consortium/NSB/application/response"
	cmn "github.com/HyperService-Consortium/NSB/common"
	"github.com/tendermint/tendermint/abci/types"
)

func (nsb *Contract) RegisteredMethod(
	env *cmn.TransactionHeader,
	frInfo *cmn.AccountInfo,
	toInfo *cmn.AccountInfo,
	funcName string,
	args []byte) *types.ResponseDeliverTx {
	switch funcName {
	case "validateMerkleProof":
		return nsb.validateMerkleProof(args)
	case "getMerkleProof":
		return nsb.getMerkleProof(args)
	case "addBlockCheck":
		return nsb.addBlockCheck(args)
	default:
		return response.InvalidFuncTypeError(response.MethodMissing)
	}
}
