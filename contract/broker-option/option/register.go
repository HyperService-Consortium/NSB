package option

import (
	"encoding/json"
	"github.com/HyperServiceOne/NSB/math"
	. "github.com/HyperServiceOne/NSB/common/contract_response"
	cmn "github.com/HyperServiceOne/NSB/common"
)

type Option struct {
	env *cmn.ContractEnvironment
}


type ArgsCreateNewContract struct {
	Owner       []byte        `json:"owner"`
	StrikePrice *math.Uint256 `json:"strike_price"`
}


type ArgsUpdateStake struct {
	Value *math.Uint256 `json:"1"`
}


type ArgsBuyOption struct {
	Proposal *math.Uint256 `json:"1"`
}


func MustUnmarshal(data []byte, load interface{}) {
	err := json.Unmarshal(data, load)
	if err != nil {
		panic(DecodeJsonError(err))
	}
}

func RigisteredMethod(contractEnvironment *cmn.ContractEnvironment) *cmn.ContractCallBackInfo {
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
	default:
		return InvalidFunctionType(contractEnvironment.FuncName)
	}
}


func CreateNewContract(contractEnvironment *cmn.ContractEnvironment) (*cmn.ContractCallBackInfo) {
	var args ArgsCreateNewContract
	MustUnmarshal(contractEnvironment.Args, &args)

	var option = &Option{env: contractEnvironment}
	return option.NewContract(args.Owner, args.StrikePrice)
}
