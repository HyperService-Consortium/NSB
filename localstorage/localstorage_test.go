package localstorage

import (
	"bytes"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"testing"
)

var (
	accountAddress1          = []byte("www")
	accountAddress2          = []byte("HH")
	acc1_history1            []byte
	acc1_history2            []byte
	acc2_history1            []byte
	acc2_history2            []byte
	map1                     = "myMap1"
	map2                     = "myMap2"
	key1                     = []byte("ke")
	key2                     = []byte("keke")
	key3                     = []byte("kekey")
	acc1_map1value1_history1 = []byte("value1111")
	acc1_map1value1_history2 = []byte("value1112")
	acc1_map1value2_history1 = []byte("value1121")
	acc1_map1value2_history2 = []byte("value1122")
	acc1_map1value3_history1 = []byte("value1131")
	acc1_map1value3_history2 = []byte("value1132")
	acc2_map1value1_history1 = []byte("value1211")
	acc2_map1value1_history2 = []byte("value1212")
	acc2_map1value2_history1 = []byte("value1221")
	acc2_map1value2_history2 = []byte("value1222")
	acc2_map1value3_history1 = []byte("valueA")
	acc2_map1value3_history2 = []byte("valueB")

	acc1_map2value1_history1 = []byte("value2111")
	acc1_map2value1_history2 = []byte("value2112")
	acc1_map2value2_history1 = []byte("value2121")
	acc1_map2value2_history2 = []byte("value2122")
	acc1_map2value3_history1 = []byte("value2131")
	acc1_map2value3_history2 = []byte("value2132")
	acc2_map2value1_history1 = []byte("value2211")
	acc2_map2value1_history2 = []byte("value2212")
	acc2_map2value2_history1 = []byte("value2221")
	acc2_map2value2_history2 = []byte("value2222")
	acc2_map2value3_history1 = []byte("valueA")
	acc2_map2value3_history2 = []byte("valueB")
	testdb                   *leveldb.DB
	err                      error
)

func TestOpenDB(t *testing.T) {
	testdb, err = leveldb.OpenFile("./testdb", nil)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestCreateLocalStorage(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage([]byte("tester"), "00", testdb)
	if err != nil {
		t.Error(err)
		return
	}
	var expVal = []byte("value")
	var getVal []byte
	err = storage.tryUpdate("myMap", []byte("key"), expVal)
	if err != nil {
		t.Error(err)
		return
	}
	getVal, err = storage.tryGet("myMap", []byte("key"))
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, expVal) {
		t.Error("not equal")
		return
	}
	var stoargeRoot []byte
	stoargeRoot, err = storage.Commit()
	if err != nil {
		t.Error(err)
		return
	}

	storage, err = NewLocalStorage([]byte("tester"), stoargeRoot, testdb)
	if err != nil {
		t.Error(err)
		return
	}
	getVal, err = storage.tryGet("myMap", []byte("key"))
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, expVal) {
		t.Error("not equal")
		return
	}
}

func TestAcc1History1(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, "00", testdb)
	if err != nil {
		t.Error(err)
		return
	}

	err = storage.tryUpdate(map1, key1, acc1_map1value1_history1)
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map1, key2, acc1_map1value2_history1)
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map1, key3, acc1_map1value3_history1)
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map2, key1, acc1_map2value1_history1)
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map2, key2, acc1_map2value2_history1)
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map2, key3, acc1_map2value3_history1)
	if err != nil {
		t.Error(err)
		return
	}

	var getVal []byte
	getVal, err = storage.tryGet(map1, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map1value1_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map1value2_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map1value3_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map2value1_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map2value2_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map2value3_history1) {
		t.Error("no equal")
		return
	}

	acc1_history1, err = storage.Commit()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestAcc1History1FromDB(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, acc1_history1, testdb)
	if err != nil {
		t.Error(err)
		return
	}

	var getVal []byte
	getVal, err = storage.tryGet(map1, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map1value1_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map1value2_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map1value3_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map2value1_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map2value2_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map2value3_history1) {
		t.Error("no equal")
		return
	}

	acc1_history1, err = storage.Commit()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestAcc2History1(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, "00", testdb)
	if err != nil {
		t.Error(err)
		return
	}

	err = storage.tryUpdate(map1, key1, acc2_map1value1_history1)
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map1, key2, acc2_map1value2_history1)
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map1, key3, acc2_map1value3_history1)
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map2, key1, acc2_map2value1_history1)
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map2, key2, acc2_map2value2_history1)
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map2, key3, acc2_map2value3_history1)
	if err != nil {
		t.Error(err)
		return
	}

	var getVal []byte
	getVal, err = storage.tryGet(map1, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map1value1_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map1value2_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map1value3_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map2value1_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map2value2_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map2value3_history1) {
		t.Error("no equal")
		return
	}

	acc2_history1, err = storage.Commit()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestAcc1History1FromDBAfterAcc2History1(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, acc1_history1, testdb)
	if err != nil {
		t.Error(err)
		return
	}

	var getVal []byte
	getVal, err = storage.tryGet(map1, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map1value1_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map1value2_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map1value3_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map2value1_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map2value2_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map2value3_history1) {
		t.Error("no equal")
		return
	}

	acc1_history1, err = storage.Commit()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestAcc2History1FromDB(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, acc2_history1, testdb)
	if err != nil {
		t.Error(err)
		return
	}

	var getVal []byte
	getVal, err = storage.tryGet(map1, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map1value1_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map1value2_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map1value3_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map2value1_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map2value2_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map2value3_history1) {
		t.Error("no equal")
		return
	}
}

