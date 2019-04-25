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
)

const ( // Transaction
	CodeInvalidTxInputFormat ResponseCode = 200 + iota
	CodeInvalidTxType
)

cosnt ( // ISC
	codeISCExecFail ResponseCode = 300 + iota
)