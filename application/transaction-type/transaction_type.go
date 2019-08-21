package transactiontype

type Type = uint8

const (
	// skip the default value "0"

	Validators Type = iota + 1
	SendTransaction
	SystemCall
	CreateContract
)
