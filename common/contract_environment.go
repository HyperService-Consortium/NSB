package common

import (
	"github.com/HyperService-Consortium/NSB/localstorage"
	"github.com/HyperService-Consortium/NSB/math"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type ContractEnvironment struct {
	Storage         *localstorage.LocalStorage
	From            []byte
	ContractAddress []byte
	FuncName        string
	Args            []byte
	Value           *math.Uint256
	BN              uip.Storage
}
