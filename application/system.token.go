package nsb

import (
	"fmt"
	"github.com/HyperService-Consortium/NSB/util"

	"github.com/HyperService-Consortium/NSB/application/response"
	cmn "github.com/HyperService-Consortium/NSB/common"
	"github.com/HyperService-Consortium/NSB/math"
	"github.com/tendermint/tendermint/abci/types"
)

/*
 * storage := actionMap
 */
type ArgsSetBalance struct {
	Value *math.Uint256 `json:"1"`
}

type ArgsTransfer struct {
	Value *math.Uint256 `json:"1"`
}

func (nsb *NSBApplication) TokenRigisteredMethod(
	env *cmn.TransactionHeader,
	frInfo *cmn.AccountInfo,
	toInfo *cmn.AccountInfo,
	funcName string,
	args []byte,
) *types.ResponseDeliverTx {
	switch funcName {
	case "setBalance":
		var uargs ArgsSetBalance
		util.MustUnmarshal(args, &uargs)
		return nsb.setBalance(frInfo, uargs.Value)
	case "transfer":
		var uargs ArgsTransfer
		util.MustUnmarshal(args, &uargs)
		return nsb.transfer(frInfo, toInfo, uargs.Value)
	default:
		return response.InvalidFuncTypeError(response.MethodMissing)
	}
}

func (nsb *NSBApplication) setBalance(accInfo *cmn.AccountInfo, value *math.Uint256) *types.ResponseDeliverTx {
	accInfo.Balance = value
	return response.ExecOK()
}

func (nsb *NSBApplication) transfer(
	frInfo *cmn.AccountInfo,
	toInfo *cmn.AccountInfo,
	value *math.Uint256,
) *types.ResponseDeliverTx {

	// resist frInfo == toInfo
	if frInfo == toInfo {
		return response.ExecOK()
	}

	checkErr := frInfo.Balance.Sub(value)
	if checkErr {
		return response.ExecContractError(fmt.Errorf("'from' account has no enough token to substract"))
	}
	checkErr = toInfo.Balance.Add(value)
	if checkErr {
		return response.ExecContractError(fmt.Errorf("'to' account's balance overflow"))
	}
	return response.ExecOK()
}
