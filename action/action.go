package action

import (
	"errors"

	"github.com/HyperService-Consortium/NSB/util"
	signaturetype "github.com/Myriad-Dreamin/go-uip/const/signature_type"
)

type Action struct {
	Type      uint32 `json:"type"`
	Signature []byte `json:"signature"`
	Content   []byte `json:"content"`
}

var (
	errShortLen   = errors.New("the length of bytes is too short")
	errMissType   = errors.New("unknown type of action signature")
	unknownAction = &Action{
		Type:      signaturetype.Unknown,
		Signature: nil,
		Content:   nil,
	}
)

func Concat(aType uint8, signature, content []byte) []byte {
	return util.ConcatBytes([]byte{aType}, signature, content)
}

func (action *Action) Concat() []byte {
	return util.ConcatBytes(util.Uint32ToBytes(action.Type), action.Signature, action.Content)
}

func (action *Action) TryRecoverFromConcation(concatBytes []byte) (err error) {
	if len(concatBytes) < 4 {
		action = unknownAction
		return errShortLen
	}
	switch util.BytesToUint32(concatBytes[0:4]) {
	case signaturetype.Secp256k1:
		if len(concatBytes) < 69 {
			action = unknownAction
			return errShortLen
		}
		action.Type = signaturetype.Secp256k1
		action.Signature = concatBytes[4:69]
		action.Content = concatBytes[69:]
	case signaturetype.Ed25519:
		if len(concatBytes) < 68 {
			action = unknownAction
			return errShortLen
		}
		action.Type = signaturetype.Ed25519
		action.Signature = concatBytes[4:68]
		action.Content = concatBytes[68:]
	default:
		action = unknownAction
		return errMissType
	}
	return
}

func TryRecoverFromConcation(concatBytes []byte) (action *Action, err error) {
	action = &Action{}
	err = action.TryRecoverFromConcation(concatBytes)
	return
}

func (action *Action) RecoverFromConcation(concatBytes []byte) {
	switch util.BytesToUint32(concatBytes[0:4]) {
	case signaturetype.Secp256k1:
		action.Type = signaturetype.Secp256k1
		action.Signature = concatBytes[4:69]
		action.Content = concatBytes[69:]
	case signaturetype.Ed25519:
		action.Type = signaturetype.Ed25519
		action.Signature = concatBytes[4:68]
		action.Content = concatBytes[68:]
	default:
		action = unknownAction
		return
	}
	return
}

func RecoverFromConcation(concatBytes []byte) (action *Action) {
	action = &Action{}
	action.RecoverFromConcation(concatBytes)
	return
}

func NewAction(aType uint32, signature, content []byte) (action *Action, err error) {
	action = &Action{}
	switch aType {
	case signaturetype.Secp256k1:
		if len(signature) != 65 {
			return unknownAction, errShortLen
		}
		action.Type = signaturetype.Secp256k1
		action.Signature = signature
		action.Content = content
		return
	case signaturetype.Ed25519:
		if len(signature) != 64 {
			return unknownAction, errShortLen
		}
		action.Type = signaturetype.Ed25519
		action.Signature = signature
		action.Content = content
		return
	default:
		return unknownAction, errMissType
	}
}

func (action *Action) Verify() bool {
	// TODO: Verify
	return true
}
