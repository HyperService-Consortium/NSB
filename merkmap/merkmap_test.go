package merkmap

import (
	"encoding/hex"
	"fmt"
	"github.com/HyperService-Consortium/go-mpt"
	"testing"
)

func TestSlotfromString(t *testing.T) {
	merkmap, err := NewMerkMap("./testdb", trie.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"), "010203040506070801020304050607080102030405060708")
	if err != nil {
		t.Error(err)
		return
	}
	defer merkmap.Close()
	fmt.Println(merkmap)
}

func TestSlotfromBytes(t *testing.T) {
	merkmap, err := NewMerkMap("./testdb", trie.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"), []byte("\x01\x02\x03\x04\x05\x06\x07\x08"))
	if err != nil {
		t.Error(err)
		return
	}
	defer merkmap.Close()
	fmt.Println(merkmap)
	merkmap.TryUpdate([]byte("key"), []byte("value"))
	var bt []byte
	bt, err = merkmap.TryGet([]byte("key"))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(bt)
	bt, err = merkmap.Commit(nil)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(hex.EncodeToString(bt[:]))
}

func TestMapfromDB(t *testing.T) {
	merkmap, err := NewMerkMap("./testdb", trie.HexToHash("b9037005f71046feb3f05592da79d1c6ef38b6c470cda1d6ee6d41e9300fd51d"), []byte("\x01\x02\x03\x04\x05\x06\x07\x08"))
	if err != nil {
		t.Error(err)
		return
	}
	defer merkmap.Close()
	var bt []byte
	bt, err = merkmap.TryGet([]byte("key"))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(bt)
	bt, err = merkmap.Commit(nil)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(hex.EncodeToString(bt[:]))
}
