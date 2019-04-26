package localstorage


import (
	"github.com/Myriad-Dreamin/NSB/util"
)


func (sto *LocalStorage) SetBytes(variName string, value []byte) error {
	return sto.variSlotMap.TryUpdate([]byte(variName), value)
}


func (sto *LocalStorage) SetString(variName string, value string) error {
	return sto.variSlotMap.TryUpdate([]byte(variName), []byte(value))
}


func (sto *LocalStorage) SetUint64(variName string, value uint64) error {
	return sto.variSlotMap.TryUpdate([]byte(variName), util.Uint64ToBytes(value))
}


func (sto *LocalStorage) SetInt64(variName string, value int64) error {
	return sto.variSlotMap.TryUpdate([]byte(variName), util.Int64ToBytes(value))
}


func (sto *LocalStorage) SetAny(variName string, value Bytable) error {
	return sto.variSlotMap.TryUpdate([]byte(variName), value.Bytes())
}


func (sto *LocalStorage) GetBytes(variName string) ([]byte, error) {
	return sto.variSlotMap.TryGet([]byte(variName))
}

func (sto *LocalStorage) GetString(variName string) (string, error) {
	bt, err := sto.variSlotMap.TryGet([]byte(variName))
	return string(bt), err
}

func (sto *LocalStorage) GetUint64(variName string) (uint64, error) {
	bt, err := sto.variSlotMap.TryGet([]byte(variName))
	if err != nil {
		return 0, err
	}
	if len(bt) != 8 {
		return 0, errors.New("Decode Error: the length of getting value is not 8")
	}
	return util.BytesToUint64(bt), err
}

func (sto *LocalStorage) GetString(variName string) (int64, error) {
	bt, err := sto.variSlotMap.TryGet([]byte(variName))
	if err != nil {
		return 0, err
	}
	if len(bt) != 8 {
		return 0, errors.New("Decode Error: the length of getting value is not 8")
	}
	return util.BytesToInt64(bt), err
}

