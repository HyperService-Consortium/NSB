package option

import (
	"encoding/json"
	"errors"
	cmn "github.com/HyperService-Consortium/NSB/common"
	response "github.com/HyperService-Consortium/NSB/common/contract_response"
	"github.com/HyperService-Consortium/NSB/localstorage"
	"github.com/HyperService-Consortium/NSB/math"
	trie "github.com/HyperService-Consortium/go-mpt"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"github.com/syndtr/goleveldb/leveldb"
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
var baseStrikePrice = math.NewUint256FromHexString("2")

func createContract(t *testing.T) *cmn.ContractEnvironment {
	g0Root := createRoot(t, user0, contract0)
	g0Root.Value = math.NewUint256FromHexString("b")
	option := &Option{env: g0Root}
	option.NewContract(user0, baseStrikePrice)
	sugar.HandlerError(option.env.Storage.Commit())
	assert.EqualValues(t, user0, option.GetOwner())
	assert.EqualValues(t, math.NewUint256FromBytes([]byte{10}), option.GetMinStake())
	assert.EqualValues(t, baseStrikePrice, option.GetStrikePrice())
	assert.EqualValues(t, math.NewUint256FromHexString("b"), option.GetRemainingFund())
	return g0Root
}

func TestOption_NewContract(t *testing.T) {
	type fields struct {
		env *cmn.ContractEnvironment
	}
	type args struct {
		Owner       []byte        `json:"1"`
		StrikePrice *math.Uint256 `json:"2"`
	}
	_2 := math.NewUint256FromHexString("2")
	_11 := math.NewUint256FromHexString("b")
	g0Root := createRoot(t, user0, contract0)
	g0Root.Value = _11
	gNotEnoughValueRoot := createRoot(t, user0, contract0)
	gNotEnoughValueRoot.Value = _2
	g1Root := createRoot(t, user0, contract0)
	g1Root.Value = _11
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *cmn.ContractCallBackInfo
		wantErr  interface{}
		callback func(t *testing.T, option *Option)
	}{
		{
			name: "easy",
			fields: fields{
				env: g0Root},
			args: args{
				user0,
				_2,
			},
			want: &cmn.ContractCallBackInfo{
				CodeResponse: uint32(codeOK),
				Data: sugar.HandlerError(json.Marshal(NewContractReply{
					Address: contract0,
					Price:   _2.Bytes(),
				})).([]byte),
				Value: _11,
			},
			callback: func(t *testing.T, option *Option) {
				sugar.HandlerError(option.env.Storage.Commit())
				assert.EqualValues(t, user0, option.GetOwner())
				assert.EqualValues(t, math.NewUint256FromBytes([]byte{10}), option.GetMinStake())
				assert.EqualValues(t, _2, option.GetStrikePrice())
				assert.EqualValues(t, _11, option.GetRemainingFund())
			},
		},
		{
			name: "not enough value",
			fields: fields{
				env: gNotEnoughValueRoot},
			args: args{
				user0,
				_2,
			},
			wantErr: errDescNotEnoughValue,
		},
		{
			name: "default owner",
			fields: fields{
				env: g1Root},
			args: args{
				nil,
				_2,
			},
			want: &cmn.ContractCallBackInfo{
				CodeResponse: uint32(codeOK),
				Data: sugar.HandlerError(json.Marshal(NewContractReply{
					Address: contract0,
					Price:   _2.Bytes(),
				})).([]byte),
				Value: _11,
			},
			callback: func(t *testing.T, option *Option) {
				sugar.HandlerError(option.env.Storage.Commit())
				assert.EqualValues(t, user0, option.GetOwner())
				assert.EqualValues(t, math.NewUint256FromBytes([]byte{10}), option.GetMinStake())
				assert.EqualValues(t, _2, option.GetStrikePrice())
				assert.EqualValues(t, _11, option.GetRemainingFund())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var packet ArgsCreateNewContract
			packet.Owner = tt.args.Owner
			packet.StrikePrice = tt.args.StrikePrice

			var unpacked ArgsCreateNewContract
			sugar.HandlerError0(
				json.Unmarshal(sugar.HandlerError(
					json.Marshal(packet)).([]byte), &unpacked))

			tt.args.Owner = unpacked.Owner
			tt.args.StrikePrice = unpacked.StrikePrice

			option := &Option{
				env: tt.fields.env,
			}

			defer func() {
				if err := recover(); err != nil {
					if !assert.EqualValues(t, tt.wantErr, err) {
						t.Errorf("NewContract() = %v, wantErr %v", err, tt.wantErr)
					}
				}
			}()

			if got := option.NewContract(tt.args.Owner, tt.args.StrikePrice); !assert.EqualValues(t, tt.want, got) {
				t.Errorf("NewContract() = %v, want %v", got, tt.want)
			}
			if tt.callback != nil {
				tt.callback(t, option)
			}
		})
	}
}

