package nsbcli

import (
	"encoding/json"
	transactiontype "github.com/HyperService-Consortium/NSB/application/transaction-type"
	ISC "github.com/HyperService-Consortium/NSB/contract/isc"
	"github.com/HyperService-Consortium/NSB/grpc/nsbrpc"
	uip "github.com/HyperService-Consortium/go-uip/uip"
)

func (nsb *NSBClient) FreezeInfo(
	user uip.Signer, contractAddress []byte,
	tid uint64,
) ([]byte, error) {
	// fmt.Println(string(buf.Bytes()))
	fap, err := nsb.freezeInfo(tid)
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

	// fmt.Println(PretiStruct(ret), err)
	return nil, nil
}

func (nsb *NSBClient) freezeInfo(
	tid uint64,
) (*nsbrpc.FAPair, error) {
	var args ISC.ArgsFreezeInfo
	args.Tid = tid
	b, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	var fap = new(nsbrpc.FAPair)
	fap.FuncName = "FreezeInfo"
	fap.Args = b
	// fmt.Println(PretiStruct(args), b)
	return fap, nil
}
