package math

import (
	"fmt"
	"testing"
)

func TestS(t *testing.T) {
	x, err := NewUint256FromUint64(1)
	if err != nil {
		t.Error(err)
		return 
	}
	fmt.Println(x.String())
}

func TestLS(t *testing.T) {
	var u, v uint64
	for idx := uint64(0); idx <= 10000; idx++ {
		for idy := uint64(0); idy <= 100000; idy++ {
			u = idx << 1
			v = idy << 1
		}
	}
	fmt.Println(u, v)
}

func TestRS(t *testing.T) {
	var u, v uint64
	for idx := uint64(0); idx <= 10000; idx++ {
		for idy := uint64(0); idy <= 100000; idy++ {
			u = idx >> 1
			v = idy << 1
		}
	}
	fmt.Println(u, v)
}

func TestAnd(t *testing.T) {
	var u, v uint64
	for idx := uint64(0); idx <= 10000; idx++ {
		for idy := uint64(0); idy <= 100000; idy++ {
			u = idx & idy
			v = idy & idx
		}
	}
	fmt.Println(u, v)
}

func TestOr(t *testing.T) {
	var u, v uint64
	for idx := uint64(0); idx <= 10000; idx++ {
		for idy := uint64(0); idy <= 100000; idy++ {
			u = idx | idy
			v = idy | idx
		}
	}
	fmt.Println(u, v)
}

func TestXor(t *testing.T) {
	var u, v uint64
	for idx := uint64(0); idx <= 10000; idx++ {
		for idy := uint64(0); idy <= 100000; idy++ {
			u = idx ^ idy
			v = idy ^ idx
		}
	}
	fmt.Println(u, v)
}
