package localstorage

import (
	"github.com/Myriad-Dreamin/NSB/merkmap"
	"encoding/json"
)


type JsonBytesMap struct {
	merk *merkmap.MerkMap
}

func (sto *LocalStorage) NewJsonBytesMap(mapName string) (*JsonBytesMap) {
	return &JsonBytesMap{
		merk: sto.makeStorageSlot(mapName),
	}
}

func (jsonBytesMap *JsonBytesMap) Set(Map_offset []byte, value interface{}) {
	bt, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	err := jsonBytesMap.merk.TryUpdate(Map_offset, bt)
	if err != nil {
		panic(err)
	}
	return
}

func (jsonBytesMap *JsonBytesMap) Get(Map_offset []byte, value interface{}) {
	bt, err := jsonBytesMap.merk.TryGet(Map_offset)
	if err != nil {
		panic(err)
	}

	err := json.Unmarshal(bt, value)
	if err != nil {
		panic(err)
	}
	return
}

func (jsonBytesMap *JsonBytesMap) Delete(Map_offset []byte) {
	err := jsonBytesMap.merk.TryDelete(Map_offset)
	if err != nil {
		panic(err)
	}
	return
}