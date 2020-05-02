package option

import (
	"encoding/json"
	cmn "github.com/HyperService-Consortium/NSB/common"
	. "github.com/HyperService-Consortium/NSB/common/contract_response"
	"github.com/HyperService-Consortium/NSB/math"
)

type Option struct {
	env *cmn.ContractEnvironment
}

type ArgsCreateNewContract struct {
	Owner       []byte        `json:"1"`
	StrikePrice *math.Uint256 `json:"2"`
}

type ArgsUpdateStake struct {
	Value *math.Uint256 `json:"1"`
}

type ArgsBuyOption struct {
	Proposal *math.Uint256 `json:"1"`
}

type ArgsCashSettle struct {
	GenuinePrice *math.Uint256 `json:"1"`
}

func MustUnmarshal(data []byte, load interface{}) {
	err := json.Unmarshal(data, load)
	if err != nil {
		panic(DecodeJsonError(err))
	}
}

func RegisteredMethod(contractEnvironment *cmn.ContractEnvironment) *cmn.ContractCallBackInfo {
	var option = &Option{env: contractEnvironment}

	switch contractEnvironment.FuncName {
	case "UpdateStake":
		var args ArgsUpdateStake
		MustUnmarshal(contractEnvironment.Args, &args)
		return option.UpdateStake(args.Value)
	case "StakeFund":
		return option.StakeFund()
	case "BuyOption":
		var args ArgsBuyOption
		MustUnmarshal(contractEnvironment.Args, &args)
		return option.BuyOption(args.Proposal)
	case "CashSettle":
		var args ArgsCashSettle
		MustUnmarshal(contractEnvironment.Args, &args)
		return option.CashSettle(args.GenuinePrice)
	default:
		return InvalidFunctionType(contractEnvironment.FuncName)
	}
}

func CreateNewContract(contractEnvironment *cmn.ContractEnvironment) *cmn.ContractCallBackInfo {
	var args ArgsCreateNewContract
	MustUnmarshal(contractEnvironment.Args, &args)

	var option = &Option{env: contractEnvironment}
	return option.NewContract(args.Owner, args.StrikePrice)
}
