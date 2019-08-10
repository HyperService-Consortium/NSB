package delegate

import (
	"encoding/hex"
	"fmt"
	cmn "github.com/HyperServiceOne/NSB/common"
	. "github.com/HyperServiceOne/NSB/common/contract_response"
	"github.com/HyperServiceOne/NSB/math"
)

var count int = 0

const NIN_PROPOSAL_COUNT = 5;

type Delegate struct {
	env *cmn.ContractEnvironment
}

func (delegate *Delegate) NewContract(_delegates [][]byte, district string, totalVotes *math.Uint256)(*cmn.ContractCallBackInfo){

	delegate.env.Storage.NewBytesMap("isDelegate")
	delegate.env.Storage.NewInt64Map("delegates")
	for i,x:= range _delegates{
		AssertFalse(len(x)==0,"delegate is not null" )
		delegate.env.Storage.NewBytesMap("isDelegate").Set(x,[]byte("true"))
		delegate.env.Storage.NewInt64Map("delegates").Set(int64(i),x)
		count += 1
	}

	delegate.env.Storage.SetBytes("totalVotes", totalVotes.Bytes())
	delegate.env.Storage.SetString("district", district)

	return &cmn.ContractCallBackInfo{
		CodeResponse: uint32(codeOK),
		Info: fmt.Sprintf(
			"create success , this contract is deploy at %v",
			hex.EncodeToString(delegate.env.ContractAddress),
		),
	}
}

func (delegate *Delegate) Vote() (*cmn.ContractCallBackInfo){
	cher:= delegate.env.Storage.NewBytesMap("isDelegate").Get(delegate.env.From)
	AssertTrue(string(cher)=="true", "delegate Only")

	str:= delegate.env.Storage.NewBytesMap("isDelegateVoted").Get(delegate.env.From)
	if string(str)== "false"{
		delegate.env.Storage.NewBytesMap("isDelegateVoted").Set(delegate.env.From,[]byte("true"))
		totalVotes := math.NewUint256FromBytes(delegate.env.Storage.GetBytes("totalVotes"))
		totalVotes.Add(math.NewUint256FromString("1", 10))
		delegate.env.Storage.SetBytes("totalVotes", totalVotes.Bytes())
	}

	return &cmn.ContractCallBackInfo{
		CodeResponse: uint32(codeOK),
		Info: fmt.Sprintf(
			"Total votes is now %v",
			hex.EncodeToString(delegate.env.Storage.GetBytes("totalVotes")),
		),
	}
}

func (delegate *Delegate) ResetVote() (*cmn.ContractCallBackInfo){
	cher:= delegate.env.Storage.NewBytesMap("isDelegate").Get(delegate.env.From)
	AssertTrue(string(cher)=="true", "delegate Only")

	delegate.env.Storage.NewBytesMap("isDelegateVoted").Set(delegate.env.From,[]byte("false"))

	return &cmn.ContractCallBackInfo{
		CodeResponse: uint32(codeOK),
		Info: fmt.Sprintf(
			"delegate %v has been reset to false",
			hex.EncodeToString(delegate.env.From),
		),
	}
}
