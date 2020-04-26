package util

import (
	"reflect"
	"testing"
)

func TestUtil(t *testing.T) {
	var uint1, uint2, uint3 uint64
	var int1, int2, int3 int64
	uint1 = 233
	uint2 = 23333
	uint3 = 233333333
	uint3 *= 233333333
	int1 = -233
	int2 = 2333
	int3 = 233333333
	int3 *= 233333333
	if uint1 != BytesToUint64(Uint64ToBytes(uint1)) {
		t.Error("no equal")
		return
	}
	if uint2 != BytesToUint64(Uint64ToBytes(uint2)) {
		t.Error("no equal")
		return
	}
	if uint3 != BytesToUint64(Uint64ToBytes(uint3)) {
		t.Error("no equal")
		return
	}

	if int1 != BytesToInt64(Int64ToBytes(int1)) {
		t.Error("no equal")
		return
	}
	if int2 != BytesToInt64(Int64ToBytes(int2)) {
		t.Error("no equal")
		return
	}
	if int3 != BytesToInt64(Int64ToBytes(int3)) {
		t.Error("no equal")
		return
	}
}

func TestConvertBytes(t *testing.T) {
	type args struct {
		node string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "odd_length", args: args{"0x0"}, want: []byte{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertBytes(tt.args.node)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}
