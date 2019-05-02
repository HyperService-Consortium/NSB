package option

import (
	"fmt"
	"bytes"
	"encoding/hex"
	"encoding/json"
	"github.com/Myriad-Dreamin/NSB/math"
	cmn "github.com/Myriad-Dreamin/NSB/common"
	"github.com/Myriad-Dreamin/NSB/contract/isc/transaction"
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
	err := json.Unmarshal(env.Data, &req)
	if err != nil {
		panic(DecodeJsonError(err))
	}
}


func RigisteredMethod(env *cmn.ContractEnvironment) *cmn.ContractCallBackInfo {
	var req RequestCallISC
	MustUnmarshal(env.Data, &req)
	switch req.FuncName {
	case "updateStake":
		return updateStake(req.Args)
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

func CreateNewContract(env *cmn.ContractEnvironment) (*cmn.ContractCallBackInfo) {
	var args ArgsCreateNewContract
	MustUnmarshal(env.Data, &args)

	env.Storage.SetBytes("remainingFund", env.Value.Bytes())
	env.Storage.SetBytes("strikePrice", args.StrikePrice.Bytes())

	if len(args.Owner) {
		env.Storage.SetBytes("owner", env.From)
	} else {
		env.Storage.SetBytes("owner", args.Owner)
	}


	return &cmn.ContractCallBackInfo{
		CodeResponse: uint32(CodeOK),
		Info: fmt.Sprintf("create success , this contract is deploy at %v", hex.EncodeToString(env.ContractAddress)),
	}
}


type ArgsUpdateStake struct {
	Value *math.Uint256 `json:"1"`
}

func updateStake(env *cmn.ContractEnvironment) (*cmn.ContractCallBackInfo) {
	var args ArgsUpdateStake
	MustUnmarshal(env.Data, &args)

	env.Storage.SetBytes("strikePrice", args.Value.Bytes())

	return ExecOK()
}

func stakeFund(env *cmn.ContractEnvironment) (*cmn.ContractCallBackInfo) {
	var args ArgsStakeFund
	MustUnmarshal(env.Data, &args)

	remainingFund := math.NewUint256FromBytes(env.Storage.GetBytes("remainingFund"))
	checkErr := remainingFund.Add(env.Value)
	if checkErr {
		return ExecOK()
	} else {
		return OverFlowError("remainingFund overflow")
	}
}


type ArgsBuyOption struct {
	Value *math.Uint256 `json:"1"`
}


func buyOption(env *cmn.ContractEnvironment) (*cmn.ContractCallBackInfo) {
	var args ArgsBuyOption
	MustUnmarshal(env.Data, &args)

	var price = optionPrice(proposal)
	AssertTrue(env.Value.Comp(price) >= 0, "sending value must bigger than price")
	env.Storage.NewJsonBytesMap("optionBuyers").Set(env.From, ValidBuyer{true, false})

	return ExecOK()
}

func optionPrice(proposal *math.Uint256) *math.Uint256 {
	var factor, price = math.NewUint256FromString("5"), math.NewUint256FromBytes(env.Storage.GetBytes("strikePrice"))
	if proposal.Comp(price) >= 0 {
		absValue, checkErr := SubUint256(proposal, strikePrice)
		checkErr = (checkErr || factor.Mul(absValue))
	} else {
		absValue, checkErr := SubUint256(strikePrice, proposal)
		checkErr = (checkErr || factor.Mul(absValue))
	}
	AssertFalse(checkErr, ArithmaticError("when calculating option price"))
	return factor
}


