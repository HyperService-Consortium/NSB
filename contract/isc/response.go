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
)

func DecodeJsonError(err error) *cmn.ContractCallBackInfo {
	return &cmn.ContractCallBackInfo{
		CodeResponse: int32(CodeDecodeJsonError),
		Log: fmt.Sprintf("DecodeJsonError: %v", err),
	}
}

func OverFlowError(info string) *cmn.ContractCallBackInfo {
	return &cmn.ContractCallBackInfo{
		CodeResponse: int32(CodeOverFlowError),
		Log: "Overflow ...",
		Info: info,
	}
}
