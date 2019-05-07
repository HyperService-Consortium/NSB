package localstorage

import (
	"bytes"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"testing"
)

/*
acc1   history1   history2   history3   histr0y4
key1     value1     value2         []         []
key2         []     value1     value2         []

acc2   history1   history2   history3   history4
key1         []     value1         []         []
key2     value1     value2     value2         []
*/

var (
	map0              = "map"
	acc1_key1_history = [][]byte{
		[]byte("value1"),
		[]byte("value2"),
		[]byte(""),
		[]byte(""),
	}

	acc1_key2_history = [][]byte{
		[]byte(""),
		[]byte("value1"),
		[]byte("value2"),
		[]byte(""),
	}

	acc2_key1_history = [][]byte{
		[]byte(""),
		[]byte("value1"),
		[]byte(""),
		[]byte(""),
	}

	acc2_key2_history = [][]byte{
		[]byte("value1"),
		[]byte("value2"),
		[]byte("value2"),
		[]byte(""),
	}

	acc1_history3 []byte
	acc1_history4 []byte
	acc2_history3 []byte
	acc2_history4 []byte
)

func CheckAcc1History1(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, acc1_history1, testdb)
	if err != nil {
		t.Error(err)
		return
	}
	var getVal []byte
	getVal, err = storage.tryGet(map0, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_key1_history[0]) {
		t.Error("no equal")
	}
	getVal, err = storage.tryGet(map0, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_key2_history[0]) {
		t.Error("no equal")
	}
}

func CheckAcc1History2(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, acc1_history2, testdb)
	if err != nil {
		t.Error(err)
		return
	}
	var getVal []byte
	getVal, err = storage.tryGet(map0, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_key1_history[1]) {
		t.Error("no equal")
	}
	getVal, err = storage.tryGet(map0, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_key2_history[1]) {
		t.Error("no equal")
	}
}

func CheckAcc1History3(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, acc1_history3, testdb)
	if err != nil {
		t.Error(err)
		return
	}
	var getVal []byte
	getVal, err = storage.tryGet(map0, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_key1_history[2]) {
		t.Error("no equal")
	}
	getVal, err = storage.tryGet(map0, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_key2_history[2]) {
		t.Error("no equal")
	}
}

func CheckAcc1History4(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, acc1_history4, testdb)
	if err != nil {
		t.Error(err)
		return
	}
	var getVal []byte
	getVal, err = storage.tryGet(map0, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_key1_history[3]) {
		t.Error("no equal")
	}
	getVal, err = storage.tryGet(map0, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc1_key2_history[3]) {
		t.Error("no equal")
	}
}

func CheckAcc2History1(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, acc2_history1, testdb)
	if err != nil {
		t.Error(err)
		return
	}
	var getVal []byte
	getVal, err = storage.tryGet(map0, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_key1_history[0]) {
		t.Error("no equal")
	}
	getVal, err = storage.tryGet(map0, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_key2_history[0]) {
		t.Error("no equal")
	}
}

func CheckAcc2History2(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, acc2_history2, testdb)
	if err != nil {
		t.Error(err)
		return
	}
	var getVal []byte
	getVal, err = storage.tryGet(map0, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_key1_history[1]) {
		t.Error("no equal")
	}
	getVal, err = storage.tryGet(map0, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_key2_history[1]) {
		t.Error("no equal")
	}
}

func CheckAcc2History3(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, acc2_history3, testdb)
	if err != nil {
		t.Error(err)
		return
	}
	var getVal []byte
	getVal, err = storage.tryGet(map0, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_key1_history[2]) {
		t.Error("no equal")
	}
	getVal, err = storage.tryGet(map0, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_key2_history[2]) {
		t.Error("no equal")
	}
}

func CheckAcc2History4(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, acc2_history4, testdb)
	if err != nil {
		t.Error(err)
		return
	}
	var getVal []byte
	getVal, err = storage.tryGet(map0, key1)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_key1_history[3]) {
		t.Error("no equal")
	}
	getVal, err = storage.tryGet(map0, key2)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(getVal, acc2_key2_history[3]) {
		t.Error("no equal")
	}
}

func TestConflictOpenDB(t *testing.T) {
	testdb, err = leveldb.OpenFile("./testdb", nil)
	if err != nil {
		t.Error(err)
		return
	}
}

/*
acc1   history1   history2   history3   histr0y4
key1     value1     value2         []         []
key2         []     value1     value2         []

acc2   history1   history2   history3   history4
key1         []     value1         []         []
key2     value1     value2     value2         []
*/

func TestAcc1MakeHistory1(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, "00", testdb)
	if err != nil {
		t.Error(err)
		return
	}

	err = storage.tryUpdate(map0, key1, acc1_key1_history[0])
	if err != nil {
		t.Error(err)
		return
	}
	acc1_history1, err = storage.Commit()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestHistoryA(t *testing.T) {
	CheckAcc1History1(t)
}

/*
acc1   history1   history2   history3   histr0y4
key1     value1     value2         []         []
key2         []     value1     value2         []

acc2   history1   history2   history3   history4
key1         []     value1         []         []
key2     value1     value2     value2         []
*/

func TestAcc2MakeHistory1(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, "00", testdb)
	if err != nil {
		t.Error(err)
		return
	}

	err = storage.tryUpdate(map0, key2, acc2_key2_history[0])
	if err != nil {
		t.Error(err)
		return
	}

	acc2_history1, err = storage.Commit()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestHistoryB(t *testing.T) {
	CheckAcc1History1(t)
	CheckAcc2History1(t)
}

