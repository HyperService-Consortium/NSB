package common

import (
	"github.com/Myriad-Dreamin/NSB/localstorage"
	"github.com/Myriad-Dreamin/NSB/math"
)


type ContractEnvironment struct {
	Storage *localstorage.LocalStorage
	From []byte
	fromInfo *AccountInfo
	ContractAddress []byte
	toInfo *AccountInfo
	Data []byte
	Value *math.Uint256
}