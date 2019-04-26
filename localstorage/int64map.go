package localstorage

import (
	"github.com/Myriad-Dreamin/NSB/merkmap"
	"github.com/Myriad-Dreamin/NSB/util"
)


type Int64Map struct {
	merk *merkmap.MerkMap
}

func (sto *LocalStorage) NewInt64Map(mapName string) (*Int64Map) {
	return &Int64Map{
		merk: sto.makeStorageSlot(mapName),
	}
}

func (i64map *Int64Map) Set(map_offset int64, value []byte) error {
	return i64map.merk.TryUpdate(util.Int64ToBytes(map_offset), value)
}

func (i64map *Int64Map) Get(map_offset int64) ([]byte, error) {
	return i64map.merk.TryGet(util.Int64ToBytes(map_offset))
}

func (i64map *Int64Map) Delete(map_offset int64) error {
	return i64map.merk.TryDelete(util.Int64ToBytes(map_offset))
}