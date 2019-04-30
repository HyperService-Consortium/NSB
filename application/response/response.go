package response


import (
	"fmt"
	"github.com/tendermint/tendermint/abci/types"
)


type ResponseCode uint32
const ( // base
	CodeOK ResponseCode = 0 + iota
	CodeExecFail
	CodeUnknown
	CodeMissingTxMethod
	CodeMissingContract
	CodeRequestStorageError
	CodeTODO = 99
)

const ( // Decode
	CodeDecodeJsonError ResponseCode = 100 + iota
	CodeDecodeBytesError
	CodeDecodeTxHeaderError
	CodeDecodeAccountInfoError
)

const ( // Transaction
	CodeInvalidTxInputFormat ResponseCode = 200 + iota
	CodeInvalidTxType
	CodeReTrieveTxError
	CodeDuplicateTxError
	CodeUpdateTxTrieError
)


var (
	CheckExecOK = &types.ResponseCheckTx{
		Code: uint32(CodeOK),
	}
	CheckDuplicateTxError = &types.ResponseCheckTx{
		Code: uint32(CodeDuplicateTxError),
		Log: "CheckDuplicateTxError: this transaction is already on the Transaction Trie",
	}
	CheckInvalidTxInputFormatWrongx18 = &types.ResponseCheckTx{
		Code: uint32(CodeInvalidTxInputFormat),
		Log: "CheckInvalidInputFormat: mismatch of format (TransactionHeader\\x18Transaction)",
	}
	CheckInvalidTxInputFormatWrongx19 = &types.ResponseCheckTx{
		Code: uint32(CodeInvalidTxInputFormat),
		Log: "CheckInvalidInputFormat: mismatch of format (TransactionHeader\\x19Transaction)",
	}
	CheckMissingContract = &types.ResponseCheckTx{
		Code: uint32(CodeMissingContract),
		Log: "CheckMissingContract: can't find this contract on the Account Trie. Is it deployed correctly?",
	}

	ExecOK = &types.ResponseDeliverTx{
		Code: uint32(CodeOK),
	}
	DuplicateTxError = &types.ResponseDeliverTx{
		Code: uint32(CodeDuplicateTxError),
		Log: "DuplicateTxError: this transaction is already on the Transaction Trie",
	}
	InvalidTxInputFormatWrongx18 = &types.ResponseDeliverTx{
		Code: uint32(CodeInvalidTxInputFormat),
		Log: "InvalidInputFormat: mismatch of format (TransactionHeader\\x18Transaction)",
	}
	InvalidTxInputFormatWrongx19 = &types.ResponseDeliverTx{
		Code: uint32(CodeInvalidTxInputFormat),
		Log: "InvalidInputFormat: mismatch of format (TransactionHeader\\x19Transaction)",
	}
	MissingContract = &types.ResponseDeliverTx{
		Code: uint32(CodeMissingContract),
		Log: "MissingContract: can't find this contract on the Account Trie. Is it deployed correctly?",
	}
)

func DecodeCheckTxHeaderError(err error) *types.ResponseCheckTx {
	return &types.ResponseCheckTx{
		Code: uint32(CodeDecodeTxHeaderError),
		Log: fmt.Sprintf("DecodeCheckTxHeaderError: %v", err),
	}
}

func DecodeCheckAccountInfoError(err error) *types.ResponseCheckTx {
	return &types.ResponseCheckTx{
		Code: uint32(CodeDecodeAccountInfoError),
		Log: fmt.Sprintf("DecodeCheckAccountInfoError: %v", err),
	}
}

func CheckReTrieveTxError(err error) *types.ResponseCheckTx {
	return &types.ResponseCheckTx{
		Code: uint32(CodeReTrieveTxError),
		Log: fmt.Sprintf("CheckReTrieveTxError: %v", err),
	}
}

func CheckRequestStorageError(err error) *types.ResponseCheckTx {
	return &types.ResponseCheckTx{
		Code: uint32(CodeRequestStorageError),
		Log: fmt.Sprintf("CheckRequestStorageError: %v", err),
	}
}

func CheckUpdateTxTrieError(err error) *types.ResponseCheckTx {
	return &types.ResponseCheckTx{
		Code: uint32(CodeUpdateTxTrieError),
		Log: fmt.Sprintf("CheckUpdateTxTrieError: can't update Transaction Trie, %v", err),
	}
}


func DecodeTxHeaderError(err error) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: uint32(CodeDecodeTxHeaderError),
		Log: fmt.Sprintf("DecodeTxHeaderError: %v", err),
	}
}

func DecodeAccountInfoError(err error) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: uint32(CodeDecodeAccountInfoError),
		Log: fmt.Sprintf("DecodeAccountInfoError: %v", err),
	}
}

func ReTrieveTxError(err error) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: uint32(CodeReTrieveTxError),
		Log: fmt.Sprintf("ReTrieveTxError: %v", err),
	}
}

func RequestStorageError(err error) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: uint32(CodeRequestStorageError),
		Log: fmt.Sprintf("RequestStorageError: %v", err),
	}
}

func UpdateTxTrieError(err error) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: uint32(CodeUpdateTxTrieError),
		Log: fmt.Sprintf("UpdateTxTrieError: can't update Transaction Trie, %v", err),
	}
}