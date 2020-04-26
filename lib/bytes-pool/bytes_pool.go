package bytespool

import "sync"

// BytesPool provides the source of bytes
type BytesPool struct {
	*sync.Pool
	maxBytesSize int
}

// NewBytesPool return a pool reuse bytes between gc
func NewBytesPool(maxBytesSize int) *BytesPool {
	return &BytesPool{
		Pool: &sync.Pool{
			New: MakeNewBytesFunc(maxBytesSize),
		},
		maxBytesSize: maxBytesSize,
	}
}

// Put bytes into pool
func (bp *BytesPool) Put(b []byte) {
	if bp.maxBytesSize <= len(b) {
		// ignore this bytes
		bp.Pool.Put(b)
	}
}

// Get bytes from pool
func (bp *BytesPool) Get() []byte {
	return bp.Pool.Get().([]byte)
}
