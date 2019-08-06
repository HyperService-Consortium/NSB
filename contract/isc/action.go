package isc

import (
	"github.com/tendermint/tendermint/abci/types"
)

func addAction() types.ResponseDeliverTx {
	return types.ResponseDeliverTx{
		Code: 0,
	}
}

func getAction() types.ResponseDeliverTx {
	return types.ResponseDeliverTx{
		Code: 0,
	}
}
