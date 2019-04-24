package TxState

type Type uint8
const (
	unknown Type = 0 + iota
	initing
	inited
	open
	opened
	closed
)
