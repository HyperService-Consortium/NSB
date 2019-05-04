package crypto

import (
	"crypto/sha256"
)

func Sha256(data ...[]byte) []byte {
	d := sha256.New()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}
