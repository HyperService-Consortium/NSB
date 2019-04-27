package nsb

import (
	// sdeam "github.com/Myriad-Dreamin/NSB/contract/sdeam"
	"github.com/Myriad-Dreamin/NSB/application/response"
)


func (nsb *NSBApplication) execContractFuncs(contractName []byte, byteJson []byte) *ContractCallBackInfo {
	switch string(contractName) {
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