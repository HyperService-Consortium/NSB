package nsb

import (
	"bytes"
	_ "bytes"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/HyperService-Consortium/NSB/application/response"
	"github.com/tendermint/tendermint/abci/example/code"
	"github.com/tendermint/tendermint/abci/types"
)

func (nsb *NSBApplication) Validators() (validators []types.ValidatorUpdate) {
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

func MakeValSetChangeTx(pubkey types.PubKey, power int64) []byte {
	return []byte(fmt.Sprintf("val:%X/%d", pubkey.Data, power))
}

func (nsb *NSBApplication) execValidatorTx(tx []byte) *types.ResponseDeliverTx {

	// get the pubkey and power
	pubKeyAndPower := strings.Split(string(tx), "/")
	if len(pubKeyAndPower) != 2 {
		return &types.ResponseDeliverTx{
			Code: uint32(response.CodeDecodeBytesError()),
			Log:  fmt.Sprintf("Expected 'pubkey/power'. Got %v", pubKeyAndPower)}
	}
	pubkeyS, powerS := pubKeyAndPower[0], pubKeyAndPower[1]

	// decode the pubkey
	pubkey, err := hex.DecodeString(pubkeyS)
	if err != nil {
		return &types.ResponseDeliverTx{
			Code: uint32(response.CodeDecodeBytesError()),
			Log:  fmt.Sprintf("Pubkey (%s) is invalid hex", pubkeyS)}
	}

	// decode the power
	power, err := strconv.ParseInt(powerS, 10, 64)
	if err != nil {
		return &types.ResponseDeliverTx{
			Code: uint32(response.CodeDecodeBytesError()),
			Log:  fmt.Sprintf("Power (%s) is not an int", powerS)}
	}

	// update
	return nsb.updateValidator(types.Ed25519ValidatorUpdate(pubkey, int64(power)))
}

func (nsb *NSBApplication) updateValidator(v types.ValidatorUpdate) *types.ResponseDeliverTx {
	key := []byte("val:" + string(v.PubKey.Data))
	if v.Power == 0 {
		// remove validator
		if !nsb.state.db.Has(key) {
			return response.UnauthorizedError(key)
		}
		nsb.state.db.Delete(key)
	} else {
		// add or update validator
		value := bytes.NewBuffer(make([]byte, 0))
		if err := types.WriteMessage(&v, value); err != nil {
			return &types.ResponseDeliverTx{
				Code: code.CodeTypeEncodingError,
				Log:  fmt.Sprintf("Error encoding validator: %v", err),
			}
		}
		nsb.state.db.Set(key, value.Bytes())
	}

	// // TODO: we only update the changes array if we successfully updated the tree
	nsb.ValUpdates = append(nsb.ValUpdates, v)
	return &types.ResponseDeliverTx{Code: uint32(response.CodeOK())}
}
