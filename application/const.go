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
	MethodMissing = errors.New("no corresponding function")
)
