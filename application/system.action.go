package nsb

import (
	"github.com/tendermint/tendermint/abci/types"
	cmn "github.com/Myriad-Dreamin/NSB/common"
	"github.com/Myriad-Dreamin/NSB/merkmap"
	"github.com/Myriad-Dreamin/NSB/application/response"
	"github.com/Myriad-Dreamin/NSB/util"
	"github.com/Myriad-Dreamin/NSB/crypto"
)

/*
 * storage := actionMap
 */

func (nsb *NSBApplication) ActionRigisteredMethod(
	env *cmn.TransactionHeader,
	accInfo AccountInfo,
	funcName string,
	args []byte,
) *types.ResponseDeliverTx {
	switch funcName {
	case "addAction":
		return nsb.addAction(args)
	default:
		return response.InvalidFuncTypeError(MethodMissing)
	}
}


type ArgsAddAction struct {
	ISCAddress []byte `json:"1"`
	// hexbytes
	Tid uint64 `json:"2"`
	// hexbytes
	Aid uint64 `json:"3"`
	Type uint8 `json:"4"`
	Content []byte `json:"5"`
	Signature []byte `json:"6"`
}


func actionKey(addr []byte, tid uint64, aid uint64) []byte {
	return crypto.Sha512(addr, util.Uint64ToBytes(tid), util.Uint64ToBytes(aid))
}


func (nsb *NSBApplication) addAction(bytesArgs []byte) *types.ResponseDeliverTx {
	var args ArgsAddAction
	util.MustUnmarshal(bytesArgs, &args)
	// TODO: check valid isc/tid/aid
	err := nsb.actionMap.TryUpdate(
		actionKey(args.ISCAddress, args.Tid, args.Aid),
		util.ConcatBytes(Content, Signature),
	)
	if err != nil {
		return response.ContractExecError(err)
	}
	return &types.ResponseDeliverTx {
		Code: uint32(response.CodeOK()),
		Info: "updateSuccess",
	}
}
