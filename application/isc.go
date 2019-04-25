package nsb


import (
	"github.com/Myriad-Dreamin/NSB/application/isc/transaction"
	"github.com/tendermint/tendermint/abci/types"
	"encoding/json"
	"fmt"
)


type RequestCreateISC struct {
	IscOwners          [][]byte                        `json:"isc_owners"`
	Funds              []uint32                        `json:"required_funds"`
	VesSig             []byte                          `json:ves_signature`
	TransactionIntents []transaction.TransactionIntent `json: transactionIntents`
}


func (nsb *NSBApplication) createISC(byteJson []byte) (types.ResponseDeliverTx) {
	var req RequestCreateISC
	err := json.Unmarshal(byteJson, &req)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: uint32(CodeDecodeJsonError),
		}
	}
	fmt.Print(req)
	return types.ResponseDeliverTx{
		Code: uint32(CodeOK),
	}
}

func (nsb *NSBApplication) activeISC(byteJson []byte) (types.ResponseDeliverTx) {
	return types.ResponseDeliverTx{
		Code: uint32(CodeOK),
	}
}