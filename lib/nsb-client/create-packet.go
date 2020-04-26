package nsbcli

import (
	"bytes"
	transactiontype "github.com/HyperService-Consortium/NSB/application/transaction-type"
	"github.com/HyperService-Consortium/NSB/grpc/nsbrpc"
	"github.com/HyperService-Consortium/NSB/lib/nsb-client/nsb-message"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/gogo/protobuf/proto"
	"math/rand"
	"time"
)

func (*NSBClient) Sign(s uip.Signer, txHeader *nsbrpc.TransactionHeader) *nsbrpc.TransactionHeader {
	// bug: buf.Reset()
	buf := bytes.NewBuffer(make([]byte, mxBytes))
	buf.Write(txHeader.Src)
	buf.Write(txHeader.Dst)
	buf.Write(txHeader.Data)
	buf.Write(txHeader.Value)
	buf.Write(txHeader.Nonce)
	x, err := s.Sign(buf.Bytes())
	if err != nil {
		//todo
		panic(err)
	}
	txHeader.Signature = x.Bytes()
	return txHeader
}

func (nsb *NSBClient) CreateContractPacket(
	s uip.Signer, toAddress, value []byte, pair *nsbrpc.FAPair,
) (*nsbrpc.TransactionHeader, error) {
	data, err := proto.Marshal(pair)
	if err != nil {
		return nil, err
	}
	return nsb.CreateNormalPacket(s, toAddress, data, value)
}

func (nsb *NSBClient) CreateUnsignedContractPacket(
	srcAddress, dstAddress, value []byte, pair *nsbrpc.FAPair,
) (*nsbrpc.TransactionHeader, error) {
	data, err := proto.Marshal(pair)
	if err != nil {
		return nil, err
	}
	return nsb.CreateUnsignedNormalPacket(srcAddress, dstAddress, data, value)
}

func (nsb *NSBClient) CreateUnsignedNormalPacket(
	srcAddress, dstAddress, data, value []byte,
) (*nsbrpc.TransactionHeader, error) {
	txHeader := new(nsbrpc.TransactionHeader)
	var err error

	txHeader.Data = data
	txHeader.Dst = dstAddress
	txHeader.Src = srcAddress

	nonce := make([]byte, 32)
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, err
	}
	txHeader.Nonce = nonce
	txHeader.Value = value
	return txHeader, nil
}

func (nsb *NSBClient) CreateNormalPacket(
	s uip.Signer, toAddress, data, value []byte,
) (*nsbrpc.TransactionHeader, error) {
	txHeader, err := nsb.CreateUnsignedNormalPacket(s.GetPublicKey(), toAddress, data, value)
	if err != nil {
		return nil, err
	}
	nsb.Sign(s, txHeader)
	return txHeader, nil
}

func (nsb *NSBClient) sendTx(t transactiontype.Type, txHeader *nsbrpc.TransactionHeader, err error) (*nsb_message.ResultInfo, error) {
	if err != nil {
		return nil, err
	}
	ret, err := nsb.sendContractTx(t, txHeader)
	if err != nil {
		return nil, err
	}
	// fmt.Println(PretiStruct(ret), err)
	return ret, nil
}

func (nsb *NSBClient) systemCall(txHeader *nsbrpc.TransactionHeader, err error) (*nsb_message.ResultInfo, error) {
	return nsb.sendTx(transactiontype.SystemCall, txHeader, err)
}

func (nsb *NSBClient) createContract(txHeader *nsbrpc.TransactionHeader, err error) (*nsb_message.ResultInfo, error) {
	return nsb.sendTx(transactiontype.CreateContract, txHeader, err)
}

func (nsb *NSBClient) sendTransaction(txHeader *nsbrpc.TransactionHeader, err error) (*nsb_message.ResultInfo, error) {
	return nsb.sendTx(transactiontype.SendTransaction, txHeader, err)
}

func (nsb *NSBClient) sign(user uip.Signer, txHeader *nsbrpc.TransactionHeader, err error) (*nsbrpc.TransactionHeader, error) {
	if err != nil {
		return nil, err
	}
	nsb.Sign(user, txHeader)
	return txHeader, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
