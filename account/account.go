package account

import (
	"time"
	eddsa "golang.org/x/crypto/ed25519"
	"github.com/Myriad-Dreamin/NSB/crypto"
	"github.com/syndtr/goleveldb/leveldb"
	"bytes"
)

var (
	pswd_header = []byte("\x19TendermintNSBpswd:")
	sign_header = []byte("\x19TendermintNSBsign:")
)

func init() {
	random_seed = time.Now().Unix()
}

type Account struct {
	privateKey eddsa.PrivateKey
	PublicKey eddsa.PublicKey
	// randSalt []byte
}

func NewAccount(pswd []byte) *Account {
	pri := eddsa.NewKeyFromSeed(crypto.Sha256(pswd_header, pswd))
	return &Account{
		privateKey: pri
		PublicKey: pri.Public()
	}
}

func (acc *Account) Sign(msg ...[]byte) {
	msgHash := crypto.Sha512(sign_header, msg)
	return eddsa.Sign(acc.privateKey, msgHash)
}

func (acc *Account) VerifyByRaw(signature []byte, msg ...[]byte) {
	msgHash := crypto.Sha512(sign_header, msg)
	return eddsa.Verify(acc.PublicKey, msgHash, signature)
}

func (acc *Account) VerifyByHash(signature []byte, msghash []byte) {
	return eddsa.Verify(acc.PublicKey, msgHash, signature)
}
