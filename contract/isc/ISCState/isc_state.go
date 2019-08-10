package ISCState

const (
	Initing uint8 = 0 + iota
	Inited
	Opening
	Settling
	Closed
)