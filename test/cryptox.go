package main

import (
	"fmt"
	"encoding/hex"
	"crypto/sha256"
	"golang.org/x/crypto/sha3"
)
const hextable = "0123456789abcdef"
func main() {
	tester := make([]byte, 0)
	for idx :=0; idx < 32; idx++ {
		tester = append(tester, 0);
	}
	ret := sha256.Sum256(tester)
	fmt.Println(ret)
	tester = make([]byte, 0)
	fmt.Println(tester)
	for idx :=0; idx < 32; idx++ {
		tester = append(tester, ret[idx]);
	}
	hexret := hex.EncodeToString(tester)
	fmt.Println(hexret)
}