package ISCState

type Type uint8
const (
	initing Type = 0 + iota
	inited
	opening
	settling
	closed
)