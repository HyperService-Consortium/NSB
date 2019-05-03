package common

import (
	"github.com/Myriad-Dreamin/NSB/math"
)

type ContractCallBackInfo struct {
	// type responceDeliverTx
	CodeResponse uint32
	Log          string
	Info         string
	Value        *math.Uint256
	Tags         []KVPair
}

func (cb *ContractCallBackInfo) IsErr() bool {
	return cb.CodeResponse != 0
}

func (cb *ContractCallBackInfo) IsOK() bool {
	return cb.CodeResponse == 0
}
