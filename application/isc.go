package nsb


import (
	"github.com/tendermint/tendermint/abci/types"
	"encoding/json"
)


type RequestCreateISC struct {
	IscOwners          [][]byte                  `json:"isc_owners"`
	Funds              []uint32                  `json:"required_funds"`
	VesSig             []byte                    `json:ves_signature`
	TransactionIntents []transaction.Transaction `json: transactionIntents`
}


func (nsb *NSBApplication) createISC(byteJson []byte) (types.ResponseDeliverTx) {
	var req RequestCreateISC
	err := json.Unmashal(byteJson, &req)
	if err != nil {
		return &types.ResponseDeliverTx{
			Code: CodeDecodeJsonError,
		}
	}
	fmt.Print(req)
	return &types.ResponseDeliverTx{
		Code: CodeOK,
	}
}

func (nsb *NSBApplication) getAction(byteJson []byte) (types.ResponseDeliverTx) {
	return &types.ResponseDeliverTx{
		Code: CodeOK,
	}
}