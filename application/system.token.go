package nsb

import (
	"encoding/json"
	"github.com/Myriad-Dreamin/NSB/application/response"
	cmn "github.com/Myriad-Dreamin/NSB/common"
	"github.com/Myriad-Dreamin/NSB/math"
	"github.com/tendermint/tendermint/abci/types"
)

/*
 * storage := actionMap
 */

func (nsb *NSBApplication) TokenRigisteredMethod(
	env *cmn.TransactionHeader,
	accInfo *AccountInfo,
	funcName string,
	args []byte,
) *types.ResponseDeliverTx {
	switch funcName {
	case "setBalance":
		var uargs ArgsSetBalance
		MustUnmarshal(args, &uargs)
		return nsb.setBalance(uargs.Value)
	case "getAction":
		return nsb.getAction(args)
	default:
		return response.InvalidFuncTypeError(MethodMissing)
	}
}

type ArgsSetBalance struct {
	Value *math.Uint256 `json:"1"`
}

func (nsb *NSBApplication) setBalance(value *math.Uint256) *types.ResponseDeliverTx {
	return ExecOK(value)
}
