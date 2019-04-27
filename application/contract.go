package nsb

import (
	"github.com/tendermint/tendermint/abci/types"
	sdeam "github.com/Myriad-Dreamin/NSB/contract/sdeam"
	"github.com/Myriad-Dreamin/NSB/application/response"
)


func (nsb *NSBApplication) execContractFuncs(contractName string, byteJson []byte) *ContractCallBackInfo {
	switch contractName {
	case "sdeam":
		return return ContractCallBackInfo{
			CodeResponse: uint32(response.CodeTODO),
		}// sdeam.RegistedMethod(byteJson)
	default:
		return ContractCallBackInfo{
			CodeResponse: uint32(response.CodeInvalidTxType),
		}
	}
}