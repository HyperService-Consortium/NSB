package localstorage

import (
	"github.com/HyperService-Consortium/NSB/merkmap"
	"github.com/HyperService-Consortium/NSB/util"
)

type Uint64Map struct {
	merk *merkmap.MerkMap
}

func (sto *LocalStorage) NewUint64Map(mapName string) *Uint64Map {
	return &Uint64Map{
		merk: sto.MakeStorageSlot(mapName),
	}
}

func (ui64map *Uint64Map) Set(map_offset uint64, value []byte) {
	err := ui64map.merk.TryUpdate(util.Uint64ToBytes(map_offset), value)
	if err != nil {
		panic(err)
	}
	return
}

func (ui64map *Uint64Map) Get(map_offset uint64) []byte {
	bt, err := ui64map.merk.TryGet(util.Uint64ToBytes(map_offset))
	if err != nil {
		panic(err)
	}
	return bt
}

func (ui64map *Uint64Map) Delete(map_offset uint64) {
	err := ui64map.merk.TryDelete(util.Uint64ToBytes(map_offset))
	if err != nil {
		panic(err)
	}
	return
}
