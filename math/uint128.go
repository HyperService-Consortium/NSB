package math

import (
	"errors"
	"math/big"
)

// type Uint128 struct {
// 	H uint64
// 	L uint64
// }

// func (x *Uint128) Add(y *Uint128) bool {
// 	var w = x.L
// 	x.H += y.H
// 	x.L += y.L
// 	if x.L < w {
// 		x.H++;
// 		return true
// 	}
// 	return false
// }

// func (x *Uint128) Mul(y *Uint128) (c bool) {
// 	if x.H != 0 && y.H != 0 {
// 		c = true
// 	}
// 	/*
// 	|xh * yh| overflow
// 	128                64                      0
// 						|       xl * yl        |
//     |       xh * yl     |
//     |       xl * yh     |
// 	 */
//     xh, xl := x.H * y.L , x.L * y.H
//     if c == true || xh / y.L != x.H || xl / y.H != x.L || xh + xl < xl {
//     	c = true
//     }
//     x.H = xh + xl
// 	// x.L * y.L
// 	/*
// 	128       96        64         32          0
// 	                    |        al * bl       |
// 	           |      ah * bl       |
// 	           |      ah * bh       |
// 	|      ah * bh       |

// 	*/
// 	xh, xl = x.L >> 32, x.L & 0xffffffff
// 	yh, yl := y.L >> 32, y.L & 0xffffffff
// 	// (2^{n/2} - 1) ^ 2 + 2^{n/2} - 1 = 2^n - 2^{n/2}
// 	tmp := xh * xl + ((xl * yl) >> 32)
// 	x.L = x.L * y.L
// 	// (2^{n/2} - 1) ^ 2 + 2 * (2^{n/2} - 1) = 2^n - 1
// 	xh = xh * yh + (tmp >> 32) + (((tmp & 0xffffffff) + (xl * xh)) >> 32)
// 	if c == true || x.H + xh < xh {
// 		c = true
// 	}
// 	x.H += xh
// 	return
// }

type Uint128 struct {
	b *big.Int
}

func NewUint128FromUint128(data *Uint128) *Uint128 {
	return &Uint128{
		b: new(big.Int).Set(data.b),
	}
}

func NewUint128FromBigInt(data *big.Int) *Uint128 {
	if data.BitLen() > 128 {
		return nil
	}
	return &Uint128{
		b: new(big.Int).Set(data),
	}
}

func NewUint128FromBytes(data []byte) *Uint128 {
	if len(data) > 16 {
		return nil
	}
	return &Uint128{
		b: new(big.Int).SetBytes(data),
	}
}

func NewUint128FromString(data string, base int) *Uint128 {
	b, ok := new(big.Int).SetString(data, base)
	if !ok || b.BitLen() > 128 {
		return nil
	}
	return &Uint128{
		b: b,
	}
}

func NewUint128FromHexString(data string) *Uint128 {
	b, ok := new(big.Int).SetString(data, 16)
	if !ok || b.BitLen() > 128 {
		return nil
	}
	return &Uint128{
		b: b,
	}
}

func (ui128 *Uint128) String() string {
	return ui128.b.String()
}

func (ui128 *Uint128) Bytes() []byte {
	var b, c = make([]byte, 16), ui128.b.Bytes()
	copy(b[16-len(c):], c)
	return b
}

func (ui128 *Uint128) Add(y *Uint128) bool {
	ui128.b.Add(ui128.b, y.b)
	if ui128.b.Bit(128) == 1 {
		ui128.b.SetBit(ui128.b, 128, 0)
		return true
	}
	return false
}

func (ui128 *Uint128) Sub(y *Uint128) bool {
	ui128.b.Sub(ui128.b, y.b)
	if ui128.b.Sign() == -1 {
		ui128.b.Add(ui128.b, P128)
		return true
	}
	return false
}

func (ui128 *Uint128) Mul(y *Uint128) bool {
	ui128.b.Mul(ui128.b, y.b)
	if ui128.b.BitLen() > 128 {
		ui128.b.And(ui128.b, MOD128)
		return true
	}
	return false
}

func (ui128 *Uint128) Div(y *Uint128) bool {
	var rem big.Int
	ui128.b.QuoRem(ui128.b, y.b, &rem)
	if rem.BitLen() != 0 {
		return true
	}
	return false
}

func (ui128 *Uint128) BitLen() int {
	return ui128.b.BitLen()
}

// func (ui128 *Uint128) MarshalJson() ([]byte, error) {
// 	fmt.Println("here")
// 	return ui128.b.Bytes(), nil
// }
func (ui128 *Uint128) MarshalJSON() ([]byte, error) {
	c, err := ui128.b.MarshalJSON()
	if err != nil {
		return nil, err
	}
	b := make([]byte, 2+len(c))
	b[0] = '"'
	b[len(c)+1] = '"'
	copy(b[1:], c)
	return b, nil
}

func (ui128 *Uint128) UnmarshalJSON(byteJson []byte) (err error) {
	if ui128.b == nil {
		ui128.b = new(big.Int)
	}
	if len(byteJson) >= 2 && byteJson[0] == byteJson[len(byteJson)-1] && byteJson[0] == '"' {
		err = ui128.b.UnmarshalJSON(byteJson[1 : len(byteJson)-1])
		if err != nil {
			return err
		}

		if ui128.b.BitLen() > 128 {
			return errors.New("overflow")
		}
		return nil
	}
	err = ui128.b.UnmarshalJSON(byteJson)
	if err != nil {
		return err
	}

	if ui128.b.BitLen() > 128 {
		return errors.New("overflow")
	}
	return nil
	return
}
