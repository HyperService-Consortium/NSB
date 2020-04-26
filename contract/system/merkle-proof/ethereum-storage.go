package system_merkle_proof

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/NSB/crypto"
	ethclient "github.com/HyperService-Consortium/NSB/lib/eth-client"
	"github.com/HyperService-Consortium/NSB/util"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/tidwall/gjson"
	"math/big"
	"strconv"
)

func newEthereumStorageHandler(id uip.ChainID, host string) StorageHandler {
	return ethStorageHandler{
		handler: ethclient.NewEthClient(host),
		ChainID: id,
	}
}

type ethStorageHandler struct {
	handler *ethclient.EthClient
	ChainID uip.ChainID
}

func (eth ethStorageHandler) GetTransactionProof(blockID uip.BlockID, color []byte) (uip.MerkleProof, error) {
	panic("implement me")
}

func (eth ethStorageHandler) getRawProof(contractAddress []byte, tag uint64, pos []byte) (
	_ []byte, err error) {
	var b []byte
	if tag == 0 {
		b, err = eth.handler.GetBlockByTag(ethclient.TagLatest, false)
	} else {
		b, err = eth.handler.GetBlockByNumber(tag, false)
	}
	if err != nil {
		return nil, err
	}

	res := gjson.ParseBytes(b)
	blockNumber, err := strconv.ParseUint(res.Get("number").String()[2:], 16, 64)
	if err != nil {
		return nil, err
	}

	return eth.handler.GetProofByNumberSR(
		"0x"+hex.EncodeToString(contractAddress), pos, blockNumber)
}

func (eth ethStorageHandler) getStorageBytes(contractAddress []byte, tag uint64, pos []byte) ([]byte, error) {
	values, err := eth.getStorageBytesSlice(
		contractAddress, tag, []byte(fmt.Sprintf(`["0x%s"]`, hex.EncodeToString(pos))))
	if err != nil {
		return nil, err
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("bad length of proving values, want 1, got %v", len(values))
	}
	return values[0], nil
}

func (eth ethStorageHandler) getStorageBytesSlice(
	contractAddress []byte, tag uint64, posSlice []byte) ([][]byte, error) {
	proof, err := eth.getRawProof(contractAddress, tag, posSlice)
	if err != nil {
		return nil, err
	}

	var reply ethclient.EthereumGetProofReply
	err = json.Unmarshal(proof, &reply)
	if err != nil {
		return nil, err
	}

	var values = make([][]byte, len(reply.StorageProofs))
	for i := range values {
		values[i], err = util.ConvertBytes(reply.StorageProofs[i].Value)
		if err != nil {
			return nil, err
		}
	}
	return values, nil
}

func (eth ethStorageHandler) getStorageBytesSliceByStringSlice(
	contractAddress []byte, tag uint64, posSlice []string) ([][]byte, error) {
	b, err := json.Marshal(posSlice)
	if err != nil {
		return nil, err
	}
	return eth.getStorageBytesSlice(contractAddress, tag, b)
}

func extendTo32(value []byte) (dst []byte) {
	dst = make([]byte, 32)
	copy(dst[32-len(value):], value)
	return
}

func add1InPlace(value []byte) (dst []byte) {
	dst = value
	for i := 31; i >= 0; i-- {
		if dst[i] == 255 {
			dst[i] = 0
			if i == 0 {
				return nil
			}
		} else {
			dst[i]++
			return
		}
	}
	return
}

