package common

import (
	"github.com/HyperService-Consortium/NSB/math"
)

type AccountInfo struct {
	Balance     *math.Uint256 `json:"balance"`
	CodeHash    []byte        `json:"code_hash"`
	StorageRoot []byte        `json:"storage_root"`
	Name        []byte        `json:"name"`
}
