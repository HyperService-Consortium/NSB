package ISCState
type StateType uint8
const (
	initing StateType = 0 + iota
	inited
	opening
	settling
	closed
)