package option

import (
	"fmt"
	"encoding/hex"
	"github.com/Myriad-Dreamin/NSB/math"
	. "github.com/Myriad-Dreamin/NSB/common/contract_response"
	cmn "github.com/Myriad-Dreamin/NSB/common"
)

type ValidBuyer struct {
	Valid    bool `json:"valid"`
	Executed bool `json:"executed"`
}


// func (nsb *NSBApplication) activeISC(byteJson []byte) (types.ResponseDeliverTx) {
// 	return types.ResponseDeliverTx{
// 		Code: uint32(CodeOK),
// 	}
// }

// // 0x637265617465495343197b226973635f6f776e657273223a5b22456a525765413d3d222c22456a5257654a6f3d225d2c2272657175697265645f66756e6473223a5b302c305d2c22566573536967223a22497a4d3d222c225472616e73616374696f6e496e74656e7473223a5b7b2266726f6d223a22456a525765413d3d222c22746f223a22456a5257654a6f3d222c22736571223a302c22616d74223a302c226d657461223a2249673d3d227d5d7d
// 2ecddf60bb43e12eb402949337a4a0795480f1409e76b7f9cf52ef783532da0a


func (option *Option) NewContract(owner []byte, strikePrice *math.Uint256) (*cmn.ContractCallBackInfo) {
	option.env.Storage.SetBytes("remainingFund", option.env.Value.Bytes())
	option.env.Storage.SetBytes("strikePrice", strikePrice.Bytes())

	if len(owner) == 0 {
		option.env.Storage.SetBytes("owner", option.env.From)
	} else {
		option.env.Storage.SetBytes("owner", owner)
	}


	return &cmn.ContractCallBackInfo{
		CodeResponse: uint32(codeOK),
		Info: fmt.Sprintf(
			"create success , this contract is deploy at %v",
			hex.EncodeToString(option.env.ContractAddress),
		),
	}
}


func (option *Option) UpdateStake(value *math.Uint256) (*cmn.ContractCallBackInfo) {
	option.env.Storage.SetBytes("strikePrice", value.Bytes())
	return ExecOK(nil)
}

func (option *Option) StakeFund() (*cmn.ContractCallBackInfo) {
	remainingFund := math.NewUint256FromBytes(option.env.Storage.GetBytes("remainingFund"))
	checkErr := remainingFund.Add(option.env.Value)

	if checkErr {
		return OverFlowError("remainingFund overflow")
	} else {
		option.env.Storage.SetBytes("remainingFund", remainingFund.Bytes())
		return ExecOK(option.env.Value)
	}
}


func (option *Option) BuyOption(proposal *math.Uint256) (*cmn.ContractCallBackInfo) {
	var price = option.optionPrice(proposal)

	AssertTrue(option.env.Value.Comp(price) >= 0, "sending value must bigger than price")
	option.env.Storage.NewJsonBytesMap("optionBuyers").Set(option.env.From, ValidBuyer{true, false})

	return ExecOK(nil)
}


func (option *Option) optionPrice(proposal *math.Uint256) *math.Uint256 {
	var factor = math.NewUint256FromString("5", 10)
	var price = math.NewUint256FromBytes(option.env.Storage.GetBytes("strikePrice"))
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
