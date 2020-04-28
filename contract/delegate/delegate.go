package delegate

import (
	"encoding/hex"
	"fmt"
	cmn "github.com/HyperService-Consortium/NSB/common"
	"github.com/HyperService-Consortium/NSB/common/contract_response"
	"github.com/HyperService-Consortium/NSB/math"
)

const MIN_PROPOSAL_COUNT = 5

type Delegate struct {
	env *cmn.ContractEnvironment
}

func (delegate *Delegate) NewContract(_delegates [][]byte, district string, totalVotes *math.Uint256) *cmn.ContractCallBackInfo {

	isDelegate := delegate.IsDelegate()
	delegates := delegate.Delegates()
	for _, x := range _delegates {
		response.AssertFalse(len(x) == 0, "delegate is not null")
		isDelegate.Set(x, true)
		delegates.Append(x)
	}
	if totalVotes == nil {
		totalVotes = math.NewUint256FromBytes([]byte{0})
	}

	delegate.SetTotalVotes(totalVotes)
	delegate.SetDistrict(district)

	return &cmn.ContractCallBackInfo{
		CodeResponse: uint32(CodeOK),
		Info: fmt.Sprintf(
			"create success , this contract is deploy at %v",
			hex.EncodeToString(delegate.env.ContractAddress),
		),
		Data: delegate.env.ContractAddress,
	}
}

func (delegate *Delegate) Vote() *cmn.ContractCallBackInfo {
	response.AssertTrue(delegate.IsDelegate().Get(delegate.env.From), "delegate Only")
	isDelegateVoted := delegate.IsDelegateVoted()
	if !isDelegateVoted.Get(delegate.env.From) {
		isDelegateVoted.Set(delegate.env.From, true)
		v, overflowed := math.AddUint256(math.NewUint256FromString("1", 10), delegate.GetTotalVotes())
		if overflowed {
			return &cmn.ContractCallBackInfo{
				CodeResponse: uint32(CodeOverflow),
				Log:          "totalVotes overflow",
			}
		}
		delegate.SetTotalVotes(v)
	}

	return &cmn.ContractCallBackInfo{
		CodeResponse: uint32(CodeOK),
		Data:         delegate.GetTotalVotes().Bytes(),
	}
}

func (delegate *Delegate) ResetVote() *cmn.ContractCallBackInfo {
	response.AssertTrue(delegate.IsDelegate().Get(delegate.env.From), "delegate Only")
	isDelegateVoted := delegate.IsDelegateVoted()

	if isDelegateVoted.Get(delegate.env.From) {
		isDelegateVoted.Set(delegate.env.From, false)
		v, overflowed := math.SubUint256(delegate.GetTotalVotes(), math.NewUint256FromString("1", 10))
		if overflowed {
			return &cmn.ContractCallBackInfo{
				CodeResponse: uint32(CodeUnderflow),
				Log:          "totalVotes underflow",
			}
		}
		delegate.SetTotalVotes(v)
	}

	return &cmn.ContractCallBackInfo{
		CodeResponse: uint32(CodeOK),
		Data:         delegate.GetTotalVotes().Bytes(),
	}
}
