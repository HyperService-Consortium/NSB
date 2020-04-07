package localstorage

import (
	"github.com/HyperService-Consortium/NSB/merkmap"
)

type BoolMap struct {
	merk *merkmap.MerkMap
}

func (sto *LocalStorage) NewBoolMap(mapName string) *BoolMap {
	return &BoolMap{
		merk: sto.MakeStorageSlot(mapName),
	}
}

func (bmap *BoolMap) Set(map_offset []byte, value bool) {
	if value {
		err := bmap.merk.TryUpdate(map_offset, []byte{1})
		if err != nil {
			panic(err)
		}
	} else {
		err := bmap.merk.TryUpdate(map_offset, []byte{0})
		if err != nil {
			panic(err)
		}
	}
	return
}

func (bmap *BoolMap) Get(map_offset []byte) bool {
	bt, err := bmap.merk.TryGet(map_offset)
	if err != nil {
		panic(err)
	}
	if bt == nil || len(bt) == 0 {
		return false
	}
	return bt[0] != 0
}

func (bmap *BoolMap) Delete(map_offset []byte) {
	err := bmap.merk.TryDelete(map_offset)
	if err != nil {
		panic(err)
	}
	return
}
