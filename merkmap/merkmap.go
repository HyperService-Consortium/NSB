package merkmap

import (
	_ "fmt"
	"encoding/hex"
	"bytes"
	"github.com/Myriad-Dreamin/go-mpt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/Myriad-Dreamin/NSB/merkmap/MerkMapError"
)

type MerkMap struct {
	merk *trie.Trie
	db *trie.NodeBase
	slot []byte
}

func concatBytes(lef []byte, rig []byte) []byte {
	var buff = bytes.NewBuffer(lef)
	buff.Write(rig)
	return buff.Next(len(lef) + len(rig))
}

func NewMerkMapFromDB(db *leveldb.DB, rootHash trie.Hash, slot interface{}) (mp *MerkMap, err error) {
	mp = new(MerkMap)
	mp.db, _ = trie.NewNodeBasefromDB(db)
	mp.merk, err = trie.NewTrie(rootHash, mp.db)
	if err != nil {
		return nil, err
	}
	
	switch ori_slot := slot.(type) {
	case string:
		//hexstring
		mp.slot, err = hex.DecodeString(ori_slot)
		if err != nil {
			return nil, err
		}
		if len(mp.slot) > 32 {
			return nil, MerkMapError.DecodeOverflow
		}
		mp.slot = concatBytes(make([]byte, 32-len(mp.slot)), mp.slot)
		return
	case []byte:
		// trans into [32]byte
		if len(ori_slot) > 32 {
			return nil, MerkMapError.DecodeOverflow
		}
		mp.slot = concatBytes(make([]byte, 32-len(ori_slot)), ori_slot)
		return
	case [32]byte:
		mp.slot = ori_slot[0:32]
		return
	default:
		return nil, MerkMapError.UnrecognizedType
	}
}

func NewMerkMap(dbDir string, rootHash trie.Hash, slot interface{}) (mp *MerkMap, err error) {
	var db *leveldb.DB
	db, err = leveldb.OpenFile(dbDir, nil)
	if err != nil {
		return
	}
	return NewMerkMapFromDB(db, rootHash, slot)
}


func (mp *MerkMap) location(key []byte) []byte {
	return trie.Keccak256(mp.slot, key)
}

func (mp *MerkMap) TryUpdate(key []byte, value []byte) error {
	return mp.merk.TryUpdate(mp.location(key), value)
}

func (mp *MerkMap) TryGet(key []byte) ([]byte, error) {
	return mp.merk.TryGet(mp.location(key))
}

func (mp *MerkMap) TryDelete(key []byte) error {
	return mp.merk.TryDelete(mp.location(key))
}

func (mp *MerkMap) Close() error {
	return mp.db.Close()
}