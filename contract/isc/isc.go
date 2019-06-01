package isc

import (
	"fmt"
	"bytes"
	"encoding/hex"
	"encoding/json"
	cmn "github.com/HyperServiceOne/NSB/common"
	"github.com/HyperServiceOne/NSB/math"
	"github.com/HyperServiceOne/NSB/util"
	"github.com/HyperServiceOne/NSB/contract/isc/transaction"
	"github.com/HyperServiceOne/NSB/contract/isc/ISCState"
	. "github.com/HyperServiceOne/NSB/common/contract_response"
)

// func (nsb *NSBApplication) activeISC(byteJson []byte) (types.ResponseDeliverTx) {
// 	return types.ResponseDeliverTx{
// 		Code: uint32(CodeOK),
// 	}
// }

func (iscc *ISC) IsOpening() bool {
	return iscc.env.Storage.GetUint8("iscState") == ISCState.Opening
}

func (iscc *ISC) IsActive() bool {
	return iscc.env.Storage.GetUint8("iscState") != ISCState.Closed
}

func (iscc *ISC) IsIniting() bool {
	return iscc.env.Storage.GetUint8("iscState") == ISCState.Initing
}

func (iscc *ISC) IsInited() bool {
	return iscc.env.Storage.GetUint8("iscState") == ISCState.Inited
}

func (iscc *ISC) IsSettling() bool {
	return iscc.env.Storage.GetUint8("iscState") == ISCState.Settling
}

func (iscc *ISC) NewContract(iscOwners [][]byte, funds []uint32, vesSig []byte, transactionIntents []*transaction.TransactionIntent) (*cmn.ContractCallBackInfo) {
	AssertTrue(len(iscOwners) == len(funds), "the length of owners is not equal to that of funds")
	AssertTrue(bytes.Equal(iscc.env.From, iscOwners[0]), "first owner must be sender")
	
	owners := iscc.env.Storage.NewBytesArray("owners")
	mustFunds := iscc.env.Storage.NewGeneralMap("mustFunds", util.BytesToBytesHelper, util.Uint32ToBytesHelper, util.BytesToUint32Helper)
	isOwner := iscc.env.Storage.NewBoolMap("iscOwner")

	for idx, iscOwner := range iscOwners {
		AssertTrue(isOwner.Get(iscOwner) == false, "the address is already the isc owner")
		owners.Append(iscOwner)
		mustFunds.Set(iscOwner, funds[idx])
		isOwner.Set(iscOwner, true)
	}
	iscc.env.Storage.NewBytesMap("userAcked").Set(iscc.env.From, vesSig)
	transactions := iscc.env.Storage.NewBytesArray("transactions")
	for _, transactionIntent := range transactionIntents {
		bt, err := json.Marshal(transactionIntent)
		if err != nil {
			return ExecContractError(err)
		}
		transactions.Append(bt)
	}
	return &cmn.ContractCallBackInfo{
		CodeResponse: uint32(CodeOK),
		Info: fmt.Sprintf("create success , this contract is deploy at %v", hex.EncodeToString(iscc.env.ContractAddress)),
	}
}

func (iscc *ISC) UpdateTxInfo(tid uint64, transactionIntent *transaction.TransactionIntent) (*cmn.ContractCallBackInfo) {
	AssertTrue(iscc.IsIniting(), "ISC is initialized")
	txArr := iscc.env.Storage.NewBytesArray("transactions")
	AssertTrue(txArr.Length() > tid, "transaction index overflow")
	AssertTrue(iscc.env.Storage.NewUint64Map("transactionsFrozen").Get(tid) == nil, "frozen transaction")
	fmt.Println("will set", transactionIntent.Bytes())
	txArr.Set(tid, transactionIntent.Bytes())
	return ExecOK(nil)
}

func (iscc *ISC) UpdateTxFr(tid uint64, fr []byte) (*cmn.ContractCallBackInfo) {
	AssertTrue(iscc.IsIniting(), "ISC is initialized")
	txArr := iscc.env.Storage.NewBytesArray("transactions")
	AssertTrue(txArr.Length() > tid, "transaction index overflow")
	AssertTrue(iscc.env.Storage.NewUint64Map("transactionsFrozen").Get(tid) == nil, "frozen transaction")

	txb := txArr.Get(tid)
	var tx transaction.TransactionIntent
	MustUnmarshal(txb, &tx)

	tx.Fr = fr
	txArr.Set(tid, tx.Bytes())
	return ExecOK(nil)
}

func (iscc *ISC) UpdateTxTo(tid uint64, to []byte) (*cmn.ContractCallBackInfo) {
	AssertTrue(iscc.IsIniting(), "ISC is initialized")
	txArr := iscc.env.Storage.NewBytesArray("transactions")
	AssertTrue(txArr.Length() > tid, "transaction index overflow")
	AssertTrue(iscc.env.Storage.NewUint64Map("transactionsFrozen").Get(tid) == nil, "frozen transaction")

	txb := txArr.Get(tid)
	var tx transaction.TransactionIntent
	MustUnmarshal(txb, &tx)

	tx.To = to
	txArr.Set(tid, tx.Bytes())
	return ExecOK(nil)
}

