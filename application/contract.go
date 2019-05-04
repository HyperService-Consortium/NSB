package nsb

import (
	"fmt"
	"github.com/Myriad-Dreamin/NSB/application/response"
	cmn "github.com/Myriad-Dreamin/NSB/common"
	"github.com/tendermint/tendermint/abci/types"
	// sdeam "github.com/Myriad-Dreamin/NSB/contract/sdeam"
	isc "github.com/Myriad-Dreamin/NSB/contract/isc"
)

func (nsb *NSBApplication) execContractFuncs(
	contractName string,
	contractEnv *cmn.ContractEnvironment,
) (cb *cmn.ContractCallBackInfo) {
	defer func() {
		if r := recover(); r != nil {
			switch r := r.(type) {
			case string:
				cb = &cmn.ContractCallBackInfo{
					CodeResponse: uint32(response.CodeContractPanic()),
					Log:          r,
				}
			case *cmn.ContractCallBackInfo:
				cb = r
			case error:
				cb = &cmn.ContractCallBackInfo{
					CodeResponse: uint32(response.CodeContractPanic()),
					Log:          r.Error(),
				}
			default:
				cb = &cmn.ContractCallBackInfo{
					CodeResponse: uint32(response.CodeContractPanic()),
					Log:          "unknown panic interface...",
				}
			}
		}
	}()

	switch contractName {
	case "isc":
		return isc.RigisteredMethod(contractEnv)
	case "sdeam":
		return &cmn.ContractCallBackInfo{
			CodeResponse: uint32(response.CodeTODO()),
		} // sdeam.RegistedMethod(byteJson)
	default:
		return &cmn.ContractCallBackInfo{
			CodeResponse: uint32(response.CodeInvalidTxType()),
			Log:          "unknown contractName",
		}
	}
}

func (nsb *NSBApplication) createContracts(
	contractName string,
	contractEnv *cmn.ContractEnvironment,
) (cb *cmn.ContractCallBackInfo) {
	defer func() {
		if r := recover(); r != nil {
			switch r := r.(type) {
			case string:
				cb = &cmn.ContractCallBackInfo{
					CodeResponse: uint32(response.CodeContractPanic()),
					Log:          r,
				}
			case *cmn.ContractCallBackInfo:
				cb = r
			case error:
				cb = &cmn.ContractCallBackInfo{
					CodeResponse: uint32(response.CodeContractPanic()),
					Log:          r.Error(),
				}
			default:
				cb = &cmn.ContractCallBackInfo{
					CodeResponse: uint32(response.CodeContractPanic()),
					Log:          "unknown panic interface...",
				}
			}
		}
	}()

	switch contractName {
	case "isc":
		fmt.Println(contractEnv)
		return isc.CreateNewContract(contractEnv)
	case "sdeam":
		return &cmn.ContractCallBackInfo{
			CodeResponse: uint32(response.CodeTODO()),
		} // sdeam.RegistedMethod(byteJson)
	default:
		return &cmn.ContractCallBackInfo{
			CodeResponse: uint32(response.CodeInvalidTxType()),
			Log:          "unknown contractName",
		}
	}
}

func (nsb *NSBApplication) systemCall(
	contractName string,
	env *cmn.TransactionHeader,
	accInfo *AccountInfo,
	funcName string,
	args []byte,
) (cb *types.ResponseDeliverTx) {
	defer func() {
		if r := recover(); r != nil {
			switch r := r.(type) {
			case string:
				cb = &types.ResponseDeliverTx{
					Code: uint32(response.CodeContractPanic()),
					Log:  r,
				}
			case *types.ResponseDeliverTx:
				cb = r
			case error:
				cb = &types.ResponseDeliverTx{
					Code: uint32(response.CodeContractPanic()),
					Log:  r.Error(),
				}
			default:
				cb = &types.ResponseDeliverTx{
					Code: uint32(response.CodeContractPanic()),
					Log:  "unknown panic interface...",
				}
			}
		}
	}()

	switch contractName {
	case "system.action":
		return nsb.ActionRigisteredMethod(env, accInfo, funcName, args)
	case "system.merkleproof":
		return &types.ResponseDeliverTx{
			Code: uint32(response.CodeTODO()),
		} // sdeam.RegistedMethod(byteJson)
	case "system.token":
		return &types.ResponseDeliverTx{
			Code: uint32(response.CodeTODO()),
		} // sdeam.RegistedMethod(byteJson)
	default:
		return &types.ResponseDeliverTx{
			Code: uint32(response.CodeInvalidTxType()),
			Log:  "unknown contractName",
		}
	}
}
