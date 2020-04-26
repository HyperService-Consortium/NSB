package bytespool

// PoolNewer creates bytes
type PoolNewer func() interface{}

// MakeNewBytesFunc return a fucntion to create new bytes
func MakeNewBytesFunc(maxBytesSize int) PoolNewer {
	return func() interface{} {
		return make([]byte, maxBytesSize)
	}
}
