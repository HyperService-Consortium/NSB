package isc

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	cmn "github.com/HyperServiceOne/NSB/common"
	"github.com/HyperServiceOne/NSB/contract/isc/ISCState"
	"github.com/HyperServiceOne/NSB/contract/isc/TxState"
	"github.com/HyperServiceOne/NSB/contract/isc/transaction"
	"github.com/HyperServiceOne/NSB/localstorage"
	"github.com/HyperServiceOne/NSB/math"
	"github.com/HyperServiceOne/NSB/util"
	"github.com/syndtr/goleveldb/leveldb"
)

var __x_ldb *leveldb.DB
var __x_storage *localstorage.LocalStorage
var __x_env *cmn.ContractEnvironment

func reset(t *testing.T, b []byte) []byte {
	t.Helper()
	var c []byte
	var err error
	if __x_storage != nil {
		c, err = __x_storage.Commit()
		if err != nil {
			t.Error(err)
			return nil
		}
	}
	__x_storage, err = localstorage.NewLocalStorage(b, c, __x_ldb)
	if err != nil {
		t.Error(err)
		return nil
	}
	__x_env = &cmn.ContractEnvironment{
		Storage: __x_storage,
	}
	return c
}

func resetroot(t *testing.T, b, c []byte) {
	t.Helper()
	var err error
	__x_storage, err = localstorage.NewLocalStorage(b, c, __x_ldb)
	if err != nil {
		t.Error(err)
		return
	}
	__x_env = &cmn.ContractEnvironment{
		Storage: __x_storage,
	}
	return
}

func TestMakeStorage(t *testing.T) {
	var err error
	__x_ldb, err = leveldb.OpenFile("./testdb", nil)
	if err != nil {
		t.Error(err)
		return
	}
	reset(t, []byte{0, 1, 2})
}

/*
type ArgsCreateNewContract struct {
	IscOwners          [][]byte                        `json:"isc_owners"`
	Funds              []uint32                        `json:"required_funds"`
	VesSig             []byte                          `json:"ves_signature"`
	TransactionIntents []*transaction.TransactionIntent `json:"transactionIntents"`
}
*/

/*
type TransactionIntent struct {
	Fr          []byte              `json:"from"`
	To          []byte              `json:"to"`
	Seq         *math.Uint256                `json:"seq"`
	Amt         *math.Uint256                `json:"amt"`
	Meta        []byte              `json:"meta"`
}
*/

/*
type ContractEnvironment struct {
	Storage         *localstorage.LocalStorage
	From            []byte
	ContractAddress []byte
	FuncName        string
	Args            []byte
	Value           *math.Uint256
}

DeliverTx
2 0w0
creating [105 115 99]
cmp [105 115 99] [105 115 99] cmp
Balance: 0
odeHash:
StorageRoot: , name:
 Balance: 0
odeHash:
StorageRoot: , name:isc

QAQ
*/

func TestRuntimeErrorBug(t *testing.T) {
	qwq := reset(t, []byte{0, 1, 2})
	defer func() {
		if err := recover(); err != nil {
			resetroot(t, []byte{0, 1, 2}, qwq)
			if err.(string) == "nil iscOwners is not allowed" {
				return
			}
			t.Error(err)
			return
		}
	}()

	// var raw = "7b226973635f6f776e657273223a5b2241514944222c2241514945225d2c2272657175697265645f66756e6473223a5b302c305d2c227665735f7369676e6174757265223a2241413d3d222c227472616e73616374696f6e5f696e74656e7473223a5b7b2266726f6d223a224151494441773d3d222c22746f223a224151494442413d3d222c22736571223a22222c22616d74223a224175413d222c226d657461223a6e756c6c7d5d7d"
	var raw = "7b2266756e6374696f6e5f6e616d65223a22222c2261726773223a2265794a7063324e66623364755a584a7a496a7062496b4652535551694c434a4255556c46496c3073496e4a6c63585670636d566b58325a31626d527a496a70624d43777758537769646d567a58334e705a323568644856795a534936496b4642505430694c434a30636d46756332466a64476c76626c3970626e526c626e527a496a706265794a6d636d3974496a6f695156464a52434973496e5276496a6f695156464a52454a42505430694c434a7a5a5845694f6949694c434a68625851694f694a4264554539496977696257563059534936626e56736248316466513d3d227d"
	b, err := hex.DecodeString(raw)
	if err != nil {
		t.Error(err)
		return
	}
	__x_env.From = []byte{1, 2, 3}
	__x_env.ContractAddress = []byte{1, 0, 2, 3, 4}
	__x_env.Args = b
	fmt.Println(CreateNewContract(__x_env))
}

