package merkmap

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/HyperService-Consortium/NSB/merkmap/MerkMapError"
	"github.com/HyperService-Consortium/go-mpt"
	"github.com/syndtr/goleveldb/leveldb"
)

type MerkMap struct {
	merk     *trie.Trie
	db       *trie.NodeBase
	slot     []byte
	lastRoot []byte
}

type ProofJson struct {
	Proof [][]byte `json:"proof"`
	Key   []byte   `json:"key"`
	Value []byte   `json:"value"`
	Log   string   `json:"log"`
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
	case trie.Hash:
		rootHash = rtHash.(trie.Hash)
	default:
		return nil, MerkMapError.UnrecognizedType
	}

	mp = new(MerkMap)
	mp.db, _ = trie.NewNodeBasefromDB(db)
	mp.merk, err = trie.NewTrie(rootHash, mp.db)
	if err != nil {
		return nil, err
	}
	mp.lastRoot = rootHash.Bytes()

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
		db:   mp.db,
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

func (mp *MerkMap) TryPureGet(key []byte) ([]byte, error) {
	return mp.merk.TryGet(key)
}

func (mp *MerkMap) TryDelete(key []byte) error {
	return mp.merk.TryDelete(mp.location(key))
}

func (mp *MerkMap) TryPureDelete(key []byte) error {
	return mp.merk.TryDelete(key)
}

func (mp *MerkMap) TryProve(key []byte) ([][]byte, error) {
	return mp.merk.TryProve(mp.location(key))
}

func (mp *MerkMap) MakeProof(key []byte) string {
	var proofJson ProofJson
	proofJson.Key = key
	val, err := mp.TryGet(key)
	if err != nil {
		return mp.MakeErrorProof(err)
	}
	proofJson.Value = val
	proof, err := mp.TryProve(key)
	if err != nil {
		return mp.MakeErrorProof(err)
	}
	proofJson.Proof = proof

	bt, _ := json.Marshal(proofJson)
	return string(bt)
}

func (mp *MerkMap) MakeErrorProof(err error) string {
	var proofJson ProofJson
	proofJson.Log = fmt.Sprintf("%v", err)
	bt, _ := json.Marshal(proofJson)
	return string(bt)
}

func (mp *MerkMap) MakeErrorProofFromString(str string) string {
	var proofJson ProofJson
	proofJson.Log = str
	bt, _ := json.Marshal(proofJson)
	return string(bt)
}

func (mp *MerkMap) Revert() (err error) {
	mp.merk, err = trie.NewTrie(trie.BytesToHash(mp.lastRoot), mp.db)
	return
}

func (mp *MerkMap) Commit(cb trie.LeafCallback) (root []byte, err error) {
	var rt trie.Hash
	rt, err = mp.merk.Commit(cb)
	if err != nil {
		return nil, err
	}
	mp.lastRoot = rt.Bytes()
	return mp.lastRoot, err
}

// dont use this function if its db handler comes from outside
func (mp *MerkMap) Close() error {
	return mp.db.Close()
}
