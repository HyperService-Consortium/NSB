package trie

import (
	"crypto/sha256"
	"golang.org/x/crypto/sha3"
)

type MerkleHash []byte

func Keccak256(data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}

func Sha256(data ...[]byte) []byte {
	d := sha256.New()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}