func (eth ethStorageHandler) GetStorageAt(typeID uip.TypeID, contractAddress uip.ContractAddress, pos []byte, description []byte) (uip.Variable, error) {
	var offset uint8
	if len(pos) > 33 {
		return nil, errors.New("length of bytes 'pos' should not be greater than 40")
	} else if len(pos) == 33 {
		pos, offset = pos[:32], pos[32]
		if offset > 32 {
			return nil, fmt.Errorf("invalid offset %v", offset)
		}
	} else if len(pos) > 32 {
		return nil, errors.New("length of bytes 'pos' without offset should not be greater than ")
	}
	// tag latest
	value, err := eth.getStorageBytes(contractAddress, 0, pos)
	if err != nil {
		return nil, err
	}
	if len(value) < 32 {
		value = extendTo32(value)
	}

	if typeID == value_type.String || typeID == value_type.Bytes {
		if len(value) < 32 {
			return nil, errors.New("no enough bytes to get string or bytes")
		}
		if (value[31] & 1) != 0 {
			offset = value[31] >> 1
			pos = extendTo32(pos)
			slot := crypto.Keccak256(pos)
			slotSize := (offset + 31) >> 5
			var reqs = make([]string, slotSize)
			slotSize--
			for i := uint8(0); i < slotSize; i++ {
				reqs[i] = "0x" + hex.EncodeToString(slot)
				add1InPlace(slot)
			}
			reqs[slotSize] = "0x" + hex.EncodeToString(slot)
			values, err := eth.getStorageBytesSliceByStringSlice(contractAddress, 0, reqs)
			if err != nil {
				return nil, err
			}
			value = bytes.Join(values, nil)
		} else {
			offset = value[31] >> 1
		}
		if int(offset) > len(value) {
			return nil, fmt.Errorf("string/bytes len(%v) greater than len(value) (%v)", offset, len(value))
		}
		value = value[:offset]
		offset = 0
	}

	if offset != 0 {
		if int(offset) > len(value) {
			return nil, fmt.Errorf("offset(%v) greater than len(value) (%v)", offset, len(value))
		}
		value = value[offset:]
	}

	return ethereumStorageBytesToValue(value, typeID)
	// chainID

	// hfType -> mt
	// jsonProof.RootHash -> eth.handler.getProof(contractAddress, [], "latest")
	// Key -> pos

	//err := nsb.validMerkleProofMap.TryUpdate(
	//	validateMerkleProofKey(hfType, jsonProof.RootHash, Key),
	//	util.ConcatBytes(bytesOne, Value),
	//)
}

func ethereumStorageBytesToValue(value []byte, id uip.TypeID) (uip.Variable, error) {
	switch id {
	case value_type.Uint256, value_type.Int256:
		if len(value) > 32 {
			value = value[:32]
		}
		return uip.VariableImpl{
			Type: id, Value: big.NewInt(0).SetBytes(value)}, nil
	case value_type.Uint128, value_type.Int128:
		if len(value) > 16 {
			value = value[:16]
		}
		return uip.VariableImpl{
			Type: id, Value: big.NewInt(0).SetBytes(value)}, nil
	case value_type.Uint64:
		if len(value) < 8 {
			return nil, errors.New("no enough bytes to get uint64")
		}
		return uip.VariableImpl{
			Type: id, Value: binary.BigEndian.Uint64(value)}, nil
	case value_type.Uint32:
		if len(value) < 4 {
			return nil, errors.New("no enough bytes to get uint32")
		}
		return uip.VariableImpl{
			Type: id, Value: binary.BigEndian.Uint32(value)}, nil
	case value_type.Uint16:
		if len(value) < 2 {
			return nil, errors.New("no enough bytes to get uint16")
		}
		return uip.VariableImpl{
			Type: id, Value: binary.BigEndian.Uint16(value)}, nil
	case value_type.Uint8:
		if len(value) < 1 {
			return nil, errors.New("no enough bytes to get uint8")
		}
		return uip.VariableImpl{
			Type: id, Value: value[0]}, nil
	case value_type.Int64:
		if len(value) < 8 {
			return nil, errors.New("no enough bytes to get int64")
		}
		return uip.VariableImpl{
			Type: id, Value: int64(binary.BigEndian.Uint64(value))}, nil
	case value_type.Int32:
		if len(value) < 4 {
			return nil, errors.New("no enough bytes to get int32")
		}
		return uip.VariableImpl{
			Type: id, Value: int32(binary.BigEndian.Uint32(value))}, nil
	case value_type.Int16:
		if len(value) < 2 {
			return nil, errors.New("no enough bytes to get int16")
		}
		return uip.VariableImpl{
			Type: id, Value: int16(binary.BigEndian.Uint16(value))}, nil
	case value_type.Int8:
		if len(value) < 1 {
			return nil, errors.New("no enough bytes to get int8")
		}
		return uip.VariableImpl{
			Type: id, Value: int8(value[0])}, nil
	case value_type.Bool:
		if len(value) < 1 {
			return nil, errors.New("no enough bytes to get bool")
		}
		return uip.VariableImpl{
			Type: id, Value: value[0] > 0}, nil
	case value_type.String:
		return uip.VariableImpl{Value: string(value), Type: value_type.String}, nil
	case value_type.Bytes:
		return uip.VariableImpl{Value: value, Type: value_type.Bytes}, nil
	case value_type.Unknown:
		for i := range value {
			if value[i] != 0 {
				return nil, fmt.Errorf("can not convert %v to unknown type", hex.EncodeToString(value))
			}
		}
		return uip.VariableImpl{Value: nil, Type: value_type.Unknown}, nil
	default:
		return nil, fmt.Errorf("unknown value type %v", id)
	}
}
