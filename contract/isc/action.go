package isc


import (
	"github.com/tendermint/tendermint/abci/types"
)


func addAction() (types.ResponseDeliverTx) {
	return types.ResponseDeliverTx{
		Code: uint32(CodeOK),
	}
}

func getAction() (types.ResponseDeliverTx) {
	return types.ResponseDeliverTx{
		Code: uint32(CodeOK),
	}
}