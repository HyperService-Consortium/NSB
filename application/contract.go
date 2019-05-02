package nsb

import (
	"fmt"
	cmn "github.com/Myriad-Dreamin/NSB/common"
	"github.com/Myriad-Dreamin/NSB/application/response"
	// sdeam "github.com/Myriad-Dreamin/NSB/contract/sdeam"
	isc "github.com/Myriad-Dreamin/NSB/contract/isc"
)


var recoverFromContractPanic = func() {
	if r := recover(); r != nil {
		switch r := r.(type) {
		case string:
			return &cmn.ContractCallBackInfo {
				CodeResponse: uint32(response.CodeContractPanic()),
				Log: r,
			}
		case *cmn.ContractCallBackInfo:
			return r
		case error:
			return &cmn.ContractCallBackInfo {
				CodeResponse: uint32(response.CodeContractPanic()),
				Log: r.Error(),
			}
		default:
			return &cmn.ContractCallBackInfo {
				CodeResponse: uint32(response.CodeContractPanic()),
				Log: "unknown panic interface...",
			}
		}
	}
}

func (nsb *NSBApplication) execContractFuncs(contractName string, contractEnv *cmn.ContractEnvironment) *cmn.ContractCallBackInfo {
	defer recoverFromContractPanic()
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


func (nsb *NSBApplication) createContracts(contractName string, contractEnv *cmn.ContractEnvironment) *cmn.ContractCallBackInfo {
	defer recoverFromContractPanic()
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
