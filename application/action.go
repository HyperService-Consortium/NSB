package nsb


import (
	"github.com/tendermint/tendermint/abci/types"
)


func (nsb *NSBApplication) addAction() (types.ResponseDeliverTx) {
	return types.ResponseDeliverTx{
		Code: uint32(CodeOK),
	}
}

func (nsb *NSBApplication) getAction() (types.ResponseDeliverTx) {
	return types.ResponseDeliverTx{
		Code: uint32(CodeOK),
	}
}