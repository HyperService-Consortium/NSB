package localstorage


import (
	"github.com/Myriad-Dreamin/NSB/util"
)


func (sto *LocalStorage) SetBytes(variName string, value []byte) {
	err := sto.variSlotMap.TryUpdate([]byte(variName), value)
	if err != nil {
		panic(err)
	}
	return 
}


func (sto *LocalStorage) SetString(variName string, value string) {
	err := sto.variSlotMap.TryUpdate([]byte(variName), []byte(value))
	if err != nil {
		panic(err)
	}
	return 
}


func (sto *LocalStorage) SetUint64(variName string, value uint64) {
	err := sto.variSlotMap.TryUpdate([]byte(variName), util.Uint64ToBytes(value))
	if err != nil {
		panic(err)
	}
	return 
}


func (sto *LocalStorage) SetInt64(variName string, value int64) {
	err := sto.variSlotMap.TryUpdate([]byte(variName), util.Int64ToBytes(value))
	if err != nil {
		panic(err)
	}
	return 
}


func (sto *LocalStorage) SetAny(variName string, value Bytable) {
	err := sto.variSlotMap.TryUpdate([]byte(variName), value.Bytes())
	if err != nil {
		panic(err)
	}
	return 
}


func (sto *LocalStorage) GetBytes(variName string) []byte {
	bt, err := sto.variSlotMap.TryGet([]byte(variName))
	if err != nil {
		panic(err)
	}
	return bt
}

func (sto *LocalStorage) GetString(variName string) string {
	bt, err := sto.variSlotMap.TryGet([]byte(variName))
	if err != nil {
		panic(err)
	}
	return string(bt)
}

func (sto *LocalStorage) GetUint64(variName string) uint64 {
	bt, err := sto.variSlotMap.TryGet([]byte(variName))
	if err != nil {
		panic(err)
	}
	if len(bt) != 8 {
		panic("Decode Error: the length of getting value is not 8")
	}
	return util.BytesToUint64(bt)
}

func (sto *LocalStorage) GetInt64(variName string) int64 {
	bt, err := sto.variSlotMap.TryGet([]byte(variName))
	if err != nil {
		panic(err)
	}
	if len(bt) != 8 {
		panic("Decode Error: the length of getting value is not 8")
	}
	return util.BytesToInt64(bt)
}