func TestISCMakeISC(t *testing.T) {
	qwq := reset(t, []byte{0, 1, 2})
	defer func() {
		if err := recover(); err != nil {
			resetroot(t, []byte{0, 1, 2}, qwq)
			t.Error(err)
			return
		}
	}()

	var u, v = []byte{0, 0, 1}, []byte{0, 0, 2}
	var iscOnwers = [][]byte{u, v}
	var funds = []uint32{0, 0}
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
	var args = &ArgsCreateNewContract{
		IscOwners:          iscOnwers,
		Funds:              funds,
		VesSig:             vesSig,
		TransactionIntents: transactionIntents,
	}
	bt, err := json.Marshal(args)
	if err != nil {
		t.Error(err)
		return
	}
	__x_env.From = u
	__x_env.ContractAddress = []byte{0, 1, 2, 3}
	__x_env.Args = bt
	fmt.Println(CreateNewContract(__x_env))
	__x_storage.Commit()
}

func TestISCBadMakeISC(t *testing.T) {
	qwq := reset(t, []byte{0, 1, 2})
	defer func() {
		if err := recover(); err != nil {
			resetroot(t, []byte{0, 1, 2}, qwq)
			fmt.Println(err)
			return
		}
	}()

	var u, v = []byte{0, 0, 0}, []byte{0, 0, 2}
	var iscOnwers = [][]byte{u, v}
	var funds = []uint32{0, 0}
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
	var args = &ArgsCreateNewContract{
		IscOwners:          iscOnwers,
		Funds:              funds,
		VesSig:             vesSig,
		TransactionIntents: transactionIntents,
	}
	bt, err := json.Marshal(args)
	if err != nil {
		t.Error(err)
		return
	}
	__x_env.From = u
	__x_env.ContractAddress = []byte{0, 1, 2, 4}
	__x_env.Args = bt
	fmt.Println(CreateNewContract(__x_env))
}

func TestISCBadMakeISC2(t *testing.T) {
	qwq := reset(t, []byte{0, 1, 2})
	defer func() {
		if err := recover(); err != nil {
			resetroot(t, []byte{0, 1, 2}, qwq)
			fmt.Println(err)
			return
		}
	}()

	var u, v = []byte{0, 0, 1}, []byte{0, 0, 2}
	var iscOnwers = [][]byte{u, v}
	var funds = []uint32{0}
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
	var args = &ArgsCreateNewContract{
		IscOwners:          iscOnwers,
		Funds:              funds,
		VesSig:             vesSig,
		TransactionIntents: transactionIntents,
	}
	bt, err := json.Marshal(args)
	if err != nil {
		t.Error(err)
		return
	}
	__x_env.From = u
	__x_env.ContractAddress = []byte{0, 1, 2, 5}
	__x_env.Args = bt
	fmt.Println(CreateNewContract(__x_env))
}

func TestResetInfo(t *testing.T) {
	qwq := reset(t, []byte{0, 1, 2})
	defer func() {
		if err := recover(); err != nil {
			resetroot(t, []byte{0, 1, 2}, qwq)
			t.Error(err)
			return
		}
	}()

	var u, v = []byte{0, 0, 1}, []byte{0, 0, 2}
	var transactionIntent = &transaction.TransactionIntent{
		Fr:   u,
		To:   v,
		Seq:  math.NewUint256FromHexString("666"),
		Amt:  math.NewUint256FromHexString("666"),
		Meta: []byte{0},
	}
	var args = &ArgsUpdateTxInfo{
		Tid:               0,
		TransactionIntent: transactionIntent,
	}
	bt, err := json.Marshal(args)
	if err != nil {
		t.Error(err)
		return
	}
	__x_env.From = u
	__x_env.ContractAddress = []byte{0, 1, 2, 3}
	__x_env.FuncName = "UpdateTxInfo"
	__x_env.Args = bt
	txArr := __x_storage.NewBytesArray("transactions")
	fmt.Println("set", txArr.Get(0))

	fmt.Println(RigisteredMethod(__x_env))
	__x_storage.Commit()
	fmt.Println("set", txArr.Get(0))

	bt, err = json.Marshal(transactionIntent)
	if err != nil {
		t.Error(err)
		return
	}

	if !bytes.Equal(bt, txArr.Get(0)) {
		fmt.Println("Test Set Info not equal...")
	}
}

