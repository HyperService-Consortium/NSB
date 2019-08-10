package TxState

type Type = uint64

const (
	Unknown Type = 0 + iota
	Initing
	Inited
	Instantiating
	Instantiated
	Open
	Opened
	Closed
)
