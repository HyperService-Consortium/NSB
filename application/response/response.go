package response


import (
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
	CodeUpdateTxTreeError
)

const ( // ISC
	codeISCExecFail ResponseCode = 300 + iota
)


var (
	DuplicateTxError = types.ResponseDeliverTx{
		Code: uint32(CodeDuplicateTxError),
		Log: "DuplicateTxError: this transaction is already on the Transaction Trie",
	}
	InvalidTxInputFormatWrongx18 = types.ResponseDeliverTx{
		Code: uint32(CodeInvalidTxInputFormat),
		Log: "InvalidInputFormat: mismatch of format (TransactionHeader\\x18Transaction)",
	}
	InvalidTxInputFormatWrongx19 = types.ResponseDeliverTx{
		Code: uint32(CodeInvalidTxInputFormat),
		Log: "InvalidInputFormat: mismatch of format (TransactionHeader\\x19Transaction)",
	}
	MissingContract = types.ResponseDeliverTx{
		Code: uint32(CodeMissingContract),
		Log: "MissingContract: can't find this contract on the Account Trie. Is it deployed correctly?",
	}
	UpdateTxTreeError = types.ResponseDeliverTx{
		Code: uint32(CodeUpdateTxTreeError),
		Log: "UpdateTxTreeError: can't update Transaction Trie",
	}
)

func DecodeTxHeaderError(err error) types.ResponseDeliverTx {
	return types.ResponseDeliverTx{
		Code: uint32(CodeDecodeTxHeaderError),
		Log: fmt.Sprintf("%v", err),
	}
}

func DecodeAccountInfoError(err error) types.ResponseDeliverTx {
	return types.ResponseDeliverTx{
		Code: uint32(CodeDecodeAccountInfoError),
		Log: fmt.Sprintf("%v", err),
	}
}

func ReTrieveTxError(err error) types.ResponseDeliverTx {
	return types.ResponseDeliverTx{
		Code: uint32(CodeReTrieveTxError),
		Log: fmt.Sprintf("%v", err),
	}
}

func RequestStorageError(err error) types.ResponseDeliverTx {
	return types.ResponseDeliverTx{
		Code: uint32(CodeRequestStorageError),
		Log: fmt.Sprintf("%v", err),
	}
}