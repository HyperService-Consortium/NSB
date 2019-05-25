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
	TransactionIntents []*transaction.TransactionIntent `json:"transactionIntents"`
}

func CreateNewContract(contractEnvironment *cmn.ContractEnvironment) (*cmn.ContractCallBackInfo) {
	var args ArgsCreateNewContract
	MustUnmarshal(contractEnvironment.Args, &args)

	var iscc = &ISC{env: contractEnvironment}
	return iscc.NewContract(args.IscOwners, args.Funds, args.VesSig, args.TransactionIntents)
}

func RigisteredMethod(env *cmn.ContractEnvironment) *cmn.ContractCallBackInfo {
	switch env.FuncName {
	case "a+b":
		return SafeAdd(env.Args)
	default:
		return InvalidFunctionType(env.FuncName)
	}
}
