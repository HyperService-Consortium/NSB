package merkmap

import (
	"fmt"
	"encoding/hex"
	"github.com/Myriad-Dreamin/go-mpt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/Myriad-Dreamin/NSB/merkmap/MerkMapError"
)

type MerkMap struct {
	merk *trie.Trie
	slot []byte
}

func (mp *MerkMap) location(key []byte) []byte {
	//
	return []byte("")
}

func NewMerkMapFromDB(db *leveldb.DB, rootHash trie.Hash, slot interface{}) (mp *MerkMap, err error) {
	mp = &MerkMap{merk: trie.NewTrie(rootHash, db)}
	switch ori_slot := slot.(type) {
	case string:
		//hexstring
		mp.slot, err = hex.DecodeString(ori_slot)
		if err != nil {
			return nil, err
		}
		fmt.Println("string", ori_slot, mp.slot)
		return
	case []byte:
		// trans into [32]byte
		fmt.Println("byte", ori_slot)
		return nil, nil
	case [32]byte:
		fmt.Println("[32]byte", ori_slot)
		return nil, nil
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
