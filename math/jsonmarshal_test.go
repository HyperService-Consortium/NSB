package math

import (
	"encoding/json"
	"fmt"
	"testing"
)

type myStruct struct {
	Ui128 *Uint128
}

type myStruct2 struct {
	Ui128 Uint128
}

func TestMarshal(t *testing.T) {
	var x = NewUint128FromBytes([]byte{1, 0, 0, 0, 0, 0, 0, 0, 0})
	fmt.Println(x)

	bt, err := json.Marshal(x)
	fmt.Println(string(bt), err)
	bt, err = json.Marshal(&x)
	fmt.Println(string(bt), err)

	var mys = &myStruct{
		Ui128: x,
	}
	bt, err = json.Marshal(&mys)
	fmt.Println(string(bt), err)
	var dec myStruct
	err = json.Unmarshal(bt, &dec)
	fmt.Println(dec.Ui128, err)

	var mys2 = &myStruct2{
		Ui128: *x,
	}
	bt, err = json.Marshal(&mys2)
	fmt.Println(string(bt), err)
	var dec2 myStruct2
	err = json.Unmarshal(bt, &dec2)
	fmt.Println(&dec2.Ui128, err)
}
