package response


import (
	"fmt"
	"github.com/Myriad-Dreamin/NSB/math"
	"github.com/tendermint/tendermint/abci/types"
)


type ResponseCode uint32
const ( // base
	codeOK ResponseCode = 0 + iota
	codeExecFail
	codeUpdateBalanceIn
	codeUpdateBalanceOut
	codeUnknown
	codeMissingTxMethod
	codeMissingContract
	codeCommitTxTrieError
	codeCommitAccTrieError
	codeUpdateTxTrieError
	codeUpdateAccTrieError
	codeRequestStorageError
	codeEncodeAccountInfoError
	codeTODO = 99
)

const ( // Decode
	codeDecodeJsonError ResponseCode = 100 + iota
	codeDecodeBytesError
	codeDecodeTxHeaderError
	codeDecodeAccountInfoError
	codeDecodeBalanceError
)

const ( // Transaction
	codeInvalidTxInputFormat ResponseCode = 200 + iota
	codeInvalidTxType
	codeReTrieveTxError
	codeDuplicateTxError
)

const ( // Contract
	codeContractPanic ResponseCode = 300 + iota
	codeInvalidFuncType
	codeExecContractError
	codeInsufficientBalanceToTransfer
	codeBalanceOverflow
)


var (
	ExecOK = &types.ResponseDeliverTx{
		Code: uint32(codeOK),
	}
	DuplicateTxError = &types.ResponseDeliverTx{
		Code: uint32(codeDuplicateTxError),
		Log: "DuplicateTxError: this transaction is already on the Transaction Trie",
	}
	InvalidTxInputFormatWrongx18 = &types.ResponseDeliverTx{
		Code: uint32(codeInvalidTxInputFormat),
		Log: "InvalidInputFormat: mismatch of format (TransactionHeader\\x18Transaction)",
	}
	InvalidTxInputFormatWrongx19 = &types.ResponseDeliverTx{
		Code: uint32(codeInvalidTxInputFormat),
		Log: "InvalidInputFormat: mismatch of format (TransactionHeader\\x19Transaction)",
	}
	MissingContract = &types.ResponseDeliverTx{
		Code: uint32(codeMissingContract),
		Log: "MissingContract: can't find this contract on the Account Trie. Is it deployed correctly?",
	}
)


func DecodeJsonError(err error) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: uint32(codeDecodeJsonError),
		Log: fmt.Sprintf("DecodeJsonError: %v", err),
	}
}

func DecodeTxHeaderError(err error) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: uint32(codeDecodeTxHeaderError),
		Log: fmt.Sprintf("DecodeTxHeaderError: %v", err),
	}
}

func DecodeAccountInfoError(err error) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: uint32(codeDecodeAccountInfoError),
		Log: fmt.Sprintf("DecodeAccountInfoError: %v", err),
	}
}

func ReTrieveTxError(err error) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: uint32(codeReTrieveTxError),
		Log: fmt.Sprintf("ReTrieveTxError: %v", err),
	}
}

func RequestStorageError(err error) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: uint32(codeRequestStorageError),
		Log: fmt.Sprintf("RequestStorageError: %v", err),
	}
}

func UpdateTxTrieError(err error) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: uint32(codeUpdateTxTrieError),
		Log: fmt.Sprintf("UpdateTxTrieError: can't update Transaction Trie, %v", err),
	}
}

func UpdateAccTrieError(err error) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: uint32(codeUpdateAccTrieError),
		Log: fmt.Sprintf("UpdateAccTrieError: can't update Account Trie, %v", err),
	}
}

func CommitTxTrieError(err error) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: uint32(codeCommitTxTrieError),
		Log: fmt.Sprintf("CommitTxTrieError: can't Commit Transaction Trie, %v", err),
	}
}

func CommitAccTrieError(err error) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: uint32(codeCommitAccTrieError),
		Log: fmt.Sprintf("CommitAccTrieError: can't Commit Account Trie, %v", err),
	}
}

func EncodeAccountInfoError(err error) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: uint32(codeEncodeAccountInfoError),
		Log: fmt.Sprintf("EncodeAccountInfoError: %v", err),
	}
}

func InvalidFuncTypeError(err error) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: uint32(codeInvalidFuncType),
		Log: fmt.Sprintf("InvalidFunctionType: %v", err),
	}
}

func ExecContractError(err error) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: uint32(codeExecContractError),
		Log: fmt.Sprintf("ExecContractError: %v", err),
	}
}

func InsufficientBalanceToTransfer(userName string) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: uint32(codeInsufficientBalanceToTransfer),
		Log: fmt.Sprintf("BalanceError: the %v's balance is insufficient", userName),
	}
}

func BalanceOverflow(userName string) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: uint32(codeBalanceOverflow),
		Log: fmt.Sprintf("BalanceError: the %v's balance overflowed", userName),
	}
}

func UpdateBalanceIn(value *math.Uint256) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: uint32(codeUpdateBalanceIn),
		Data: value.Bytes(),
	}
}

func UpdateBalanceOut(value *math.Uint256) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: uint32(codeUpdateBalanceOut),
		Data: value.Bytes(),
	}
}

func DecodeBalanceError() *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: uint32(codeDecodeBalanceError),
		Log: "BalanceError: cannot decode from bytes",
	}
}


func CodeOK() ResponseCode {return codeOK}
func CodeContractPanic() ResponseCode {return codeContractPanic}
func CodeUpdateBalanceIn() ResponseCode {return codeUpdateBalanceIn}
func CodeUpdateBalanceOut() ResponseCode {return codeUpdateBalanceOut}
func CodeTODO() ResponseCode {return codeTODO}
func CodeInvalidTxType() ResponseCode {return codeInvalidTxType}
func CodeDecodeBytesError() ResponseCode {return codeDecodeBytesError}
