package delegate

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	cmn "github.com/HyperService-Consortium/NSB/common"
	"github.com/HyperService-Consortium/NSB/localstorage"
	"github.com/HyperService-Consortium/NSB/math"
	trie "github.com/HyperService-Consortium/go-mpt"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"github.com/syndtr/goleveldb/leveldb"
	"reflect"
	"testing"
)

type StorageKey struct {
	chainID         uip.ChainID
	typeID          uip.TypeID
	contractAddress string
	pos             string
	description     string
}

type storageImpl struct {
	externalStorage map[StorageKey]uip.Variable
}

func (c *storageImpl) GetTransactionProof(chainID uip.ChainID, blockID uip.BlockID, color []byte) (uip.MerkleProof, error) {
	panic("implement me")
}

func (c *storageImpl) GetStorageAt(chainID uip.ChainID, typeID uip.TypeID,
	contractAddress uip.ContractAddress, pos []byte, description []byte) (uip.Variable, error) {
	if c.externalStorage == nil {
		c.externalStorage = make(map[StorageKey]uip.Variable)
	}
	if x, ok := c.externalStorage[StorageKey{
		chainID:         chainID,
		typeID:          typeID,
		contractAddress: string(contractAddress),
		pos:             string(pos),
		description:     string(description),
	}]; ok {
		return x, nil
	} else {
		return nil, errors.New("no found")
	}
}

func (c *storageImpl) ProvideExternalStorageAt(chainID uip.ChainID, typeID uip.TypeID,
	contractAddress uip.ContractAddress, pos []byte, description []byte, ref uip.Variable) {
	if c.externalStorage == nil {
		c.externalStorage = make(map[StorageKey]uip.Variable)
	}
	c.externalStorage[StorageKey{
		chainID:         chainID,
		typeID:          typeID,
		contractAddress: string(contractAddress),
		pos:             string(pos),
		description:     string(description),
	}] = ref
}

var __x_ldb *leveldb.DB

func createRoot(t *testing.T, b, c []byte) *cmn.ContractEnvironment {
	t.Helper()
	openDB(t)
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
		BN:              &storageImpl{},
	}
	return env
}

func openDB(t *testing.T) {
	var err error
	if __x_ldb == nil {
		__x_ldb, err = leveldb.OpenFile("./testdb", nil)
		if err != nil {
			t.Error(err)
			return
		}
	}
}

func closeDB() {

	if __x_ldb != nil {
		sugar.HandlerError0(__x_ldb.Close())
		__x_ldb = nil
	}
}

func TestMain(m *testing.M) {
	m.Run()
	closeDB()
}

var user0 = []byte{1}
var user1 = []byte{0}
var contract0 = []byte{2}

