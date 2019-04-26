package util

import (
	"testing"
)

func TestUtil(t *testing.T) {
	var uint1,uint2,uint3 uint64
	var int1,int2,int3 int64
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