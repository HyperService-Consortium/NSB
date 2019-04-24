package nsb

import (
	"github.com/tendermint/tendermint/version"
)

var (
	stateKey = []byte("NSBStateKey")
	actionHeader = []byte("NACHD:")
)


const (
	NSBVersion version.Protocol = 0x1
)