package common

type KVPair interface {
	// must be bytes
	Key() []byte
	// must be bytes
	Value() []byte
}

type KVPairInstance struct {
	key   []byte
	value []byte
}

func (kv *KVPairInstance) Key() []byte {
	return kv.key
}

func (kv *KVPairInstance) Value() []byte {
	return kv.value
}

func MakeKVPair(key []byte, value []byte) KVPair {
	return &KVPairInstance{
		key:   key,
		value: value,
	}
}
