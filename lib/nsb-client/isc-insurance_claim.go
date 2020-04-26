package nsbcli

import (
	"encoding/binary"
	transactiontype "github.com/HyperService-Consortium/NSB/application/transaction-type"
	"github.com/HyperService-Consortium/NSB/grpc/nsbrpc"
	"github.com/HyperService-Consortium/NSB/lib/nsb-client/nsb-message"
	uip "github.com/HyperService-Consortium/go-uip/uip"
)

func (nsb *NSBClient) InsuranceClaim(
	user uip.Signer, contractAddress []byte,
	tid, aid uint64,
) (*nsb_message.DeliverTx, error) {
	// fmt.Println(string(buf.Bytes()))
	fap, err := nsb.insuranceClaim(tid, aid)
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

func (nsb *NSBClient) insuranceClaim(
	tid, aid uint64,
) (*nsbrpc.FAPair, error) {
	var fap = new(nsbrpc.FAPair)
	fap.FuncName = "InsuranceClaim"
	fap.Args = make([]byte, 16)
	binary.BigEndian.PutUint64(fap.Args[0:8], tid)
	binary.BigEndian.PutUint64(fap.Args[8:], aid)
	// fmt.Println(PretiStruct(args), b)
	return fap, nil
}
