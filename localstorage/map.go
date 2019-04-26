package localstorage

import (
	"github.com/Myriad-Dreamin/NSB/merkmap"
)


type Bytable interface {
	Bytes() []byte
}

type Map struct {
	merk *merkmap.MerkMap
}

func (sto *LocalStorage) NewMap(mapName string) (*Map) {
	return &Map{
		merk: sto.makeStorageSlot(mapName),
	}
}

func (udMap *Map) Set(Map_offset Bytable, value []byte) error {
	return udMap.merk.TryUpdate(Map_offset.Bytes(), value)
}

func (udMap *Map) Get(Map_offset Bytable) ([]byte, error) {
	return udMap.merk.TryGet(Map_offset.Bytes())
}

func (udMap *Map) Delete(Map_offset Bytable) error {
	return udMap.merk.TryDelete(Map_offset.Bytes())
}