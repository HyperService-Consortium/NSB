package localstorage

import (
	_ "bytes"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"testing"
	"unsafe"
)

func TestReinterpret(t *testing.T) {
	var bt = []byte{3, 2, 0, 0, 0, 0, 0, 0}
	fmt.Println(*(*uint64)(unsafe.Pointer(&bt[0])))

	var xx = uint64(515)
	var btx = []byte{0,0,0,0,0,0,0,0}
	*(*uint64)(unsafe.Pointer(&btx[0])) = *(*uint64)(unsafe.Pointer(&xx))
	fmt.Println(btx)
	// 0000 0011 0000 0010
}

func TestOpenDBBbytes(t *testing.T) {
	testdb, err = leveldb.OpenFile("./testdb", nil)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestBytesArraySetAndGet(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, "00", testdb)

	barr := storage.NewBytesArray("testArr")
	barr.Append([]byte("abcx"))
	barr.Append([]byte("cbax"))
	barr.Append([]byte("2333"))
	barr.Append([]byte("3211"))
	barr.Append([]byte("5555"))
	barr.Append([]byte("1231"))
	barr.Set(1, []byte("changed"))

	
	acc1_history1, err = storage.Commit()
	if err != nil {
		t.Error(err)
		return
	}

	storage, err = NewLocalStorage(accountAddress1, acc1_history1, testdb)
	barr = storage.NewBytesArray("testArr")
	fmt.Println(string(barr.Get(0)))
	fmt.Println(string(barr.Get(1)))
	fmt.Println(string(barr.Get(2)))
	fmt.Println(string(barr.Get(3)))
	fmt.Println(string(barr.Get(4)))
	fmt.Println(string(barr.Get(5)))
	fmt.Println(barr.length)
}

func TestCloseDBBbytes(t *testing.T) {
	testdb.Close()
}
