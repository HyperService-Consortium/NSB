package account

import (
	"time"
	eddsa "golang.org/x/crypto/ed25519"
	"github.com/Myriad-Dreamin/NSB/crypto"
)

var (
	seed_header = []byte("\x19TendermintNSBseed:")
	sign_header = []byte("\x19TendermintNSBsign:")
	sign_header_arr = [][]byte{sign_header}
	random_seed = time.Now().Unix()
)

type Account struct {
	PrivateKey eddsa.PrivateKey
	PublicKey eddsa.PublicKey
	// randSalt []byte
}


func NewAccount(seed []byte) *Account {
	pri := eddsa.NewKeyFromSeed(crypto.Sha256(seed_header, seed))
	return &Account{
		PrivateKey: pri,
		PublicKey: pri.Public().(eddsa.PublicKey),
	}
}

func ReadAccount(PrivateKey []byte) *Account {
	return &Account{
		PrivateKey: eddsa.PrivateKey(PrivateKey),
		PublicKey: eddsa.PrivateKey(PrivateKey).Public().(eddsa.PublicKey),
	}
}

func MakeMsgHash(msg ...[]byte) []byte {
	return crypto.Sha512(append(sign_header_arr, msg...)...)
}


func (acc *Account) Sign(msg ...[]byte) []byte {
	msgHash := crypto.Sha512(append(sign_header_arr, msg...)...)
	return eddsa.Sign(acc.PrivateKey, msgHash)
}

func (acc *Account) SignHash(msgHash []byte) []byte {
	return eddsa.Sign(acc.PrivateKey, msgHash)
}


func (acc *Account) VerifyByRaw(signature []byte, msg ...[]byte) bool {
	msgHash := crypto.Sha512(append(sign_header_arr, msg...)...)
	return eddsa.Verify(acc.PublicKey, msgHash, signature)
}

func (acc *Account) VerifyByHash(signature []byte, msgHash []byte) bool {
	return eddsa.Verify(acc.PublicKey, msgHash, signature)
}

// func VerifyTransaction()