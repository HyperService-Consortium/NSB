package math

import (
	"github.com/HyperService-Consortium/go-uip/lib/math"
	"math/big"
)

type Uint128 = math.Uint128

func NewUint128FromUint128(data *Uint128) *Uint128 {
	return math.NewUint128FromUint128(data)
}

func NewUint128FromBytes(data []byte) *Uint128 {
	return math.NewUint128FromBytes(data)
}

func NewUint128FromString(data string, base int) *Uint128 {
	return math.NewUint128FromString(data, base)
}

func NewUint128FromHexString(data string) *Uint128 {
	return math.NewUint128FromHexString(data)
}

func NewUint128FromBigInt(data *big.Int) *Uint128 {
	return math.NewUint128FromBigInt(data)
}
