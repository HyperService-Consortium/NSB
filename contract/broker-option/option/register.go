package option


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


func MustUnmarshal(data []byte, load interface{}) {
	err := json.Unmarshal(data, &load)
	if err != nil {
		panic(DecodeJsonError(err))
	}
}

func RigisteredMethod(contractEnvironment *cmn.ContractEnvironment) *cmn.ContractCallBackInfo {
	var option = &Option{env: contractEnvironment}
	switch env.FuncName {
	case "UpdateStake":
		var args ArgsUpdateStake
		MustUnmarshal(bytesArgs, &args)
		return option.UpdateStake(args.Value)
	case "StakeFund":
		return option.StakeFund()
	case "BuyOption":
		var args ArgsBuyOption
		MustUnmarshal(bytesArgs, &args)
		return option.BuyOption(args.Proposal)
	default:
		return InvalidFunctionType(env.FuncName)
	}
}


func CreateNewContract(contractEnvironment *cmn.ContractEnvironment) (*cmn.ContractCallBackInfo) {
	var args ArgsCreateNewContract
	MustUnmarshal(contractEnvironment.Args, &args)

	contractEnvironment.Storage.SetBytes("remainingFund", contractEnvironment.Value.Bytes())
	contractEnvironment.Storage.SetBytes("strikePrice", args.StrikePrice.Bytes())

	if len(args.Owner) == 0 {
		contractEnvironment.Storage.SetBytes("owner", contractEnvironment.From)
	} else {
		contractEnvironment.Storage.SetBytes("owner", args.Owner)
	}


	return &cmn.ContractCallBackInfo{
		CodeResponse: uint32(codeOK),
		Info: fmt.Sprintf("create success , this contract is deploy at %v", hex.EncodeToString(env.ContractAddress)),
	}
}
