package localstorage

import (
	"github.com/Myriad-Dreamin/NSB/merkmap"
)


type UserDefinedType interface {
	Bytes() []byte
}

type UserDefinedMap struct {
	merk *merkmap.MerkMap
}

func (sto *LocalStorage) NewUserDefinedMap(mapName string) (*UserDefinedMap) {
	return &UserDefinedMap{
		merk: sto.makeStorageSlot(mapName),
	}
}

func (udmap *UserDefinedMap) Set(map_offset UserDefinedType, value []byte) error {
	return udmap.merk.TryUpdate(map_offset.Bytes(), value)
}

func (udmap *UserDefinedMap) Get(map_offset UserDefinedType) ([]byte, error) {
	return udmap.merk.TryGet(map_offset.Bytes())
}

func (udmap *UserDefinedMap) Delete(map_offset UserDefinedType) error {
	return udmap.merk.TryDelete(map_offset.Bytes())
}