package localstorage

import (
	"github.com/HyperService-Consortium/NSB/crypto"
	"github.com/HyperService-Consortium/NSB/merkmap"
	"github.com/syndtr/goleveldb/leveldb"
)

type CommitEvent func()

type LocalStorage struct {
	accountAddress []byte
	statedb        *leveldb.DB
	variSlotMap    *merkmap.MerkMap
	slotMapCache   map[string]*merkmap.MerkMap
	events         []CommitEvent
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

func (sto *LocalStorage) MakeStorageSlot(slotName string) *merkmap.MerkMap {
	if slotMap, ok := sto.slotMapCache[slotName]; ok {
		return slotMap
	}
	slotMap := sto.variSlotMap.ArrangeSlot(crypto.Sha256(sto.accountAddress, []byte(slotName)))
	sto.slotMapCache[slotName] = slotMap
	return slotMap
}

func (sto *LocalStorage) tryUpdate(slotName string, map_offset []byte, value []byte) error {
	return sto.MakeStorageSlot(slotName).TryUpdate(map_offset, value)
}

func (sto *LocalStorage) tryGet(slotName string, map_offset []byte) ([]byte, error) {
	return sto.MakeStorageSlot(slotName).TryGet(map_offset)
}

func (sto *LocalStorage) tryDelete(slotName string, map_offset []byte) error {
	return sto.MakeStorageSlot(slotName).TryDelete(map_offset)
}

func (sto *LocalStorage) Commit() (root []byte, err error) {
	for _, ev := range sto.events {
		ev()
	}
	return sto.variSlotMap.Commit(nil)
}
