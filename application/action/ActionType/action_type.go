package ActionType

type Type uint8;
const (
	EthereumAction Type = 0 + iota
	NebulasAction
	TendermintAction
)