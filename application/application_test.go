package nsb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"

	transactiontype "github.com/HyperServiceOne/NSB/application/transaction-type"
	isc "github.com/HyperServiceOne/NSB/contract/isc"
	"github.com/HyperServiceOne/NSB/contract/isc/transaction"
	nsbrpc "github.com/HyperServiceOne/NSB/grpc/nsbrpc"
	log "github.com/HyperServiceOne/NSB/log"
	"github.com/HyperServiceOne/NSB/math"
	signaturer "github.com/Myriad-Dreamin/go-uip/signaturer"
	"github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/abci/types"
)

func TestCreateContract(t *testing.T) {

	var pri, nonce, bytesBuf = make([]byte, 64), make([]byte, 32), make([]byte, 65536)
	for idx := 0; idx < 64; idx++ {
		pri[idx] = uint8(idx)
	}
	var signer = signaturer.NewTendermintNSBSigner(pri)

	var err error
	var u, v = []byte{0, 0, 1}, []byte{0, 0, 2}
	var iscOnwers = [][]byte{signer.GetPublicKey(), u, v}
	var funds = []uint32{0, 0, 0}
	var vesSig = []byte{0}
	var transactionIntents = []*transaction.TransactionIntent{
		&transaction.TransactionIntent{
			Fr:   u,
			To:   v,
			Seq:  math.NewUint256FromHexString("10"),
			Amt:  math.NewUint256FromHexString("10"),
			Meta: []byte{0},
		},
	}
	var args = &isc.ArgsCreateNewContract{
		IscOwners:          iscOnwers,
		Funds:              funds,
		VesSig:             vesSig,
		TransactionIntents: transactionIntents,
	}
	var txHeader nsbrpc.TransactionHeader

	var fap FAPair
	fap.FuncName = "isc"
	fap.Args, err = json.Marshal(args)
	if err != nil {
		t.Error(err)
		return
	}

	txHeader.Data, err = proto.Marshal(&fap)
	if err != nil {
		t.Error(err)
		return
	}

	txHeader.Src = signer.GetPublicKey()

	_, err = rand.Read(nonce)
	if err != nil {
		t.Error(err)
		return
	}

	// bytesBuf[0] = transactiontype.CreateContract
	var buf = bytes.NewBuffer(bytesBuf)
	buf.Reset()

	txHeader.Nonce = math.NewUint256FromBytes(nonce).Bytes()
	txHeader.Value = math.NewUint256FromBytes([]byte{0}).Bytes()
	buf.Reset()

	buf.Write(txHeader.Src)
	buf.Write(txHeader.Dst)
	buf.Write(txHeader.Data)
	buf.Write(txHeader.Value)
	buf.Write(txHeader.Nonce)
	txHeader.Signature = signer.Sign(buf.Bytes()).Bytes()
	b, err := proto.Marshal(&txHeader)
	if err != nil {
		t.Error(err)
		return
	}

	bytesBuf[0] = transactiontype.CreateContract

	copy(bytesBuf[1:], b)

	logger, err := log.NewZapColorfulDevelopmentSugarLogger()
	if err != nil {
		t.Error(err)
		return
	}

	nsb, err := NewNSBApplication(logger, "./data/")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(nsb.DeliverTx(types.RequestDeliverTx{
		Tx: bytesBuf[:1+len(b)],
	}))
	var err2 error
	err, err2 = nsb.Stop()
	if err != nil {
		t.Error(err)
		return
	}
	if err2 != nil {
		t.Error(err2)
		return
	}
}

func TestSetBalance(t *testing.T) {

	var pri, nonce, bytesBuf = make([]byte, 64), make([]byte, 32), make([]byte, 65536)
	for idx := 0; idx < 64; idx++ {
		pri[idx] = uint8(idx)
	}
	var signer = signaturer.NewTendermintNSBSigner(pri)

	var err error
	var args = &ArgsTransfer{
		Value: math.NewUint256FromBytes([]byte{1}),
	}
	var txHeader nsbrpc.TransactionHeader

	var fap FAPair
	fap.FuncName = "system.token@setBalance"
	fap.Args, err = json.Marshal(args)
	if err != nil {
		t.Error(err)
		return
	}

	txHeader.Data, err = proto.Marshal(&fap)
	if err != nil {
		t.Error(err)
		return
	}

	txHeader.Src = signer.GetPublicKey()
	txHeader.Dst = make([]byte, 32)

	_, err = rand.Read(nonce)
	if err != nil {
		t.Error(err)
		return
	}

	// bytesBuf[0] = transactiontype.CreateContract
	var buf = bytes.NewBuffer(bytesBuf)
	buf.Reset()

	txHeader.Nonce = math.NewUint256FromBytes(nonce).Bytes()
	txHeader.Value = math.NewUint256FromBytes([]byte{1}).Bytes()
	buf.Reset()

	buf.Write(txHeader.Src)
	buf.Write(txHeader.Dst)
	buf.Write(txHeader.Data)
	buf.Write(txHeader.Value)
	buf.Write(txHeader.Nonce)
	txHeader.Signature = signer.Sign(buf.Bytes()).Bytes()
	b, err := proto.Marshal(&txHeader)
	if err != nil {
		t.Error(err)
		return
	}

	bytesBuf[0] = transactiontype.SystemCall

	copy(bytesBuf[1:], b)

	logger, err := log.NewZapColorfulDevelopmentSugarLogger()
	if err != nil {
		t.Error(err)
		return
	}

	nsb, err := NewNSBApplication(logger, "./data/")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(nsb.DeliverTx(types.RequestDeliverTx{
		Tx: bytesBuf[:1+len(b)],
	}))
	fmt.Println(nsb.Commit())
	var err2 error
	err, err2 = nsb.Stop()
	if err != nil {
		t.Error(err)
		return
	}
	if err2 != nil {
		t.Error(err2)
		return
	}
}

