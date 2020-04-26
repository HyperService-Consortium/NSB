package nsbcli

import (
	"encoding/json"
	"fmt"
	transactiontype "github.com/HyperService-Consortium/NSB/application/transaction-type"
	"github.com/HyperService-Consortium/NSB/contract/delegate"
	"github.com/HyperService-Consortium/NSB/contract/isc"
	"github.com/HyperService-Consortium/NSB/grpc/nsbrpc"
	bytespool "github.com/HyperService-Consortium/NSB/lib/bytes-pool"
	"github.com/HyperService-Consortium/NSB/lib/nsb-client/nsb-message"
	"github.com/HyperService-Consortium/NSB/lib/request"
	"github.com/HyperService-Consortium/NSB/math"
	"github.com/HyperService-Consortium/go-uip/signaturer"
	"github.com/HyperService-Consortium/go-uip/uip"
	"io"
	"reflect"
	"testing"
)

var nc = NewNSBClient("121.89.200.234:26657")
var signer = HandlerError(signaturer.NewTendermintNSBSigner(make([]byte, 64))).(*signaturer.TendermintNSBSigner)

func PrettifyStruct(i interface{}) string {
	je, _ := json.MarshalIndent(i, "", "\t")
	return string(je)
}

func PrintStruct(i interface{}) {
	fmt.Println(PrettifyStruct(i))
}

func HandlerError(i interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return i
}

type fields struct {
	handler    *request.Client
	bufferPool *bytespool.BytesPool
}

func getNormalField() fields {
	return fields{
		handler:    nc.handler,
		bufferPool: nc.bufferPool,
	}
}

