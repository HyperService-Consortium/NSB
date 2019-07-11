package TxState

type Type = uint64

const (
	Unknown Type = 0 + iota
	Initing
	Inited
	Open
	Opened
	Closed
)