func TestAcc1History2(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, "00", testdb)
	if err != nil {
		t.Error(err)
		return
	}

	err = storage.tryUpdate(map1, key1, acc1_map1value1_history2)
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map1, key2, acc1_map1value2_history2)
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map1, key3, acc1_map1value3_history2)
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map2, key1, acc1_map2value1_history2)
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map2, key2, acc1_map2value2_history2)
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map2, key3, acc1_map2value3_history2)
	if err != nil {
		t.Error(err)
		return
	}

	var getVal []byte
	getVal, err = storage.tryGet(map1, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map1value1_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map1value2_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map1value3_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map2value1_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map2value2_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map2value3_history2) {
		t.Error("no equal")
		return
	}

	acc1_history2, err = storage.Commit()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestAcc1History2FromDB(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, acc1_history2, testdb)
	if err != nil {
		t.Error(err)
		return
	}

	var getVal []byte
	getVal, err = storage.tryGet(map1, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map1value1_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map1value2_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map1value3_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map2value1_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map2value2_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map2value3_history2) {
		t.Error("no equal")
		return
	}

	acc1_history2, err = storage.Commit()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestAcc2History2(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, "00", testdb)
	if err != nil {
		t.Error(err)
		return
	}

	err = storage.tryUpdate(map1, key1, acc2_map1value1_history2)
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map1, key2, acc2_map1value2_history2)
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map1, key3, acc2_map1value3_history2)
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map2, key1, acc2_map2value1_history2)
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map2, key2, acc2_map2value2_history2)
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map2, key3, acc2_map2value3_history2)
	if err != nil {
		t.Error(err)
		return
	}

	var getVal []byte
	getVal, err = storage.tryGet(map1, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map1value1_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map1value2_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map1value3_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map2value1_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map2value2_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map2value3_history2) {
		t.Error("no equal")
		return
	}

	acc2_history2, err = storage.Commit()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestAcc1History2FromDBAfterAcc2History2(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, acc1_history2, testdb)
	if err != nil {
		t.Error(err)
		return
	}

	var getVal []byte
	getVal, err = storage.tryGet(map1, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map1value1_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map1value2_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map1value3_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map2value1_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map2value2_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map2value3_history2) {
		t.Error("no equal")
		return
	}

	acc1_history2, err = storage.Commit()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestAcc2History2FromDB(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, acc2_history2, testdb)
	if err != nil {
		t.Error(err)
		return
	}

	var getVal []byte
	getVal, err = storage.tryGet(map1, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map1value1_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map1value2_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map1value3_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map2value1_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map2value2_history2) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map2value3_history2) {
		t.Error("no equal")
		return
	}
}

func TestAcc1History1FromDBAfterAcc1History2(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, acc1_history1, testdb)
	if err != nil {
		t.Error(err)
		return
	}

	var getVal []byte
	getVal, err = storage.tryGet(map1, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map1value1_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map1value2_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map1value3_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map2value1_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map2value2_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_map2value3_history1) {
		t.Error("no equal")
		return
	}

	acc1_history1, err = storage.Commit()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestAcc2History1FromDBAfterAcc2History2(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, acc2_history1, testdb)
	if err != nil {
		t.Error(err)
		return
	}

	var getVal []byte
	getVal, err = storage.tryGet(map1, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map1value1_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map1value2_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map1, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map1value3_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map2value1_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map2value2_history1) {
		t.Error("no equal")
		return
	}

	getVal, err = storage.tryGet(map2, key3)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_map2value3_history1) {
		t.Error("no equal")
		return
	}
}

func TestResult(t *testing.T) {
	fmt.Println("acc1_history1", acc1_history1)
	fmt.Println("acc1_history2", acc1_history2)
	fmt.Println("acc2_history1", acc2_history1)
	fmt.Println("acc2_history2", acc2_history2)
}

func TestCloseDB(t *testing.T) {
	testdb.Close()
}
