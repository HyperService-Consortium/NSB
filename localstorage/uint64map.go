package localstorage

import (
	"github.com/HyperServiceOne/NSB/merkmap"
	"github.com/HyperServiceOne/NSB/util"
)

type Uint64Map struct {
	merk *merkmap.MerkMap
}

func (sto *LocalStorage) NewUint64Map(mapName string) *Uint64Map {
	return &Uint64Map{
		merk: sto.makeStorageSlot(mapName),
	}
}

func (ui64map *Uint64Map) Set(map_offset uint64, value []byte) error {
	return ui64map.merk.TryUpdate(util.Uint64ToBytes(map_offset), value)
}

func (ui64map *Uint64Map) Get(map_offset uint64) ([]byte, error) {
	return ui64map.merk.TryGet(util.Uint64ToBytes(map_offset))
}

func (ui64map *Uint64Map) Delete(map_offset uint64) error {
	return ui64map.merk.TryDelete(util.Uint64ToBytes(map_offset))
}
