package nsb

import (
	"github.com/tendermint/tendermint/abci/types"
	sdeam "github.com/Myriad-Dreamin/NSB/contract/sdeam"
)

func (nsb *NSBApplication) foundContracts(contractName string, byteJson []byte) {
	switch contractName {
	case "sdeam":
		return sdeam.RegistedMethod(byteJson)
	default:
		return types.ResponseDeliverTx{Code: uint32(CodeInvalidTxType)}
	}
}