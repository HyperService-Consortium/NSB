package math

import (
	"fmt"
	"testing"
)

func TestUint128_A(t *testing.T) {
	var x = NewUint128FromBytes([]byte{1, 0, 0, 0, 0, 0, 0, 0, 0})
	c := x.Mul(x)
	fmt.Println(c, x.String())
}

func TestUint128_B(t *testing.T) {
	var x = NewUint128FromBytes([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
	var y = NewUint128FromBytes([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
	var z = NewUint128FromBigInt(BigOne)
	fmt.Println(x.String(), x.BitLen())
	fmt.Println(y.String(), y.BitLen())
	c := x.Mul(x)
	fmt.Println(c, x.String(), x.BitLen())
	c = x.Add(y)
	fmt.Println(c, x.String(), x.BitLen())
	c = x.Add(y)
	fmt.Println(c, x.String(), x.BitLen())
	c = x.Add(z)
	fmt.Println(c, x.String(), x.BitLen())
}

func TestUint128_C(t *testing.T) {
	var x = NewUint128FromBigInt(MOD128)
	var y = NewUint128FromBigInt(BigOne)
	fmt.Println(x.String(), x.BitLen())
	fmt.Println(y.String(), y.BitLen())
	c := x.Add(y)
	fmt.Println(c, x.String(), x.BitLen())
}

func TestUint128_D(t *testing.T) {
	var x = NewUint128FromBytes([]byte{0})
	var y = NewUint128FromBytes([]byte{1})
	fmt.Println(x.String(), x.BitLen())
	fmt.Println(y.String(), y.BitLen())
	c := x.Sub(y)
	fmt.Println(c, x.String(), x.BitLen())
}

func TestUint128_E(t *testing.T) {
	var x = NewUint128FromBytes([]byte{0})
	var y = NewUint128FromBytes([]byte{0})
	fmt.Println(x.String(), x.BitLen())
	fmt.Println(y.String(), y.BitLen())
	c := x.Sub(y)
	fmt.Println(c, x.String(), x.BitLen())
}

func TestUint128_F(t *testing.T) {
	var x = NewUint128FromBytes([]byte{4})
	var y = NewUint128FromBytes([]byte{3})
	fmt.Println(x.String(), x.BitLen())
	fmt.Println(y.String(), y.BitLen())
	c := x.Div(y)
	fmt.Println(c, x.String(), x.BitLen())
}

func TestUint128_G(t *testing.T) {
	var x = NewUint128FromBytes([]byte{6})
	var y = NewUint128FromBytes([]byte{3})
	fmt.Println(x.String(), x.BitLen())
	fmt.Println(y.String(), y.BitLen())
	c := x.Div(y)
	fmt.Println(c, x.String(), x.BitLen())
}
