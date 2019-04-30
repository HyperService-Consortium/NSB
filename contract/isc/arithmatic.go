package isc

import (
	"encoding/json"
	"github.com/Myriad-Dreamin/NSB/math"
)

type ArgsSafeAdd struct {
	A *math.Uint256 `json:"a"`
	B *math.Uint256 `json:"b"`
}

func SafeAdd(JsonParas []byte) *cmn.ContractCallBackInfo {
	var args ArgsSafeAdd
	err := json.Unmarshal(JsonParas, args)
	if err != nil {
		return DecodeJsonError(err)
	}
	// -------------
	overflowCheck := A.Add(B)
	if overflowCheck {
		return overflowError("Arithmetic Overflow occurred while executing A+B")
	}
}