package common

type KVPair interface {
	// must be bytes
	Key() []byte
	// must be bytes
	Value() []byte
}

type KVPairInstance struct {}

func MakeKVPair(key []byte, value []byte) KVPair {
	return KVPairInstance{
		Key: func() []byte {return key},
		Value: func() []byte {return value},
	}
}