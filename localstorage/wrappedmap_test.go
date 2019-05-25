package localstorage

import (
	"bytes"
	"github.com/syndtr/goleveldb/leveldb"
	"testing"
)

type MyString string

func (mstr MyString) Bytes() []byte {
	return []byte(string(mstr))
}

type MyBytes []byte

func (mbt MyBytes) Bytes() []byte {
	return []byte(mbt)
}

func TestCreateMap(t *testing.T) {
	var testdb *leveldb.DB
	var err error
	testdb, err = leveldb.OpenFile("./testdb", nil)
	if err != nil {
		t.Error(err)
		return
	}
	defer testdb.Close()
	var storage *LocalStorage
	storage, err = NewLocalStorage([]byte("tester"), "00", testdb)
	if err != nil {
		t.Error(err)
		return
	}

	udmap := storage.NewMap("myMap")

	var mykey = MyString("www")
	var mykey2 = MyBytes("wwwww")
	var expVal = []byte("exp")

	udmap.Set(mykey, expVal)
	udmap.Set(mykey2, expVal)

	var getVal []byte
	getVal = udmap.Get(mykey)
	if !bytes.Equal(getVal, expVal) {
		t.Error("no equal")
		return
	}
	getVal = udmap.Get(mykey2)
	if !bytes.Equal(getVal, expVal) {
		t.Error("no equal")
		return
	}
}
