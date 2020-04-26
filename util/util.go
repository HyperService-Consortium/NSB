package util

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"
)

func BytesToBytes(bt []byte) []byte {
	return bt
}

func BytesToBytesHelper(bt interface{}) []byte {
	return bt.([]byte)
}

func BytesToBytesHelperR(bt []byte) interface{} {
	return bt
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

func Int32ToBytes(i int32) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

func BytesToInt32(buf []byte) int32 {
	return int32(binary.BigEndian.Uint32(buf))
}

func Int32ToBytesHelper(i interface{}) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i.(int32)))
	return buf
}

func BytesToInt32Helper(buf []byte) interface{} {
	return int32(binary.BigEndian.Uint32(buf))
}

func Uint64ToBytes(i uint64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToUint64(buf []byte) uint64 {
	return uint64(binary.BigEndian.Uint64(buf))
}

func Uint32ToBytes(i uint32) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

func BytesToUint32(buf []byte) uint32 {
	return uint32(binary.BigEndian.Uint32(buf))
}

func Uint32ToBytesHelper(i interface{}) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i.(uint32)))
	return buf
}

func BytesToUint32Helper(buf []byte) interface{} {
	return uint32(binary.BigEndian.Uint32(buf))
}

func StringToBytes(str string) []byte {
	return []byte(str)
}

func BytesToString(bt []byte) string {
	return string(bt)
}

func ErrorToString(err error) string {
	return fmt.Sprintf("%v", err)
}

func ConcatBytes(dat ...[]byte) []byte {
	var buff bytes.Buffer
	var totlen int
	for _, btdat := range dat {
		buff.Write(btdat)
		totlen += len(btdat)
	}
	return buff.Next(totlen)
}

func ConvertBytes(node string) ([]byte, error) {
	if strings.HasPrefix(node, "0x") {
		node = node[2:]
	}
	if len(node)&1 == 1 {
		node = "0" + node
	}
	return hex.DecodeString(node)
}

func ConvertBytesSlice(proof []string) (ret [][]byte, err error) {
	if len(proof) == 0 {
		return [][]byte{}, nil
	}
	ret = make([][]byte, len(proof))
	for i := range proof {
		ret[i], err = ConvertBytes(proof[i])
		if err != nil {
			return nil, err
		}
	}
	return
}
