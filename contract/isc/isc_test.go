package isc

import (
	trie "github.com/HyperService-Consortium/go-mpt"
	"testing"

	cmn "github.com/HyperService-Consortium/NSB/common"
	"github.com/HyperService-Consortium/NSB/localstorage"
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

func createRoot(t *testing.T, b, c []byte) *cmn.ContractEnvironment {
	t.Helper()
	setupStorage(t, b)
	var err error
	storage, err := localstorage.NewLocalStorage(c, trie.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"), __x_ldb)
	if err != nil {
		t.Error(err)
		return nil
	}
	env := &cmn.ContractEnvironment{
		Storage:         storage,
		From:            b,
		ContractAddress: c,
		BN:              &storageImpl{},
	}
	return env
}

func setupStorage(t *testing.T, b []byte) {
	if __x_ldb != nil {
		reset(t, b)
		return
	}
	var err error
	__x_ldb, err = leveldb.OpenFile("./testdb", nil)
	if err != nil {
		t.Error(err)
		return
	}
	reset(t, b)
}

//func TestMakeStorage(t *testing.T) {
//	setupStorage(t, []byte{0, 1, 2})
//}

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

//func TestISCMakeISC(t *testing.T) {
//	qwq := reset(t, []byte{0, 1, 2})
//	defer func() {
//		if err := recover(); err != nil {
//			resetRoot(t, []byte{0, 1, 2}, qwq)
//			t.Error(err)
//			return
//		}
//	}()
//
//	var u, v = []byte{0, 0, 1}, []byte{0, 0, 2}
//	var iscOnwers = [][]byte{u, v}
//	var funds = []uint32{0, 0}
//	var vesSig = []byte{0}
//	var transactionIntents = []*transaction.TransactionIntent{
//		&transaction.TransactionIntent{
//			Fr:   u,
//			To:   v,
//			Seq:  math.NewUint256FromHexString("10"),
//			Amt:  math.NewUint256FromHexString("10"),
//			Meta: []byte{0},
//		},
//	}
//	var args = &ArgsCreateNewContract{
//		IscOwners:          iscOnwers,
//		Funds:              funds,
//		VesSig:             vesSig,
//		TransactionIntents: transactionIntents,
//	}
//	bt, err := json.Marshal(args)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	__x_env.From = u
//	__x_env.ContractAddress = []byte{0, 1, 2, 3}
//	__x_env.Args = bt
//	fmt.Println(CreateNewContract(__x_env))
//	__x_storage.Commit()
//}
//
//func TestISCBadMakeISC(t *testing.T) {
//	qwq := reset(t, []byte{0, 1, 2})
//	defer func() {
//		if err := recover(); err != nil {
//			resetRoot(t, []byte{0, 1, 2}, qwq)
//			fmt.Println(err)
//			return
//		}
//	}()
//
//	var u, v = []byte{0, 0, 0}, []byte{0, 0, 2}
//	var iscOnwers = [][]byte{u, v}
//	var funds = []uint32{0, 0}
//	var vesSig = []byte{0}
//	var transactionIntents = []*transaction.TransactionIntent{
//		&transaction.TransactionIntent{
//			Fr:   u,
//			To:   v,
//			Seq:  math.NewUint256FromHexString("10"),
//			Amt:  math.NewUint256FromHexString("10"),
//			Meta: []byte{0},
//		},
//	}
//	var args = &ArgsCreateNewContract{
//		IscOwners:          iscOnwers,
//		Funds:              funds,
//		VesSig:             vesSig,
//		TransactionIntents: transactionIntents,
//	}
//	bt, err := json.Marshal(args)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	__x_env.From = u
//	__x_env.ContractAddress = []byte{0, 1, 2, 4}
//	__x_env.Args = bt
//	fmt.Println(CreateNewContract(__x_env))
//}
//
//func TestISCBadMakeISC2(t *testing.T) {
//	qwq := reset(t, []byte{0, 1, 2})
//	defer func() {
//		if err := recover(); err != nil {
//			resetRoot(t, []byte{0, 1, 2}, qwq)
//			fmt.Println(err)
//			return
//		}
//	}()
//
//	var u, v = []byte{0, 0, 1}, []byte{0, 0, 2}
//	var iscOnwers = [][]byte{u, v}
//	var funds = []uint32{0}
//	var vesSig = []byte{0}
//	var transactionIntents = []*transaction.TransactionIntent{
//		&transaction.TransactionIntent{
//			Fr:   u,
//			To:   v,
//			Seq:  math.NewUint256FromHexString("10"),
//			Amt:  math.NewUint256FromHexString("10"),
//			Meta: []byte{0},
//		},
//	}
//	var args = &ArgsCreateNewContract{
//		IscOwners:          iscOnwers,
//		Funds:              funds,
//		VesSig:             vesSig,
//		TransactionIntents: transactionIntents,
//	}
//	bt, err := json.Marshal(args)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	__x_env.From = u
//	__x_env.ContractAddress = []byte{0, 1, 2, 5}
//	__x_env.Args = bt
//	fmt.Println(CreateNewContract(__x_env))
//}
//
//func TestResetInfo(t *testing.T) {
//	qwq := reset(t, []byte{0, 1, 2})
//	defer func() {
//		if err := recover(); err != nil {
//			resetRoot(t, []byte{0, 1, 2}, qwq)
//			t.Error(err)
//			return
//		}
//	}()
//
//	var u, v = []byte{0, 0, 1}, []byte{0, 0, 2}
//	var transactionIntent = &transaction.TransactionIntent{
//		Fr:   u,
//		To:   v,
//		Seq:  math.NewUint256FromHexString("666"),
//		Amt:  math.NewUint256FromHexString("666"),
//		Meta: []byte{0},
//	}
//	var args = &ArgsUpdateTxInfo{
//		Tid:               0,
//		TransactionIntent: transactionIntent,
//	}
//	bt, err := json.Marshal(args)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	__x_env.From = u
//	__x_env.ContractAddress = []byte{0, 1, 2, 3}
//	__x_env.FuncName = "UpdateTxInfo"
//	__x_env.Args = bt
//	txArr := __x_storage.NewBytesArray("transactions")
//	fmt.Println("set", txArr.Get(0))
//
//	fmt.Println(RegisteredMethod(__x_env))
//	__x_storage.Commit()
//	fmt.Println("set", txArr.Get(0))
//
//	bt, err = json.Marshal(transactionIntent)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	if !bytes.Equal(bt, txArr.Get(0)) {
//		fmt.Println("Test Set Info not equal...")
//	}
//}
//
//func TestFreezeInfo(t *testing.T) {
//	qwq := reset(t, []byte{0, 1, 2})
//	defer func() {
//		if err := recover(); err != nil {
//			resetRoot(t, []byte{0, 1, 2}, qwq)
//			t.Error(err)
//			return
//		}
//	}()
//
//	var u = []byte{0, 0, 1}
//	var args = &ArgsFreezeInfo{
//		Tid: 0,
//	}
//	bt, err := json.Marshal(args)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	__x_env.From = u
//	__x_env.ContractAddress = []byte{0, 1, 2, 3}
//	__x_env.FuncName = "FreezeInfo"
//	__x_env.Args = bt
//
//	fmt.Println(RegisteredMethod(__x_env))
//	__x_storage.Commit()
//
//	if __x_storage.NewUint64Map("transactionsFrozen").Get(0) == nil {
//		fmt.Println("not frozen")
//	}
//	fmt.Println(__x_storage.GetUint8("iscState") == ISCState.Inited)
//}
//
//func TestResetInfoAfterAllFrozen(t *testing.T) {
//	qwq := reset(t, []byte{0, 1, 2})
//	defer func() {
//		if err := recover(); err != nil {
//			resetRoot(t, []byte{0, 1, 2}, qwq)
//			fmt.Println(err)
//			return
//		}
//	}()
//
//	var u, v = []byte{0, 0, 1}, []byte{0, 0, 2}
//	var transactionIntent = &transaction.TransactionIntent{
//		Fr:   u,
//		To:   v,
//		Seq:  math.NewUint256FromHexString("666"),
//		Amt:  math.NewUint256FromHexString("666"),
//		Meta: []byte{0},
//	}
//	var args = &ArgsUpdateTxInfo{
//		Tid:               0,
//		TransactionIntent: transactionIntent,
//	}
//	bt, err := json.Marshal(args)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	__x_env.From = u
//	__x_env.ContractAddress = []byte{0, 1, 2, 3}
//	__x_env.FuncName = "UpdateTxInfo"
//	__x_env.Args = bt
//	txArr := __x_storage.NewBytesArray("transactions")
//	fmt.Println("set", txArr.Get(0))
//
//	fmt.Println(RegisteredMethod(__x_env))
//	__x_storage.Commit()
//	fmt.Println("set", txArr.Get(0))
//
//	bt, err = json.Marshal(transactionIntent)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	if !bytes.Equal(bt, txArr.Get(0)) {
//		fmt.Println("Test Set Info not equal...")
//	}
//}
//
//func TestUserAcks(t *testing.T) {
//	qwq := reset(t, []byte{0, 1, 2})
//	defer func() {
//		if err := recover(); err != nil {
//			resetRoot(t, []byte{0, 1, 2}, qwq)
//			t.Error(err)
//			return
//		}
//	}()
//
//	var u, v = []byte{0, 0, 1}, []byte{0, 0, 2}
//	var args = &ArgsUserAck{
//		Address:   u,
//		Signature: []byte("test..."),
//	}
//	bt, err := json.Marshal(args)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	__x_env.From = u
//	__x_env.ContractAddress = []byte{0, 1, 2, 3}
//	__x_env.FuncName = "UserAck"
//	__x_env.Args = bt
//
//	fmt.Println(RegisteredMethod(__x_env))
//	__x_storage.Commit()
//
//	args.Address = v
//	bt, err = json.Marshal(args)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	__x_env.From = v
//	__x_env.Args = bt
//
//	fmt.Println(RegisteredMethod(__x_env))
//	__x_storage.Commit()
//
//	fmt.Println(__x_storage.GetUint8("iscState") == ISCState.Opening)
//}
//
//func TestProcessTransaction(t *testing.T) {
//	qwq := reset(t, []byte{0, 1, 2})
//	defer func() {
//		if err := recover(); err != nil {
//			resetRoot(t, []byte{0, 1, 2}, qwq)
//			t.Error(err)
//			return
//		}
//	}()
//
//	var u, v = []byte{0, 0, 1}, []byte{0, 0, 2}
//
//	__x_env.From = v
//	__x_env.ContractAddress = []byte{0, 1, 2, 3}
//	__x_env.FuncName = "InsuranceClaim"
//	__x_env.Args = append(append(make([]byte, 0, 16), util.Uint64ToBytes(0)...), util.Uint64ToBytes(TxState.Instantiating)...)
//
//	fmt.Println(RegisteredMethod(__x_env))
//	__x_storage.Commit()
//
//	__x_env.From = u
//	__x_env.ContractAddress = []byte{0, 1, 2, 3}
//	__x_env.FuncName = "InsuranceClaim"
//	__x_env.Args = append(append(make([]byte, 0, 16), util.Uint64ToBytes(0)...), util.Uint64ToBytes(TxState.Instantiated)...)
//
//	fmt.Println(RegisteredMethod(__x_env))
//	__x_storage.Commit()
//
//	__x_env.From = v
//	__x_env.ContractAddress = []byte{0, 1, 2, 3}
//	__x_env.FuncName = "InsuranceClaim"
//	__x_env.Args = append(append(make([]byte, 0, 16), util.Uint64ToBytes(0)...), util.Uint64ToBytes(TxState.Open)...)
//
//	fmt.Println(RegisteredMethod(__x_env))
//	__x_storage.Commit()
//
//	__x_env.From = u
//	__x_env.Args = append(append(make([]byte, 0, 16), util.Uint64ToBytes(0)...), util.Uint64ToBytes(TxState.Opened)...)
//
//	fmt.Println(RegisteredMethod(__x_env))
//	__x_storage.Commit()
//
//	__x_env.From = v
//	__x_env.Args = append(append(make([]byte, 0, 16), util.Uint64ToBytes(0)...), util.Uint64ToBytes(TxState.Closed)...)
//
//	fmt.Println(RegisteredMethod(__x_env))
//	__x_storage.Commit()
//
//	fmt.Println(__x_storage.GetUint8("iscState") == ISCState.Opening)
//	fmt.Println(__x_storage.GetUint8("iscState") == ISCState.Settling)
//}
//
//func TestCloseContract(t *testing.T) {
//	qwq := reset(t, []byte{0, 1, 2})
//	defer func() {
//		if err := recover(); err != nil {
//			resetRoot(t, []byte{0, 1, 2}, qwq)
//			t.Error(err)
//			return
//		}
//	}()
//
//	var u = []byte{0, 0, 1}
//
//	__x_env.From = u
//	__x_env.ContractAddress = []byte{0, 1, 2, 3}
//	__x_env.FuncName = "SettleContract"
//
//	fmt.Println(RegisteredMethod(__x_env))
//	__x_storage.Commit()
//
//	// __x_env.From = v
//	//
//	// fmt.Println(RegisteredMethod(__x_env))
//	// __x_storage.Commit()
//
//	fmt.Println(__x_storage.GetUint8("iscState") == ISCState.Opening)
//	fmt.Println(__x_storage.GetUint8("iscState") == ISCState.Settling)
//	fmt.Println(__x_storage.GetUint8("iscState") == ISCState.Closed)
//}
//
//func TestCloseContractTwice(t *testing.T) {
//	qwq := reset(t, []byte{0, 1, 2})
//	defer func() {
//		if err := recover(); err != nil {
//			resetRoot(t, []byte{0, 1, 2}, qwq)
//			fmt.Println(err)
//			return
//		}
//	}()
//
//	var v = []byte{0, 0, 2}
//
//	__x_env.From = v
//	__x_env.ContractAddress = []byte{0, 1, 2, 3}
//	__x_env.FuncName = "SettleContract"
//
//	fmt.Println(RegisteredMethod(__x_env))
//
//	fmt.Println(__x_storage.GetUint8("iscState") == ISCState.Opening)
//	fmt.Println(__x_storage.GetUint8("iscState") == ISCState.Settling)
//	fmt.Println(__x_storage.GetUint8("iscState") == ISCState.Closed)
//}

//func TestUnmakeStorage(t *testing.T) {
//	__x_ldb.Close()
//}
