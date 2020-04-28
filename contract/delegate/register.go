package delegate

import (
	"encoding/json"
	cmn "github.com/HyperService-Consortium/NSB/common"
	"github.com/HyperService-Consortium/NSB/common/contract_response"
	"github.com/HyperService-Consortium/NSB/math"
)

type ArgsCreateNewContract struct {
	Delegates  [][]byte      `json:"1"`
	District   string        `json:"2"`
	TotalVotes *math.Uint256 `json:"3"`
}

type ArgsAddDelegate struct {
	NewDelegate []byte `json:"1"`
}

type ArgsRemoveDelegate struct {
	RemovedDelegate []byte `json:"1"`
}

func MustUnmarshal(data []byte, load interface{}) {
	err := json.Unmarshal(data, load)
	if err != nil {
		panic(response.DecodeJsonError(err))
	}
}

func RegisteredMethod(contractEnvironment *cmn.ContractEnvironment) *cmn.ContractCallBackInfo {
	var delegate = &Delegate{env: contractEnvironment}
	switch contractEnvironment.FuncName {
	case "Vote":
		return delegate.Vote()
	case "ResetVote":
		return delegate.ResetVote()
	default:
		return response.InvalidFunctionType(contractEnvironment.FuncName)
	}
}

func CreateNewContract(contractEnvironment *cmn.ContractEnvironment) *cmn.ContractCallBackInfo {
	var args ArgsCreateNewContract
	MustUnmarshal(contractEnvironment.Args, &args)

	var delegate = &Delegate{env: contractEnvironment}
	return delegate.NewContract(args.Delegates, args.District, args.TotalVotes)
}
