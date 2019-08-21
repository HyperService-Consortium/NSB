package main

import "C"
import "unsafe"
import wallet "github.com/HyperService-Consortium/NSB/account"

//export CDLL_NewLevelDBHandler
func CDLL_NewLevelDBHandler(dbpath *C.char) C.int {
	return C.int(wallet.NewLevelDBHandler((*wallet.Export_C_Char)(dbpath)))
}

//export CDLL_CloseDB
func CDLL_CloseDB(dbptr C.int) {
	wallet.CloseDB((wallet.Export_C_Int)(dbptr))
}

//export CDLL_PreCheckWallet
func CDLL_PreCheckWallet(wltptr C.int) C.int {
	return C.int(wallet.PreCheckWallet((wallet.Export_C_Int)(wltptr)))
}

//export CDLL_NewWalletHandlerFromDB
func CDLL_NewWalletHandlerFromDB(dbptr C.int, wltname *C.char) C.int {
	return C.int(wallet.NewWalletHandlerFromDB((wallet.Export_C_Int)(dbptr), (*wallet.Export_C_Char)(wltname)))
}

//export CDLL_NewWalletHandler
func CDLL_NewWalletHandler(dbptr C.int, wltname *C.char) C.int {
	return C.int(wallet.NewWalletHandler((wallet.Export_C_Int)(dbptr), (*wallet.Export_C_Char)(wltname)))
}

//export CDLL_WalletAddress
func CDLL_WalletAddress(wltptr C.int, idx C.int) unsafe.Pointer {
	return wallet.WalletAddress((wallet.Export_C_Int)(wltptr), (wallet.Export_C_Int)(idx))
}

//export CDLL_WalletSign
func CDLL_WalletSign(wltptr C.int, idx C.int, msg unsafe.Pointer, msgSize C.int) unsafe.Pointer {
	return wallet.WalletSign((wallet.Export_C_Int)(wltptr), (wallet.Export_C_Int)(idx), msg, (wallet.Export_C_Int)(msgSize))
}

//export CDLL_WalletSignHash
func CDLL_WalletSignHash(wltptr C.int, idx C.int, msgHash unsafe.Pointer) unsafe.Pointer {
	return wallet.WalletSignHash((wallet.Export_C_Int)(wltptr), (wallet.Export_C_Int)(idx), msgHash)
}

//export CDLL_WalletVerifyByRaw
func CDLL_WalletVerifyByRaw(wltptr C.int, idx C.int, signature unsafe.Pointer, msg unsafe.Pointer, msgSize C.int) C.int {
	return C.int(wallet.WalletVerifyByRaw((wallet.Export_C_Int)(wltptr), (wallet.Export_C_Int)(idx), signature, msg, (wallet.Export_C_Int)(msgSize)))
}

//export CDLL_WalletVerifyByHash
func CDLL_WalletVerifyByHash(wltptr C.int, idx C.int, signature unsafe.Pointer, msgHash unsafe.Pointer) C.int {
	return C.int(wallet.WalletVerifyByHash((wallet.Export_C_Int)(wltptr), (wallet.Export_C_Int)(idx), signature, msgHash))
}

func main() {}
