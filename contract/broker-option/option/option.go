package option

import (
	"fmt"
	"encoding/hex"
	"encoding/json"
	"github.com/Myriad-Dreamin/NSB/math"
	. "github.com/Myriad-Dreamin/NSB/common/contract_response"
	cmn "github.com/Myriad-Dreamin/NSB/common"
)


var (
	env *cmn.ContractEnvironment
)


type ValidBuyer struct {
	Valid    bool `json:"valid"`
	Executed bool `json:"executed"`
}

type RequestCallISC struct {
	FuncName string `json:"function_name"`
	Args     []byte `json:"args"`
}

func MustUnmarshal(data []byte, load interface{}) {
	err := json.Unmarshal(data, &load)
	if err != nil {
		panic(DecodeJsonError(err))
	}
}



func RigisteredMethod(contractEnvironment *cmn.ContractEnvironment) *cmn.ContractCallBackInfo {
	var req RequestCallISC
	MustUnmarshal(contractEnvironment.Data, &req)
	env = contractEnvironment
	switch req.FuncName {
	case "updateStake":
		return updateStake(req.Args)
	case "stakeFund":
		return stakeFund(req.Args)
	case "buyOption":
		return buyOption(req.Args)
	default:
		return InvalidFunctionType(req.FuncName)
	}
}

// func (nsb *NSBApplication) activeISC(byteJson []byte) (types.ResponseDeliverTx) {
// 	return types.ResponseDeliverTx{
// 		Code: uint32(CodeOK),
// 	}
// }

type ArgsCreateNewContract struct {
	Owner       []byte        `json:"1"`
	StrikePrice *math.Uint256 `json:"2"`
}
// // 0x637265617465495343197b226973635f6f776e657273223a5b22456a525765413d3d222c22456a5257654a6f3d225d2c2272657175697265645f66756e6473223a5b302c305d2c22566573536967223a22497a4d3d222c225472616e73616374696f6e496e74656e7473223a5b7b2266726f6d223a22456a525765413d3d222c22746f223a22456a5257654a6f3d222c22736571223a302c22616d74223a302c226d657461223a2249673d3d227d5d7d
// 2ecddf60bb43e12eb402949337a4a0795480f1409e76b7f9cf52ef783532da0a

func CreateNewContract(contractEnvironment *cmn.ContractEnvironment) (*cmn.ContractCallBackInfo) {
	var args ArgsCreateNewContract
	MustUnmarshal(contractEnvironment.Data, &args)

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


type ArgsUpdateStake struct {
	Value *math.Uint256 `json:"1"`
}

func updateStake(bytesArgs []byte) (*cmn.ContractCallBackInfo) {
	var args ArgsUpdateStake
	MustUnmarshal(bytesArgs, &args)

	env.Storage.SetBytes("strikePrice", args.Value.Bytes())

	return ExecOK()
}

func stakeFund(bytesArgs []byte) (*cmn.ContractCallBackInfo) {
	
	remainingFund := math.NewUint256FromBytes(env.Storage.GetBytes("remainingFund"))
	checkErr := remainingFund.Add(env.Value)
	if checkErr {
		return ExecOK()
	} else {
		return OverFlowError("remainingFund overflow")
	}
}


type ArgsBuyOption struct {
	Proposal *math.Uint256 `json:"1"`
}


func buyOption(bytesArgs []byte) (*cmn.ContractCallBackInfo) {
	var args ArgsBuyOption
	MustUnmarshal(bytesArgs, &args)

	var price = optionPrice(args.Proposal)
	AssertTrue(env.Value.Comp(price) >= 0, "sending value must bigger than price")
	env.Storage.NewJsonBytesMap("optionBuyers").Set(env.From, ValidBuyer{true, false})

	return ExecOK()
}

func optionPrice(proposal *math.Uint256) *math.Uint256 {
	var factor, price = math.NewUint256FromString("5", 10), math.NewUint256FromBytes(env.Storage.GetBytes("strikePrice"))
	if proposal.Comp(price) >= 0 {
		absValue, checkErr := math.SubUint256(proposal, price)
		checkErr = (checkErr || factor.Mul(absValue))
		AssertFalse(checkErr, ArithmeticError("when calculating option price"))
	} else {
		absValue, checkErr := math.SubUint256(price, proposal)
		checkErr = (checkErr || factor.Mul(absValue))
		AssertFalse(checkErr, ArithmeticError("when calculating option price"))
	}
	return factor
}
