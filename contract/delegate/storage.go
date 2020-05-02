package delegate

import (
	"github.com/HyperService-Consortium/NSB/localstorage"
	"github.com/HyperService-Consortium/go-uip/lib/math"
)

func (delegate Delegate) IsDelegate() *localstorage.BoolMap {
	return delegate.env.Storage.NewBoolMap("isDelegate")
}

func (delegate Delegate) Delegates() *localstorage.BytesArray {
	return delegate.env.Storage.NewBytesArray("Delegates")
}

func (delegate Delegate) IsDelegateVoted() *localstorage.BoolMap {
	return delegate.env.Storage.NewBoolMap("isDelegateVoted")
}

func (delegate Delegate) SetTotalVotes(totalVotes *math.Uint256) {
	delegate.env.Storage.SetUint256("totalVotes", totalVotes)
}

func (delegate Delegate) GetTotalVotes() *math.Uint256 {
	return delegate.env.Storage.GetUint256("totalVotes")
}

func (delegate Delegate) SetDistrict(district string) {
	delegate.env.Storage.SetString("district", district)
}

func (delegate Delegate) GetDistrict() string {
	return delegate.env.Storage.GetString("district")
}
