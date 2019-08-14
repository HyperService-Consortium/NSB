package crypto

import (
	"crypto/sha256"
	"hash"
)

// type Hasher interface {

// }

func Sha256(data ...[]byte) []byte {
	d := sha256.New()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}

func Sha256Hash(data ...[]byte) hash.Hash {
	d := sha256.New()
	for _, b := range data {
		d.Write(b)
	}
	return d
}
