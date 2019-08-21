package common

import (
	"github.com/HyperService-Consortium/NSB/math"
)

type TransactionHeader struct {
	From            []byte        `json:"from"`
	ContractAddress []byte        `json:"to"`
	Data            []byte        `json:"data"`
	Value           *math.Uint256 `json:"value"`
	Nonce           *math.Uint256 `json:"nonce"`
	Signature       []byte        `json:"signature"`
}
