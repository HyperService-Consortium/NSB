package nsb

import (
	"fmt"
	cmn "github.com/Myriad-Dreamin/NSB/common"
	"github.com/Myriad-Dreamin/NSB/application/response"
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
				cb = &cmn.ContractCallBackInfo {
					CodeResponse: uint32(response.CodeContractPanic()),
					Log: r,
				}
			case *cmn.ContractCallBackInfo:
				cb =  r
			case error:
				cb = &cmn.ContractCallBackInfo {
					CodeResponse: uint32(response.CodeContractPanic()),
					Log: r.Error(),
				}
			default:
				cb = &cmn.ContractCallBackInfo {
					CodeResponse: uint32(response.CodeContractPanic()),
					Log: "unknown panic interface...",
				}
			}
		}
	}()
	switch contractName {
	case "isc":
		return isc.RigisteredMethod(contractEnv)
	case "sdeam":
		return &cmn.ContractCallBackInfo {
			CodeResponse: uint32(response.CodeTODO()),
		}// sdeam.RegistedMethod(byteJson)
	default:
		return &cmn.ContractCallBackInfo {
			CodeResponse: uint32(response.CodeInvalidTxType()),
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
				cb = &cmn.ContractCallBackInfo {
					CodeResponse: uint32(response.CodeContractPanic()),
					Log: r,
				}
			case *cmn.ContractCallBackInfo:
				cb =  r
			case error:
				cb = &cmn.ContractCallBackInfo {
					CodeResponse: uint32(response.CodeContractPanic()),
					Log: r.Error(),
				}
			default:
				cb = &cmn.ContractCallBackInfo {
					CodeResponse: uint32(response.CodeContractPanic()),
					Log: "unknown panic interface...",
				}
			}
		}
	}()
	switch contractName {
	case "isc":
		fmt.Println(contractEnv)
		return isc.CreateNewContract(contractEnv)
	case "sdeam":
		return &cmn.ContractCallBackInfo {
			CodeResponse: uint32(response.CodeTODO()),
		}// sdeam.RegistedMethod(byteJson)
	default:
		return &cmn.ContractCallBackInfo {
			CodeResponse: uint32(response.CodeInvalidTxType()),
		}
	}
}
