package crypto

import (
	"crypto/sha512"
)

func Sha512(data ...[]byte) []byte {
	d := sha512.New()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}
