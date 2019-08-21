package account

import (
	"C"
	"fmt"
	"github.com/HyperService-Consortium/NSB/crypto"
	"github.com/syndtr/goleveldb/leveldb"
	eddsa "golang.org/x/crypto/ed25519"
	"unsafe"
)

const (
	// max number of onload-dbs
	MXONLOADDB = 64
)

const (
	CodeInvalidDBPtr = -1 - iota
	CodeInvalidWalletPtr
	CodeInvalidIndex
	CodeIOError
)

var (
	//onload-db array
	dbPacket = make([]*leveldb.DB, 0, MXONLOADDB)
	//dbfree reference
	dbpi = 0
	//active wallet array
	wltPacket = make([]*Wallet, 0)
)

type Export_C_Char C.char
type Export_C_Int C.int

//export NewLevelDBHandler
func NewLevelDBHandler(dbpath *Export_C_Char) C.int {
	db, err := leveldb.OpenFile(C.GoString((*C.char)(dbpath)), nil)
	if err != nil {
		fmt.Println("link error")
		fmt.Println(err)
		return C.int(CodeInvalidDBPtr)
	} else {
		dbPacket = append(dbPacket, db)
		dbpi++
		return C.int(dbpi - 1)
	}
}

//export CloseDB
func CloseDB(dbptr Export_C_Int) {
	db := dbPacket[dbptr]
	db.Close()
	dbPacket[dbptr] = nil
}

func getDB(dbptr int) *leveldb.DB {
	if dbptr >= dbpi || dbptr < 0 {
		return nil
	}
	return dbPacket[dbptr]
}

func getWallet(wltptr int) *Wallet {
	if wltptr >= len(wltPacket) || wltptr < 0 {
		return nil
	}
	return wltPacket[wltptr]
}

//export PreCheckWallet
func PreCheckWallet(wltptr Export_C_Int) C.int {
	if int(wltptr) >= len(wltPacket) || wltptr < 0 {
		return 0
	}
	if wltPacket[wltptr] != nil {
		return 1
	} else {
		return 0
	}
}

//export NewWalletHandlerFromDB
func NewWalletHandlerFromDB(dbptr Export_C_Int, wltname *Export_C_Char) C.int {
	db, name := getDB(int(dbptr)), C.GoString((*C.char)(wltname))
	if db == nil {
		return C.int(CodeInvalidDBPtr)
	}
	wlt, err := ReadWallet(db, name)
	if err != nil {
		fmt.Println(err)
		return C.int(CodeIOError)
	}
	wltPacket = append(wltPacket, wlt)
	return C.int(len(wltPacket) - 1)
}

//export NewWalletHandler
func NewWalletHandler(dbptr Export_C_Int, wltname *Export_C_Char) C.int {
	db, name := getDB(int(dbptr)), C.GoString((*C.char)(wltname))
	if db == nil {
		return C.int(CodeInvalidDBPtr)
	}
	wltPacket = append(wltPacket, NewWallet(db, name))
	return C.int(len(wltPacket) - 1)
}

//export WalletAddress
func WalletAddress(wltptr Export_C_Int, idx Export_C_Int) unsafe.Pointer {
	wlt := getWallet(int(wltptr))
	if wlt == nil {
		return unsafe.Pointer(nil)
	}
	if len(wlt.Acc) <= int(idx) || idx < 0 {
		return unsafe.Pointer(nil)
	}

	return C.CBytes([]byte(wlt.Acc[idx].PublicKey))
}

//export WalletSign
func WalletSign(wltptr Export_C_Int, idx Export_C_Int, msg unsafe.Pointer, msgSize Export_C_Int) unsafe.Pointer {
	wlt := getWallet(int(wltptr))
	if wlt == nil {
		return unsafe.Pointer(nil)
	}
	if len(wlt.Acc) <= int(idx) || idx < 0 {
		return unsafe.Pointer(nil)
	}
	msgHash := crypto.Sha512(sign_header, C.GoBytes(msg, C.int(msgSize)))
	signature := eddsa.Sign(wlt.Acc[idx].PrivateKey, msgHash)
	return C.CBytes(signature)
}

//export WalletSignHash
func WalletSignHash(wltptr Export_C_Int, idx Export_C_Int, msgHash unsafe.Pointer) unsafe.Pointer {
	wlt := getWallet(int(wltptr))
	if wlt == nil {
		return unsafe.Pointer(nil)
	}
	if len(wlt.Acc) <= int(idx) || idx < 0 {
		return unsafe.Pointer(nil)
	}
	signature := eddsa.Sign(wlt.Acc[idx].PrivateKey, C.GoBytes(msgHash, 64))
	return C.CBytes(signature)
}

//export WalletVerifyByRaw
func WalletVerifyByRaw(wltptr Export_C_Int, idx Export_C_Int, signature unsafe.Pointer, msg unsafe.Pointer, msgSize Export_C_Int) C.int {
	wlt := getWallet(int(wltptr))
	if wlt == nil {
		return CodeInvalidWalletPtr
	}
	if len(wlt.Acc) <= int(idx) || idx < 0 {
		return CodeInvalidIndex
	}
	msgHash := crypto.Sha512(sign_header, C.GoBytes(msg, C.int(msgSize)))
	checkOK := eddsa.Verify(wlt.Acc[idx].PublicKey, msgHash, C.GoBytes(signature, 64))
	if checkOK {
		return 1
	} else {
		return 0
	}
}

//export WalletVerifyByHash
func WalletVerifyByHash(wltptr Export_C_Int, idx Export_C_Int, signature unsafe.Pointer, msgHash unsafe.Pointer) C.int {
	wlt := getWallet(int(wltptr))
	if wlt == nil {
		return CodeInvalidWalletPtr
	}
	if len(wlt.Acc) <= int(idx) || idx < 0 {
		return CodeInvalidIndex
	}
	checkOK := eddsa.Verify(wlt.Acc[idx].PublicKey, C.GoBytes(msgHash, 64), C.GoBytes(signature, 64))
	if checkOK {
		return 1
	} else {
		return 0
	}
}
