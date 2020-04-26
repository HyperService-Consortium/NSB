package nsbcli

import (
	"encoding/json"
	transactiontype "github.com/HyperService-Consortium/NSB/application/transaction-type"
	"github.com/HyperService-Consortium/NSB/lib/nsb-client/nsb-message"

	ISC "github.com/HyperService-Consortium/NSB/contract/isc"
	"github.com/HyperService-Consortium/NSB/grpc/nsbrpc"
	uip "github.com/HyperService-Consortium/go-uip/uip"
)

func (nsb *NSBClient) UserAck(
	user uip.Signer, contractAddress []byte,
	address, signature []byte,
) (*nsb_message.DeliverTx, error) {
	// fmt.Println(string(buf.Bytes()))
	fap, err := nsb.userAck(address, signature)
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
	return &ret.DeliverTx, nil
}

func (nsb *NSBClient) userAck(
	address, signature []byte,
) (*nsbrpc.FAPair, error) {

	var args ISC.ArgsUserAck
	args.Address = address
	args.Signature = signature
	b, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	var fap = new(nsbrpc.FAPair)
	fap.FuncName = "UserAck"
	fap.Args = b
	// fmt.Println(PretiStruct(args), b)
	return fap, err
}
