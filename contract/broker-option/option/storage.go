package option

import (
	"github.com/HyperService-Consortium/NSB/localstorage"
	"github.com/HyperService-Consortium/NSB/math"
)

func (option Option) GetRemainingFund() *math.Uint256 {
	return option.env.Storage.GetUint256("remainingFund")
}

func (option Option) SetRemainingFund(fund *math.Uint256) {
	option.env.Storage.SetUint256("remainingFund", fund)
}

func (option Option) GetStrikePrice() *math.Uint256 {
	return option.env.Storage.GetUint256("strikePrice")
}

func (option Option) SetStrikePrice(price *math.Uint256) {
	option.env.Storage.SetUint256("strikePrice", price)
}

func (option Option) GetOwner() []byte {
	return option.env.Storage.GetBytes("owner")
}

func (option Option) SetOwner(owner []byte) {
	option.env.Storage.SetBytes("owner", owner)
}

func (option Option) GetMinStake() *math.Uint256 {
	return option.env.Storage.GetUint256("minStake")
}

func (option Option) setMinStake(minStake *math.Uint256) {
	option.env.Storage.SetUint256("minStake", minStake)
}

type MappingOptionBuyerWrapper struct {
	bm *localstorage.JsonBytesMap
}

func (w MappingOptionBuyerWrapper) Get(key []byte) ValidBuyer {
	var b ValidBuyer
	w.bm.Get(key, &b)
	return b
}

func (w MappingOptionBuyerWrapper) Set(key []byte, value ValidBuyer) {
	w.bm.Set(key, value)
}

func (option Option) GetOptionBuyers() MappingOptionBuyerWrapper {
	return MappingOptionBuyerWrapper{
		option.env.Storage.NewJsonBytesMap("optionBuyers")}
}
