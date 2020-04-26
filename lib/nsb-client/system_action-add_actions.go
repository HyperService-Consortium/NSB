package nsbcli

import (
	"encoding/json"
	appl "github.com/HyperService-Consortium/NSB/application"
	transactiontype "github.com/HyperService-Consortium/NSB/application/transaction-type"
	"github.com/HyperService-Consortium/NSB/grpc/nsbrpc"
	uip "github.com/HyperService-Consortium/go-uip/uip"
)

func (nsb *NSBClient) AddActions(
	user uip.Signer, toAddress []byte, predictNumbers int,
) *AddActionsBatcher {
	return &AddActionsBatcher{
		nc:        nsb,
		user:      user,
		toAddress: toAddress,
		argss:     make([]appl.ArgsAddAction, 0, predictNumbers),
	}
}

type AddActionsBatcher struct {
	nc        *NSBClient
	user      uip.Signer
	toAddress []byte
	argss     []appl.ArgsAddAction
}

func (batcher *AddActionsBatcher) Insert(
	iscAddress []byte, tid uint64, aid uint64, stype uint32,
	content []byte, signature []byte,
) *AddActionsBatcher {
	batcher.argss = append(batcher.argss, appl.ArgsAddAction{
		ISCAddress: iscAddress,
		Tid:        tid,
		Aid:        aid,
		Type:       stype,
		Content:    content,
		Signature:  signature,
	})
	return batcher
}

func (batcher *AddActionsBatcher) Commit() ([]byte, error) {

	// fmt.Println(string(buf.Bytes()))
	fap, err := batcher.nc.addActions(batcher.argss)
	if err != nil {
		return nil, err
	}
	txHeader, err := batcher.nc.CreateContractPacket(batcher.user, batcher.toAddress, []byte{0}, fap)
	if err != nil {
		return nil, err
	}
	ret, err := batcher.nc.sendContractTx(transactiontype.SystemCall, txHeader)
	if err != nil {
		return nil, err
	}
	// fmt.Println(PretiStruct(ret), err)
	return ret.DeliverTx.Data, nil
}

func (nsb *NSBClient) addActions(argss []appl.ArgsAddAction) (*nsbrpc.FAPair, error) {
	var args appl.ArgsAddActions
	args.Args = argss

	var fap = new(nsbrpc.FAPair)
	var err error
	fap.FuncName = "system.action@addActions"
	fap.Args, err = json.Marshal(args)
	if err != nil {
		return nil, err
	}
	return fap, err
}
