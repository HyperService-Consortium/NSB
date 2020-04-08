package localstorage

import (
	"github.com/HyperService-Consortium/NSB/merkmap"
)

type StringMap struct {
	merk *merkmap.MerkMap
}

func (sto *LocalStorage) NewStringMap(mapName string) *StringMap {
	return &StringMap{
		merk: sto.MakeStorageSlot(mapName),
	}
}

func (smap *StringMap) Set(map_offset string, value []byte) {
	err := smap.merk.TryUpdate([]byte(map_offset), value)
	if err != nil {
		panic(err)
	}
	return
}

func (smap *StringMap) Get(map_offset string) []byte {
	bt, err := smap.merk.TryGet([]byte(map_offset))
	if err != nil {
		panic(err)
	}
	return bt
}

func (smap *StringMap) Delete(map_offset string) {
	err := smap.merk.TryDelete([]byte(map_offset))
	if err != nil {
		panic(err)
	}
	return
}
