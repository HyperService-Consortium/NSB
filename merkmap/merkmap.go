package merkmap

import (
	"fmt"
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

func NewMerkMapFromDB(db *leveldb.DB, slot interface{}) (mp *MerkMap, err error) {
	switch slot := slot.(type) {
	case string:
		//hexstring
		fmt.Println(slot)
		return nil, nil
	case []byte:
		// trans into [32]byte
		fmt.Println(slot)
		return nil, nil
	case [32]byte:
		fmt.Println(slot)
		return nil, nil
	default:
		return nil, MerkMapError.UnrecognizedType
	}
}

func NewMerkMap(dbDir string, slot interface{}) (mp *MerkMap, err error) {
	var db *leveldb.DB
	db, err = leveldb.OpenFile(dbDir, nil)
	if err != nil {
		return
	}
	return NewMerkMapFromDB(db, slot)
}
