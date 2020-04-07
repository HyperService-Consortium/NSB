package localstorage

import (
	"github.com/HyperService-Consortium/NSB/merkmap"
)

type ToBytesHelper func(interface{}) []byte
type FromBytesHelper func([]byte) interface{}

type GeneralMap struct {
	merk                     *merkmap.MerkMap
	leftToBytesHelperFunc    ToBytesHelper
	rightToBytesHelperFunc   ToBytesHelper
	rightFromBytesHelperFunc FromBytesHelper
}

func (sto *LocalStorage) NewGeneralMap(
	mapName string,
	leftToBytesHelperFunc func(interface{}) []byte,
	rightToBytesHelperFunc func(interface{}) []byte,
	rightFromBytesHelperFunc func([]byte) interface{},
) *GeneralMap {
	return &GeneralMap{
		merk:                     sto.MakeStorageSlot(mapName),
		leftToBytesHelperFunc:    leftToBytesHelperFunc,
		rightToBytesHelperFunc:   rightToBytesHelperFunc,
		rightFromBytesHelperFunc: rightFromBytesHelperFunc,
	}
}

func (udMap *GeneralMap) Set(Map_offset interface{}, value interface{}) {
	err := udMap.merk.TryUpdate(udMap.leftToBytesHelperFunc(Map_offset), udMap.rightToBytesHelperFunc(value))
	if err != nil {
		panic(err)
	}
	return
}

func (udMap *GeneralMap) Get(Map_offset interface{}) interface{} {
	bt, err := udMap.merk.TryGet(udMap.leftToBytesHelperFunc(Map_offset))
	if err != nil {
		panic(err)
	}
	return udMap.rightFromBytesHelperFunc(bt)
}

func (udMap *GeneralMap) Delete(Map_offset interface{}) {
	err := udMap.merk.TryDelete(udMap.leftToBytesHelperFunc(Map_offset))
	if err != nil {
		panic(err)
	}
	return
}
