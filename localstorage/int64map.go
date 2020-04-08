package localstorage

import (
	"github.com/HyperService-Consortium/NSB/merkmap"
	"github.com/HyperService-Consortium/NSB/util"
)

type Int64Map struct {
	merk *merkmap.MerkMap
}

func (sto *LocalStorage) NewInt64Map(mapName string) *Int64Map {
	return &Int64Map{
		merk: sto.MakeStorageSlot(mapName),
	}
}

func (i64map *Int64Map) Set(map_offset int64, value []byte) {
	err := i64map.merk.TryUpdate(util.Int64ToBytes(map_offset), value)
	if err != nil {
		panic(err)
	}
	return
}

func (i64map *Int64Map) Get(map_offset int64) []byte {
	bt, err := i64map.merk.TryGet(util.Int64ToBytes(map_offset))
	if err != nil {
		panic(err)
	}
	return bt
}

func (i64map *Int64Map) Delete(map_offset int64) {
	err := i64map.merk.TryDelete(util.Int64ToBytes(map_offset))
	if err != nil {
		panic(err)
	}
	return
}
