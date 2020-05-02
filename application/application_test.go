package nsb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/op-intent/instruction"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"

	transactiontype "github.com/HyperService-Consortium/NSB/application/transaction-type"
	"github.com/HyperService-Consortium/NSB/contract/isc"
	"github.com/HyperService-Consortium/NSB/grpc/nsbrpc"
	. "github.com/HyperService-Consortium/NSB/log"
	"github.com/HyperService-Consortium/NSB/math"
	signaturetype "github.com/HyperService-Consortium/go-uip/const/signature_type"
	"github.com/HyperService-Consortium/go-uip/signaturer"
	"github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/abci/types"

	"golang.org/x/crypto/ed25519"
)

func createISCTestPacket(t *testing.T, signer, uu, vv uip.Signer) (tx types.RequestDeliverTx, ok bool) {
	var u, v = uu.GetPublicKey(), vv.GetPublicKey()
	var iscOnwers = [][]byte{signer.GetPublicKey(), u, v}
	var funds = []uint64{0, 0, 0}
	var vesSig = []byte{0}

	var txBuf = bytes.NewBuffer(nil)
	sugar.HandlerError0(instruction.EncodeInstruction(&instruction.TransactionIntent{
		Src: u,
		Dst: v,
		//Seq:  math.NewUint256FromHexString("10"),
		Amt:  "10",
		Meta: []byte{0},
	}, txBuf))

	var transactionIntents = [][]byte{
		txBuf.Bytes(),
	}
	var args = &isc.ArgsCreateNewContract{
		IscOwners:          iscOnwers,
		Funds:              funds,
		VesSig:             vesSig,
		TransactionIntents: transactionIntents,
	}

	var fap FAPair
	var err error
	fap.FuncName = "isc"
	fap.Args, err = json.Marshal(args)
	if err != nil {
		t.Error(err)
		return
	}

	b, ok := pack(t, signer, fap)

	return types.RequestDeliverTx{
		Tx: b,
	}, ok
}

func pack(t *testing.T, signer uip.Signer, fap FAPair) (b []byte, ok bool) {
	var nonce, bytesBuf = make([]byte, 32), make([]byte, 65536)

	var (
		txHeader nsbrpc.TransactionHeader
		err      error
	)

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
	signature, err := signer.Sign(buf.Bytes())
	if err != nil {
		t.Error(err)
		return
	}
	txHeader.Signature = signature.Bytes()
	b, err = proto.Marshal(&txHeader)
	if err != nil {
		t.Error(err)
		return
	}

	bytesBuf[0] = transactiontype.CreateContract

	copy(bytesBuf[1:], b)
	return bytesBuf[:1+len(b)], true
}

func createSigner(t *testing.T) uip.Signer {
	var pri = make([]byte, 32)
	for idx := 0; idx < 32; idx++ {
		pri[idx] = uint8(idx)
	}

	var signer, err = signaturer.NewTendermintNSBSigner(ed25519.NewKeyFromSeed(pri))
	if err != nil {
		t.Error(err)
		return nil
	}
	return signer
}

func createU(t *testing.T) uip.Signer {
	uu, err := signaturer.NewTendermintNSBSigner(
		ed25519.NewKeyFromSeed(append(make([]byte, 31), 1)))
	if err != nil {
		t.Error(err)
		return nil
	}
	return uu
}

func TestCreateContract(t *testing.T) {
	var nonce, bytesBuf = make([]byte, 32), make([]byte, 65536)
	signer := createSigner(t)
	assert.NotNil(t, signer)
	uu := createU(t)
	assert.NotNil(t, uu)
	vv := createV(t)
	assert.NotNil(t, vv)

	nsb := createApplication(t, "./data/")
	if nsb == nil {
		return
	}

	tx, ok := createISCTestPacket(t, signer, uu, vv)
	if !ok {
		return
	}
	ret := nsb.DeliverTx(tx)

	fmt.Println(ret)

	var argss = &ArgsAddAction{
		ISCAddress: ret.Data,
		Tid:        0,
		Aid:        3,
		Type:       signaturetype.Ed25519,
		Signature:  sugar.HandlerError(uu.Sign([]byte("123"))).(uip.Signature).Bytes(),
		Content:    []byte("123"),
	}
	var txHeader nsbrpc.TransactionHeader
	var fap FAPair
	var err error
	fap.FuncName = "system.action@addAction"
	fap.Args, err = json.Marshal(argss)
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
	buf := bytes.NewBuffer(bytesBuf)
	buf.Reset()

	txHeader.Nonce = math.NewUint256FromBytes(nonce).Bytes()
	txHeader.Value = math.NewUint256FromBytes([]byte{0}).Bytes()
	buf.Reset()

	buf.Write(txHeader.Src)
	buf.Write(txHeader.Dst)
	buf.Write(txHeader.Data)
	buf.Write(txHeader.Value)
	buf.Write(txHeader.Nonce)
	txHeader.Signature = sugar.HandlerError(signer.Sign(buf.Bytes())).(uip.Signature).Bytes()
	b, err := proto.Marshal(&txHeader)
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

func createV(t *testing.T) uip.Signer {

	vv, err := signaturer.NewTendermintNSBSigner(ed25519.NewKeyFromSeed(append(make([]byte, 31), 2)))
	if err != nil {
		t.Error(err)
		return nil
	}

	return vv
}

func createApplication(t *testing.T, s string) *NSBApplication {
	logger, err := NewZapColorfulDevelopmentSugarLogger()
	if err != nil {
		t.Error(err)
		return nil
	}

	nsb, err := NewNSBApplication(logger, s)
	if err != nil {
		t.Error(err)
		return nil
	}
	return nsb
}

func TestSetBalance(t *testing.T) {

	var pri, nonce, bytesBuf = make([]byte, 64), make([]byte, 32), make([]byte, 65536)
	for idx := 0; idx < 64; idx++ {
		pri[idx] = uint8(idx)
	}
	var signer, err = signaturer.NewTendermintNSBSigner(pri)
	if err != nil {
		t.Fatal(err)
	}

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
	txHeader.Signature = sugar.HandlerError(signer.Sign(buf.Bytes())).(uip.Signature).Bytes()

	b, err := proto.Marshal(&txHeader)
	if err != nil {
		t.Error(err)
		return
	}

	bytesBuf[0] = transactiontype.SystemCall

	copy(bytesBuf[1:], b)

	nsb := createApplication(t, "./data")
	if nsb == nil {
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
	var signer, err = signaturer.NewTendermintNSBSigner(pri)
	if err != nil {
		t.Fatal(err)
	}

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
	txHeader.Signature = sugar.HandlerError(signer.Sign(buf.Bytes())).(uip.Signature).Bytes()
	b, err := proto.Marshal(&txHeader)
	if err != nil {
		t.Error(err)
		return
	}

	bytesBuf[0] = transactiontype.SystemCall

	copy(bytesBuf[1:], b)

	nsb := createApplication(t, "./data")
	if nsb == nil {
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
	txHeader.Signature = sugar.HandlerError(signer.Sign(buf.Bytes())).(uip.Signature).Bytes()
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