func TestOption_UpdateStake(t *testing.T) {
	type fields struct {
		env *cmn.ContractEnvironment
	}
	type args struct {
		Value *math.Uint256 `json:"1"`
	}
	_14 := math.NewUint256FromBytes([]byte{14})
	g0Root := createContract(t)
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *cmn.ContractCallBackInfo
		wantErr  interface{}
		callback func(t *testing.T, option *Option)
	}{
		{
			name: "easy",
			fields: fields{
				env: g0Root},
			args: args{
				_14,
			},
			want: &cmn.ContractCallBackInfo{
				CodeResponse: uint32(codeOK),
				Data:         _14.Bytes(),
			},
			callback: func(t *testing.T, option *Option) {
				sugar.HandlerError(option.env.Storage.Commit())
				assert.EqualValues(t, _14, option.GetStrikePrice())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var packet ArgsUpdateStake
			packet.Value = tt.args.Value

			var unpacked ArgsUpdateStake
			sugar.HandlerError0(
				json.Unmarshal(sugar.HandlerError(
					json.Marshal(packet)).([]byte), &unpacked))

			tt.args.Value = unpacked.Value

			option := &Option{
				env: tt.fields.env,
			}

			defer func() {
				if err := recover(); err != nil {
					if !assert.EqualValues(t, tt.wantErr, err) {
						t.Errorf("UpdateStake() = %v, wantErr %v", err, tt.wantErr)
					}
				}
			}()

			if got := option.UpdateStake(tt.args.Value); !assert.EqualValues(t, tt.want, got) {
				t.Errorf("UpdateStake() = %v, want %v", got, tt.want)
			}
			if tt.callback != nil {
				tt.callback(t, option)
			}
		})
	}
}

func TestOption_StakeFund(t *testing.T) {
	type fields struct {
		env *cmn.ContractEnvironment
	}
	_14 := math.NewUint256FromBytes([]byte{14})
	g0Root := createContract(t)
	g0Root.Value = _14
	tests := []struct {
		name     string
		fields   fields
		want     *cmn.ContractCallBackInfo
		wantErr  interface{}
		callback func(t *testing.T, option *Option)
	}{
		{
			name: "easy",
			fields: fields{
				env: g0Root},
			want: response.ExecOK(g0Root.Value),
			callback: func(t *testing.T, option *Option) {
				sugar.HandlerError(option.env.Storage.Commit())
				res, overflowed := math.AddUint256(
					_14, math.NewUint256FromBytes([]byte{11}))
				assert.False(t, overflowed)
				assert.EqualValues(t,
					res, option.GetRemainingFund())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			option := &Option{
				env: tt.fields.env,
			}

			defer func() {
				if err := recover(); err != nil {
					if !assert.EqualValues(t, tt.wantErr, err) {
						t.Errorf("StakeFund() = %v, wantErr %v", err, tt.wantErr)
					}
				}
			}()

			if got := option.StakeFund(); !assert.EqualValues(t, tt.want, got) {
				t.Errorf("StakeFund() = %v, want %v", got, tt.want)
			}
			if tt.callback != nil {
				tt.callback(t, option)
			}
		})
	}
}

