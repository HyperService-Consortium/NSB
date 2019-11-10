package common

import (
	"github.com/HyperService-Consortium/NSB/math"
)

type ContractCallBackInfo struct {
	// type responceDeliverTx
	CodeResponse uint32
	Log          string
	Info         string
	Data         []byte
	Value        *math.Uint256
	OutFlag      bool
	Tags         []KVPair
}

func (cb *ContractCallBackInfo) IsErr() bool {
	return cb.CodeResponse != 0
}

func (cb *ContractCallBackInfo) IsOK() bool {
	return cb.CodeResponse == 0
}
