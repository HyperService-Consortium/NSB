package prover

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/HyperService-Consortium/NSB/crypto"
	"github.com/HyperService-Consortium/NSB/lib/eth-client"
	"github.com/HyperService-Consortium/NSB/localstorage"
	"github.com/HyperService-Consortium/NSB/util"
	trie "github.com/HyperService-Consortium/go-mpt"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"github.com/syndtr/goleveldb/leveldb"
	"strconv"
	"testing"
)

func TestProof(t *testing.T) {
	sr := "0xf3a367bc23612e791c234ff0ed6184901612ccf37078002bec76bd2658016f7c"
	var b = []byte(`{"address":"0x4b3a59cd1183ab81b3c31b5a22bce26adf928ac2","accountProof":["0xf9011180a0aaa943c508c847ca2dee7a9720293022bf530e31d5cd9821c1b1f86fde3f46c98080a062226959fca97b0cf948b8e46510fab662957b7e194f75c3a1348b31dbcd8f4ea0814fb016bb65113e7e9b80254b615caf41d6dcd7aeece6987dc28781f3762ea9a033ca2637b94b6eb8df0e1deffd6f4b3dc7db6bde632eb95c98e1fc1cbb0ee76ba01bf195448d818d363a2c26a71a38c6fc68e71d9c1d72a15aa4b790b77bc2854b8080a024ce20059a270602a655acffc296b727eba52429c4af42e46c45d254c16584578080a0fccd61908610067d25b182b7483e96c2a4e1bf794f069198d47a6914f9b93e53a0cd7b5289efdb13341bb2bc32fc69704363ddaf7c78044d403055f50814ad07a48080","0xf873a03b1958a7563602365b69feca4d6a3d7107d8256887e95f95724108e85867a325b850f84e168a0881e4200bba365bff97a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a0c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"],"balance":"0x881e4200bba365bff97","codeHash":"0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470","nonce":"0x16","storageHash":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","storageProof":[]}`)

	var reply ethclient.EthereumGetProofReply
	sugar.HandlerError0(json.Unmarshal(b, &reply))
	proof := sugar.HandlerError(
		util.ConvertBytesSlice(reply.AccountProof)).([][]byte)
	pv := sugar.HandlerError(GetMerkleProofValueWithValidateMPTSecure(
		sugar.HandlerError(util.ConvertBytes(sr)).([]byte), proof,
		hex.EncodeToString(crypto.Keccak256(
			(sugar.HandlerError(hex.DecodeString(
				"4b3a59cd1183ab81b3c31b5a22bce26adf928ac2"))).([]byte))))).([]byte)
	assert.EqualValues(t, "f84e168a0881e4200bba365bff97a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a0c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470", hex.EncodeToString(pv))
}

func TestStorageProof(t *testing.T) {
	sr := "0x462a2fa96c11321df9d87f6e826e735d8283b1e197183327c3fab07a0f879850"
	var b = []byte(`["0xf8718080a0b9ab932a547db848304805051f57d3e0382654a7b845382cf78775cc23b3d71e80a005630604c35da9cb7ad58b005ba9f93efd48ef3c262a22740512b0daccb8eb2f808080808080a0f4984a11f61a2921456141df88de6e1a710d28681b91af794c5a721e47839cd78080808080", "0xe2a0390decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56310"]`)

	var rawProof []string
	sugar.HandlerError0(json.Unmarshal(b, &rawProof))
	proof := sugar.HandlerError(
		util.ConvertBytesSlice(rawProof)).([][]byte)
	pv := sugar.HandlerError(GetMerkleProofValueWithValidateMPTSecure(
		sugar.HandlerError(util.ConvertBytes(sr)).([]byte), proof,
		hex.EncodeToString(crypto.Keccak256(
			(sugar.HandlerError(hex.DecodeString(
				"0000000000000000000000000000000000000000000000000000000000000000"))).([]byte))))).([]byte)
	fmt.Println(pv)
	//assert.EqualValues(t, "f84e168a0881e4200bba365bff97a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a0c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470", hex.EncodeToString(pv))
}

var __x_ldb *leveldb.DB