func TestOption_BuyOption(t *testing.T) {
	type fields struct {
		env *cmn.ContractEnvironment
	}
	type args struct {
		Proposal *math.Uint256
	}
	_5 := math.NewUint256FromBytes([]byte{5})
	exactValue, underflow := math.SubUint256(_5, baseStrikePrice)
	assert.False(t, underflow)
	assert.False(t, exactValue.Mul(_5))

	g0Root := createContract(t)
	g0Root.Value = exactValue
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *cmn.ContractCallBackInfo
		wantErr  interface{}
		callback func(t *testing.T, option *Option)
	}{
		{
			name: "easy",
			fields: fields{
				env: g0Root},
			args: args{
				_5,
			},
			want: response.ExecOK(nil),
			callback: func(t *testing.T, option *Option) {
				sugar.HandlerError(option.env.Storage.Commit())
				assert.EqualValues(t, ValidBuyer{Valid: true, Executed: false}, option.GetOptionBuyers().Get(user0))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var packet ArgsBuyOption
			packet.Proposal = tt.args.Proposal

			var unpacked ArgsBuyOption
			sugar.HandlerError0(
				json.Unmarshal(sugar.HandlerError(
					json.Marshal(packet)).([]byte), &unpacked))

			tt.args.Proposal = unpacked.Proposal

			option := &Option{
				env: tt.fields.env,
			}

			defer func() {
				if err := recover(); err != nil {
					if !assert.EqualValues(t, tt.wantErr, err) {
						t.Errorf("BuyOption() = %v, wantErr %v", err, tt.wantErr)
					}
				}
			}()

			if got := option.BuyOption(tt.args.Proposal); !assert.EqualValues(t, tt.want, got) {
				t.Errorf("BuyOption() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createContractWithValidBuyer(t *testing.T) *cmn.ContractEnvironment {
	_5 := math.NewUint256FromBytes([]byte{5})
	g0Root := createContract(t)

	exactValue, underflow := math.SubUint256(_5, baseStrikePrice)
	assert.False(t, underflow)
	assert.False(t, exactValue.Mul(_5))

	g0Root.Value = exactValue
	option := &Option{
		env: g0Root,
	}
	assert.EqualValues(t, response.ExecOK(nil), option.BuyOption(_5))
	_, err := g0Root.Storage.Commit()
	assert.NoError(t, err)
	g0Root.Value = nil
	return g0Root
}

func TestOption_CashSettle(t *testing.T) {
	type fields struct {
		env *cmn.ContractEnvironment
	}
	type args struct {
		GenuinePrice *math.Uint256 `json:"1"`
	}
	_2 := math.NewUint256FromBytes([]byte{2})
	_3 := math.NewUint256FromBytes([]byte{3})
	_5 := math.NewUint256FromBytes([]byte{5})
	g0Root := createContractWithValidBuyer(t)
	g0Root.Value = _3
	gNotEnoughValueRoot := createContractWithValidBuyer(t)
	gNotEnoughValueRoot.Value = _2

	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *cmn.ContractCallBackInfo
		wantErr  interface{}
		callback func(t *testing.T, option *Option)
	}{
		{
			name: "easy",
			fields: fields{
				env: g0Root},
			args: args{
				_5,
			},
			want: response.ExecOK(math.NewUint256FromBytes([]byte{3})),
			callback: func(t *testing.T, option *Option) {
				sugar.HandlerError(option.env.Storage.Commit())
				assert.EqualValues(t, ValidBuyer{Valid: true, Executed: true}, option.GetOptionBuyers().Get(user0))
			},
		},
		{
			name: "not enough value",
			fields: fields{
				env: gNotEnoughValueRoot},
			args: args{
				_5,
			},
			wantErr: ValueLessThanProfit,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var packet ArgsCashSettle
			packet.GenuinePrice = tt.args.GenuinePrice

			var unpacked ArgsCashSettle
			sugar.HandlerError0(
				json.Unmarshal(sugar.HandlerError(
					json.Marshal(packet)).([]byte), &unpacked))

			tt.args.GenuinePrice = unpacked.GenuinePrice

			option := &Option{
				env: tt.fields.env,
			}

			defer func() {
				if err := recover(); err != nil {
					if !assert.EqualValues(t, tt.wantErr, err) {
						t.Errorf("CashSettle() = %v, wantErr %v", err, tt.wantErr)
					}
				}
			}()

			if got := option.CashSettle(tt.args.GenuinePrice); !assert.EqualValues(t, tt.want, got) {
				t.Errorf("CashSettle() = %v, want %v", got, tt.want)
			}
			if tt.callback != nil {
				tt.callback(t, option)
			}
		})
	}
}
