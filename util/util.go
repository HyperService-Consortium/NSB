package util


import (
    "fmt"
    "encoding/binary"
    "bytes"
)

func Int64ToBytes(i int64) []byte {
    var buf = make([]byte, 8)
    binary.BigEndian.PutUint64(buf, uint64(i))
    return buf
}

func BytesToInt64(buf []byte) int64 {
    return int64(binary.BigEndian.Uint64(buf))
}

func Uint64ToBytes(i uint64) []byte {
    var buf = make([]byte, 8)
    binary.BigEndian.PutUint64(buf, uint64(i))
    return buf
}

func BytesToUint64(buf []byte) uint64 {
    return uint64(binary.BigEndian.Uint64(buf))
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