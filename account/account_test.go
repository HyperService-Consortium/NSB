package account

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAccount(t *testing.T) {
	acc := NewAccount([]byte("account:myd;pawd:123456"))
	raw := []byte("{from:acc.pri, to: another, tag: tanne}")
	rawhash := MakeMsgHash(raw)
	signed_raw := acc.Sign(raw)

	fmt.Println(rawhash, signed_raw)

	if !acc.VerifyByRaw(signed_raw, raw) {
		t.Error("verify-raw error")
		return
	}

	if !acc.VerifyByHash(signed_raw, rawhash) {
		t.Error("verify-hash error")
		return
	}

	if !bytes.Equal(signed_raw, acc.SignHash(rawhash)) {
		t.Error("signature error")
		return
	}
}

func TestCWallet(t *testing.T) {
	bt, err := hex.DecodeString("68f865764d2705554ff95a356433580474e88eef19b5886304cf3491522562bd9a701b3ed117634371dad81cdcaeba61e3e2074f4e47a70294ae5fe1b2d8ded4")
	if err != nil {
		t.Error(err)
		return
	}
	acc := ReadAccount(bt)
	if "9a701b3ed117634371dad81cdcaeba61e3e2074f4e47a70294ae5fe1b2d8ded4" != hex.EncodeToString(acc.PublicKey) {
		t.Error("public err")
		return
	}
	fmt.Println(hex.EncodeToString(acc.Sign([]byte("\x10\x00"))))
}
