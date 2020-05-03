package system_merkle_proof

import (
	"fmt"
	"github.com/HyperService-Consortium/NSB/application/nsb_proto"
	"github.com/HyperService-Consortium/NSB/merkmap"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/storage"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type NSBClient interface {
	GetStorageProofAt(
		args nsb_proto.ArgsGetStorageAt) (*merkmap.ProofJson, error)
}

var newNSBClient = func(host string) NSBClient {
	panic("should register")
}

func SetNewNSBClient(f func(host string) NSBClient) {
	newNSBClient = f
}

func newTendermintNSBStorageHandler(host string) StorageHandler {
	return TendermintNSBStorageHandler{
		Handler: newNSBClient(host),
	}
}

type TendermintNSBStorageHandler struct {
	Handler NSBClient
}

func (t TendermintNSBStorageHandler) GetTransactionProof(blockID uip.BlockID, color []byte) (uip.MerkleProof, error) {
	panic("implement me")
}

func (t TendermintNSBStorageHandler) GetStorageAt(typeID uip.TypeID, contractAddress uip.ContractAddress, pos []byte, description []byte) (uip.Variable, error) {
	resp, err := t.Handler.GetStorageProofAt(
		nsb_proto.ArgsGetStorageAt{
			Address: contractAddress,
			Key:     description,
			Slot:    pos,
		})
	if err != nil {
		return nil, err
	}

	return tendermintNSBStorageBytesToValue(resp.Value, typeID)
}

func tendermintNSBStorageBytesToValue(value []byte, id uip.TypeID) (uip.Variable, error) {
	var decoder interface {
		Decode([]byte) (interface{}, error)
	}
	switch id {
	case value_type.Uint8:
		decoder = storage.Uint8
	case value_type.Uint16:
		decoder = storage.Uint16
	case value_type.Uint32:
		decoder = storage.Uint32
	case value_type.Uint64:
		decoder = storage.Uint32
	case value_type.Uint128:
		decoder = storage.Uint128
	case value_type.Uint256:
		decoder = storage.Uint256
	case value_type.Int8:
		decoder = storage.Int8
	case value_type.Int16:
		decoder = storage.Int16
	case value_type.Int32:
		decoder = storage.Int32
	case value_type.Int64:
		decoder = storage.Int32
	case value_type.Int128:
		decoder = storage.Int128
	case value_type.Int256:
		decoder = storage.Int256
	case value_type.String:
		decoder = storage.String
	case value_type.Bytes:
		decoder = storage.Bytes
	case value_type.Bool:
		decoder = storage.Bool
	case value_type.Unknown:
		if len(value) == 0 {
			return &uip.VariableImpl{Type: id, Value: nil}, nil
		}
	}
	if decoder == nil {
		return nil, fmt.Errorf("invalid value type: %v", id)
	}
	return refWrapper(id).wrapValue(decoder.Decode(value))
}

type refWrapper uip.TypeID

func (id refWrapper) wrapValue(value interface{}, err error) (uip.Variable, error) {
	if err != nil {
		return nil, err
	}
	return uip.VariableImpl{Value: value, Type: uip.TypeID(id)}, nil
}
