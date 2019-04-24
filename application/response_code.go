package nsb


type ResponseCode uint32
const (
	CodeOK ResponseCode = 0 + iota
	CodeFail
	CodeUnknown
	CodeMissing
	CodeTODO
)
