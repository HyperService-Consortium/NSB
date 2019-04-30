package merkmap

import (
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
	lastRoot []byte
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

func NewMerkMapFromDB(db *leveldb.DB, rtHash interface{}, slot interface{}) (mp *MerkMap, err error) {
	
	var rootHash trie.Hash
	switch rtHash.(type) {
	case string:
		rootHash = trie.HexToHash(rtHash.(string))
	case []byte:
		rootHash = trie.BytesToHash(rtHash.([]byte))
	case [32]byte:
		rootHash = trie.BytesToHash(rtHash.([]byte)[:])
	default:
		return nil, MerkMapError.UnrecognizedType
	}
	
	mp = new(MerkMap)
	mp.db, _ = trie.NewNodeBasefromDB(db)
	mp.merk, err = trie.NewTrie(rootHash, mp.db)
	if err != nil {
		return nil, err
	}
	mp.lastRoot = rootHash

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
	case []byte:
		// trans into [32]byte
		if len(ori_slot) > 32 {
			return nil, MerkMapError.DecodeOverflow
		}
		mp.slot = concatBytes(make([]byte, 32-len(ori_slot)), ori_slot)
	case [32]byte:
		mp.slot = ori_slot[0:32]
	default:
		return nil, MerkMapError.UnrecognizedType
	}
	return
}

func NewMerkMap(dbDir string, rootHash interface{}, slot interface{}) (mp *MerkMap, err error) {
	var db *leveldb.DB
	db, err = leveldb.OpenFile(dbDir, nil)
	if err != nil {
		return
	}
	return NewMerkMapFromDB(db, rootHash, slot)
}


func (mp *MerkMap) ArrangeSlot(newSlot []byte) *MerkMap {
	return &MerkMap{
		merk: mp.merk,
		db: mp.db,
		slot: newSlot,
	}
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

func (mp *MerkMap) Revert() (err error) {
	mp.merk, err = trie.NewTrie(mp.lastRoot, mp.db)
	return
}

func (mp *MerkMap) Commit(cb trie.LeafCallback) (root []byte, err error) {
	var rt trie.Hash
	rt, err = mp.merk.Commit(cb)
	mp.lastRoot = rt.Bytes()
	return mp.lastRoot, err
}

// dont use this function if its db handler comes from outside
func (mp *MerkMap) Close() error {
	return mp.db.Close()
}