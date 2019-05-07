package math

import (
	"fmt"
	"testing"
)

func TestUint256_A(t *testing.T) {
	var x = NewUint256FromBytes([]byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	c := x.Mul(x)
	fmt.Println(c, x.String())
}

func TestUint256_B(t *testing.T) {
	var x = NewUint256FromBytes([]byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	var y = NewUint256FromUint256(x)
	var z = NewUint256FromBigInt(BigOne)
	x.Sub(z)
	y.Sub(z)
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

func TestUint256_C(t *testing.T) {
	var x = NewUint256FromBigInt(MOD256)
	var y = NewUint256FromBigInt(BigOne)
	fmt.Println(x.String(), x.BitLen())
	fmt.Println(y.String(), y.BitLen())
	c := x.Add(y)
	fmt.Println(c, x.String(), x.BitLen())
}

func TestUint256_D(t *testing.T) {
	var x = NewUint256FromBytes([]byte{0})
	var y = NewUint256FromBytes([]byte{1})
	fmt.Println(x.String(), x.BitLen())
	fmt.Println(y.String(), y.BitLen())
	c := x.Sub(y)
	fmt.Println(c, x.String(), x.BitLen())
}

func TestUint256_E(t *testing.T) {
	var x = NewUint256FromBytes([]byte{0})
	var y = NewUint256FromBytes([]byte{0})
	fmt.Println(x.String(), x.BitLen())
	fmt.Println(y.String(), y.BitLen())
	c := x.Sub(y)
	fmt.Println(c, x.String(), x.BitLen())
}

func TestUint256_F(t *testing.T) {
	var x = NewUint256FromBytes([]byte{4})
	var y = NewUint256FromBytes([]byte{3})
	fmt.Println(x.String(), x.BitLen())
	fmt.Println(y.String(), y.BitLen())
	c := x.Div(y)
	fmt.Println(c, x.String(), x.BitLen())
}

func TestUint256_G(t *testing.T) {
	var x = NewUint256FromBytes([]byte{6})
	var y = NewUint256FromBytes([]byte{3})
	fmt.Println(x.String(), x.BitLen())
	fmt.Println(y.String(), y.BitLen())
	c := x.Div(y)
	fmt.Println(c, x.String(), x.BitLen())
}

func TestUint256_H(t *testing.T) {
	var x = NewUint256FromBytes(nil)
	fmt.Println(x);
}
