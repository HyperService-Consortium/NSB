package account


import (
	"github.com/syndtr/goleveldb/leveldb"
	"encoding/hex"
	"errors"
	eddsa "golang.org/x/crypto/ed25519"
)

type Wallet struct {
	db *leveldb.DB
	Acc *Account
}

func NewWallet(db *leveldb.DB, acc *Account) *Wallet {
	return &Wallet{
		db: db,
		Acc: acc,
	}
}

func ReadWallet(db *leveldb.DB, pubkey interface{}) (*Wallet, error) {
	if db == nil {
		return nil, errors.New("nil database pointer")
	}
	switch pubkey := pubkey.(type) {
	case string:
		pri, err := hex.DecodeString(pubkey)
		if err != nil {
			return nil, err
		}
		pri, err = db.Get(pri, nil)
		if err != nil {
			return nil, err
		}
		return &Wallet{
			db: db,
			Acc: ReadAccount(pri),
		}, nil
	case []byte:
		pri, err := db.Get(pubkey, nil)
		if err != nil {
			return nil, err
		}
		return &Wallet{
			db: db,
			Acc: ReadAccount(pri),
		}, nil
	case eddsa.PublicKey:
		pri, err := db.Get([]byte(pubkey), nil)
		if err != nil {
			return nil, err
		}
		return &Wallet{
			db: db,
			Acc: ReadAccount(pri),
		}, nil
	default:
		return nil, errors.New("invalid public key type")
	}
}

func (wt *Wallet) Save() error {
	return wt.db.Put(wt.Acc.PublicKey, wt.Acc.PrivateKey, nil)
}