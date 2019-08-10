package isc

import (
	"encoding/json"
	"errors"

	cmn "github.com/HyperServiceOne/NSB/common"
	. "github.com/HyperServiceOne/NSB/common/contract_response"
	"github.com/HyperServiceOne/NSB/contract/isc/transaction"
	"github.com/HyperServiceOne/NSB/util"
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
	IscOwners          [][]byte                         `json:"isc_owners"`
	Funds              []uint32                         `json:"required_funds"`
	VesSig             []byte                           `json:"ves_signature"`
	TransactionIntents []*transaction.TransactionIntent `json:"transaction_intents"`
}

type ArgsUpdateTxInfo struct {
	Tid               uint64                         `json:"tid"`
	TransactionIntent *transaction.TransactionIntent `json:"transaction_intent"`
}

type ArgsUpdateTxFr struct {
	Tid uint64 `json:"tid"`
	Fr  []byte `json:"from"`
}

type ArgsFreezeInfo struct {
	Tid uint64 `json:"tid"`
}

type ArgsUserAck struct {
	Address   []byte `json:"address"`
	Signature []byte `json:"signature"`
}

type ArgsInsuranceClaim struct {
	Tid uint64
	Aid uint64
}

func CreateNewContract(contractEnvironment *cmn.ContractEnvironment) *cmn.ContractCallBackInfo {
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
		return iscc.UserAck(args.Address, args.Signature)
	case "InsuranceClaim":
		return iscc.InsuranceClaim(util.BytesToUint64(env.Args[0:8]), util.BytesToUint64(env.Args[8:16]))
	case "SettleContract":
		if env.Args != nil {
			return ExecContractError(errors.New("this function must have no input"))
		}
		return iscc.SettleContract()
	default:
		return InvalidFunctionType(env.FuncName)
	}
}
