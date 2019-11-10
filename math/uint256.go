package math

import (
	"errors"
	"math/big"
)

type Uint256 struct {
	b *big.Int
}

func NewUint256FromUint256(data *Uint256) *Uint256 {
	return &Uint256{
		b: new(big.Int).Set(data.b),
	}
}

func NewUint256FromBigInt(data *big.Int) *Uint256 {
	if data.BitLen() > 256 {
		return nil
	}
	return &Uint256{
		b: new(big.Int).Set(data),
	}
}

func NewUint256FromBytes(data []byte) *Uint256 {
	if len(data) > 32 {
		return nil
	}
	return &Uint256{
		b: new(big.Int).SetBytes(data),
	}
}

func NewUint256FromString(data string, base int) *Uint256 {
	b, ok := new(big.Int).SetString(data, base)
	if !ok || b.BitLen() > 256 {
		return nil
	}
	return &Uint256{
		b: b,
	}
}

func NewUint256FromHexString(data string) *Uint256 {
	b, ok := new(big.Int).SetString(data, 16)
	if !ok || b.BitLen() > 256 {
		return nil
	}
	return &Uint256{
		b: b,
	}
}

func (ui256 *Uint256) String() string {
	return ui256.b.String()
}

func (ui256 *Uint256) Bytes() []byte {
	var b, c = make([]byte, 32), ui256.b.Bytes()
	copy(b[32-len(c):], c)
	return b
}

func (ui256 *Uint256) Add(y *Uint256) bool {
	ui256.b.Add(ui256.b, y.b)
	if ui256.b.Bit(256) == 1 {
		ui256.b.SetBit(ui256.b, 256, 0)
		return true
	}
	return false
}

func AddUint256(x *Uint256, y *Uint256) (ret *Uint256, check bool) {
	ret = NewUint256FromUint256(x)
	check = ret.Add(y)
	return
}

func (ui256 *Uint256) Sub(y *Uint256) bool {
	ui256.b.Sub(ui256.b, y.b)
	if ui256.b.Sign() == -1 {
		ui256.b.Add(ui256.b, P256)
		return true
	}
	return false
}

func SubUint256(x *Uint256, y *Uint256) (ret *Uint256, check bool) {
	ret = NewUint256FromUint256(x)
	check = ret.Sub(y)
	return
}

func (ui256 *Uint256) Mul(y *Uint256) bool {
	ui256.b.Mul(ui256.b, y.b)
	if ui256.b.BitLen() > 256 {
		ui256.b.And(ui256.b, MOD256)
		return true
	}
	return false
}

func MulUint256(x *Uint256, y *Uint256) (ret *Uint256, check bool) {
	ret = NewUint256FromUint256(x)
	check = ret.Mul(y)
	return
}

func (ui256 *Uint256) Div(y *Uint256) bool {
	var rem big.Int
	// div0? ...
	ui256.b.QuoRem(ui256.b, y.b, &rem)
	if rem.BitLen() != 0 {
		return true
	}
	return false
}

func DivUint256(x *Uint256, y *Uint256) (ret *Uint256, check bool) {
	ret = NewUint256FromUint256(x)
	check = ret.Div(y)
	return
}

func (ui256 *Uint256) Comp(y *Uint256) int {
	return new(big.Int).Sub(ui256.b, y.b).Sign()
}

func (ui256 *Uint256) BitLen() int {
	return ui256.b.BitLen()
}

func (ui256 *Uint256) MarshalJSON() ([]byte, error) {
	c, err := ui256.b.MarshalJSON()
	if err != nil {
		return nil, err
	}
	b := make([]byte, 2+len(c))
	b[0] = '"'
	b[len(c)+1] = '"'
	copy(b[1:], c)
	return b, nil
}

func (ui256 *Uint256) UnmarshalJSON(byteJson []byte) (err error) {
	if ui256.b == nil {
		ui256.b = new(big.Int)
	}
	if len(byteJson) >= 2 && byteJson[0] == byteJson[len(byteJson)-1] && byteJson[0] == '"' {
		err = ui256.b.UnmarshalJSON(byteJson[1 : len(byteJson)-1])
		if err != nil {
			return err
		}

		if ui256.b.BitLen() > 256 {
			return errors.New("overflow")
		}
		return nil
	}
	err = ui256.b.UnmarshalJSON(byteJson)
	if err != nil {
		return err
	}

	if ui256.b.BitLen() > 256 {
		return errors.New("overflow")
	}
	return nil
	return
}
