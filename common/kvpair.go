package common


type KVPair interface {
	// must be bytes
	Key() []byte
	// must be bytes
	Value() []byte
}
