package nsb

import (
	"encoding/json"
	"github.com/Myriad-Dreamin/NSB/application/response"
	cmn "github.com/Myriad-Dreamin/NSB/common"
	"github.com/Myriad-Dreamin/NSB/crypto"
	"github.com/Myriad-Dreamin/NSB/util"
	"github.com/tendermint/tendermint/abci/types"
)

/*
 * storage := actionMap
 */

func MustUnmarshal(data []byte, load interface{}) {
	err := json.Unmarshal(data, &load)
	if err != nil {
		panic(response.DecodeJsonError(err))
	}
}

func (nsb *NSBApplication) ActionRigisteredMethod(
	env *cmn.TransactionHeader,
	accInfo *AccountInfo,
	funcName string,
	args []byte,
) *types.ResponseDeliverTx {
	switch funcName {
	case "addAction":
		return nsb.addAction(args)
	case "getAction":
		return nsb.getAction(args)
	default:
		return response.InvalidFuncTypeError(MethodMissing)
	}
}

type ArgsAddAction struct {
	ISCAddress []byte `json:"1"`
	// hexbytes
	Tid uint64 `json:"2"`
	// hexbytes
	Aid       uint64 `json:"3"`
	Type      uint8  `json:"4"`
	Content   []byte `json:"5"`
	Signature []byte `json:"6"`
}

func actionKey(addr []byte, tid uint64, aid uint64) []byte {
	return crypto.Sha512(addr, util.Uint64ToBytes(tid), util.Uint64ToBytes(aid))
}

func (nsb *NSBApplication) addAction(bytesArgs []byte) *types.ResponseDeliverTx {
	var args ArgsAddAction
	MustUnmarshal(bytesArgs, &args)
	// TODO: check valid isc/tid/aid
	err := nsb.actionMap.TryUpdate(
		actionKey(args.ISCAddress, args.Tid, args.Aid),
		util.ConcatBytes([]byte{args.Type}, args.Content, args.Signature),
	)
	if err != nil {
		return response.ContractExecError(err)
	}
	return &types.ResponseDeliverTx{
		Code: uint32(response.CodeOK()),
		Info: "updateSuccess",
	}
}

type ArgsGetAction struct {
	ISCAddress []byte `json:"1"`
	// hexbytes
	Tid uint64 `json:"2"`
	// hexbytes
	Aid uint64 `json:"3"`
}

func (nsb *NSBApplication) getAction(bytesArgs []byte) *types.ResponseDeliverTx {
	var args ArgsGetAction
	MustUnmarshal(bytesArgs, &args)
	// TODO: check valid isc/tid/aid
	bt, err := nsb.actionMap.TryGet(actionKey(args.ISCAddress, args.Tid, args.Aid))
	if err != nil {
		return response.ContractExecError(err)
	}
	return &types.ResponseDeliverTx{
		Code: uint32(response.CodeOK()),
		Data: bt,
	}
}
