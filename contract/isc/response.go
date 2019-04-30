package isc

import (
	cmn "github.com/Myriad-Dreamin/NSB/common"
)


const (
	CodeOK = 0
)

const ( // ISC
	codeISCExecFail int32 = 1000 + iota
	CodeDecodeJsonError
	CodeOverFlowError
	CodeInvalidFunctionType
)

func DecodeJsonError(err error) *cmn.ContractCallBackInfo {
	return &cmn.ContractCallBackInfo{
		CodeResponse: uint32(CodeDecodeJsonError),
		Log: fmt.Sprintf("DecodeJsonError: %v", err),
	}
}

func OverFlowError(info string) *cmn.ContractCallBackInfo {
	return &cmn.ContractCallBackInfo{
		CodeResponse: uint32(CodeOverFlowError),
		Log: "Overflow ...",
		Info: info,
	}
}

func InvalidFunctionType(info string) *cmn.ContractCallBackInfo {
	return &cmn.ContractCallBackInfo{
		CodeResponse: uint32(CodeInvalidFunctionType),
		Log: "InvalidFunctionType: Unrecognized function",
		Info: info,
	}
}