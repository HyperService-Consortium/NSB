package math

import (
	"math/big"
)

var (
	BigOne = big.NewInt(1)
	P128   = new(big.Int).SetBit(new(big.Int), 128, 1)
	MOD128 = new(big.Int).Sub(P128, BigOne)
	P256   = new(big.Int).SetBit(new(big.Int), 256, 1)
	MOD256 = new(big.Int).Sub(P256, BigOne)
)
