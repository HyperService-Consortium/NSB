package localstorage

import (
	"github.com/HyperServiceOne/NSB/crypto"
	"github.com/HyperServiceOne/NSB/merkmap"
	"github.com/syndtr/goleveldb/leveldb"
)

type LocalStorage struct {
	accountAddress []byte
	statedb        *leveldb.DB
	variSlotMap    *merkmap.MerkMap
	slotMapCache   map[string]*merkmap.MerkMap
}

func NewLocalStorage(accountAddress []byte, storageRoot interface{}, db *leveldb.DB) (*LocalStorage, error) {
	emptySlot := crypto.Sha256(accountAddress)
	variSlotMap, err := merkmap.NewMerkMapFromDB(db, storageRoot, emptySlot)
	if err != nil {
		return nil, err
	}
	return &LocalStorage{
		accountAddress: append(accountAddress),
		statedb:        db,
		variSlotMap:    variSlotMap,
		slotMapCache:   map[string]*merkmap.MerkMap{"": variSlotMap},
	}, nil
}

func (sto *LocalStorage) makeStorageSlot(slotName string) *merkmap.MerkMap {
	if slotMap, ok := sto.slotMapCache[slotName]; ok {
		return slotMap
	}
	slotMap := sto.variSlotMap.ArrangeSlot(crypto.Sha256(sto.accountAddress, []byte(slotName)))
	sto.slotMapCache[slotName] = slotMap
	return slotMap
}

func (sto *LocalStorage) tryUpdate(slotName string, map_offset []byte, value []byte) error {
	return sto.makeStorageSlot(slotName).TryUpdate(map_offset, value)
}

func (sto *LocalStorage) tryGet(slotName string, map_offset []byte) ([]byte, error) {
	return sto.makeStorageSlot(slotName).TryGet(map_offset)
}

func (sto *LocalStorage) tryDelete(slotName string, map_offset []byte) error {
	return sto.makeStorageSlot(slotName).TryDelete(map_offset)
}

func (sto *LocalStorage) Commit() (root []byte, err error) {
	return sto.variSlotMap.Commit(nil)
}
