package account

import (
	"fmt"
	"testing"
	"bytes"
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
