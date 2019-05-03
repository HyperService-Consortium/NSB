package nsb

import (
	"errors"
	"github.com/tendermint/tendermint/version"
)

const (
	NSBVersion version.Protocol = 0x1
)

var (
	
	stateKey = []byte("NSBStateKey")
	actionHeader = []byte("NACHD:")
)

var (
	MethodMissing = errors.New("no corresponding function")
)
