package common


type ContractCallBackInfo struct {
	// type responceDeliverTx
	CodeResponse uint32
	Log string
	Info string
	Tags []KVPair
}

func (cb *ContractCallBackInfo) IsErr() bool {
	return cb.CodeResponse != 0
}

func (cb *ContractCallBackInfo) IsOK() bool {
	return cb.CodeResponse == 0
}