func TestNSBClient_GetAbciInfo(t *testing.T) {
	tests := []struct {
		name    string
		client  *NSBClient
		wantErr bool
	}{
		{name: "test-connection", client: nc, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := nc.GetAbciInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAbciInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Fatalf("GetAbciInfo() got = %v, want not nil", got)
			}
			if tt.wantErr == false {
				rt := got.LastBlockAppHash
				if len(rt) != 32 {
					t.Fatalf("GetAbciInfo().state_root got = %v, length %v", rt, len(rt))
				}
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("GetAbciInfo() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestNSBClient_BroadcastTxAsync(t *testing.T) {
	type args struct {
		body []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.BroadcastTxAsync(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("BroadcastTxAsync() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BroadcastTxAsync() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_sendContractTx(t *testing.T) {
	type args struct {
		transType uint8
		txContent *nsbrpc.TransactionHeader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nsb_message.ResultInfo
		wantErr bool
	}{
		//{name: "test_easy", fields: getNormalField(), args: args{
		//	transType: transactiontype.CreateContract,
		//	txContent: HandlerError(nc.CreateContractPacket(
		//		signer,
		//		nil,
		//		nil,
		//		&nsbrpc.FAPair{
		//			FuncName: "option",
		//			Args: HandlerError(json.Marshal(&delegate.ArgsCreateNewContract{
		//				Delegates:  [][]byte{{0}},
		//				District:   "",
		//				TotalVotes: math.NewUint256FromHexString("1111"),
		//			})).([]byte),
		//		})).(*nsbrpc.TransactionHeader),
		//}, want: nil, wantErr: false},
		//iscOwners:       [][]byte{ctx.s},
		//			funds:           []uint64{0},
		//			instructions:    funcSetA(),
		//			rawInstructions: encodeInstructions(funcSetA()),
		{name: "test_isc", fields: getNormalField(), args: args{
			transType: transactiontype.CreateContract,
			txContent: HandlerError(nc.CreateContractPacket(
				signer,
				nil,
				nil,
				&nsbrpc.FAPair{
					FuncName: "isc",
					Args: HandlerError(json.Marshal(&isc.ArgsCreateNewContract{
						IscOwners:          [][]byte{signer.GetPublicKey()},
						Funds:              []uint64{0},
						VesSig:             nil,
						TransactionIntents: nil,
					})).([]byte),
				})).(*nsbrpc.TransactionHeader),
		}, want: nil, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.sendContractTx(tt.args.transType, tt.args.txContent)
			if (err != nil) != tt.wantErr {
				t.Errorf("sendContractTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			PrintStruct(got)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("sendContractTx() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestNSBClient_BroadcastTxCommit(t *testing.T) {
	type args struct {
		body []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nsb_message.ResultInfo
		wantErr bool
	}{
		// todo
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.BroadcastTxCommit(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("BroadcastTxCommit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			PrintStruct(got)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("BroadcastTxCommit() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestNSBClient_BroadcastTxCommitReturnBytes(t *testing.T) {
	type args struct {
		body []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.BroadcastTxCommitReturnBytes(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("BroadcastTxCommitReturnBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BroadcastTxCommitReturnBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_GetBlock(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nsb_message.BlockInfo
		wantErr bool
	}{
		{name: "test_get_block", fields: getNormalField(), args: args{id: 1}, want: nil, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.GetBlock(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			PrintStruct(got)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("GetBlock() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestNSBClient_GetBlockResults(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nsb_message.BlockResultsInfo
		wantErr bool
	}{
		{name: "test_get_block_result", fields: getNormalField(), args: args{id: 1}, want: nil, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.GetBlockResults(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBlockResults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			PrintStruct(got)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("GetBlock() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestNSBClient_GetBlocks(t *testing.T) {
	type args struct {
		rangeL int64
		rangeR int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nsb_message.BlocksInfo
		wantErr bool
	}{
		{name: "test_get_blocks", fields: getNormalField(), args: args{rangeL: 1, rangeR: 2}, want: nil, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.GetBlocks(tt.args.rangeL, tt.args.rangeR)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBlocks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			PrintStruct(got)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("GetBlock() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestNSBClient_GetCommitInfo(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nsb_message.CommitInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.GetCommitInfo(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCommitInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCommitInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_GetConsensusParamsInfo(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nsb_message.ConsensusParamsInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.GetConsensusParamsInfo(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConsensusParamsInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConsensusParamsInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_GetConsensusState(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		want    *nsb_message.ConsensusStateInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.GetConsensusState()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConsensusState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConsensusState() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_GetGenesis(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		want    *nsb_message.GenesisInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.GetGenesis()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGenesis() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGenesis() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_GetHealth(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.GetHealth()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHealth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHealth() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_GetNetInfo(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		want    *nsb_message.NetInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.GetNetInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNetInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNetInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_GetNumUnconfirmedTxs(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		want    *nsb_message.NumUnconfirmedTxsInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.GetNumUnconfirmedTxs()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNumUnconfirmedTxs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNumUnconfirmedTxs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_GetProof(t *testing.T) {
	type args struct {
		txHeader []byte
		subQuery string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *StorageProofResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.GetProof(tt.args.txHeader, tt.args.subQuery)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProof() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProof() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_GetStatus(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		want    *nsb_message.StatusInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.GetStatus()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_GetTransaction(t *testing.T) {
	type args struct {
		hash string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.GetTransaction(tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTransaction() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_GetUnconfirmedTxs(t *testing.T) {
	type args struct {
		limit int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nsb_message.NumUnconfirmedTxsInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.GetUnconfirmedTxs(tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUnconfirmedTxs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUnconfirmedTxs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_GetValidators(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nsb_message.ValidatorsInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.GetValidators(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetValidators() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValidators() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_preloadJSONResponse(t *testing.T) {
	type args struct {
		bb io.ReadCloser
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.preloadJSONResponse(tt.args.bb)
			if (err != nil) != tt.wantErr {
				t.Errorf("preloadJSONResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preloadJSONResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_sendContractTxAsync(t *testing.T) {
	type args struct {
		transType uint8
		txContent *nsbrpc.TransactionHeader
		option    *AsyncOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.sendContractTxAsync(tt.args.transType, tt.args.txContent, tt.args.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("sendContractTxAsync() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sendContractTxAsync() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decorateHost(t *testing.T) {
	type args struct {
		host string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test_easy", args: args{host: "127.0.0.1"}, want: "http://127.0.0.1"},
		{name: "test_easy_with_port", args: args{host: "127.0.0.1:26657"}, want: "http://127.0.0.1:26657"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decorateHost(tt.args.host); got != tt.want {
				t.Errorf("decorateHost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_AddBlockCheck(t *testing.T) {
	type args struct {
		user      uip.Signer
		toAddress []byte
		chainID   uint64
		blockID   []byte
		rootHash  []byte
		rcType    uint8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nsb_message.ResultInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.AddBlockCheck(tt.args.user, tt.args.toAddress, tt.args.chainID, tt.args.blockID, tt.args.rootHash, tt.args.rcType)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddBlockCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddBlockCheck() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_AddMerkleProof(t *testing.T) {
	type args struct {
		user       uip.Signer
		toAddress  []byte
		merkletype uint16
		rootHash   []byte
		proof      []byte
		key        []byte
		value      []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nsb_message.ResultInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.AddMerkleProof(tt.args.user, tt.args.toAddress, tt.args.merkletype, tt.args.rootHash, tt.args.proof, tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddMerkleProof() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddMerkleProof() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_BroadcastTxAsync1(t *testing.T) {
	type args struct {
		body []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.BroadcastTxAsync(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("BroadcastTxAsync() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BroadcastTxAsync() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestNSBClient_BroadcastTxCommit1(t *testing.T) {
//	type args struct {
//		body []byte
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *ResultInfo
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			nc := &NSBClient{
//				handler:    tt.fields.handler,
//				bufferPool: tt.fields.bufferPool,
//			}
//			got, err := nc.BroadcastTxCommit(tt.args.body)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("BroadcastTxCommit() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("BroadcastTxCommit() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestNSBClient_BroadcastTxCommitReturnBytes1(t *testing.T) {
//	type args struct {
//		body []byte
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    []byte
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			nc := &NSBClient{
//				handler:    tt.fields.handler,
//				bufferPool: tt.fields.bufferPool,
//			}
//			got, err := nc.BroadcastTxCommitReturnBytes(tt.args.body)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("BroadcastTxCommitReturnBytes() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("BroadcastTxCommitReturnBytes() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestNSBClient_CreateContractPacket(t *testing.T) {

	type args struct {
		s         uip.Signer
		toAddress []byte
		value     []byte
		pair      *nsbrpc.FAPair
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nsbrpc.TransactionHeader
		wantErr bool
	}{
		{name: "test_get_block_result", fields: getNormalField(), args: args{
			s:         signer,
			toAddress: nil,
			value:     nil,
			pair: &nsbrpc.FAPair{
				FuncName: "option",
				Args: HandlerError(json.Marshal(&delegate.ArgsCreateNewContract{
					Delegates:  nil,
					District:   "",
					TotalVotes: math.NewUint256FromHexString("1111"),
				})).([]byte),
			},
		}, want: nil, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NSBClient{
				handler:    tt.fields.handler,
				bufferPool: tt.fields.bufferPool,
			}
			got, err := nc.CreateContractPacket(tt.args.s, tt.args.toAddress, tt.args.value, tt.args.pair)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateContractPacket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			PrintStruct(got)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("CreateContractPacket() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
