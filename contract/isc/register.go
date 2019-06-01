package isc

import (
	"github.com/HyperServiceOne/NSB/contract/isc/transaction"
	"encoding/json"
	. "github.com/HyperServiceOne/NSB/common/contract_response"
	cmn "github.com/HyperServiceOne/NSB/common"
)

func MustUnmarshal(data []byte, load interface{}) {
	err := json.Unmarshal(data, load)
	if err != nil {
		panic(DecodeJsonError(err))
	}
}

type ISC struct {
	env *cmn.ContractEnvironment
}

type ArgsCreateNewContract struct {
	IscOwners          [][]byte                        `json:"isc_owners"`
	Funds              []uint32                        `json:"required_funds"`
	VesSig             []byte                          `json:"ves_signature"`
	TransactionIntents []*transaction.TransactionIntent `json:"transaction_intents"`
}

type ArgsUpdateTxInfo struct {
	Tid uint64 `json:"tid"`
	TransactionIntent *transaction.TransactionIntent `json:"transaction_intent"`
}

type ArgsUpdateTxFr struct {
	Tid uint64 `json:"tid"`
	Fr []byte `json:"from"`
}

type ArgsFreezeInfo struct {
	Tid uint64 `json:"tid"`
}

type ArgsUserAck struct {
	Signature []byte `json:"signature"`
}


func CreateNewContract(contractEnvironment *cmn.ContractEnvironment) (*cmn.ContractCallBackInfo) {
	var args ArgsCreateNewContract
	MustUnmarshal(contractEnvironment.Args, &args)

	var iscc = &ISC{env: contractEnvironment}
	return iscc.NewContract(args.IscOwners, args.Funds, args.VesSig, args.TransactionIntents)
}

func RigisteredMethod(env *cmn.ContractEnvironment) *cmn.ContractCallBackInfo {
	var iscc = &ISC{env: env}
	switch env.FuncName {
	case "UpdateTxInfo":
		var args ArgsUpdateTxInfo
		MustUnmarshal(env.Args, &args)
		return iscc.UpdateTxInfo(args.Tid, args.TransactionIntent)
	case "UpdateTxFr":
		var args ArgsUpdateTxFr
		MustUnmarshal(env.Args, &args)
		return iscc.UpdateTxFr(args.Tid, args.Fr)
	case "FreezeInfo":
		var args ArgsFreezeInfo
		MustUnmarshal(env.Args, &args)
		return iscc.FreezeInfo(args.Tid)
	case "UserAck":
		var args ArgsUserAck
		MustUnmarshal(env.Args, &args)
		return iscc.UserAck(args.Signature)
	default:
		return InvalidFunctionType(env.FuncName)
	}
}
