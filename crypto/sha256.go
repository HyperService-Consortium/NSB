package crypto

import (
	"crypto/sha256"
	"crypto/sha512"
)

func Sha256(data ...[]byte) []byte {
	d := sha256.New()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}

func Sha512(data ...[]byte) []byte {
	d := sha512.New()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}