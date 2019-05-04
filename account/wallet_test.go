package account

import (
	"testing"
	// "bytes"
	// "encoding/hex"
	// "github.com/syndtr/goleveldb/leveldb"
)

func AbortedTestWallet(t *testing.T) {
	// db, err := leveldb.OpenFile("./testdb", nil)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// defer db.Close()
	// acc := NewAccount([]byte("account:myd;pawd:123456"))
	// wt := NewWallet(db, acc)
	// wt.Save()

	// var wtFromDB *Wallet
	// wtFromDB, err = ReadWallet(db, acc.PublicKey)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// if !bytes.Equal(wtFromDB.Acc.PrivateKey, wt.Acc.PrivateKey) {
	// 	t.Error("no equal")
	// 	return
	// }

	// wtFromDB, err = ReadWallet(db, hex.EncodeToString(acc.PublicKey))
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// if !bytes.Equal(wtFromDB.Acc.PrivateKey, wt.Acc.PrivateKey) {
	// 	t.Error("no equal")
	// 	return
	// }
}

func TestWallet(t *testing.T) {

}
