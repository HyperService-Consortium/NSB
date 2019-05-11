package common

type KVPair interface {
	// must be bytes
	Key() []byte
	// must be bytes
	Value() []byte
}

type KVPairInstance struct {}


func MakeKVPair(key []byte, value []byte) KVPair {
	var mk = func() []byte {
		return key
	}
	var mv = func() []byte {
		return value
	}
	return KVPairInstance{
		Key: mk,
		Value: mv,
	}
}