func TestFreezeInfo(t *testing.T) {
	qwq := reset(t, []byte{0, 1, 2})
	defer func() {
		if err := recover(); err != nil {
			resetroot(t, []byte{0, 1, 2}, qwq)
			t.Error(err)
			return
		}
	}()

	var u = []byte{0, 0, 1}
	var args = &ArgsFreezeInfo{
		Tid: 0,
	}
	bt, err := json.Marshal(args)
	if err != nil {
		t.Error(err)
		return
	}
	__x_env.From = u
	__x_env.ContractAddress = []byte{0, 1, 2, 3}
	__x_env.FuncName = "FreezeInfo"
	__x_env.Args = bt

	fmt.Println(RigisteredMethod(__x_env))
	__x_storage.Commit()

	if __x_storage.NewUint64Map("transactionsFrozen").Get(0) == nil {
		fmt.Println("not frozen")
	}
	fmt.Println(__x_storage.GetUint8("iscState") == ISCState.Inited)
}

func TestResetInfoAfterAllFrozen(t *testing.T) {
	qwq := reset(t, []byte{0, 1, 2})
	defer func() {
		if err := recover(); err != nil {
			resetroot(t, []byte{0, 1, 2}, qwq)
			fmt.Println(err)
			return
		}
	}()

	var u, v = []byte{0, 0, 1}, []byte{0, 0, 2}
	var transactionIntent = &transaction.TransactionIntent{
		Fr:   u,
		To:   v,
		Seq:  math.NewUint256FromHexString("666"),
		Amt:  math.NewUint256FromHexString("666"),
		Meta: []byte{0},
	}
	var args = &ArgsUpdateTxInfo{
		Tid:               0,
		TransactionIntent: transactionIntent,
	}
	bt, err := json.Marshal(args)
	if err != nil {
		t.Error(err)
		return
	}
	__x_env.From = u
	__x_env.ContractAddress = []byte{0, 1, 2, 3}
	__x_env.FuncName = "UpdateTxInfo"
	__x_env.Args = bt
	txArr := __x_storage.NewBytesArray("transactions")
	fmt.Println("set", txArr.Get(0))

	fmt.Println(RigisteredMethod(__x_env))
	__x_storage.Commit()
	fmt.Println("set", txArr.Get(0))

	bt, err = json.Marshal(transactionIntent)
	if err != nil {
		t.Error(err)
		return
	}

	if !bytes.Equal(bt, txArr.Get(0)) {
		fmt.Println("Test Set Info not equal...")
	}
}

func TestUserAcks(t *testing.T) {
	qwq := reset(t, []byte{0, 1, 2})
	defer func() {
		if err := recover(); err != nil {
			resetroot(t, []byte{0, 1, 2}, qwq)
			t.Error(err)
			return
		}
	}()

	var u, v = []byte{0, 0, 1}, []byte{0, 0, 2}
	var args = &ArgsUserAck{
		Address:   u,
		Signature: []byte("test..."),
	}
	bt, err := json.Marshal(args)
	if err != nil {
		t.Error(err)
		return
	}
	__x_env.From = u
	__x_env.ContractAddress = []byte{0, 1, 2, 3}
	__x_env.FuncName = "UserAck"
	__x_env.Args = bt

	fmt.Println(RigisteredMethod(__x_env))
	__x_storage.Commit()

	args.Address = v
	bt, err = json.Marshal(args)
	if err != nil {
		t.Error(err)
		return
	}
	__x_env.From = v
	__x_env.Args = bt

	fmt.Println(RigisteredMethod(__x_env))
	__x_storage.Commit()

	fmt.Println(__x_storage.GetUint8("iscState") == ISCState.Opening)
}

