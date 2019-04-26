package nsb


type ResponseCode uint32
const ( // base
	CodeOK ResponseCode = 0 + iota
	CodeExecFail
	CodeUnknown
	CodeMissingTxMethod
	CodeMissingContract
	CodeTODO = 99
)

const ( // Decode
	CodeDecodeJsonError ResponseCode = 100 + iota
	CodeDecodeTxHeaderError
)

const ( // Transaction
	CodeInvalidTxInputFormat ResponseCode = 200 + iota
	CodeInvalidTxType
)

cosnt ( // ISC
	codeISCExecFail ResponseCode = 300 + iota
)


var (
	invalidTxInputFormatWrongx18 = types.ResponseDeliverTx{
		Code: uint32(CodeInvalidTxInputFormat),
		Log: "InvalidInputFormat: mismatch of format (TransactionHeader\\x18Transaction)",
	}
)

func DecodeTxHeaderError(err error) types.ResponseDeliverTx {
	return return types.ResponseDeliverTx{
		Code: uint32(CodeDecodeTxHeaderError),
		Log: fmt.Sprintf("%v", err),
	}
}