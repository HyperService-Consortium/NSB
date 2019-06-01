package isc

import (
	"github.com/HyperServiceOne/NSB/localstorage"
	"github.com/HyperServiceOne/NSB/math"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/HyperServiceOne/NSB/contract/isc/transaction"
	"github.com/HyperServiceOne/NSB/contract/isc/ISCState"
	cmn "github.com/HyperServiceOne/NSB/common"
	"fmt"
	"bytes"
	"testing"
	"encoding/json"
)


var __x_ldb *leveldb.DB
var __x_storage *localstorage.LocalStorage
var __x_env *cmn.ContractEnvironment


func TestMakeStorage(t *testing.T) {
	var err error
	__x_ldb, err = leveldb.OpenFile("./testdb", nil)
	if err != nil {
		t.Error(err)
		return
	}
	__x_storage, err = localstorage.NewLocalStorage([]byte{0,1,2}, []byte{}, __x_ldb)
	if err != nil {
		t.Error(err)
		return
	}
	__x_env = &cmn.ContractEnvironment{
		Storage: __x_storage,
	}
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
*/

func TestISCMakeISC(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			__x_storage.Revert()
			t.Error(err)
			return
		}
	}()


	var u, v = []byte{0, 0, 1}, []byte{0, 0, 2};
	var iscOnwers = [][]byte{u, v}
	var funds = []uint32{0, 0}
	var vesSig = []byte{0}
	var transactionIntents = []*transaction.TransactionIntent{
		&transaction.TransactionIntent{
			Fr: u,
			To: v,
			Seq: math.NewUint256FromHexString("10"),
			Amt: math.NewUint256FromHexString("10"),
			Meta: []byte{0},
		},
	}
	var args = &ArgsCreateNewContract{
		IscOwners: iscOnwers,
		Funds: funds,
		VesSig: vesSig,
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
	defer func() {
		if err := recover(); err != nil {
			__x_storage.Revert()
			fmt.Println(err)
			return
		}
	}()

	var u, v = []byte{0, 0, 0}, []byte{0, 0, 2};
	var iscOnwers = [][]byte{u, v}
	var funds = []uint32{0, 0}
	var vesSig = []byte{0}
	var transactionIntents = []*transaction.TransactionIntent{
		&transaction.TransactionIntent{
			Fr: u,
			To: v,
			Seq: math.NewUint256FromHexString("10"),
			Amt: math.NewUint256FromHexString("10"),
			Meta: []byte{0},
		},
	}
	var args = &ArgsCreateNewContract{
		IscOwners: iscOnwers,
		Funds: funds,
		VesSig: vesSig,
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
	defer func() {
		if err := recover(); err != nil {
			__x_storage.Revert()
			fmt.Println(err)
			return
		}
	}()
	
	var u, v = []byte{0, 0, 1}, []byte{0, 0, 2};
	var iscOnwers = [][]byte{u, v}
	var funds = []uint32{0}
	var vesSig = []byte{0}
	var transactionIntents = []*transaction.TransactionIntent{
		&transaction.TransactionIntent{
			Fr: u,
			To: v,
			Seq: math.NewUint256FromHexString("10"),
			Amt: math.NewUint256FromHexString("10"),
			Meta: []byte{0},
		},
	}
	var args = &ArgsCreateNewContract{
		IscOwners: iscOnwers,
		Funds: funds,
		VesSig: vesSig,
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
	defer func() {
		if err := recover(); err != nil {
			__x_storage.Revert()
			fmt.Println(err)
			return
		}
	}()
	

	var u, v = []byte{0, 0, 1}, []byte{0, 0, 2};
	var transactionIntent = &transaction.TransactionIntent{
		Fr: u,
		To: v,
		Seq: math.NewUint256FromHexString("666"),
		Amt: math.NewUint256FromHexString("666"),
		Meta: []byte{0},
	}
	var args = &ArgsUpdateTxInfo {
		Tid: 0,
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

	if ! bytes.Equal(bt, txArr.Get(0)) {
		fmt.Println("Test Set Info not equal...")
	}
}

func TestFreezeInfo(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			__x_storage.Revert()
			fmt.Println(err)
			return
		}
	}()
	

	var u = []byte{0, 0, 1}
	var args = &ArgsFreezeInfo {
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
	defer func() {
		if err := recover(); err != nil {
			__x_storage.Revert()
			fmt.Println(err)
			return
		}
	}()
	

	var u, v = []byte{0, 0, 1}, []byte{0, 0, 2};
	var transactionIntent = &transaction.TransactionIntent{
		Fr: u,
		To: v,
		Seq: math.NewUint256FromHexString("666"),
		Amt: math.NewUint256FromHexString("666"),
		Meta: []byte{0},
	}
	var args = &ArgsUpdateTxInfo {
		Tid: 0,
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

	if ! bytes.Equal(bt, txArr.Get(0)) {
		fmt.Println("Test Set Info not equal...")
	}
}


func TestUserAcks(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			__x_storage.Revert()
			fmt.Println(err)
			return
		}
	}()
	

	var u, v = []byte{0, 0, 1}, []byte{0, 0, 2}
	var args = &ArgsUserAck {
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
	
	
	__x_env.From = v

	fmt.Println(RigisteredMethod(__x_env))
	__x_storage.Commit()
	
	fmt.Println(__x_storage.GetUint8("iscState") == ISCState.Opening)
}

func TestUnmakeStorage(t *testing.T) {
	__x_ldb.Close()
}