func TestProcessTransaction(t *testing.T) {
	qwq := reset(t, []byte{0, 1, 2})
	defer func() {
		if err := recover(); err != nil {
			resetroot(t, []byte{0, 1, 2}, qwq)
			t.Error(err)
			return
		}
	}()

	var u, v = []byte{0, 0, 1}, []byte{0, 0, 2}

	__x_env.From = v
	__x_env.ContractAddress = []byte{0, 1, 2, 3}
	__x_env.FuncName = "InsuranceClaim"
	__x_env.Args = append(append(make([]byte, 0, 16), util.Uint64ToBytes(0)...), util.Uint64ToBytes(TxState.Instantiating)...)

	fmt.Println(RigisteredMethod(__x_env))
	__x_storage.Commit()

	__x_env.From = u
	__x_env.ContractAddress = []byte{0, 1, 2, 3}
	__x_env.FuncName = "InsuranceClaim"
	__x_env.Args = append(append(make([]byte, 0, 16), util.Uint64ToBytes(0)...), util.Uint64ToBytes(TxState.Instantiated)...)

	fmt.Println(RigisteredMethod(__x_env))
	__x_storage.Commit()

	__x_env.From = v
	__x_env.ContractAddress = []byte{0, 1, 2, 3}
	__x_env.FuncName = "InsuranceClaim"
	__x_env.Args = append(append(make([]byte, 0, 16), util.Uint64ToBytes(0)...), util.Uint64ToBytes(TxState.Open)...)

	fmt.Println(RigisteredMethod(__x_env))
	__x_storage.Commit()

	__x_env.From = u
	__x_env.Args = append(append(make([]byte, 0, 16), util.Uint64ToBytes(0)...), util.Uint64ToBytes(TxState.Opened)...)

	fmt.Println(RigisteredMethod(__x_env))
	__x_storage.Commit()

	__x_env.From = v
	__x_env.Args = append(append(make([]byte, 0, 16), util.Uint64ToBytes(0)...), util.Uint64ToBytes(TxState.Closed)...)

	fmt.Println(RigisteredMethod(__x_env))
	__x_storage.Commit()

	fmt.Println(__x_storage.GetUint8("iscState") == ISCState.Opening)
	fmt.Println(__x_storage.GetUint8("iscState") == ISCState.Settling)
}

func TestCloseContract(t *testing.T) {
	qwq := reset(t, []byte{0, 1, 2})
	defer func() {
		if err := recover(); err != nil {
			resetroot(t, []byte{0, 1, 2}, qwq)
			t.Error(err)
			return
		}
	}()

	var u = []byte{0, 0, 1}

	__x_env.From = u
	__x_env.ContractAddress = []byte{0, 1, 2, 3}
	__x_env.FuncName = "SettleContract"

	fmt.Println(RigisteredMethod(__x_env))
	__x_storage.Commit()

	// __x_env.From = v
	//
	// fmt.Println(RigisteredMethod(__x_env))
	// __x_storage.Commit()

	fmt.Println(__x_storage.GetUint8("iscState") == ISCState.Opening)
	fmt.Println(__x_storage.GetUint8("iscState") == ISCState.Settling)
	fmt.Println(__x_storage.GetUint8("iscState") == ISCState.Closed)
}

func TestCloseContractTwice(t *testing.T) {
	qwq := reset(t, []byte{0, 1, 2})
	defer func() {
		if err := recover(); err != nil {
			resetroot(t, []byte{0, 1, 2}, qwq)
			fmt.Println(err)
			return
		}
	}()

	var v = []byte{0, 0, 2}

	__x_env.From = v
	__x_env.ContractAddress = []byte{0, 1, 2, 3}
	__x_env.FuncName = "SettleContract"

	fmt.Println(RigisteredMethod(__x_env))

	fmt.Println(__x_storage.GetUint8("iscState") == ISCState.Opening)
	fmt.Println(__x_storage.GetUint8("iscState") == ISCState.Settling)
	fmt.Println(__x_storage.GetUint8("iscState") == ISCState.Closed)
}

func TestUnmakeStorage(t *testing.T) {
	__x_ldb.Close()
}
