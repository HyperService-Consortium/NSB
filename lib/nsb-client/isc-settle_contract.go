package nsbcli

import (
	transactiontype "github.com/HyperService-Consortium/NSB/application/transaction-type"
	"github.com/HyperService-Consortium/NSB/lib/nsb-client/nsb-message"

	"github.com/HyperService-Consortium/NSB/grpc/nsbrpc"
	uip "github.com/HyperService-Consortium/go-uip/uip"
)

func (nsb *NSBClient) SettleContract(
	user uip.Signer, contractAddress []byte,
) (*nsb_message.DeliverTx, error) {
	// fmt.Println(string(buf.Bytes()))
	fap, err := nsb.settleContract()
	if err != nil {
		return nil, err
	}
	txHeader, err := nsb.CreateContractPacket(user, contractAddress, []byte{0}, fap)
	if err != nil {
		return nil, err
	}
	ret, err := nsb.sendContractTx(transactiontype.SendTransaction, txHeader)
	if err != nil {
		return nil, err
	}
	// fmt.Println(PretiStruct(ret), err)
	return &ret.DeliverTx, nil
}

func (nsb *NSBClient) settleContract() (*nsbrpc.FAPair, error) {
	var fap = new(nsbrpc.FAPair)
	fap.FuncName = "SettleContract"
	// fmt.Println(PretiStruct(args), b)
	return fap, nil
}
