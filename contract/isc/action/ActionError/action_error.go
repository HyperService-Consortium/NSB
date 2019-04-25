package ActionError

import "errors"

var (
	UnrecognizedType = errors.New("ActionError: Unrecognized Merkle Proof Type")
)