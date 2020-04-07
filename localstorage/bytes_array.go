package localstorage

import (
	"fmt"
	"unsafe"

	"github.com/HyperService-Consortium/NSB/merkmap"
)

type BytesArray struct {
	merk                       *merkmap.MerkMap
	length                     uint64
	__x_fast_bytes_interpreter []byte
}

func (sto *LocalStorage) NewBytesArray(arrName string) *BytesArray {
	barr := &BytesArray{
		merk:                       sto.MakeStorageSlot(arrName),
		length:                     sto.GetUint64(arrName),
		__x_fast_bytes_interpreter: []byte{0, 0, 0, 0, 0, 0, 0, 0},
	}
	sto.events = append(sto.events, func() {
		sto.SetUint64(arrName, barr.length)
	})
	return barr
}

func (barr *BytesArray) Length() uint64 {
	return barr.length
}

func (barr *BytesArray) Set(arr_offset uint64, value []byte) {
	fmt.Println("setting!!!!", arr_offset, value)
	*(*uint64)(unsafe.Pointer(&barr.__x_fast_bytes_interpreter[0])) = *(*uint64)(unsafe.Pointer(&arr_offset))
	err := barr.merk.TryUpdate(barr.__x_fast_bytes_interpreter, value)
	if err != nil {
		panic(err)
	}
	return
}

func (barr *BytesArray) Get(arr_offset uint64) []byte {
	*(*uint64)(unsafe.Pointer(&barr.__x_fast_bytes_interpreter[0])) = *(*uint64)(unsafe.Pointer(&arr_offset))
	bt, err := barr.merk.TryGet(barr.__x_fast_bytes_interpreter)
	if err != nil {
		panic(err)
	}
	return bt
}

func (barr *BytesArray) Delete(arr_offset uint64) {
	*(*uint64)(unsafe.Pointer(&barr.__x_fast_bytes_interpreter[0])) = *(*uint64)(unsafe.Pointer(&arr_offset))
	err := barr.merk.TryDelete(barr.__x_fast_bytes_interpreter)
	if err != nil {
		panic(err)
	}
	return
}

func (barr *BytesArray) Append(value []byte) {
	*(*uint64)(unsafe.Pointer(&barr.__x_fast_bytes_interpreter[0])) = *(*uint64)(unsafe.Pointer(&barr.length))
	err := barr.merk.TryUpdate(barr.__x_fast_bytes_interpreter, value)
	if err != nil {
		panic(err)
	}
	barr.length++
}
