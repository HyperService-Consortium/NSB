package nsbcli

import (
	"encoding/json"
	appl "github.com/HyperService-Consortium/NSB/application"
	transactiontype "github.com/HyperService-Consortium/NSB/application/transaction-type"
	"github.com/HyperService-Consortium/NSB/grpc/nsbrpc"
	uip "github.com/HyperService-Consortium/go-uip/uip"
)

func (nsb *NSBClient) GetAction(
	user uip.Signer, toAddress []byte,
	iscAddress []byte, tid uint64, aid uint64,
) ([]byte, error) {
	// fmt.Println(string(buf.Bytes()))
	fap, err := nsb.getAction(iscAddress, tid, aid)
	if err != nil {
		return nil, err
	}
	txHeader, err := nsb.CreateContractPacket(user, toAddress, []byte{0}, fap)
	if err != nil {
		return nil, err
	}
	ret, err := nsb.sendContractTx(transactiontype.SystemCall, txHeader)
	if err != nil {
		return nil, err
	}
	// fmt.Println(PretiStruct(ret), err)
	return ret.DeliverTx.Data, nil
}

func (nsb *NSBClient) getAction(
	iscAddress []byte, tid uint64, aid uint64,
) (*nsbrpc.FAPair, error) {
	var args appl.ArgsGetAction
	args.ISCAddress = iscAddress
	args.Tid = tid
	args.Aid = aid
	b, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	var fap = new(nsbrpc.FAPair)
	fap.FuncName = "system.action@getAction"
	fap.Args = b
	// fmt.Println(PretiStruct(args), b)
	return fap, err
}
