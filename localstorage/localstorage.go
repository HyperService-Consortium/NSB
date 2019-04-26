package localstorage

import (
	"bytes"
	"github.com/Myriad-Dreamin/NSB/merkmap"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/Myriad-Dreamin/NSB/crypto"
)

var (
	emptySlot = [32]byte{}
)

type LocalStorage struct {
	accountAddress []byte
	statedb *leveldb.DB
	emptySlotMap *merkmap.MerkMap
	slotMapCache map[string]*merkmap.MerkMap
}


func concatBytes(dat ...[]byte) []byte {
	var buff bytes.Buffer
	var totlen int
	for _, btdat := range dat {
		buff.Write(btdat)
		totlen += len(btdat)
	}
	return buff.Next(totlen)
}

func NewLocalStorage(accountAddress []byte, storageRoot interface{}, db *leveldb.DB) (*LocalStorage, error) {
	emptySlotMap, err := merkmap.NewMerkMapFromDB(db, storageRoot, emptySlot)
	if err != nil {
		return nil, err
	}
	return &LocalStorage{
		accountAddress: append(accountAddress),
		statedb: db,
		emptySlotMap: emptySlotMap,
		slotMapCache: map[string]*merkmap.MerkMap{},
	}, nil
}

func (sto *LocalStorage) MakeStorageSlot(slotName string) *merkmap.MerkMap {
	if slotMap, ok := sto.slotMapCache[slotName]; ok {
		return slotMap
	}
	slotMap := sto.emptySlotMap.ArrangeSlot(crypto.Sha256(sto.accountAddress, []byte(slotName)))
	sto.slotMapCache[slotName] = slotMap
	return slotMap
}

func (sto *LocalStorage) TryUpdate(slotName  string, map_offset []byte, value []byte) error {
	return sto.MakeStorageSlot(slotName).TryUpdate(map_offset, value)
}

func (sto *LocalStorage) TryGet(slotName  string, map_offset []byte) ([]byte, error) {
	return sto.MakeStorageSlot(slotName).TryGet(map_offset)
}

func (sto *LocalStorage) TryDelete(slotName  string, map_offset []byte) error {
	return sto.MakeStorageSlot(slotName).TryDelete(map_offset)
}

func (sto *LocalStorage) Commit() (root []byte, err error) {
	return sto.emptySlotMap.Commit(nil)
}