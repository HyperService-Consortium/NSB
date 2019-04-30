package common

type ContractEnvironment struct {
	Storage *localstorage.LocalStorage
	From []byte
	fromInfo *AccountInfo
	ContractAddress []byte
	toInfo *AccountInfo
	Data []byte
	Value *math.Uint256
}