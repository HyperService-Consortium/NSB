package math

import (
	"github.com/HyperService-Consortium/go-uip/lib/math"
	"math/big"
)

type Uint256 = math.Uint256

func NewUint256FromUint256(data *Uint256) *Uint256 {
	return math.NewUint256FromUint256(data)
}

func NewUint256FromBytes(data []byte) *Uint256 {
	return math.NewUint256FromBytes(data)
}

func NewUint256FromString(data string, base int) *Uint256 {
	return math.NewUint256FromString(data, base)
}

func NewUint256FromHexString(data string) *Uint256 {
	return math.NewUint256FromHexString(data)
}

func NewUint256FromBigInt(data *big.Int) *Uint256 {
	return math.NewUint256FromBigInt(data)
}

func AddUint256(x *Uint256, y *Uint256) (ret *Uint256, check bool) {
	ret = NewUint256FromUint256(x)
	check = ret.Add(y)
	return
}

func SubUint256(x *Uint256, y *Uint256) (ret *Uint256, check bool) {
	ret = NewUint256FromUint256(x)
	check = ret.Sub(y)
	return
}