func (iscc *ISC) UpdateTxSeq(tid uint64, seq *math.Uint256) (*cmn.ContractCallBackInfo) {
	AssertTrue(iscc.IsIniting(), "ISC is initialized")
	txArr := iscc.env.Storage.NewBytesArray("transactions")
	AssertTrue(txArr.Length() > tid, "transaction index overflow")
	AssertTrue(iscc.env.Storage.NewUint64Map("transactionsFrozen").Get(tid) == nil, "frozen transaction")

	txb := txArr.Get(tid)
	var tx transaction.TransactionIntent
	MustUnmarshal(txb, &tx)

	tx.Seq = seq
	txArr.Set(tid, tx.Bytes())
	return ExecOK(nil)
}

func (iscc *ISC) UpdateTxAmt(tid uint64, amt *math.Uint256) (*cmn.ContractCallBackInfo) {
	AssertTrue(iscc.IsIniting(), "ISC is initialized")
	txArr := iscc.env.Storage.NewBytesArray("transactions")
	AssertTrue(txArr.Length() > tid, "transaction index overflow")
	AssertTrue(iscc.env.Storage.NewUint64Map("transactionsFrozen").Get(tid) == nil, "frozen transaction")

	txb := txArr.Get(tid)
	var tx transaction.TransactionIntent
	MustUnmarshal(txb, &tx)

	tx.Amt = amt
	txArr.Set(tid, tx.Bytes())
	return ExecOK(nil)
}

func (iscc *ISC) UpdateTxMeta(tid uint64, meta []byte) (*cmn.ContractCallBackInfo) {
	AssertTrue(iscc.IsIniting(), "ISC is initialized")
	txArr := iscc.env.Storage.NewBytesArray("transactions")
	AssertTrue(txArr.Length() > tid, "transaction index overflow")
	AssertTrue(iscc.env.Storage.NewUint64Map("transactionsFrozen").Get(tid) == nil, "frozen transaction")
	
	txb := txArr.Get(tid)
	var tx transaction.TransactionIntent
	MustUnmarshal(txb, &tx)

	tx.Meta = meta
	txArr.Set(tid, tx.Bytes())
	return ExecOK(nil)
}

func (iscc *ISC) FreezeInfo(tid uint64) (*cmn.ContractCallBackInfo) {
	AssertTrue(iscc.IsIniting(), "ISC is initialized")
	transFrozen := iscc.env.Storage.NewUint64Map("transactionsFrozen")
	if transFrozen.Get(tid) == nil {
		transFrozen.Set(tid, []byte{1})
		newf := iscc.env.Storage.GetUint64("freezedInfoCount") + 1
		if newf == iscc.env.Storage.NewBytesArray("transactions").Length() {
			iscc.env.Storage.SetUint8("iscState", ISCState.Inited)
		}
		iscc.env.Storage.SetUint64("freezedInfoCount", newf)
	}
	return ExecOK(nil)
}

// func (iscc *ISC) freezeAllInfo() (*cmn.ContractCallBackInfo) {
// 	iscc.env.Storage.NewUint64Map("transactionsFrozen").Set(tid, []byte{1})
// }

func (iscc *ISC) UserAck(signature []byte) (*cmn.ContractCallBackInfo) {
	// sign session
	AssertTrue(iscc.IsInited(), "ISC is initializing")
	userAcked := iscc.env.Storage.NewBytesMap("userAcked")
	if userAcked.Get(iscc.env.From) == nil {
		userAcked.Set(iscc.env.From, []byte{1})
		newf := iscc.env.Storage.GetUint64("userAckCount") + 1
		fmt.Println(newf, iscc.env.Storage.NewBytesArray("owners").Length())
		if newf == iscc.env.Storage.NewBytesArray("owners").Length() {
			iscc.env.Storage.SetUint8("iscState", ISCState.Opening)
		}
		iscc.env.Storage.SetUint64("userAckCount", newf)
	}
	
	return ExecOK(nil)
}


func (iscc *ISC) resetAckState() {
	iscc.env.Storage.SetUint8("iscState", ISCState.Initing)

	iscc.env.Storage.SetUint64("freezedInfoCount", 0)
	transFrozen := iscc.env.Storage.NewUint64Map("transactionsFrozen")
	txArr := iscc.env.Storage.NewBytesArray("transactions")
	for idx := uint64(0); idx < txArr.Length(); idx++ {
		transFrozen.Delete(idx)
	}

	iscc.env.Storage.SetUint64("userAckCount", 0)
	userAcked := iscc.env.Storage.NewBytesMap("userAcked")
	ownerArr := iscc.env.Storage.NewBytesArray("owners")
	for idx := uint64(0); idx < ownerArr.Length(); idx++ {
		userAcked.Delete(ownerArr.Get(idx))
	}
}


func (iscc *ISC) UserRefuse(signature []byte) (*cmn.ContractCallBackInfo) {
	AssertTrue(iscc.IsInited(), "ISC is initializing")
	iscc.resetAckState()
	return ExecOK(nil)
}

func (iscc *ISC) InsuranceClaim() (*cmn.ContractCallBackInfo) {
	AssertTrue(iscc.IsOpening(), "ISC is not open yet")
	return ExecOK(nil)
}

func (iscc *ISC) SettleContract() (*cmn.ContractCallBackInfo) {
	AssertTrue(iscc.IsSettling(), "ISC is not settling yet")
	return ExecOK(nil)
}
