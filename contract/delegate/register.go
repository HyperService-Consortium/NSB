package delegate

import (
	"encoding/json"
	cmn "github.com/HyperServiceOne/NSB/common"
	. "github.com/HyperServiceOne/NSB/common/contract_response"
	"github.com/HyperServiceOne/NSB/math"
)

type ArgsCreateNewContract struct {
	Delegates       [][]byte        `json:"1"`
	District        string          `json:"2"`
	TotalVotes      *math.Uint256   `json:"3"`
}

type ArgsAddDelegate struct {
	NewDelegate	[]byte			`json:"1"`
}

type ArgsRemoveDelegate struct {
	RemovedDelegate	  []byte		`json:"1"`
}

func MustUnmarshal(data []byte, load interface{}) {
	err := json.Unmarshal(data, &load)
	if err != nil {
		panic(DecodeJsonError(err))
	}
}

func RigisteredMethod(contractEnvironment *cmn.ContractEnvironment) *cmn.ContractCallBackInfo {
	var delegate = &Delegate{env: contractEnvironment}
	switch contractEnvironment.FuncName {
	case "Vote":
		return delegate.Vote()
	case "ResetVote":
		return delegate.ResetVote()
	default:
		return InvalidFunctionType(contractEnvironment.FuncName)
	}
}


func CreateNewContract(contractEnvironment *cmn.ContractEnvironment) (*cmn.ContractCallBackInfo) {
	var args ArgsCreateNewContract
	MustUnmarshal(contractEnvironment.Args, &args)

	var delegate = &Delegate{env: contractEnvironment}
	return delegate.NewContract(args.Delegates, args.District, args.TotalVotes)
}

