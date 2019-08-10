package action

import (
	"errors"
	"github.com/HyperServiceOne/NSB/action/ActionType"
	"github.com/HyperServiceOne/NSB/util"
)

type Action struct {
	Type      uint8  `json:"type"`
	Signature []byte `json:"signature"`
	Content   []byte `json:"content"`
}

var (
	errShortLen   = errors.New("the length of bytes is too short")
	errMissType   = errors.New("unknown type of action signature")
	unknownAction = &Action{
		Type:      ActionType.Unknown,
		Signature: nil,
		Content:   nil,
	}
)

func Concat(aType uint8, signature, content []byte) []byte {
	return util.ConcatBytes([]byte{aType}, signature, content)
}

func (action *Action) Concat() []byte {
	return util.ConcatBytes([]byte{action.Type}, action.Signature, action.Content)
}

func (action *Action) TryRecoverFromConcation(concatBytes []byte) (err error) {
	if len(concatBytes) == 0 {
		action = unknownAction
		return errShortLen
	}
	switch concatBytes[0] {
	case ActionType.Secp256k1:
		if len(concatBytes) < 66 {
			action = unknownAction
			return errShortLen
		}
		action.Type = ActionType.Secp256k1
		action.Signature = concatBytes[1:66]
		action.Content = concatBytes[66:]
	case ActionType.Ed25519:
		if len(concatBytes) < 65 {
			action = unknownAction
			return errShortLen
		}
		action.Type = ActionType.Ed25519
		action.Signature = concatBytes[1:65]
		action.Content = concatBytes[65:]
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
	switch concatBytes[0] {
	case ActionType.Secp256k1:
		action.Type = ActionType.Secp256k1
		action.Signature = concatBytes[1:66]
		action.Content = concatBytes[66:]
	case ActionType.Ed25519:
		action.Type = ActionType.Ed25519
		action.Signature = concatBytes[1:65]
		action.Content = concatBytes[65:]
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

func NewAction(aType uint8, signature, content []byte) (action *Action, err error) {
	action = &Action{}
	switch aType {
	case ActionType.Secp256k1:
		if len(signature) != 65 {
			return unknownAction, errShortLen
		}
		action.Type = ActionType.Secp256k1
		action.Signature = signature
		action.Content = content
		return
	case ActionType.Ed25519:
		if len(signature) != 64 {
			return unknownAction, errShortLen
		}
		action.Type = ActionType.Ed25519
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
