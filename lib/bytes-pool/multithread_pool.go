package bytespool

import "sync"

const (
	maxCacheP  = 100
	maxCacheP2 = 1000
)

// MultiThreadBytesPool provides the source of bytes
type MultiThreadBytesPool struct {
	*sync.Pool
	maxBytesSize int
	cacheP       [][]byte
	csp          int
	mutex        sync.Mutex
}

// NewMultiThreadBytesPool return a pool reuse bytes between gc
func NewMultiThreadBytesPool(maxBytesSize int) *MultiThreadBytesPool {
	mbp := &MultiThreadBytesPool{
		Pool: &sync.Pool{
			New: MakeNewBytesFunc(maxBytesSize),
		},
		cacheP:       make([][]byte, maxCacheP),
		csp:          maxCacheP,
		maxBytesSize: maxBytesSize,
	}
	for i := 0; i < maxCacheP; i++ {
		mbp.cacheP[i] = make([]byte, maxBytesSize)
	}
	return mbp
}

// Put bytes into pool
func (bp *MultiThreadBytesPool) Put(b []byte) {
	if bp.maxBytesSize <= len(b) {
		if bp.csp < maxCacheP {
			bp.mutex.Lock()
			if bp.csp < maxCacheP {
				bp.cacheP[bp.csp] = b
				bp.csp++
				bp.mutex.Unlock()
			} else {
				bp.mutex.Unlock()
				bp.Pool.Put(b)
			}
		} else {
			bp.Pool.Put(b)
		}
	} // else ignore this bytes
}

// Get bytes from pool
func (bp *MultiThreadBytesPool) Get() []byte {
	if bp.csp != 0 {
		bp.mutex.Lock()
		if bp.csp != 0 {
			bp.csp--
			b := bp.cacheP[bp.csp]
			bp.mutex.Unlock()
			return b
		} else {
			bp.mutex.Unlock()
			return bp.Pool.Get().([]byte)
		}
	} else {
		return bp.Pool.Get().([]byte)
	}
}

// MultiThreadBytesPool provides the source of bytes
type BMultiThreadBytesPool struct {
	*sync.Pool
	maxBytesSize int
	cacheP       [][]byte
	csp          int
	mutex        sync.Mutex
}

// NewBMultiThreadBytesPool return a pool reuse bytes between gc
func NewBMultiThreadBytesPool(maxBytesSize int) *BMultiThreadBytesPool {
	mbp := &BMultiThreadBytesPool{
		Pool: &sync.Pool{
			New: MakeNewBytesFunc(maxBytesSize),
		},
		cacheP:       make([][]byte, maxCacheP2),
		csp:          maxCacheP2,
		maxBytesSize: maxBytesSize,
	}
	for i := 0; i < maxCacheP2; i++ {
		mbp.cacheP[i] = make([]byte, maxBytesSize)
	}
	return mbp
}

// Put bytes into pool
func (bp *BMultiThreadBytesPool) Put(b []byte) {
	if bp.maxBytesSize <= len(b) {
		if bp.csp < maxCacheP2 {
			bp.mutex.Lock()
			if bp.csp < maxCacheP2 {
				bp.cacheP[bp.csp] = b
				bp.csp++
				bp.mutex.Unlock()
			} else {
				bp.mutex.Unlock()
				bp.Pool.Put(b)
			}
		} else {
			bp.Pool.Put(b)
		}
	} // else ignore this bytes
}

// Get bytes from pool
func (bp *BMultiThreadBytesPool) Get() []byte {
	if bp.csp != 0 {
		bp.mutex.Lock()
		if bp.csp != 0 {
			bp.csp--
			b := bp.cacheP[bp.csp]
			bp.mutex.Unlock()
			return b
		} else {
			bp.mutex.Unlock()
			return bp.Pool.Get().([]byte)
		}
	} else {
		return bp.Pool.Get().([]byte)
	}
}
