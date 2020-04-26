package nsbcli

import (
	"github.com/HyperService-Consortium/NSB/grpc/nsbrpc"
	"github.com/HyperService-Consortium/NSB/lib/nsb-client/nsb-message"
	"github.com/HyperService-Consortium/go-uip/uip"
	"io"
	"reflect"
	"testing"
)

func TestNSBClient_CreateISC(t *testing.T) {
	type args struct {
		user                    uip.Signer
		funds                   []uint64
		iscOwners               [][]byte
		bytesTransactionIntents [][]byte
		vesSig                  []byte
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
			got, err := nc.CreateISC(tt.args.user, tt.args.funds, tt.args.iscOwners, tt.args.bytesTransactionIntents, tt.args.vesSig)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateISC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateISC() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_CreateNormalPacket(t *testing.T) {
	type args struct {
		s         uip.Signer
		toAddress []byte
		data      []byte
		value     []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nsbrpc.TransactionHeader
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
			got, err := nc.CreateNormalPacket(tt.args.s, tt.args.toAddress, tt.args.data, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateNormalPacket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateNormalPacket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_FreezeInfo(t *testing.T) {
	type args struct {
		user            uip.Signer
		contractAddress []byte
		tid             uint64
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
			got, err := nc.FreezeInfo(tt.args.user, tt.args.contractAddress, tt.args.tid)
			if (err != nil) != tt.wantErr {
				t.Errorf("FreezeInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FreezeInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_GetAbciInfo2(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		want    *nsb_message.AbciInfoResponse
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
			got, err := nc.GetAbciInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAbciInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAbciInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_GetBlock1(t *testing.T) {
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
		// TODO: Add test cases.
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBlock() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_GetBlockResults1(t *testing.T) {
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
		// TODO: Add test cases.
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBlockResults() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_GetBlocks1(t *testing.T) {
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
		// TODO: Add test cases.
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBlocks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_GetCommitInfo1(t *testing.T) {
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

func TestNSBClient_GetConsensusParamsInfo1(t *testing.T) {
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

func TestNSBClient_GetConsensusState1(t *testing.T) {
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

func TestNSBClient_GetGenesis1(t *testing.T) {
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

func TestNSBClient_GetHealth1(t *testing.T) {
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

func TestNSBClient_GetNetInfo1(t *testing.T) {
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

func TestNSBClient_GetNumUnconfirmedTxs1(t *testing.T) {
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

func TestNSBClient_GetProof1(t *testing.T) {
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

func TestNSBClient_GetStatus1(t *testing.T) {
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

func TestNSBClient_GetTransaction1(t *testing.T) {
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

func TestNSBClient_GetUnconfirmedTxs1(t *testing.T) {
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

func TestNSBClient_GetValidators1(t *testing.T) {
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

func TestNSBClient_InsuranceClaim(t *testing.T) {
	type args struct {
		user            uip.Signer
		contractAddress []byte
		tid             uint64
		aid             uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nsb_message.DeliverTx
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
			got, err := nc.InsuranceClaim(tt.args.user, tt.args.contractAddress, tt.args.tid, tt.args.aid)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsuranceClaim() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsuranceClaim() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_SettleContract(t *testing.T) {
	type args struct {
		user            uip.Signer
		contractAddress []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nsb_message.DeliverTx
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
			got, err := nc.SettleContract(tt.args.user, tt.args.contractAddress)
			if (err != nil) != tt.wantErr {
				t.Errorf("SettleContract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SettleContract() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_UserAck(t *testing.T) {
	type args struct {
		user            uip.Signer
		contractAddress []byte
		address         []byte
		signature       []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nsb_message.DeliverTx
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
			got, err := nc.UserAck(tt.args.user, tt.args.contractAddress, tt.args.address, tt.args.signature)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserAck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserAck() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_addBlockCheck(t *testing.T) {
	type args struct {
		chainID  uint64
		blockID  []byte
		rootHash []byte
		rtType   uint8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nsbrpc.FAPair
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
			got, err := nc.addBlockCheck(tt.args.chainID, tt.args.blockID, tt.args.rootHash, tt.args.rtType)
			if (err != nil) != tt.wantErr {
				t.Errorf("addBlockCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addBlockCheck() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_addMerkleProof(t *testing.T) {
	type args struct {
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
		want    *nsbrpc.FAPair
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
			got, err := nc.addMerkleProof(tt.args.merkletype, tt.args.rootHash, tt.args.proof, tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("addMerkleProof() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addMerkleProof() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_createISC(t *testing.T) {
	type args struct {
		funds                   []uint64
		iscOwners               [][]byte
		bytesTransactionIntents [][]byte
		vesSig                  []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nsbrpc.FAPair
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
			got, err := nc.createISC(tt.args.funds, tt.args.iscOwners, tt.args.bytesTransactionIntents, tt.args.vesSig)
			if (err != nil) != tt.wantErr {
				t.Errorf("createISC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createISC() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_freezeInfo(t *testing.T) {
	type args struct {
		tid uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nsbrpc.FAPair
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
			got, err := nc.freezeInfo(tt.args.tid)
			if (err != nil) != tt.wantErr {
				t.Errorf("freezeInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("freezeInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_insuranceClaim(t *testing.T) {
	type args struct {
		tid uint64
		aid uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nsbrpc.FAPair
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
			got, err := nc.insuranceClaim(tt.args.tid, tt.args.aid)
			if (err != nil) != tt.wantErr {
				t.Errorf("insuranceClaim() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insuranceClaim() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_preloadJSONResponse1(t *testing.T) {
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

func TestNSBClient_sendContractTx1(t *testing.T) {
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
		// TODO: Add test cases.
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sendContractTx() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_sendContractTxAsync1(t *testing.T) {
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

func TestNSBClient_settleContract(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		want    *nsbrpc.FAPair
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
			got, err := nc.settleContract()
			if (err != nil) != tt.wantErr {
				t.Errorf("settleContract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("settleContract() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNSBClient_userAck(t *testing.T) {
	type args struct {
		address   []byte
		signature []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nsbrpc.FAPair
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
			got, err := nc.userAck(tt.args.address, tt.args.signature)
			if (err != nil) != tt.wantErr {
				t.Errorf("userAck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userAck() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAsyncOption1(t *testing.T) {
	tests := []struct {
		name string
		want *AsyncOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAsyncOption(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAsyncOption() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNSBClient1(t *testing.T) {
	type args struct {
		host string
	}
	tests := []struct {
		name string
		args args
		want *NSBClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNSBClient(tt.args.host); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNSBClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decorateHost1(t *testing.T) {
	type args struct {
		host string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decorateHost(tt.args.host); got != tt.want {
				t.Errorf("decorateHost() = %v, want %v", got, tt.want)
			}
		})
	}
}
