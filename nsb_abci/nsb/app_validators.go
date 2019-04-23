package nsb

import (
	"fmt"
	"bytes"
	"github.com/tendermint/tendermint/abci/types"
)

func (nsb *NSBApplication) Validators() (validators []types.ValidatorUpdate) {
	itr := nsb.state.db.Iterator(nil, nil)
	for ; itr.Valid(); itr.Next() {
		if isValidatorTx(itr.Key()) {
			validator := new(types.ValidatorUpdate)
			err := types.ReadMessage(bytes.NewBuffer(itr.Value()), validator)
			if err != nil {
				panic(err)
			}
			validators = append(validators, *validator)
		}
	}
	return
}

func MakeValSetChangeTx(pubkey types.PubKey, power int64) []byte {
	return []byte(fmt.Sprintf("val:%X/%d", pubkey.Data, power))
}

func isValidatorTx(tx []byte) bool {
	return true// strings.HasPrefix(string(tx), ValidatorSetChangePrefix)
}

func (nsb *NSBApplication) updateValidator(v types.ValidatorUpdate) types.ResponseDeliverTx {
	return types.ResponseDeliverTx{Code: CodeOK}
}