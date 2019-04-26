package localstorage

import (
	"testing"
	"github.com/syndtr/goleveldb/leveldb"
	"bytes"
)

type MyString string

func (mstr MyString) Bytes() []byte {
	return []byte(string(mstr))
}

type MyBytes []byte

func (mbt MyBytes) Bytes() []byte {
	return []byte(mbt)
}

func TestCreateUserDefinedMap(t *testing.T) {
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

	udmap := storage.NewUserDefinedMap("myMap")


	var mykey = MyString("www")
	var mykey2 = MyBytes("wwwww")
	var expVal = []byte("exp")
	
	err = udmap.Set(mykey, expVal)
	if err != nil {
		t.Error(err)
		return
	}
	err = udmap.Set(mykey2, expVal)
	if err != nil {
		t.Error(err)
		return
	}

	var getVal []byte
	getVal, err = udmap.Get(mykey)
	if !bytes.Equal(getVal, expVal) {
		t.Error("no equal")
		return
	}
	getVal, err = udmap.Get(mykey2)
	if !bytes.Equal(getVal, expVal) {
		t.Error("no equal")
		return
	}
}