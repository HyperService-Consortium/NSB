package math

import (
	"math/big"
	// "encoding/binary"
	"encoding/json"
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
	return ui256.b.Bytes()
}

func (ui256 *Uint256) Add(y *Uint256) bool {
	ui256.b.Add(ui256.b, y.b)
	if ui256.b.Bit(256) == 1 {
		ui256.b.SetBit(ui256.b, 256, 0)
		return true
	}
	return false
}

func (ui256 *Uint256) Sub(y *Uint256) bool {
	ui256.b.Sub(ui256.b, y.b)
	if ui256.b.Sign() == -1 {
		ui256.b.Add(ui256.b, P256)
		return true
	}
	return false
}

func (ui256 *Uint256) Mul(y *Uint256) bool {
	ui256.b.Mul(ui256.b, y.b)
	if ui256.b.BitLen() > 256 {
		ui256.b.And(ui256.b, MOD256)
		return true
	}
	return false
}

func (ui256 *Uint256) Div(y *Uint256) bool {
	var rem big.Int
	ui256.b.QuoRem(ui256.b, y.b, &rem)
	if rem.BitLen() != 0 {
		return true
	}
	return false
}

func (ui256 *Uint256) Comp(y *Uint256) int {
	return new(big.Int).Sub(ui256.b, y.b).Sign()
}

func (ui256 *Uint256) BitLen() int {
	return ui256.b.BitLen()
}


func (ui256 *Uint256) MarshalJSON() ([]byte, error) {
	return json.Marshal(ui256.b.Bytes())
}


func (ui256 *Uint256) UnmarshalJSON(byteJson []byte) (err error) {
	var bt []byte
	err = json.Unmarshal(byteJson, &bt)
	if err != nil {
		return
	}
	ui256.b = new(big.Int).SetBytes(bt)
	return nil
}