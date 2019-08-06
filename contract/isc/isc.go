package isc

import (
	"bytes"
	"encoding/json"
	"fmt"

	cmn "github.com/HyperServiceOne/NSB/common"
	. "github.com/HyperServiceOne/NSB/common/contract_response"
	"github.com/HyperServiceOne/NSB/contract/isc/ISCState"
	TxState "github.com/HyperServiceOne/NSB/contract/isc/TxState"
	"github.com/HyperServiceOne/NSB/contract/isc/transaction"
	"github.com/HyperServiceOne/NSB/math"
	"github.com/HyperServiceOne/NSB/util"
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

func (iscc *ISC) NewContract(iscOwners [][]byte, funds []uint32, vesSig []byte, transactionIntents []*transaction.TransactionIntent) *cmn.ContractCallBackInfo {
	AssertTrue(len(iscOwners) != 0, "nil iscOwners is not allowed")
	AssertTrue(len(iscOwners) == len(funds), "the length of owners is not equal to that of funds")
	AssertTrue(bytes.Equal(iscc.env.From, iscOwners[0]), "first owner must be sender")

	owners := iscc.env.Storage.NewBytesArray("owners")
	mustFunds := iscc.env.Storage.NewGeneralMap("mustFunds", util.BytesToBytesHelper, util.Uint32ToBytesHelper, util.BytesToUint32Helper)
	isOwner := iscc.env.Storage.NewBoolMap("iscOwner")
	var AidMap = iscc.env.Storage.NewUint64Map("AidMap")

	for idx, iscOwner := range iscOwners {
		AssertTrue(isOwner.Get(iscOwner) == false, "the address is already the isc owner")
		owners.Append(iscOwner)
		mustFunds.Set(iscOwner, funds[idx])
		isOwner.Set(iscOwner, true)
	}
	iscc.env.Storage.NewBytesMap("userAcked").Set(iscc.env.From, vesSig)
	iscc.env.Storage.SetUint64("userAckCount", 1)
	transactions := iscc.env.Storage.NewBytesArray("transactions")
	for idx, transactionIntent := range transactionIntents {
		bt, err := json.Marshal(transactionIntent)
		if err != nil {
			return ExecContractError(err)
		}
		transactions.Append(bt)
		// todo : avoid int -> uint64
		AidMap.Set(uint64(idx), util.Uint64ToBytes(TxState.Initing))
	}
	return &cmn.ContractCallBackInfo{
		CodeResponse: CodeOK(),
		Info:         "success",
		Data:         iscc.env.ContractAddress,
	}
}

func (iscc *ISC) UpdateTxInfo(tid uint64, transactionIntent *transaction.TransactionIntent) *cmn.ContractCallBackInfo {
	AssertTrue(iscc.IsIniting(), "ISC is initialized")
	txArr := iscc.env.Storage.NewBytesArray("transactions")
	AssertTrue(txArr.Length() > tid, "transaction index overflow")
	AssertTrue(util.BytesToUint64(iscc.env.Storage.NewUint64Map("AidMap").Get(tid)) == TxState.Initing, "frozen transaction")
	fmt.Println("will set", transactionIntent.Bytes())
	txArr.Set(tid, transactionIntent.Bytes())
	return ExecOK(nil)
}

func (iscc *ISC) UpdateTxFr(tid uint64, fr []byte) *cmn.ContractCallBackInfo {
	AssertTrue(iscc.IsIniting(), "ISC is initialized")
	txArr := iscc.env.Storage.NewBytesArray("transactions")
	AssertTrue(txArr.Length() > tid, "transaction index overflow")
	AssertTrue(util.BytesToUint64(iscc.env.Storage.NewUint64Map("AidMap").Get(tid)) == TxState.Initing, "frozen transaction")

	txb := txArr.Get(tid)
	var tx transaction.TransactionIntent
	MustUnmarshal(txb, &tx)

	tx.Fr = fr
	txArr.Set(tid, tx.Bytes())
	return ExecOK(nil)
}

func (iscc *ISC) UpdateTxTo(tid uint64, to []byte) *cmn.ContractCallBackInfo {
	AssertTrue(iscc.IsIniting(), "ISC is initialized")
	txArr := iscc.env.Storage.NewBytesArray("transactions")
	AssertTrue(txArr.Length() > tid, "transaction index overflow")
	AssertTrue(util.BytesToUint64(iscc.env.Storage.NewUint64Map("AidMap").Get(tid)) == TxState.Initing, "frozen transaction")

	txb := txArr.Get(tid)
	var tx transaction.TransactionIntent
	MustUnmarshal(txb, &tx)

	tx.To = to
	txArr.Set(tid, tx.Bytes())
	return ExecOK(nil)
}

func (iscc *ISC) UpdateTxSeq(tid uint64, seq *math.Uint256) *cmn.ContractCallBackInfo {
	AssertTrue(iscc.IsIniting(), "ISC is initialized")
	txArr := iscc.env.Storage.NewBytesArray("transactions")
	AssertTrue(txArr.Length() > tid, "transaction index overflow")
	AssertTrue(util.BytesToUint64(iscc.env.Storage.NewUint64Map("AidMap").Get(tid)) == TxState.Initing, "frozen transaction")

	txb := txArr.Get(tid)
	var tx transaction.TransactionIntent
	MustUnmarshal(txb, &tx)

	tx.Seq = seq
	txArr.Set(tid, tx.Bytes())
	return ExecOK(nil)
}

func (iscc *ISC) UpdateTxAmt(tid uint64, amt *math.Uint256) *cmn.ContractCallBackInfo {
	AssertTrue(iscc.IsIniting(), "ISC is initialized")
	txArr := iscc.env.Storage.NewBytesArray("transactions")
	AssertTrue(txArr.Length() > tid, "transaction index overflow")
	AssertTrue(util.BytesToUint64(iscc.env.Storage.NewUint64Map("AidMap").Get(tid)) == TxState.Initing, "frozen transaction")

	txb := txArr.Get(tid)
	var tx transaction.TransactionIntent
	MustUnmarshal(txb, &tx)

	tx.Amt = amt
	txArr.Set(tid, tx.Bytes())
	return ExecOK(nil)
}

func (iscc *ISC) UpdateTxMeta(tid uint64, meta []byte) *cmn.ContractCallBackInfo {
	AssertTrue(iscc.IsIniting(), "ISC is initialized")
	txArr := iscc.env.Storage.NewBytesArray("transactions")
	AssertTrue(txArr.Length() > tid, "transaction index overflow")
	AssertTrue(util.BytesToUint64(iscc.env.Storage.NewUint64Map("AidMap").Get(tid)) == TxState.Initing, "frozen transaction")

	txb := txArr.Get(tid)
	var tx transaction.TransactionIntent
	MustUnmarshal(txb, &tx)

	tx.Meta = meta
	txArr.Set(tid, tx.Bytes())
	return ExecOK(nil)
}

func (iscc *ISC) FreezeInfo(tid uint64) *cmn.ContractCallBackInfo {
	AssertTrue(iscc.IsIniting(), "ISC is initialized")
	var AidMap = iscc.env.Storage.NewUint64Map("AidMap")
	if util.BytesToUint64(AidMap.Get(tid)) == TxState.Initing {
		AidMap.Set(tid, util.Uint64ToBytes(TxState.Inited))
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

func (iscc *ISC) UserAck(fr, signature []byte) *cmn.ContractCallBackInfo {
	// sign session
	AssertTrue(iscc.IsInited(), "ISC is initializing")
	userAcked := iscc.env.Storage.NewBytesMap("userAcked")
	if userAcked.Get(fr) == nil {
		userAcked.Set(fr, []byte{1})
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
	var AidMap = iscc.env.Storage.NewUint64Map("AidMap")
	txArr := iscc.env.Storage.NewBytesArray("transactions")
	for idx := uint64(0); idx < txArr.Length(); idx++ {
		AidMap.Set(idx, util.Uint64ToBytes(TxState.Initing))
	}

	iscc.env.Storage.SetUint64("userAckCount", 0)
	userAcked := iscc.env.Storage.NewBytesMap("userAcked")
	ownerArr := iscc.env.Storage.NewBytesArray("owners")
	for idx := uint64(0); idx < ownerArr.Length(); idx++ {
		userAcked.Delete(ownerArr.Get(idx))
	}
}

func (iscc *ISC) UserRefuse(signature []byte) *cmn.ContractCallBackInfo {
	AssertTrue(iscc.IsInited(), "ISC is initializing")
	iscc.resetAckState()
	return ExecOK(nil)
}

func (iscc *ISC) InsuranceClaim(tid, aid uint64) *cmn.ContractCallBackInfo {
	AssertTrue(iscc.IsOpening(), "ISC is not open yet")
	var storing_tid = iscc.env.Storage.GetUint64("tid")
	AssertTrue(storing_tid == tid, "this transaction is not active")
	var AidMap = iscc.env.Storage.NewUint64Map("AidMap")
	var storing_aid = util.BytesToUint64(AidMap.Get(tid))
	// fmt.Println(storing_aid+1, aid)
	AssertTrue(storing_aid+1 == aid, "this action is not active")
	AidMap.Set(tid, util.Uint64ToBytes(aid))
	if aid == TxState.Closed {
		tid++
		iscc.env.Storage.SetUint64("tid", tid)
		if tid == iscc.env.Storage.NewBytesArray("transactions").Length() {
			iscc.env.Storage.SetUint8("iscState", ISCState.Settling)
		}
	}
	return ExecOK(nil)
}

func (iscc *ISC) SettleContract() *cmn.ContractCallBackInfo {
	AssertTrue(iscc.IsSettling(), "ISC is not settling yet")
	iscc.env.Storage.SetUint8("iscState", ISCState.Closed)
	return ExecOK(nil)
}

func (iscc *ISC) GetOwners() *cmn.ContractCallBackInfo {
	owners := iscc.env.Storage.NewBytesArray("owners")
	var ret [][]byte
	for idx := uint64(0); idx < owners.Length(); idx++ {
		ret = append(ret, owners.Get(idx))
	}

	b, err := json.Marshal(ret)
	if err != nil {
		return ExecContractError(err)
	}

	return &cmn.ContractCallBackInfo{
		CodeResponse: CodeOK(),
		Data:         b,
	}
}

func (iscc *ISC) IsOwner(address []byte) *cmn.ContractCallBackInfo {
	isOwner := iscc.env.Storage.NewBoolMap("iscOwner")
	b, err := json.Marshal(isOwner.Get(address))
	if err != nil {
		return ExecContractError(err)
	}

	return &cmn.ContractCallBackInfo{
		CodeResponse: CodeOK(),
		Data:         b,
	}
}

func (iscc *ISC) GetState() *cmn.ContractCallBackInfo {
	return &cmn.ContractCallBackInfo{
		CodeResponse: CodeOK(),
		Data:         []byte{iscc.env.Storage.GetUint8("iscState")},
	}
}
