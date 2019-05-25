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
		AssertTrue(isOwner.Get(iscOwner) == false, "the address is already isc owner")
		owners.Append(iscOwner)
		mustFunds.Set(iscOwner, funds[idx])
		isOwner.Set(iscOwner, true)
	}
	iscc.env.Storage.NewBytesMap("userAcked").Set(iscc.env.From, vesSig)
	transactions := iscc.env.Storage.NewBytesArray("transactions")
	for transactionIntent := range transactionIntents {
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

func (iscc *ISC) UpdateTxInfo(tid int64, transactionIntent *transaction.TransactionIntent) (*cmn.ContractCallBackInfo) {
	AssertTrue(iscc.IsIniting(), "ISC is initialized")
	AssertTrue(iscc.env.Storage.NewInt64Map("transactionsFrozen").Get(tid)[0] == 0, "frozen transaction")

	iscc.env.Storage.NewInt64Map("transactions").Set(tid, transactionIntent.Bytes())
	return ExecOK(nil)
}

func (iscc *ISC) UpdateTxFr(tid int64, fr []byte) (*cmn.ContractCallBackInfo) {
	AssertTrue(iscc.IsIniting(), "ISC is initialized")
	AssertTrue(iscc.env.Storage.NewInt64Map("transactionsFrozen").Get(tid)[0] == 0, "frozen transaction")

	txArr, tb := iscc.env.Storage.NewInt64Map("transactions"), tid
	txb := txArr.Get(tb)
	var tx transaction.TransactionIntent
	MustUnmarshal(txb, &tx)

	tx.Fr = fr
	txArr.Set(tb, tx.Bytes())
	return ExecOK(nil)
}

func (iscc *ISC) UpdateTxTo(tid int64, to []byte) (*cmn.ContractCallBackInfo) {
	AssertTrue(iscc.IsIniting(), "ISC is initialized")
	AssertTrue(iscc.env.Storage.NewInt64Map("transactionsFrozen").Get(tid)[0] == 0, "frozen transaction")

	txArr, tb := iscc.env.Storage.NewInt64Map("transactions"), tid
	txb := txArr.Get(tb)
	var tx transaction.TransactionIntent
	MustUnmarshal(txb, &tx)

	tx.To = to
	txArr.Set(tb, tx.Bytes())
	return ExecOK(nil)
}

func (iscc *ISC) UpdateTxSeq(tid int64, seq *math.Uint256) (*cmn.ContractCallBackInfo) {
	AssertTrue(iscc.IsIniting(), "ISC is initialized")
	AssertTrue(iscc.env.Storage.NewInt64Map("transactionsFrozen").Get(tid)[0] == 0, "frozen transaction")

	txArr, tb := iscc.env.Storage.NewInt64Map("transactions"), tid
	txb := txArr.Get(tb)
	var tx transaction.TransactionIntent
	MustUnmarshal(txb, &tx)

	tx.Seq = seq
	txArr.Set(tb, tx.Bytes())
	return ExecOK(nil)
}

func (iscc *ISC) UpdateTxAmt(tid int64, amt *math.Uint256) (*cmn.ContractCallBackInfo) {
	AssertTrue(iscc.IsIniting(), "ISC is initialized")
	AssertTrue(iscc.env.Storage.NewInt64Map("transactionsFrozen").Get(tid)[0] == 0, "frozen transaction")

	txArr, tb := iscc.env.Storage.NewInt64Map("transactions"), tid
	txb := txArr.Get(tb)
	var tx transaction.TransactionIntent
	MustUnmarshal(txb, &tx)

	tx.Amt = amt
	txArr.Set(tb, tx.Bytes())
	return ExecOK(nil)
}

func (iscc *ISC) UpdateTxMeta(tid int64, meta []byte) (*cmn.ContractCallBackInfo) {
	AssertTrue(iscc.IsIniting(), "ISC is initialized")
	AssertTrue(iscc.env.Storage.NewInt64Map("transactionsFrozen").Get(tid)[0] == 0, "frozen transaction")
	
	txArr, tb := iscc.env.Storage.NewInt64Map("transactions"), tid
	txb := txArr.Get(tb)
	var tx transaction.TransactionIntent
	MustUnmarshal(txb, &tx)

	tx.Meta = meta
	txArr.Set(tb, tx.Bytes())
	return ExecOK(nil)
}

func (iscc *ISC) FreezeInfo(tid int64) (*cmn.ContractCallBackInfo) {
	AssertTrue(iscc.IsIniting(), "ISC is initialized")
	iscc.env.Storage.NewInt64Map("transactionsFrozen").Set(tid, []byte{1})
	return ExecOK(nil)
}

// func (iscc *ISC) freezeAllInfo() (*cmn.ContractCallBackInfo) {
// 	iscc.env.Storage.NewInt64Map("transactionsFrozen").Set(tid, []byte{1})
// }

func (iscc *ISC) UserAck(signature []byte) (*cmn.ContractCallBackInfo) {
	// sign session
	AssertTrue(iscc.IsIniting(), "ISC is initialized")
	iscc.env.Storage.NewBytesMap("userAcked").Set(iscc.env.From, []byte{1})
	return ExecOK(nil)
}

func (iscc *ISC) UserRefuse(signature []byte) (*cmn.ContractCallBackInfo) {
	AssertTrue(iscc.IsIniting(), "ISC is initialized")
	iscc.env.Storage.NewBytesMap("userAcked").Set(iscc.env.From, []byte{1})
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

func (iscc *ISC) ReturnFunds() (*cmn.ContractCallBackInfo) {
	AssertTrue(iscc.IsActive() == false, "ISC is not closed yet")
	return ExecOK(nil)
}

