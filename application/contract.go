package nsb

import (
	"github.com/HyperServiceOne/NSB/application/response"
	cmn "github.com/HyperServiceOne/NSB/common"
	"github.com/tendermint/tendermint/abci/types"

	// sdeam "github.com/HyperServiceOne/NSB/contract/sdeam"
	opt "github.com/HyperServiceOne/NSB/contract/broker-option/option"
	dlg "github.com/HyperServiceOne/NSB/contract/delegate"
	isc "github.com/HyperServiceOne/NSB/contract/isc"
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
	case "option":
		return opt.RigisteredMethod(contractEnv)
	case "delegate":
		return dlg.RigisteredMethod(contractEnv)
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
		return isc.CreateNewContract(contractEnv)
	case "sdeam":
		return &cmn.ContractCallBackInfo{
			CodeResponse: uint32(response.CodeTODO()),
		} // sdeam.RegistedMethod(byteJson)
	case "option":
		return opt.CreateNewContract(contractEnv)
	case "delegate":
		return dlg.CreateNewContract(contractEnv)
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
	frInfo *AccountInfo,
	toInfo *AccountInfo,
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
		return nsb.ActionRigisteredMethod(env, frInfo, toInfo, funcName, args)
	case "system.merkleproof":
		return nsb.MerkleProofRigisteredMethod(env, frInfo, toInfo, funcName, args)
	case "system.token":
		return nsb.TokenRigisteredMethod(env, frInfo, toInfo, funcName, args)
	default:
		return &types.ResponseDeliverTx{
			Code: uint32(response.CodeInvalidTxType()),
			Log:  "unknown contractName",
		}
	}
}
