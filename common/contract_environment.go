package common

import (
	"github.com/Myriad-Dreamin/NSB/merkmap"
	"github.com/Myriad-Dreamin/NSB/localstorage"
	"github.com/Myriad-Dreamin/NSB/math"
)


type ContractEnvironment struct {
	Storage *localstorage.LocalStorage
	From []byte
	ContractAddress []byte
	Data []byte
	Value *math.Uint256
}

type SystemContractEnvironment struct {
	Storage *merkmap.MerkMap
	From []byte
	To []byte
	Data []byte
	Value *math.Uint256
}
