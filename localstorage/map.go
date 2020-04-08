package localstorage

import (
	"github.com/HyperService-Consortium/NSB/merkmap"
)

type Bytable interface {
	Bytes() []byte
}

type Map struct {
	merk *merkmap.MerkMap
}

func (sto *LocalStorage) NewMap(mapName string) *Map {
	return &Map{
		merk: sto.MakeStorageSlot(mapName),
	}
}

func (udMap *Map) Set(Map_offset Bytable, value []byte) {
	err := udMap.merk.TryUpdate(Map_offset.Bytes(), value)
	if err != nil {
		panic(err)
	}
	return
}

func (udMap *Map) Get(Map_offset Bytable) []byte {
	bt, err := udMap.merk.TryGet(Map_offset.Bytes())
	if err != nil {
		panic(err)
	}
	return bt
}

func (udMap *Map) Delete(Map_offset Bytable) {
	err := udMap.merk.TryDelete(Map_offset.Bytes())
	if err != nil {
		panic(err)
	}
	return
}
