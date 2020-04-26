package bytespool

import (
	"runtime"
	"testing"
)

func TestBytesPool(t *testing.T) {
	var bp = NewBytesPool(20)
	bp.Put(bp.Get())
	x := make(chan bool, 700)

	for i := 0; i < 666666; i++ {
		go func() {
			s := bp.Get()
			bp.Put(s)
			x <- true
		}()
	}
	for i := 0; i < 666666; i++ {
		<-x
	}
}

func TestBBytesPool(t *testing.T) {
	var bp = NewMultiThreadBytesPool(20)
	bp.Put(bp.Get())
	x := make(chan bool, 666666)

	for i := 0; i < 666666; i++ {
		go func() {
			s := bp.Get()
			bp.Put(s)
			x <- true
		}()
	}
	for i := 0; i < 666666; i++ {
		<-x
	}
}

func TestBBBytesPool(t *testing.T) {
	var bp = NewBMultiThreadBytesPool(20)
	bp.Put(bp.Get())
	x := make(chan bool, 666666)

	for i := 0; i < 666666; i++ {
		go func() {
			s := bp.Get()
			bp.Put(s)
			x <- true
		}()
	}
	for i := 0; i < 666666; i++ {
		<-x
	}
}

func BenchmarkGetSet(b *testing.B) {
	b.StopTimer()
	runtime.GC()
	var bp = NewBytesPool(20)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s := bp.Get()
		bp.Put(s)
	}
}

func BenchmarkMGetSet(b *testing.B) {
	b.StopTimer()
	runtime.GC()
	var bp = NewMultiThreadBytesPool(20)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s := bp.Get()
		bp.Put(s)
	}
}

func BenchmarkMGetSet2(b *testing.B) {
	b.StopTimer()
	runtime.GC()
	var bp = NewBMultiThreadBytesPool(20)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s := bp.Get()
		bp.Put(s)
	}
}

func BenchmarkGetGetSetSet(b *testing.B) {
	b.StopTimer()
	runtime.GC()
	var bp = NewBytesPool(20)
	const testN = 350
	var s [testN][]byte
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < testN; j++ {
			s[j] = bp.Get()
		}
		for j := 0; j < testN; j++ {
			bp.Put(s[j])
		}
	}
}

func BenchmarkMGetGetSetSet(b *testing.B) {
	b.StopTimer()
	runtime.GC()
	var bp = NewMultiThreadBytesPool(20)
	const testN = 350
	var s [testN][]byte
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < testN; j++ {
			s[j] = bp.Get()
		}
		for j := 0; j < testN; j++ {
			bp.Put(s[j])
		}
	}
}

func BenchmarkMGetGetSetSet2(b *testing.B) {
	b.StopTimer()
	runtime.GC()
	var bp = NewBMultiThreadBytesPool(20)
	const testN = 350
	var s [testN][]byte
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < testN; j++ {
			s[j] = bp.Get()
		}
		for j := 0; j < testN; j++ {
			bp.Put(s[j])
		}
	}
}

func BenchmarkGSet(b *testing.B) {

	b.StopTimer()
	runtime.GC()
	var bp = NewBytesPool(20)
	x := make(chan bool, 30000)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			s := bp.Get()
			bp.Put(s)
			x <- true
		}()
	}
	for i := 0; i < b.N; i++ {
		<-x
	}
}

func BenchmarkMGSet(b *testing.B) {

	b.StopTimer()
	runtime.GC()
	var bp = NewMultiThreadBytesPool(20)
	x := make(chan bool, 30000)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			s := bp.Get()
			bp.Put(s)
			x <- true
		}()
	}
	for i := 0; i < b.N; i++ {
		<-x
	}
}

func BenchmarkMGSet2(b *testing.B) {

	b.StopTimer()
	runtime.GC()
	var bp = NewBMultiThreadBytesPool(20)
	x := make(chan bool, 30000)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			s := bp.Get()
			bp.Put(s)
			x <- true
		}()
	}
	for i := 0; i < b.N; i++ {
		<-x
	}
}

/*
=== RUN   TestBytesPool
--- PASS: TestBytesPool (3.49s)
=== RUN   TestBBytesPool
--- PASS: TestBBytesPool (0.35s)
=== RUN   TestBBBytesPool
--- PASS: TestBBBytesPool (0.35s)
goos: windows
goarch: amd64
pkg: github.com/HyperService-Consortium/object-pool/bytes-pool
BenchmarkGetSet-12              20000000                74.1 ns/op            32 B/op          1 allocs/op
BenchmarkMGetSet-12             50000000                37.1 ns/op             0 B/op          0 allocs/op
BenchmarkMGetSet2-12            50000000                36.3 ns/op             0 B/op          0 allocs/op
BenchmarkGetGetSetSet-12           50000             36457 ns/op           11203 B/op        350 allocs/op
BenchmarkMGetGetSetSet-12          50000             30004 ns/op            8001 B/op        250 allocs/op
BenchmarkMGetGetSetSet2-12        100000             13741 ns/op               0 B/op          0 allocs/op
BenchmarkGSet-12                 1000000              2628 ns/op             258 B/op          2 allocs/op
BenchmarkMGSet-12                1000000              1358 ns/op              93 B/op          0 allocs/op
BenchmarkMGSet2-12               1000000              1114 ns/op              93 B/op          0 allocs/op
PASS
*/
