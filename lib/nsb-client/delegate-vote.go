package nsbcli

import (
	transactiontype "github.com/HyperService-Consortium/NSB/application/transaction-type"
	"github.com/HyperService-Consortium/NSB/grpc/nsbrpc"
	nsb_message "github.com/HyperService-Consortium/NSB/lib/nsb-client/nsb-message"
	uip "github.com/HyperService-Consortium/go-uip/uip"
)

func (nsb *NSBClient) DelegateVote(
	user uip.Signer, contractAddress []byte,
) (*nsb_message.ResultInfo, error) {
	fap, err := nsb.delegateVote()
	if err != nil {
		return nil, err
	}
	txHeader, err := nsb.CreateContractPacket(user, contractAddress, []byte{0}, fap)
	if err != nil {
		return nil, err
	}
	res, err := nsb.sendContractTx(transactiontype.SendTransaction, txHeader)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (nsb *NSBClient) delegateVote() (*nsbrpc.FAPair, error) {
	var fap = new(nsbrpc.FAPair)
	fap.FuncName = "Vote"
	fap.Args = nil
	return fap, nil
}
