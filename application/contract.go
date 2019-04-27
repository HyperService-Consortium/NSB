package nsb

import (
	// sdeam "github.com/Myriad-Dreamin/NSB/contract/sdeam"
	"fmt"
	"github.com/Myriad-Dreamin/NSB/application/response"
)


func (nsb *NSBApplication) execContractFuncs(contractName []byte, contractEnv ContractEnvironment) *ContractCallBackInfo {
	switch string(contractName) {
	case "isc":
		fmt.Println(contractEnv)
		return &ContractCallBackInfo{
			CodeResponse: uint32(response.CodeTODO),
		}
	case "sdeam":
		return &ContractCallBackInfo{
			CodeResponse: uint32(response.CodeTODO),
		}// sdeam.RegistedMethod(byteJson)
	default:
		return &ContractCallBackInfo{
			CodeResponse: uint32(response.CodeInvalidTxType),
		}
	}
}


func (nsb *NSBApplication) createContracts(contractName []byte, contractEnv ContractEnvironment) *ContractCallBackInfo {
	switch string(contractName) {
	case "isc":
		fmt.Println(contractEnv)
		return &ContractCallBackInfo{
			CodeResponse: uint32(response.CodeTODO),
		}
	case "sdeam":
		return &ContractCallBackInfo{
			CodeResponse: uint32(response.CodeTODO),
		}// sdeam.RegistedMethod(byteJson)
	default:
		return &ContractCallBackInfo{
			CodeResponse: uint32(response.CodeInvalidTxType),
		}
	}
}
