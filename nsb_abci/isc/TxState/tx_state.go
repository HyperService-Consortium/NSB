package TxState
type StateType uint8
const (
	unknown StateType = 0 + iota
	initing
	inited
	open
	opened
	closed
)
