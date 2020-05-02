package nsbcli

import (
	"encoding/json"
	transactiontype "github.com/HyperService-Consortium/NSB/application/transaction-type"
	"github.com/HyperService-Consortium/NSB/contract/delegate"
	"github.com/HyperService-Consortium/NSB/grpc/nsbrpc"
	"github.com/HyperService-Consortium/NSB/math"
	"github.com/HyperService-Consortium/go-uip/uip"
)

func (nsb *NSBClient) CreateDelegate(
	user uip.Signer,
	delegates [][]byte, district string, totalVotes *math.Uint256,
) ([]byte, error) {
	fap, err := nsb.createDelegate(delegates, district, totalVotes)
	if err != nil {
		return nil, err
	}
	txHeader, err := nsb.CreateContractPacket(user, nil, []byte{0}, fap)
	if err != nil {
		return nil, err
	}
	ret, err := nsb.sendContractTx(transactiontype.CreateContract, txHeader)

	if err != nil {
		return nil, err
	}

	return ret.DeliverTx.Data, nil
}

func (nsb *NSBClient) createDelegate(
	delegates [][]byte, district string, totalVotes *math.Uint256,
) (*nsbrpc.FAPair, error) {
	var args delegate.ArgsCreateNewContract
	args.Delegates = delegates
	args.District = district
	args.TotalVotes = totalVotes
	b, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	var fap = new(nsbrpc.FAPair)
	fap.FuncName = "delegate"
	fap.Args = b
	return fap, err
}
