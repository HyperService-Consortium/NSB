package nsb

import (
	"github.com/tendermint/tendermint/abci/types"
)

var _ types.Application = new(NSBApplication)
