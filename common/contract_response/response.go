package response

import (
	"fmt"

	cmn "github.com/HyperService-Consortium/NSB/common"
	"github.com/HyperService-Consortium/NSB/math"
)

const (
	codeOK = 0
)

const (
	codeDecodeJsonError = 1000 + iota
	codeOverFlowError
	codeArithmeticError
	codeInvalidFunctionType
	codeExecContractError
)

func ExecOK(value *math.Uint256) *cmn.ContractCallBackInfo {
	return &cmn.ContractCallBackInfo{
		CodeResponse: uint32(codeOK),
		Value:        value,
		OutFlag:      false,
	}
}

func ExecContractError(err error) *cmn.ContractCallBackInfo {
	return &cmn.ContractCallBackInfo{
		CodeResponse: uint32(codeExecContractError),
		Log:          fmt.Sprintf("ExecContractError: %v", err),
	}
}

func OKAndTransfer(value *math.Uint256) *cmn.ContractCallBackInfo {
	return &cmn.ContractCallBackInfo{
		CodeResponse: uint32(codeOK),
		Value:        value,
		OutFlag:      true,
	}
}

func DecodeJsonError(err error) *cmn.ContractCallBackInfo {
	return &cmn.ContractCallBackInfo{
		CodeResponse: uint32(codeDecodeJsonError),
		Log:          fmt.Sprintf("DecodeJsonError: %v", err),
	}
}

func OverFlowError(info string) *cmn.ContractCallBackInfo {
	return &cmn.ContractCallBackInfo{
		CodeResponse: uint32(codeOverFlowError),
		Log:          "Overflow ...",
		Info:         info,
	}
}

func InvalidFunctionType(info string) *cmn.ContractCallBackInfo {
	return &cmn.ContractCallBackInfo{
		CodeResponse: uint32(codeInvalidFunctionType),
		Log:          "InvalidFunctionType: Unrecognized function",
		Info:         info,
	}
}

func ArithmeticError(info string) *cmn.ContractCallBackInfo {
	return &cmn.ContractCallBackInfo{
		CodeResponse: uint32(codeArithmeticError),
		Log:          "ArithmeticError",
		Info:         info,
	}
}

func AssertTrue(assertValue bool, AssertInfo interface{}) {
	if assertValue == false {
		panic(AssertInfo)
	}
	return
}

func AssertFalse(assertValue bool, AssertInfo interface{}) {
	if assertValue == true {
		panic(AssertInfo)
	}
	return
}

func AssertNil(assertObj interface{}, AssertInfo interface{}) {
	if assertObj != nil {
		panic(AssertInfo)
	}
	return
}

func AssertNotNil(assertObj interface{}, AssertInfo interface{}) {
	if assertObj == nil {
		panic(AssertInfo)
	}
	return
}

func AssertNoErr(assertErr error, AssertInfo interface{}) {
	if assertErr != nil {
		panic(AssertInfo)
	}
	return
}

func CodeOK() uint32 { return codeOK }