func TestDelegate_NewContract(t *testing.T) {
	type fields struct {
		env *cmn.ContractEnvironment
	}
	type args struct {
		delegates  [][]byte
		district   string
		totalVotes *math.Uint256
	}
	_2 := math.NewUint256FromHexString("2")

	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *cmn.ContractCallBackInfo
		wantErr  interface{}
		callback func(t *testing.T, delegate *Delegate)
	}{
		{
			name: "easy",
			fields: fields{
				env: createRoot(t, user0, contract0)},
			args: args{
				[][]byte{user0},
				"chain2",
				_2,
			},
			want: &cmn.ContractCallBackInfo{
				CodeResponse: uint32(CodeOK),
				Info: fmt.Sprintf(
					"create success , this contract is deploy at %v",
					hex.EncodeToString(contract0),
				),
				Data: contract0,
			},
			callback: func(t *testing.T, delegate *Delegate) {
				sugar.HandlerError(delegate.env.Storage.Commit())
				isDelegate := delegate.IsDelegate()
				assert.True(t, isDelegate.Get(user0))
				assert.True(t, !isDelegate.Get(user1))
				delegates := delegate.Delegates()
				assert.EqualValues(t, 1, delegates.Length())
				assert.EqualValues(t, user0, delegates.Get(0))
				assert.EqualValues(t, _2, delegate.GetTotalVotes())
				assert.EqualValues(t, "chain2", delegate.GetDistrict())
			},
		},
		{
			name: "nil delegate",
			fields: fields{
				env: createRoot(t, user0, contract0)},
			args: args{
				[][]byte{{1}, nil},
				"chain2",
				math.NewUint256FromHexString("2"),
			},
			want:    nil,
			wantErr: "delegate is not null",
		},
		{
			name: "nil votes",
			fields: fields{
				env: createRoot(t, user0, contract0)},
			args: args{
				[][]byte{{1}},
				"chain2",
				nil,
			},
			want: &cmn.ContractCallBackInfo{
				CodeResponse: uint32(CodeOK),
				Info: fmt.Sprintf(
					"create success , this contract is deploy at %v",
					hex.EncodeToString(contract0),
				),
				Data: contract0,
			},
			callback: func(t *testing.T, delegate *Delegate) {
				sugar.HandlerError(delegate.env.Storage.Commit())
				isDelegate := delegate.IsDelegate()
				assert.True(t, isDelegate.Get(user0))
				assert.True(t, !isDelegate.Get(user1))
				delegates := delegate.Delegates()
				assert.EqualValues(t, 1, delegates.Length())
				assert.EqualValues(t, user0, delegates.Get(0))
				assert.EqualValues(t,
					math.NewUint256FromBytes([]byte{0}), delegate.GetTotalVotes())
				assert.EqualValues(t, "chain2", delegate.GetDistrict())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var packet ArgsCreateNewContract
			packet.Delegates = tt.args.delegates
			packet.District = tt.args.district
			packet.TotalVotes = tt.args.totalVotes

			var unpacked ArgsCreateNewContract
			sugar.HandlerError0(
				json.Unmarshal(sugar.HandlerError(
					json.Marshal(packet)).([]byte), &unpacked))

			tt.args.delegates = unpacked.Delegates
			tt.args.district = unpacked.District
			tt.args.totalVotes = unpacked.TotalVotes

			delegate := &Delegate{
				env: tt.fields.env,
			}

			defer func() {
				if err := recover(); err != nil {
					if !assert.EqualValues(t, tt.wantErr, err) {
						t.Errorf("NewContract() = %v, wantErr %v", err, tt.wantErr)
					}
				}
			}()

			if got := delegate.NewContract(tt.args.delegates, tt.args.district, tt.args.totalVotes); !assert.EqualValues(t, tt.want, got) {
				t.Errorf("NewContract() = %v, want %v", got, tt.want)
			}
			if tt.callback != nil {
				tt.callback(t, delegate)
			}
		})
	}
}

