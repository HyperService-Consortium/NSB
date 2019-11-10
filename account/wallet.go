package account

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/HyperService-Consortium/NSB/crypto"
	"github.com/syndtr/goleveldb/leveldb"
	eddsa "golang.org/x/crypto/ed25519"
)

var (
	prikeyHeader = []byte("Private Key: ")
	pubkeyHeader = []byte("Public Key: ")
	newlineByte  = byte('\n')
)

type Wallet struct {
	db   *leveldb.DB
	name string
	Acc  []*Account `json:"accs"`
}

func NewWallet(db *leveldb.DB, name string) *Wallet {
	return &Wallet{
		db:   db,
		name: name,
	}
}

func ReadWallet(db *leveldb.DB, name string) (*Wallet, error) {
	if db == nil {
		return nil, errors.New("nil database pointer")
	}

	bt, err := db.Get([]byte(name), nil)
	if err != nil {
		return nil, err
	}
	var wlt Wallet
	err = json.Unmarshal(bt, &wlt)
	if err != nil {
		return nil, err
	}
	wlt.db = db
	wlt.name = name
	return &wlt, nil
	// switch pubkey := pubkey.(type) {
	// case string:
	// 	pri, err := hex.DecodeString(pubkey)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	pri, err = db.Get(pri, nil)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return &Wallet{
	// 		db: db,
	// 		Acc: ReadAccount(pri),
	// 	}, nil
	// case []byte:
	// 	pri, err := db.Get(pubkey, nil)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return &Wallet{
	// 		db: db,
	// 		Acc: ReadAccount(pri),
	// 	}, nil
	// case eddsa.PublicKey:
	// 	pri, err := db.Get([]byte(pubkey), nil)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return &Wallet{
	// 		db: db,
	// 		Acc: ReadAccount(pri),
	// 	}, nil
	// default:
	// 	return nil, errors.New("invalid public key type")
	// }
}

func WalletExist(db *leveldb.DB, name string) (bool, error) {
	if db == nil {
		return false, errors.New("nil database pointer")
	}

	bt, err := db.Get([]byte(name), nil)
	
	if err != nil {
		if err.Error() == "leveldb: not found" {
			return false, nil
		}
		return false, err
	}
	return len(bt) != 0, nil

}

func (wlt *Wallet) AppendAccount(acc *Account) {
	wlt.Acc = append(wlt.Acc, acc)
}

func (wlt *Wallet) String() string {
	if len(wlt.Acc) == 0 {
		return "<empty>"
	}
	var ret bytes.Buffer
	var str string
	for _, acc := range wlt.Acc {

		ret.Write(prikeyHeader)
		str = hex.EncodeToString(acc.PrivateKey)
		ret.WriteString(str)
		ret.WriteByte(newlineByte)

		ret.Write(pubkeyHeader)
		str = hex.EncodeToString(acc.PublicKey)
		ret.WriteString(str)
		ret.WriteByte(newlineByte)
	}
	return ret.String()
}

func (wlt *Wallet) Sign(idx int, msg ...[]byte) []byte {
	msgHash := crypto.Sha512(append(sign_header_arr, msg...)...)
	return eddsa.Sign(wlt.Acc[idx].PrivateKey, msgHash)
}

func (wlt *Wallet) SignHash(idx int, msgHash []byte) []byte {
	return eddsa.Sign(wlt.Acc[idx].PrivateKey, msgHash)
}

func (wlt *Wallet) VerifyByRaw(idx int, signature []byte, msg ...[]byte) bool {
	msgHash := crypto.Sha512(append(sign_header_arr, msg...)...)
	return eddsa.Verify(wlt.Acc[idx].PublicKey, msgHash, signature)
}

func (wlt *Wallet) VerifyByHash(idx int, signature []byte, msgHash []byte) bool {
	return eddsa.Verify(wlt.Acc[idx].PublicKey, msgHash, signature)
}

func (wlt *Wallet) Save() error {
	bt, err := json.Marshal(wlt)
	if err != nil {
		return err
	}
	return wlt.db.Put([]byte(wlt.name), bt, nil)
}
