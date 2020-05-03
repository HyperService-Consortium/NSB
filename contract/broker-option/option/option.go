package option

import (
	"encoding/json"
	cmn "github.com/HyperService-Consortium/NSB/common"
	. "github.com/HyperService-Consortium/NSB/common/contract_response"
	"github.com/HyperService-Consortium/NSB/math"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
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

type NewContractReply struct {
	Address []byte `json:"address"`
	Price   []byte `json:"price"`
}

const errDescNotEnoughValue = "Remaining fund should be more than min stake"

func (option *Option) NewContract(owner []byte, strikePrice *math.Uint256) *cmn.ContractCallBackInfo {

	if option.env.Value == nil {
		option.env.Value = math.NewUint256FromHexString("0")
	}

	var minStake = math.NewUint256FromBytes([]byte{10})
	AssertTrue(option.env.Value.Comp(minStake) >= 0, errDescNotEnoughValue)
	option.setMinStake(minStake)
	option.SetRemainingFund(option.env.Value)

	option.SetStrikePrice(strikePrice)

	if len(owner) == 0 {
		option.SetOwner(option.env.From)
	} else {
		option.SetOwner(owner)
	}

	return &cmn.ContractCallBackInfo{
		CodeResponse: uint32(codeOK),
		Data: sugar.HandlerError(json.Marshal(NewContractReply{
			Address: option.env.ContractAddress,
			Price:   strikePrice.Bytes(),
		})).([]byte),
		Value: option.env.Value,
	}
}

func (option *Option) UpdateStake(value *math.Uint256) *cmn.ContractCallBackInfo {
	option.SetStrikePrice(value)
	return &cmn.ContractCallBackInfo{
		CodeResponse: uint32(codeOK),
		Data:         value.Bytes(),
	}
}

func (option *Option) StakeFund() *cmn.ContractCallBackInfo {
	remainingFund := option.GetRemainingFund()
	checkErr := remainingFund.Add(option.env.Value)

	if checkErr {
		return OverFlowError("remainingFund overflow")
	} else {
		option.SetRemainingFund(remainingFund)
		return ExecOK(option.env.Value)
	}
}

func (option *Option) BuyOption(proposal *math.Uint256) *cmn.ContractCallBackInfo {
	var price = option.optionPrice(proposal)

	AssertTrue(option.env.Value.Comp(price) >= 0, "sending value must be bigger than or equal to price")
	option.GetOptionBuyers().
		Set(option.env.From, ValidBuyer{true, false})

	return ExecOK(nil)
}

func (option *Option) optionPrice(proposal *math.Uint256) *math.Uint256 {
	var factor = math.NewUint256FromString("5", 10)
	var price = option.GetStrikePrice()
	if proposal.Comp(price) >= 0 {
		absValue, checkErr := math.SubUint256(proposal, price)
		checkErr = checkErr || factor.Mul(absValue)
		AssertFalse(checkErr, ArithmeticError("when calculating option price"))
	} else {
		absValue, checkErr := math.SubUint256(price, proposal)
		checkErr = checkErr || factor.Mul(absValue)
		AssertFalse(checkErr, ArithmeticError("when calculating option price"))
	}
	return factor
}

const ValueLessThanProfit = "sending value must be bigger than or equal to profit"

func (option *Option) CashSettle(genuinePrice *math.Uint256) *cmn.ContractCallBackInfo {

	validBuyers := option.GetOptionBuyers()
	validBuyer := validBuyers.Get(option.env.From)
	AssertTrue(validBuyer.Valid, "Valid should be true")
	AssertFalse(validBuyer.Executed, "Executed should be false")

	var strikePrice = option.GetStrikePrice()
	if genuinePrice.Comp(strikePrice) > 0 {
		profit, underflow := math.SubUint256(genuinePrice, strikePrice)
		AssertFalse(underflow, ArithmeticError("underflow"))
		AssertTrue(option.env.Value.Comp(profit) >= 0,
			ValueLessThanProfit)

		validBuyer.Executed = true
		validBuyers.Set(option.env.From, validBuyer)
		return ExecOK(math.NewUint256FromUint256(profit))
	}
	return ExecOK(nil)
}
