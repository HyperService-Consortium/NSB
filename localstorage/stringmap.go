package localstorage

import (
	"github.com/Myriad-Dreamin/NSB/merkmap"
)

type StringMap struct {
	merk *merkmap.MerkMap
}

func (sto *LocalStorage) NewStringMap(mapName string) (*StringMap) {
	return &StringMap{
		merk: sto.makeStorageSlot(mapName),
	}
}

func (smap *StringMap) Set(map_offset string, value []byte) error {
	return smap.merk.TryUpdate([]byte(map_offset), value)
}

func (smap *StringMap) Get(map_offset string) ([]byte, error) {
	return smap.merk.TryGet([]byte(map_offset))
}

func (smap *StringMap) Delete(map_offset string) error {
	return smap.merk.TryDelete([]byte(map_offset))
}