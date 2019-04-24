package merkmap

import (
	"fmt"
	"testing"
	"github.com/Myriad-Dreamin/go-mpt"
)

func TestMerkMap(t *testing.T) {
	merkmap, err := NewMerkMap("./testdb", trie.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"), "00")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(merkmap)
}