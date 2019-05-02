package response


import (
	"fmt"
	"github.com/tendermint/tendermint/abci/types"
)


type ResponseCode uint32
const ( // base
	codeOK ResponseCode = 0 + iota
	codeExecFail
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
)

const ( // Transaction
	codeInvalidTxInputFormat ResponseCode = 200 + iota
	codeInvalidTxType
	codeReTrieveTxError
	codeDuplicateTxError
)

const ( // Contract
	codeContractPanic ResponseCode = 200 + iota
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
