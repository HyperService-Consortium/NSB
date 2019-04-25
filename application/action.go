package nsb


import (
	"github.com/tendermint/tendermint/abci/types"
)


func (nsb *NSBApplication) addAction() (types.ResponseDeliverTx) {
	// itr := nsb.state.db.Iterator(nil, nil)
	// for ; itr.Valid(); itr.Next() {
	// 	if isValidatorTx(itr.Key()) {
	// 		validator := new(types.ValidatorUpdate)
	// 		err := types.ReadMessage(bytes.NewBuffer(itr.Value()), validator)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		validators = append(validators, *validator)
	// 	}
	// }
	return
}

func (nsb *NSBApplication) getAction() (types.ResponseDeliverTx) {
	// itr := nsb.state.db.Iterator(nil, nil)
	// for ; itr.Valid(); itr.Next() {
	// 	if isValidatorTx(itr.Key()) {
	// 		validator := new(types.ValidatorUpdate)
	// 		err := types.ReadMessage(bytes.NewBuffer(itr.Value()), validator)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		validators = append(validators, *validator)
	// 	}
	// }
	return
}