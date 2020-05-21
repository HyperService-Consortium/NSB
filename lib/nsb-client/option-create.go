package nsbcli

import (
	"encoding/json"
	transactiontype "github.com/HyperService-Consortium/NSB/application/transaction-type"
	"github.com/HyperService-Consortium/NSB/contract/broker-option/option"
	"github.com/HyperService-Consortium/NSB/grpc/nsbrpc"
	"github.com/HyperService-Consortium/NSB/math"
	"github.com/HyperService-Consortium/go-uip/uip"
)

func (nsb *NSBClient) CreateOptionContract(
	user uip.Signer, value *math.Uint256,
	owner []byte, strikePrice *math.Uint256,
) ([]byte, error) {
	fap, err := nsb.createOptionContract(owner, strikePrice)
	if err != nil {
		return nil, err
	}
	txHeader, err := nsb.CreateContractPacket(user, nil, value.Bytes(), fap)
	if err != nil {
		return nil, err
	}
	ret, err := nsb.sendContractTx(transactiontype.CreateContract, txHeader)

	if err != nil {
		return nil, err
	}

	return ret.DeliverTx.Data, nil
}

func (nsb *NSBClient) createOptionContract(
	owner []byte, strikePrice *math.Uint256,
) (*nsbrpc.FAPair, error) {
	var args option.ArgsCreateNewContract
	args.Owner = owner
	args.StrikePrice = strikePrice
	b, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	var fap = new(nsbrpc.FAPair)
	fap.FuncName = "option"
	fap.Args = b
	return fap, err
}