func TestDelegate_Vote(t *testing.T) {
	type fields struct {
		env *cmn.ContractEnvironment
	}
	_2 := math.NewUint256FromHexString("2")

	var createRootVote = func() *cmn.ContractEnvironment {
		storage := createRoot(t, user0, contract0)

		assert.EqualValues(t, 0,
			(&Delegate{storage}).NewContract(
				[][]byte{user0},
				"chain2",
				_2).CodeResponse)
		sugar.HandlerError(storage.Storage.Commit())
		assert.EqualValues(t, false, (&Delegate{storage}).
			IsDelegateVoted().Get(user0))

		return storage
	}

	commonRoot := createRootVote()

	tests := []struct {
		name     string
		fields   fields
		want     *cmn.ContractCallBackInfo
		callback func(t *testing.T, delegate *Delegate)
	}{
		{
			name: "easy",
			fields: fields{
				env: commonRoot},
			want: &cmn.ContractCallBackInfo{
				CodeResponse: uint32(CodeOK),
				Data:         math.NewUint256FromBytes([]byte{3}).Bytes(),
			},
			callback: func(t *testing.T, delegate *Delegate) {
				sugar.HandlerError(delegate.env.Storage.Commit())
				isDelegate := delegate.IsDelegate()
				assert.True(t, isDelegate.Get(user0))
				assert.True(t, !isDelegate.Get(user1))
				delegates := delegate.Delegates()
				assert.EqualValues(t, 1, delegates.Length())
				assert.EqualValues(t, user0, delegates.Get(0))
				assert.EqualValues(t, math.NewUint256FromBytes([]byte{3}), delegate.GetTotalVotes())
				assert.EqualValues(t, "chain2", delegate.GetDistrict())
				assert.EqualValues(t, true, delegate.
					IsDelegateVoted().Get(user0))
			},
		},
		{
			name: "redo vote",
			fields: fields{
				env: commonRoot},
			want: &cmn.ContractCallBackInfo{
				CodeResponse: uint32(CodeOK),
				Data:         math.NewUint256FromBytes([]byte{3}).Bytes(),
			},
			callback: func(t *testing.T, delegate *Delegate) {
				assert.EqualValues(t, math.NewUint256FromBytes([]byte{3}), delegate.GetTotalVotes())
				assert.EqualValues(t, "chain2", delegate.GetDistrict())
				assert.EqualValues(t, true, delegate.
					IsDelegateVoted().Get(user0))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			delegate := &Delegate{
				env: tt.fields.env,
			}
			if got := delegate.Vote(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vote() = %v, want %v", got, tt.want)
			}
			if tt.callback != nil {
				tt.callback(t, delegate)
			}
		})
	}
}

func TestDelegate_ResetVote(t *testing.T) {
	type fields struct {
		env *cmn.ContractEnvironment
	}
	_2 := math.NewUint256FromHexString("2")

	var createRootVote = func() *cmn.ContractEnvironment {
		storage := createRoot(t, user0, contract0)
		delegate := (&Delegate{storage})
		assert.EqualValues(t, 0,
			delegate.NewContract(
				[][]byte{user0},
				"chain2",
				_2).CodeResponse)
		sugar.HandlerError(storage.Storage.Commit())

		assert.EqualValues(t, math.NewUint256FromBytes([]byte{2}), delegate.GetTotalVotes())
		assert.EqualValues(t, false, delegate.
			IsDelegateVoted().Get(user0))

		delegate.Vote()
		sugar.HandlerError(storage.Storage.Commit())

		assert.EqualValues(t, math.NewUint256FromBytes([]byte{3}), delegate.GetTotalVotes())
		assert.EqualValues(t, true, delegate.
			IsDelegateVoted().Get(user0))

		return storage
	}

	tests := []struct {
		name     string
		fields   fields
		want     *cmn.ContractCallBackInfo
		callback func(t *testing.T, delegate *Delegate)
	}{
		{
			name: "easy",
			fields: fields{
				env: createRootVote()},
			want: &cmn.ContractCallBackInfo{
				CodeResponse: uint32(CodeOK),
				Data:         math.NewUint256FromBytes([]byte{2}).Bytes(),
			},
			callback: func(t *testing.T, delegate *Delegate) {
				sugar.HandlerError(delegate.env.Storage.Commit())
				isDelegate := delegate.IsDelegate()
				assert.True(t, isDelegate.Get(user0))
				assert.True(t, !isDelegate.Get(user1))
				delegates := delegate.Delegates()
				assert.EqualValues(t, user0, delegates.Get(0))
				assert.EqualValues(t, math.NewUint256FromBytes([]byte{2}), delegate.GetTotalVotes())
				assert.EqualValues(t, "chain2", delegate.GetDistrict())
				assert.EqualValues(t, false, delegate.
					IsDelegateVoted().Get(user0))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			delegate := &Delegate{
				env: tt.fields.env,
			}
			if got := delegate.ResetVote(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResetVote() = %v, want %v", got, tt.want)
			}
			if tt.callback != nil {
				tt.callback(t, delegate)
			}
		})
	}
}
