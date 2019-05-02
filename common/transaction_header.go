package common


import (
	"github.com/Myriad-Dreamin/NSB/math"
)

type TransactionHeader struct {
	From            []byte        `json:"from"`
	ContractAddress []byte        `json:"to"`
	JsonParas       []byte        `json:"data"`
	Value           *math.Uint256 `json:"value"`
	Nonce           *math.Uint256 `json:"nonce"`
	Signature       []byte        `json:"signature"`
}
