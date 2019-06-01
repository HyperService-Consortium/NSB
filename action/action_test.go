package action

import (
	"bytes"
	"github.com/HyperServiceOne/NSB/action/ActionType"
	"testing"
)

func TestSecp256k1(t *testing.T) {
	aType := ActionType.Secp256k1
	signature := []byte(
		"\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x00",
	)
	content := make([]byte, 1, 1)
	action, err := NewAction(aType, signature, content)
	if err != nil {
		t.Error(err)
		return
	}
	if action.Type != aType {
		t.Errorf("actionType mismatch")
		return
	}
	if !bytes.Equal(action.Signature, signature) {
		t.Errorf("signature mismatch")
		return
	}
	if !bytes.Equal(action.Content, content) {
		t.Errorf("content mismatch")
		return
	}

	concatBytes := action.Concat()
	actionf, err := TryRecoverFromConcation(concatBytes)
	if err != nil {
		t.Error(err)
		return
	}
	if action.Type != actionf.Type {
		t.Errorf("actionType mismatch")
		return
	}
	if !bytes.Equal(action.Signature, actionf.Signature) {
		t.Errorf("signature mismatch\n%v\n%v\n", action.Signature, actionf.Signature)
		return
	}
	if !bytes.Equal(action.Content, actionf.Content) {
		t.Errorf("content mismatch")
		return
	}

	actionf = RecoverFromConcation(concatBytes)
	if action.Type != actionf.Type {
		t.Errorf("actionType mismatch")
		return
	}
	if !bytes.Equal(action.Signature, actionf.Signature) {
		t.Errorf("signature mismatch\n%v\n%v\n", action.Signature, actionf.Signature)
		return
	}
	if !bytes.Equal(action.Content, actionf.Content) {
		t.Errorf("content mismatch")
		return
	}

}

func TestSecp256k1_SHORTLEN(t *testing.T) {
	aType := ActionType.Secp256k1
	signature := []byte(
		"\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef",
	)
	content := make([]byte, 1, 1)
	action, err := NewAction(aType, signature, content)

	if action != unknownAction {
		t.Errorf("action mismatch")
		return
	}
	if err != errShortLen {
		t.Errorf("errShortLen must be detected, but no error here")
		return
	}
}

func TestEd25519(t *testing.T) {
	aType := ActionType.Ed25519
	signature := []byte(
		"\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef",
	)
	content := make([]byte, 1, 1)
	action, err := NewAction(aType, signature, content)
	if err != nil {
		t.Error(err)
		return
	}
	if action.Type != aType {
		t.Errorf("actionType mismatch")
		return
	}
	if !bytes.Equal(action.Signature, signature) {
		t.Errorf("signature mismatch")
		return
	}
	if !bytes.Equal(action.Content, content) {
		t.Errorf("content mismatch")
		return
	}

	concatBytes := action.Concat()
	actionf, err := TryRecoverFromConcation(concatBytes)
	if err != nil {
		t.Error(err)
		return
	}
	if action.Type != actionf.Type {
		t.Errorf("actionType mismatch")
		return
	}
	if !bytes.Equal(action.Signature, actionf.Signature) {
		t.Errorf("signature mismatch\n%v\n%v\n", action.Signature, actionf.Signature)
		return
	}
	if !bytes.Equal(action.Content, actionf.Content) {
		t.Errorf("content mismatch")
		return
	}

	actionf = RecoverFromConcation(concatBytes)
	if action.Type != actionf.Type {
		t.Errorf("actionType mismatch")
		return
	}
	if !bytes.Equal(action.Signature, actionf.Signature) {
		t.Errorf("signature mismatch\n%v\n%v\n", action.Signature, actionf.Signature)
		return
	}
	if !bytes.Equal(action.Content, actionf.Content) {
		t.Errorf("content mismatch")
		return
	}
}

func TestEd25519_SHORTLEN(t *testing.T) {
	aType := ActionType.Ed25519
	signature := []byte(
		"\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd",
	)
	content := make([]byte, 1, 1)
	action, err := NewAction(aType, signature, content)

	if action != unknownAction {
		t.Errorf("action mismatch")
		return
	}
	if err != errShortLen {
		t.Errorf("errShortLen must be detected, but no error here")
		return
	}
}

func TestUnknownActionType(t *testing.T) {
	aType := uint8(255)
	signature := []byte(
		"\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd\xef\x01\x23\x45\x67\x89\xab\xcd",
	)
	content := make([]byte, 1, 1)
	action, err := NewAction(aType, signature, content)

	if action != unknownAction {
		t.Errorf("action mismatch")
		return
	}
	if err != errMissType {
		t.Errorf("errMissType must be detected, but no error here")
		return
	}
}
