package nsbcli

import (
	"encoding/json"
	transactiontype "github.com/HyperService-Consortium/NSB/application/transaction-type"
	"github.com/HyperService-Consortium/NSB/contract/isc"

	"github.com/HyperService-Consortium/NSB/grpc/nsbrpc"
	uip "github.com/HyperService-Consortium/go-uip/uip"
)

func (nsb *NSBClient) ISCGetPC(
	user uip.Signer, contractAddress []byte,
) (uint64, error) {
	// fmt.Println(string(buf.Bytes()))
	fap, err := nsb.iscGetPC()
	if err != nil {
		return 0, err
	}
	txHeader, err := nsb.CreateContractPacket(user, contractAddress, []byte{0}, fap)
	if err != nil {
		return 0, err
	}
	ret, err := nsb.sendContractTx(transactiontype.SendTransaction, txHeader)
	if err != nil {
		return 0, err
	}

	var reply isc.GetPCReply
	err = json.Unmarshal(ret.DeliverTx.Data, &reply)
	if err != nil {
		return 0, err
	}
	return reply.PC, nil
}

func (nsb *NSBClient) iscGetPC() (*nsbrpc.FAPair, error) {
	var fap = new(nsbrpc.FAPair)
	fap.FuncName = "GetPC"
	return fap, nil
}
