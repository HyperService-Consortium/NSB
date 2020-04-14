package common

import (
	"encoding/hex"
	"fmt"
	"github.com/HyperService-Consortium/NSB/math"
)

type AccountInfo struct {
	Balance     *math.Uint256 `json:"balance"`
	CodeHash    []byte        `json:"code_hash"`
	StorageRoot []byte        `json:"storage_root"`
	Name        []byte        `json:"name"`
}

func (accInfo *AccountInfo) String() string {
	return fmt.Sprintf(
		"Balance: %v\nodeHash: %v\nStorageRoot: %v, name:%v\n",
		accInfo.Balance.String(),
		hex.EncodeToString(accInfo.CodeHash),
		hex.EncodeToString(accInfo.StorageRoot),
		string(accInfo.Name),
	)
}
