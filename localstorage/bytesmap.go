package localstorage

import (
	"github.com/HyperService-Consortium/NSB/merkmap"
)

type BytesMap struct {
	merk *merkmap.MerkMap
}

func (sto *LocalStorage) NewBytesMap(mapName string) *BytesMap {
	return &BytesMap{
		merk: sto.MakeStorageSlot(mapName),
	}
}

func (bmap *BytesMap) Set(map_offset []byte, value []byte) {
	err := bmap.merk.TryUpdate(map_offset, value)
	if err != nil {
		panic(err)
	}
	return
}

func (bmap *BytesMap) Get(map_offset []byte) []byte {
	bt, err := bmap.merk.TryGet(map_offset)
	if err != nil {
		panic(err)
	}
	return bt
}

func (bmap *BytesMap) Delete(map_offset []byte) {
	err := bmap.merk.TryDelete(map_offset)
	if err != nil {
		panic(err)
	}
	return
}
