package response


import (
	"fmt"
	cmn "github.com/Myriad-Dreamin/NSB/common"
)


const (
	codeOK = 0
)


const (
	codeDecodeJsonError = 1000 + iota
	codeOverFlowError
	codeArithmaticError
	codeInvalidFunctionType
)


func ExecOK() *cmn.ContractCallBackInfo {
	return &cmn.ContractCallBackInfo {
		CodeResponse: uint32(codeOK),
	}
}


func DecodeJsonError(err error) *cmn.ContractCallBackInfo {
	return &cmn.ContractCallBackInfo {
		CodeResponse: uint32(codeDecodeJsonError),
		Log: fmt.Sprintf("DecodeJsonError: %v", err),
	}
}

func OverFlowError(info string) *cmn.ContractCallBackInfo {
	return &cmn.ContractCallBackInfo {
		CodeResponse: uint32(codeOverFlowError),
		Log: "Overflow ...",
		Info: info,
	}
}

func InvalidFunctionType(info string) *cmn.ContractCallBackInfo {
	return &cmn.ContractCallBackInfo {
		CodeResponse: uint32(codeInvalidFunctionType),
		Log: "InvalidFunctionType: Unrecognized function",
		Info: info,
	}
}


func ArithmaticError(info string) *cmn.ContractCallBackInfo {
	return &cmn.ContractCallBackInfo {
		CodeResponse: uint32(codeArithmaticError),
		Log: "ArithmaticError",
		Info: info,
	}
}


func Assert(assertValue ...bool, AssertInfo interface{}) {
	for _, checkOK := range assertValue {
		if checkOK == false {
			panic(AssertInfo)
		}
	}
	return
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
	if assertValue != nil {
		panic(AssertInfo)
	}
	return
}

func AssertNotNil(assertObj interface{}, AssertInfo interface{}) {
	if assertValue == nil {
		panic(AssertInfo)
	}
	return
}

func AssertNoErr(assertObj interface{}, AssertInfo interface{}) {
	if assertValue == nil {
		panic(AssertInfo)
	}
	return
}