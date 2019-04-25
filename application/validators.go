package nsb

import (
	"fmt"
	"bytes"
	"github.com/tendermint/tendermint/abci/types"
	"strings"
	"strconv"
	"github.com/tendermint/tendermint/abci/example/code"
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

func (nsb *NSBApplication) execValidatorTx(tx []byte) types.ResponseDeliverTx {
	// tx = tx[len(ValidatorSetChangePrefix):]

	//get the pubkey and power
	pubKeyAndPower := strings.Split(string(tx), "/")
	if len(pubKeyAndPower) != 2 {
		return types.ResponseDeliverTx{
			Code: code.CodeTypeEncodingError,
			Log:  fmt.Sprintf("Expected 'pubkey/power'. Got %v", pubKeyAndPower)}
	}
	pubkeyS, powerS := pubKeyAndPower[0], pubKeyAndPower[1]

	// decode the pubkey
	pubkey, err := hex.DecodeString(pubkeyS)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: code.CodeTypeEncodingError,
			Log:  fmt.Sprintf("Pubkey (%s) is invalid hex", pubkeyS)}
	}

	// decode the power
	power, err := strconv.ParseInt(powerS, 10, 64)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: code.CodeTypeEncodingError,
			Log:  fmt.Sprintf("Power (%s) is not an int", powerS)}
	}

	// update
	return nsb.updateValidator(types.Ed25519ValidatorUpdate(pubkey, int64(power)))
}

func (nsb *NSBApplication) updateValidator(v types.ValidatorUpdate) types.ResponseDeliverTx {
	return types.ResponseDeliverTx{Code: uint32(CodeOK)}
}