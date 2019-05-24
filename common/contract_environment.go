package common

import (
	"github.com/HyperServiceOne/NSB/localstorage"
	"github.com/HyperServiceOne/NSB/math"
)

type ContractEnvironment struct {
	Storage         *localstorage.LocalStorage
	From            []byte
	ContractAddress []byte
	FuncName        string
	Args            []byte
	Value           *math.Uint256
}
