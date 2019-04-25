package merkmap

import (
	"fmt"
	"testing"
	"github.com/Myriad-Dreamin/go-mpt"
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

func TestSlotfromBytes(t2 *testing.T) {
	merkmap, err := NewMerkMap("./testdb", trie.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"), []byte("\x01\x02\x03\x04\x05\x06\x07\x08"))
	if err != nil {
		t2.Error(err)
		return
	}
	defer merkmap.Close()
	fmt.Println(merkmap)
}