package prover

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/HyperService-Consortium/NSB/crypto"
	"github.com/HyperService-Consortium/NSB/lib/eth-client"
	"github.com/HyperService-Consortium/NSB/util"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProof(t *testing.T) {
	sr := "0xf3a367bc23612e791c234ff0ed6184901612ccf37078002bec76bd2658016f7c"
	var b = []byte(`{"address":"0x4b3a59cd1183ab81b3c31b5a22bce26adf928ac2","accountProof":["0xf9011180a0aaa943c508c847ca2dee7a9720293022bf530e31d5cd9821c1b1f86fde3f46c98080a062226959fca97b0cf948b8e46510fab662957b7e194f75c3a1348b31dbcd8f4ea0814fb016bb65113e7e9b80254b615caf41d6dcd7aeece6987dc28781f3762ea9a033ca2637b94b6eb8df0e1deffd6f4b3dc7db6bde632eb95c98e1fc1cbb0ee76ba01bf195448d818d363a2c26a71a38c6fc68e71d9c1d72a15aa4b790b77bc2854b8080a024ce20059a270602a655acffc296b727eba52429c4af42e46c45d254c16584578080a0fccd61908610067d25b182b7483e96c2a4e1bf794f069198d47a6914f9b93e53a0cd7b5289efdb13341bb2bc32fc69704363ddaf7c78044d403055f50814ad07a48080","0xf873a03b1958a7563602365b69feca4d6a3d7107d8256887e95f95724108e85867a325b850f84e168a0881e4200bba365bff97a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a0c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"],"balance":"0x881e4200bba365bff97","codeHash":"0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470","nonce":"0x16","storageHash":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","storageProof":[]}`)

	var reply ethclient.EthereumGetProofReply
	sugar.HandlerError0(json.Unmarshal(b, &reply))
	proof := sugar.HandlerError(
		util.ConvertBytesSlice(reply.AccountProof)).([][]byte)
	pv := sugar.HandlerError(GetMerkleProofValueWithValidate(
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
	pv := sugar.HandlerError(GetMerkleProofValueWithValidate(
		sugar.HandlerError(util.ConvertBytes(sr)).([]byte), proof,
		hex.EncodeToString(crypto.Keccak256(
			(sugar.HandlerError(hex.DecodeString(
				"0000000000000000000000000000000000000000000000000000000000000000"))).([]byte))))).([]byte)
	fmt.Println(pv)
	//assert.EqualValues(t, "f84e168a0881e4200bba365bff97a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a0c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470", hex.EncodeToString(pv))
}

//"0x375e99a7e7a09dff26465fcc7bf4b3da846d047b"
//{
//  accountProof: ["0xf9013180a0aaa943c508c847ca2dee7a9720293022bf530e31d5cd9821c1b1f86fde3f46c98080a062226959fca97b0cf948b8e46510fab662957b7e194f75c3a1348b31dbcd8f4ea0814fb016bb65113e7e9b80254b615caf41d6dcd7aeece6987dc28781f3762ea9a033ca2637b94b6eb8df0e1deffd6f4b3dc7db6bde632eb95c98e1fc1cbb0ee76ba0223fefa472df071075927ac8dc63e7ebda62d35cd76beab461725e601ffd55f08080a024ce20059a270602a655acffc296b727eba52429c4af42e46c45d254c16584578080a0fccd61908610067d25b182b7483e96c2a4e1bf794f069198d47a6914f9b93e53a0cd7b5289efdb13341bb2bc32fc69704363ddaf7c78044d403055f50814ad07a4a06aaf02d2e3fea8b85ba14f0a5d6c679d62fde38a381dcf4865d2d5070384ae9d80", "0xf869a03eb55e7a8d1f5168335ccdeeb17b221245ff385dee619b6f8d747f7c382ce9e6b846f8440110a0462a2fa96c11321df9d87f6e826e735d8283b1e197183327c3fab07a0f879850a06c289275d9a193131968860b7607126378ef80d71d82fdd974d6dc0a1fea91e7"],
//  address: "0x375e99a7e7a09dff26465fcc7bf4b3da846d047b",
//  balance: "0x10",
//  codeHash: "0x6c289275d9a193131968860b7607126378ef80d71d82fdd974d6dc0a1fea91e7",
//  nonce: "0x1",
//  storageHash: "0x462a2fa96c11321df9d87f6e826e735d8283b1e197183327c3fab07a0f879850",
//  storageProof: [{
//      key: "0x0",
//      proof: ["0xf8718080a0b9ab932a547db848304805051f57d3e0382654a7b845382cf78775cc23b3d71e80a005630604c35da9cb7ad58b005ba9f93efd48ef3c262a22740512b0daccb8eb2f808080808080a0f4984a11f61a2921456141df88de6e1a710d28681b91af794c5a721e47839cd78080808080", "0xe2a0390decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56310"],
//      value: "0x10"
//  }]
//}