/*
acc1   history1   history2   history3   histr0y4
key1     value1     value2         []         []
key2         []     value1     value2         []

acc2   history1   history2   history3   history4
key1         []     value1         []         []
key2     value1     value2     value2         []
*/

func TestAcc2MakeHistory2(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, acc2_history1, testdb)
	if err != nil {
		t.Error(err)
		return
	}

	err = storage.tryUpdate(map0, key1, acc2_key1_history[1])
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map0, key2, acc2_key2_history[1])
	if err != nil {
		t.Error(err)
		return
	}

	acc2_history2, err = storage.Commit()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestHistoryC(t *testing.T) {
	CheckAcc1History1(t)
	CheckAcc2History1(t)
	CheckAcc2History2(t)
}

/*
acc1   history1   history2   history3   histr0y4
key1     value1     value2         []         []
key2         []     value1     value2         []

acc2   history1   history2   history3   history4
key1         []     value1         []         []
key2     value1     value2     value2         []
*/

func TestAcc1MakeHistory2(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, acc1_history1, testdb)
	if err != nil {
		t.Error(err)
		return
	}

	err = storage.tryUpdate(map0, key1, acc1_key1_history[1])
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map0, key2, acc1_key2_history[1])
	if err != nil {
		t.Error(err)
		return
	}

	acc1_history2, err = storage.Commit()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestHistoryD(t *testing.T) {
	CheckAcc1History1(t)
	CheckAcc2History1(t)
	CheckAcc2History2(t)
	CheckAcc1History2(t)
}

/*
acc1   history1   history2   history3   histr0y4
key1     value1     value2         []         []
key2         []     value1     value2         []

acc2   history1   history2   history3   history4
key1         []     value1         []         []
key2     value1     value2     value2         []
*/

func TestAcc1MakeHistory3(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, acc1_history2, testdb)
	if err != nil {
		t.Error(err)
		return
	}

	err = storage.tryDelete(map0, key1)
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map0, key2, acc1_key2_history[2])
	if err != nil {
		t.Error(err)
		return
	}

	acc1_history3, err = storage.Commit()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestHistoryE(t *testing.T) {
	CheckAcc1History1(t)
	CheckAcc2History1(t)
	CheckAcc2History2(t)
	CheckAcc1History2(t)
	CheckAcc1History3(t)
}

/*
acc1   history1   history2   history3   histr0y4
key1     value1     value2         []         []
key2         []     value1     value2         []

acc2   history1   history2   history3   history4
key1         []     value1         []         []
key2     value1     value2     value2         []
*/

func TestAcc2MakeHistory3(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, acc2_history2, testdb)
	if err != nil {
		t.Error(err)
		return
	}

	err = storage.tryDelete(map0, key1)
	if err != nil {
		t.Error(err)
		return
	}
	err = storage.tryUpdate(map0, key2, acc2_key2_history[2])
	if err != nil {
		t.Error(err)
		return
	}

	acc2_history3, err = storage.Commit()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestHistoryF(t *testing.T) {
	CheckAcc1History1(t)
	CheckAcc2History1(t)
	CheckAcc2History2(t)
	CheckAcc1History2(t)
	CheckAcc1History3(t)
	CheckAcc2History3(t)
}

/*
acc1   history1   history2   history3   histr0y4
key1     value1     value2         []         []
key2         []     value1     value2         []

acc2   history1   history2   history3   history4
key1         []     value1         []         []
key2     value1     value2     value2         []
*/

func TestAcc2MakeHistory4(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, acc2_history3, testdb)
	if err != nil {
		t.Error(err)
		return
	}

	err = storage.tryDelete(map0, key2)
	if err != nil {
		t.Error(err)
		return
	}

	acc2_history4, err = storage.Commit()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestHistoryG(t *testing.T) {
	CheckAcc1History1(t)
	CheckAcc2History1(t)
	CheckAcc2History2(t)
	CheckAcc1History2(t)
	CheckAcc1History3(t)
	CheckAcc2History3(t)
	CheckAcc2History4(t)
}

/*
acc1   history1   history2   history3   histr0y4
key1     value1     value2         []         []
key2         []     value1     value2         []

acc2   history1   history2   history3   history4
key1         []     value1         []         []
key2     value1     value2     value2         []
*/

func TestAcc1MakeHistory4(t *testing.T) {
	var storage *LocalStorage
	storage, err = NewLocalStorage(accountAddress1, acc1_history3, testdb)
	if err != nil {
		t.Error(err)
		return
	}

	err = storage.tryDelete(map0, key2)
	if err != nil {
		t.Error(err)
		return
	}

	acc1_history4, err = storage.Commit()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestHistoryH(t *testing.T) {
	CheckAcc1History1(t)
	CheckAcc2History1(t)
	CheckAcc2History2(t)
	CheckAcc1History2(t)
	CheckAcc1History3(t)
	CheckAcc2History3(t)
	CheckAcc2History4(t)
	CheckAcc1History4(t)
}

func TestConflictResult(t *testing.T) {
	fmt.Println("acc1_history1", acc1_history1)
	fmt.Println("acc1_history2", acc1_history2)
	fmt.Println("acc1_history3", acc1_history3)
	fmt.Println("acc1_history4", acc1_history4)
	fmt.Println("acc2_history1", acc2_history1)
	fmt.Println("acc2_history2", acc2_history2)
	fmt.Println("acc2_history3", acc2_history3)
	fmt.Println("acc2_history4", acc2_history4)
}

func TestConflictCloseDB(t *testing.T) {
	testdb.Close()
}
