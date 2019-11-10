package localstorage

import (
	"github.com/HyperService-Consortium/NSB/util"
	"github.com/HyperService-Consortium/NSB/math"
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

func (sto *LocalStorage) SetUint256(variName string, value *math.Uint256) {
	err := sto.variSlotMap.TryUpdate([]byte(variName), value.Bytes())
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

func (sto *LocalStorage) SetInt8(variName string, value int8) {
	err := sto.variSlotMap.TryUpdate([]byte(variName), []byte{uint8(value)})
	if err != nil {
		panic(err)
	}
	return
}

func (sto *LocalStorage) SetUint8(variName string, value uint8) {
	err := sto.variSlotMap.TryUpdate([]byte(variName), []byte{value})
	if err != nil {
		panic(err)
	}
	return
}

func (sto *LocalStorage) SetBool(variName string, value bool) {
	var err error
	if value {
		err = sto.variSlotMap.TryUpdate([]byte(variName), []byte{1})
	} else {
		err = sto.variSlotMap.TryUpdate([]byte(variName), []byte{0})
	}
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
	if len(bt) > 8 {
		panic("Decode Error: the length of getting value is more than 8")
	}
	return util.BytesToUint64(append(make([]byte, 8 - len(bt), 8), bt...))
}

func (sto *LocalStorage) GetUint256(variName string) *math.Uint256 {
	bt, err := sto.variSlotMap.TryGet([]byte(variName))
	if err != nil {
		panic(err)
	}
	if len(bt) > 32 {
		panic("Decode Error: the length of getting value is more than 32")
	}
	return math.NewUint256FromBytes(append(make([]byte, 32 - len(bt), 32), bt...))
}

func (sto *LocalStorage) GetInt64(variName string) int64 {
	bt, err := sto.variSlotMap.TryGet([]byte(variName))
	if err != nil {
		panic(err)
	}
	if len(bt) > 8 {
		panic("Decode Error: the length of getting value is more than 8")
	}
	return util.BytesToInt64(append(make([]byte, 8 - len(bt), 8), bt...))
}

func (sto *LocalStorage) GetInt8(variName string) int8 {
	bt, err := sto.variSlotMap.TryGet([]byte(variName))
	if err != nil {
		panic(err)
	}
	if len(bt) > 1 {
		panic("Decode Error: the length of getting value is more than 1")
	}
	if bt == nil || len(bt) == 0 {
		return 0
	} else {
		return int8(bt[0])
	}
}

func (sto *LocalStorage) GetUint8(variName string) uint8 {
	bt, err := sto.variSlotMap.TryGet([]byte(variName))
	if err != nil {
		panic(err)
	}
	if len(bt) > 1 {
		panic("Decode Error: the length of getting value is more than 1")
	}
	if bt == nil || len(bt) == 0 {
		return 0
	} else {
		return uint8(bt[0])
	}
}

func (sto *LocalStorage) GetBool(variName string) bool {
	bt, err := sto.variSlotMap.TryGet([]byte(variName))
	if err != nil {
		panic(err)
	}
	if len(bt) != 1 {
		panic("Decode Error: the length of getting value is not 1")
	}
	return bt[0] != 0
}
