package nsb

import (
	"encoding/json"
	"github.com/Myriad-Dreamin/NSB/application/response"
	cmn "github.com/Myriad-Dreamin/NSB/common"
	"github.com/Myriad-Dreamin/NSB/crypto"
	"github.com/Myriad-Dreamin/NSB/util"
	"github.com/Myriad-Dreamin/NSB/math"
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
	case "setBalance":
		uargs := MustUnmarshal(args, ArgsSetBalance)
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

func (nsb *NSBApplication) setBalance(value *math.Uint256) {
	return ExecOK(value)
}
