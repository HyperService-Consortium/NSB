package isc

import (
	"math/rand"
	"time"
	"github.com/Myriad-Dreamin/go-mpt"
)


type Address []byte

func NewISCAddress() (addr Address) {
	rand.Seed(time.Now().UnixNano())
	addr = Keccak256(
		[]byte(string(rand.Intn(1 << 31))),
		[]byte(string(rand.Intn(1 << 31))),
		[]byte(string(rand.Intn(1 << 31))),
		[]byte(string(rand.Intn(1 << 31))),
		[]byte(string(rand.Intn(1 << 31))),
		[]byte(string(rand.Intn(1 << 31))),
		[]byte(string(rand.Intn(1 << 31))),
		[]byte(string(rand.Intn(1 << 31))),
		[]byte(string(rand.Intn(1 << 31))),
	)

	return 
}