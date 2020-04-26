package nsb_proto

type ArgsGetStorageAt struct {
	Address []byte `json:"1"`
	Key     []byte `json:"2"`
	Slot    []byte `json:"3"`
}