func TestTransfer(t *testing.T) {
	var pri, nonce, bytesBuf = make([]byte, 64), make([]byte, 32), make([]byte, 65536)
	for idx := 0; idx < 64; idx++ {
		pri[idx] = uint8(idx)
	}
	var signer = signaturer.NewTendermintNSBSigner(pri)

	var err error
	var args = &ArgsTransfer{
		Value: math.NewUint256FromBytes([]byte{1}),
	}
	var txHeader nsbrpc.TransactionHeader

	var fap FAPair
	fap.FuncName = "system.token@setBalance"
	fap.Args, err = json.Marshal(args)
	if err != nil {
		t.Error(err)
		return
	}

	txHeader.Data, err = proto.Marshal(&fap)
	if err != nil {
		t.Error(err)
		return
	}

	txHeader.Src = signer.GetPublicKey()
	txHeader.Dst = make([]byte, 32)

	_, err = rand.Read(nonce)
	if err != nil {
		t.Error(err)
		return
	}

	// bytesBuf[0] = transactiontype.CreateContract
	var buf = bytes.NewBuffer(bytesBuf)
	buf.Reset()

	txHeader.Nonce = math.NewUint256FromBytes(nonce).Bytes()
	txHeader.Value = math.NewUint256FromBytes([]byte{1}).Bytes()
	buf.Reset()

	buf.Write(txHeader.Src)
	buf.Write(txHeader.Dst)
	buf.Write(txHeader.Data)
	buf.Write(txHeader.Value)
	buf.Write(txHeader.Nonce)
	txHeader.Signature = signer.Sign(buf.Bytes()).Bytes()
	b, err := proto.Marshal(&txHeader)
	if err != nil {
		t.Error(err)
		return
	}

	bytesBuf[0] = transactiontype.SystemCall

	copy(bytesBuf[1:], b)

	logger, err := log.NewZapColorfulDevelopmentSugarLogger()
	if err != nil {
		t.Error(err)
		return
	}

	nsb, err := NewNSBApplication(logger, "./data/")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(nsb.DeliverTx(types.RequestDeliverTx{
		Tx: bytesBuf[:1+len(b)],
	}))
	fmt.Println(nsb.Commit())

	var args2 = &ArgsTransfer{
		Value: math.NewUint256FromBytes([]byte{1}),
	}

	fap.FuncName = "system.token@transfer"
	fap.Args, err = json.Marshal(args2)
	if err != nil {
		t.Error(err)
		return
	}

	txHeader.Data, err = proto.Marshal(&fap)
	if err != nil {
		t.Error(err)
		return
	}

	txHeader.Src = signer.GetPublicKey()
	txHeader.Dst = make([]byte, 32)

	_, err = rand.Read(nonce)
	if err != nil {
		t.Error(err)
		return
	}

	// bytesBuf[0] = transactiontype.CreateContract
	buf.Reset()

	txHeader.Nonce = math.NewUint256FromBytes(nonce).Bytes()
	txHeader.Value = math.NewUint256FromBytes([]byte{1}).Bytes()
	buf.Reset()

	buf.Write(txHeader.Src)
	buf.Write(txHeader.Dst)
	buf.Write(txHeader.Data)
	buf.Write(txHeader.Value)
	buf.Write(txHeader.Nonce)
	txHeader.Signature = signer.Sign(buf.Bytes()).Bytes()
	b, err = proto.Marshal(&txHeader)
	if err != nil {
		t.Error(err)
		return
	}

	bytesBuf[0] = transactiontype.SystemCall

	copy(bytesBuf[1:], b)

	fmt.Println(nsb.DeliverTx(types.RequestDeliverTx{
		Tx: bytesBuf[:1+len(b)],
	}))
	var err2 error
	err, err2 = nsb.Stop()
	if err != nil {
		t.Error(err)
		return
	}
	if err2 != nil {
		t.Error(err2)
		return
	}
}
