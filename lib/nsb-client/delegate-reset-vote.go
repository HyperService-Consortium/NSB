package nsbcli

import (
	transactiontype "github.com/HyperService-Consortium/NSB/application/transaction-type"
	"github.com/HyperService-Consortium/NSB/grpc/nsbrpc"
	uip "github.com/HyperService-Consortium/go-uip/uip"
)

func (nsb *NSBClient) DelegateResetVote(
	user uip.Signer, contractAddress []byte,
) ([]byte, error) {
	fap, err := nsb.delegateResetVote()
	if err != nil {
		return nil, err
	}
	txHeader, err := nsb.CreateContractPacket(user, contractAddress, []byte{0}, fap)
	if err != nil {
		return nil, err
	}
	_, err = nsb.sendContractTx(transactiontype.SendTransaction, txHeader)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (nsb *NSBClient) delegateResetVote() (*nsbrpc.FAPair, error) {
	var fap = new(nsbrpc.FAPair)
	fap.FuncName = "ResetVote"
	fap.Args = nil
	return fap, nil
}