func reset(t *testing.T,
	storage *localstorage.LocalStorage, b []byte) (*localstorage.LocalStorage, []byte) {
	t.Helper()
	var c []byte
	var err error
	if storage != nil {
		c, err = storage.Commit()
		if err != nil {
			t.Error(err)
			return nil, nil
		}
	} else {
		c = trie.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421").Bytes()
	}

	storage, err = localstorage.NewLocalStorage(b, c, __x_ldb)
	if err != nil {
		t.Error(err)
		return nil, nil
	}
	return storage, c
}

func TestMain(m *testing.M) {
	m.Run()
	if __x_ldb != nil {
		sugar.HandlerError0(__x_ldb.Close())
		__x_ldb = nil
	}
}

func setupStorage(t *testing.T, storage *localstorage.LocalStorage, b []byte,
) *localstorage.LocalStorage {
	if __x_ldb != nil {
		storage, _ = reset(t, storage, b)
		return storage
	}
	var err error
	__x_ldb, err = leveldb.OpenFile("./testdb", nil)
	if err != nil {
		t.Error(err)
		return nil
	}
	storage, _ = reset(t, storage, b)
	return storage
}

var slot = NSBGetSlot([]byte{1}, []byte("test"))

func TestNSBMerkProof(t *testing.T) {
	storage := setupStorage(t, nil, []byte{1})
	merk := storage.MakeStorageSlot("test")
	sugar.HandlerError0(merk.TryUpdate([]byte("key"), []byte("value")))
	proof, err := merk.TryProve([]byte("key"))
	assert.NoError(t, err)
	assert.Equal(t, 2, len(proof))
	assert.EqualValues(t, []byte("value"), sugar.HandlerError(GetMerkleProofValueWithValidateNSBMPT(proof[0], proof[1:],
		NSBGetKeyBySlot(slot, []byte("key")))))
}

func TestNSBMerkleProof2(t *testing.T) {
	type ec struct {
		Key, Value []byte
	}
	var RandomEnsure = []ec{
		{[]byte("key"), []byte("value")},
	}

	for i := 1; i < 100; i++ {
		RandomEnsure = append(RandomEnsure, ec{
			Key:   []byte("key" + strconv.FormatInt(int64(i), 10)),
			Value: []byte("key" + strconv.FormatInt(int64(i), 10)),
		})
	}

	storage := setupStorage(t, nil, []byte{1})
	merk := storage.MakeStorageSlot("test")
	var assertOK = func(k, v []byte) {
		proof, err := merk.TryProve(k)
		assert.NoError(t, err)
		assert.EqualValues(t, v, sugar.HandlerError(GetMerkleProofValueWithValidateNSBMPT(proof[0], proof[1:],
			NSBGetKeyBySlot(slot, k))), string(k))
	}

	for i := 0; i < 100; i++ {
		sugar.HandlerError0(merk.TryUpdate(
			RandomEnsure[i].Key, RandomEnsure[i].Value))
		for j := 0; j <= i; j++ {
			assertOK(RandomEnsure[j].Key, RandomEnsure[j].Value)
		}
	}
}

func TestNSBMerkleProof3(t *testing.T) {
	type ec struct {
		Key, Value []byte
	}
	var RandomEnsure = []ec{
		{[]byte("key"), []byte("value")},
	}

	var scale = 8000

	for i := 1; i < scale; i++ {
		RandomEnsure = append(RandomEnsure, ec{
			Key:   []byte("key" + strconv.FormatInt(int64(i), 10)),
			Value: []byte("key" + strconv.FormatInt(int64(i), 10)),
		})
	}

	storage := setupStorage(t, nil, []byte{1})
	merk := storage.MakeStorageSlot("test")
	var assertOK = func(k, v []byte) {
		proof, err := merk.TryProve(k)
		assert.NoError(t, err)
		assert.EqualValues(t, v, sugar.HandlerError(GetMerkleProofValueWithValidateNSBMPT(proof[0], proof[1:],
			NSBGetKeyBySlot(slot, k))), string(k))
	}

	for i := 0; i < scale; i++ {
		sugar.HandlerError0(merk.TryUpdate(
			RandomEnsure[i].Key, RandomEnsure[i].Value))
	}
	for j := 0; j < scale; j++ {
		assertOK(RandomEnsure[j].Key, RandomEnsure[j].Value)
	}
}
