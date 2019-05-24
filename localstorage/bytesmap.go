package localstorage

import (
	"github.com/HyperServiceOne/NSB/merkmap"
)

type BytesMap struct {
	merk *merkmap.MerkMap
}

func (sto *LocalStorage) NewBytesMap(mapName string) *BytesMap {
	return &BytesMap{
		merk: sto.makeStorageSlot(mapName),
	}
}

func (bmap *BytesMap) Set(map_offset []byte, value []byte) error {
	return bmap.merk.TryUpdate(map_offset, value)
}

func (bmap *BytesMap) Get(map_offset []byte) ([]byte, error) {
	return bmap.merk.TryGet(map_offset)
}

func (bmap *BytesMap) Delete(map_offset []byte) error {
	return bmap.merk.TryDelete(map_offset)
}
