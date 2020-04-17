package system_merkle_proof

import (
	cmn "github.com/HyperService-Consortium/NSB/common"
	"github.com/HyperService-Consortium/NSB/localstorage"
	"github.com/HyperService-Consortium/NSB/merkmap"
	trie "github.com/HyperService-Consortium/go-mpt"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/stretchr/testify/assert"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/tendermint/tendermint/abci/types"
	"math/big"
	"reflect"
	"testing"
)

var __x_ldb *leveldb.DB
var __x_storage *localstorage.LocalStorage

func reset(t *testing.T, b []byte) []byte {
	t.Helper()
	var c []byte
	var err error
	if __x_storage != nil {
		c, err = __x_storage.Commit()
		if err != nil {
			t.Error(err)
			return nil
		}
	}
	__x_storage, err = localstorage.NewLocalStorage(b, c, __x_ldb)
	if err != nil {
		t.Error(err)
		return nil
	}
	return c
}

func createRoot(t *testing.T, b, c []byte) *cmn.ContractEnvironment {
	t.Helper()
	setupStorage(t, b)
	var err error
	storage, err := localstorage.NewLocalStorage(c, trie.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"), __x_ldb)
	if err != nil {
		t.Error(err)
		return nil
	}
	env := &cmn.ContractEnvironment{
		Storage:         storage,
		From:            b,
		ContractAddress: c,
	}
	return env
}

func setupStorage(t *testing.T, b []byte) {
	if __x_ldb != nil {
		reset(t, b)
		return
	}
	var err error
	__x_ldb, err = leveldb.OpenFile("./testdb", nil)
	if err != nil {
		t.Error(err)
		return
	}
	reset(t, b)
}

func TestContract_validateMerkleProof(t *testing.T) {
	type fields struct {
		validMerkleProofMap        *merkmap.MerkMap
		validOnChainMerkleProofMap *merkmap.MerkMap
	}
	type args struct {
		bytesArgs []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *types.ResponseDeliverTx
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nsb := &Contract{
				validMerkleProofMap:        tt.fields.validMerkleProofMap,
				validOnChainMerkleProofMap: tt.fields.validOnChainMerkleProofMap,
			}
			if got := nsb.validateMerkleProof(tt.args.bytesArgs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateMerkleProof() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ethereumStorageBytesToValue(t *testing.T) {
	type args struct {
		value []byte
		id    uip.TypeID
	}
	tests := []struct {
		name    string
		args    args
		want    uip.Variable
		wantErr bool
	}{
		{name: "uint256", args: args{[]byte{16}, value_type.Uint256},
			want: uip.VariableImpl{Type: value_type.Uint256, Value: big.NewInt(16)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ethereumStorageBytesToValue(tt.args.value, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ethereumStorageBytesToValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.EqualValues(t, tt.want, got) {
				t.Errorf("ethereumStorageBytesToValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}
