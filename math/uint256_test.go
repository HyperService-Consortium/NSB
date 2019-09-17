package math

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
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
	fmt.Println(x)
}

type MyStruct struct {
	R *Uint256
	S *Uint256
	V *Uint256
}

func TestUint256_I(t *testing.T) {
	var ms MyStruct

	ms.R = NewUint256FromBytes(nil)
	ms.S = nil
	ms.V = NewUint256FromBytes([]byte{1, 1, 1, 1, 1})
	fmt.Println(ms)
	bt, err := json.Marshal(ms)
	fmt.Println(string(bt), err)
	ms.R = nil
	ms.S = nil
	ms.V = nil
	err = json.Unmarshal(bt, &ms)
	fmt.Println(ms, err)
}

func TestUint256_J(t *testing.T) {
	var ms MyStruct

	ms.R = NewUint256FromBytes(nil)
	var tt = bytes.NewBuffer(make([]byte, 0, 32))
	binary.Write(tt, binary.BigEndian, int64(23333333))
	fmt.Println(tt.Bytes())
	ms.S = NewUint256FromBytes(tt.Bytes())
	ms.V = NewUint256FromBytes([]byte{1, 1, 1, 2, 1})

	fmt.Println(ms)
	fmt.Println(ms.R.Bytes())
	fmt.Println(ms.S.Bytes())
	fmt.Println(ms.V.Bytes())

	fmt.Println(NewUint256FromBytes(ms.R.Bytes()))
	fmt.Println(NewUint256FromBytes(ms.S.Bytes()))
	fmt.Println(NewUint256FromBytes(ms.V.Bytes()))

	fmt.Println(ms.R.String())
	fmt.Println(ms.S.String())
	fmt.Println(ms.V.String())

	fmt.Println(NewUint256FromString(ms.R.String(), 10))
	fmt.Println(NewUint256FromString(ms.S.String(), 10))
	fmt.Println(NewUint256FromString(ms.V.String(), 10))
}
