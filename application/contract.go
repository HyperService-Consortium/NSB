package nsb

import (
	"fmt"
	cmn "github.com/Myriad-Dreamin/NSB/common"
	"github.com/Myriad-Dreamin/NSB/application/response"
	// sdeam "github.com/Myriad-Dreamin/NSB/contract/sdeam"
	isc "github.com/Myriad-Dreamin/NSB/contract/isc"
)


func (nsb *NSBApplication) execContractFuncs(contractName string, contractEnv *cmn.ContractEnvironment) *cmn.ContractCallBackInfo {
	switch contractName {
	case "isc":
		fmt.Println(contractEnv)
		return isc.RigisteredMethod(contractEnv)
	case "sdeam":
		return &cmn.ContractCallBackInfo{
			CodeResponse: uint32(response.CodeTODO),
		}// sdeam.RegistedMethod(byteJson)
	default:
		return &cmn.ContractCallBackInfo{
			CodeResponse: uint32(response.CodeInvalidTxType),
		}
	}
}


func (nsb *NSBApplication) createContracts(contractName string, contractEnv *cmn.ContractEnvironment) *cmn.ContractCallBackInfo {
	switch contractName {
	case "isc":
		fmt.Println(contractEnv)
		return isc.CreateNewContract()
	case "sdeam":
		return &cmn.ContractCallBackInfo{
			CodeResponse: uint32(response.CodeTODO),
		}// sdeam.RegistedMethod(byteJson)
	default:
		return &cmn.ContractCallBackInfo{
			CodeResponse: uint32(response.CodeInvalidTxType),
		}
	}
}
