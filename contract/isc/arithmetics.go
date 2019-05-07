package isc

import (
	"encoding/json"
	"github.com/Myriad-Dreamin/NSB/math"
	cmn "github.com/Myriad-Dreamin/NSB/common"
	. "github.com/Myriad-Dreamin/NSB/common/contract_response"
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
	overflowCheck := args.A.Add(args.B)
	if overflowCheck {
		return OverFlowError("Arithmetic Overflow occurred while executing A+B")
	}
	return &cmn.ContractCallBackInfo{
		CodeResponse: CodeOK,
		Info: args.A.String(),
	}
}