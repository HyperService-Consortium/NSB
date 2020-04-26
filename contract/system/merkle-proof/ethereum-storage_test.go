package system_merkle_proof

import (
	"encoding/hex"
	ethclient "github.com/HyperService-Consortium/NSB/lib/eth-client"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func Test_ethStorageHandler_GetStorageAt(t *testing.T) {
	type fields struct {
		handler *ethclient.EthClient
		ChainID uip.ChainID
	}
	type args struct {
		typeID          uip.TypeID
		contractAddress uip.ContractAddress
		pos             []byte
		description     []byte
	}
	c := ethclient.NewEthClient("http://121.89.200.234:8545")
	addr := sugar.HandlerError(hex.DecodeString("198dcf509c33d3b10bf9111123e38269b34416cf")).([]byte)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uip.Variable
		wantErr bool
	}{
		{name: "test_uint128", fields: fields{
			handler: c,
			ChainID: 7,
		}, args: args{
			typeID:          value_type.Uint128,
			contractAddress: addr,
			pos:             make([]byte, 32),
			description:     nil,
		}, want: uip.VariableImpl{Type: value_type.Uint128, Value: big.NewInt(2)}},
		{name: "test_uint128_2", fields: fields{
			handler: c,
			ChainID: 7,
		}, args: args{
			typeID:          value_type.Uint128,
			contractAddress: addr,
			pos:             append(make([]byte, 32), 16),
			description:     nil,
		}, want: uip.VariableImpl{Type: value_type.Uint128, Value: big.NewInt(1)}},
		{name: "test_uint256", fields: fields{
			handler: c,
			ChainID: 7,
		}, args: args{
			typeID:          value_type.Uint256,
			contractAddress: addr,
			pos:             append(make([]byte, 31), 1),
			description:     nil,
		}, want: uip.VariableImpl{Type: value_type.Uint256, Value: big.NewInt(3)}},
		{name: "test_int64", fields: fields{
			handler: c,
			ChainID: 7,
		}, args: args{
			typeID:          value_type.Int64,
			contractAddress: addr,
			pos:             append(make([]byte, 31), []byte{2, 24}...),
			description:     nil,
		}, want: uip.VariableImpl{Type: value_type.Int64, Value: int64(4)}},
		{name: "test_int32", fields: fields{
			handler: c,
			ChainID: 7,
		}, args: args{
			typeID:          value_type.Int32,
			contractAddress: addr,
			pos:             append(make([]byte, 31), []byte{2, 20}...),
			description:     nil,
		}, want: uip.VariableImpl{Type: value_type.Int32, Value: int32(5)}},
		{name: "test_int16", fields: fields{
			handler: c,
			ChainID: 7,
		}, args: args{
			typeID:          value_type.Int16,
			contractAddress: addr,
			pos:             append(make([]byte, 31), []byte{2, 18}...),
			description:     nil,
		}, want: uip.VariableImpl{Type: value_type.Int16, Value: int16(6)}},
		{name: "test_int8", fields: fields{
			handler: c,
			ChainID: 7,
		}, args: args{
			typeID:          value_type.Int8,
			contractAddress: addr,
			pos:             append(make([]byte, 31), []byte{2, 17}...),
			description:     nil,
		}, want: uip.VariableImpl{Type: value_type.Int8, Value: int8(7)}},
		{name: "test_short_string", fields: fields{
			handler: c,
			ChainID: 7,
		}, args: args{
			typeID:          value_type.String,
			contractAddress: addr,
			pos:             append(make([]byte, 31), 3),
			description:     nil,
		}, want: uip.VariableImpl{Type: value_type.String, Value: "12345678"}},
		{name: "test_short_bytes", fields: fields{
			handler: c,
			ChainID: 7,
		}, args: args{
			typeID:          value_type.Bytes,
			contractAddress: addr,
			pos:             append(make([]byte, 31), 5),
			description:     nil,
		}, want: uip.VariableImpl{Type: value_type.Bytes, Value: []byte("12345678")}},
		{name: "test_bool", fields: fields{
			handler: c,
			ChainID: 7,
		}, args: args{
			typeID:          value_type.Bool,
			contractAddress: addr,
			pos:             append(make([]byte, 31), []byte{7, 31}...),
			description:     nil,
		}, want: uip.VariableImpl{Type: value_type.Bool, Value: true}},
		{name: "test_long_string", fields: fields{
			handler: c,
			ChainID: 7,
		}, args: args{
			typeID:          value_type.String,
			contractAddress: addr,
			pos:             append(make([]byte, 31), 4),
			description:     nil,
		}, want: uip.VariableImpl{Type: value_type.String, Value: "123456781234567812345678123456789"}},
		{name: "test_long_bytes", fields: fields{
			handler: c,
			ChainID: 7,
		}, args: args{
			typeID:          value_type.Bytes,
			contractAddress: addr,
			pos:             append(make([]byte, 31), 6),
			description:     nil,
		}, want: uip.VariableImpl{Type: value_type.Bytes, Value: []byte("12345678123456781234567812345678")}},
		{name: "test_undefined", fields: fields{
			handler: c,
			ChainID: 7,
		}, args: args{
			typeID:          value_type.Unknown,
			contractAddress: addr,
			pos:             append(make([]byte, 31), 8),
			description:     nil,
		}, want: uip.VariableImpl{Type: value_type.Unknown, Value: nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eth := ethStorageHandler{
				handler: tt.fields.handler,
				ChainID: tt.fields.ChainID,
			}
			got, err := eth.GetStorageAt(tt.args.typeID, tt.args.contractAddress, tt.args.pos, tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStorageAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.EqualValues(t, tt.want, got) {
				t.Errorf("GetStorageAt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
