package nsb

import (
	"errors"

	"github.com/tendermint/tendermint/version"
)

const (
	NSBVersion version.Protocol = 0x2
)

type slot = []byte

var (
	stateKey                       = []byte("NSBStateKey")
	actionHeader                   = []byte("NACHD:")
	accMapSlot                     = slot("acc:")
	txMapSlot                      = slot("tx:")
	actionMapSlot                  = slot("act:")
	validMerkleProofMapSlot        = slot("vlm:")
	validOnchainMerkleProofMapSlot = slot("vom:")
)

var (
	errorDecodeUint256    = errors.New("the value and nonce should be in length of 32")
	errorDecodeSrcAddress = errors.New("the src address should be in length of 32")
	errorDecodeDstAddress = errors.New("the dst address should be in length of 32 or 0")
)
