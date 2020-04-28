package nsb

import (
	"encoding/hex"
	_ "encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/HyperService-Consortium/NSB/application/nsb_proto"
	"github.com/HyperService-Consortium/NSB/common"
	"github.com/HyperService-Consortium/NSB/contract/delegate"
	"github.com/HyperService-Consortium/NSB/lib/prover"
	"github.com/HyperService-Consortium/NSB/math"
	"github.com/HyperService-Consortium/NSB/merkmap"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/abci/types"
	"testing"
)

func createNewDelegateTestPacket(t *testing.T, signer uip.Signer) (tx types.RequestDeliverTx, ok bool) {
	var packet delegate.ArgsCreateNewContract
	packet.Delegates = [][]byte{signer.GetPublicKey()}
	packet.District = "chain1"
	packet.TotalVotes = math.NewUint256FromBytes([]byte{0})

	var fap FAPair
	fap.FuncName = "delegate"
	fap.Args = sugar.HandlerError(
		json.Marshal(packet)).([]byte)

	b, ok := pack(t, signer, fap)

	return types.RequestDeliverTx{
		Tx: b,
	}, ok
}

func assertEqualProof(t *testing.T, nsb *NSBApplication,
	addr []byte, key []byte, value []byte) {

	queryResp := nsb.Query(types.RequestQuery{
		Data: addr,
		Path: QueryKeyGetAccInfo,
	})

	assert.Equal(t, uint32(0), queryResp.Code)

	var accInfo common.AccountInfo
	sugar.HandlerError0(json.Unmarshal(
		[]byte(queryResp.Info), &accInfo))

	proofResp := nsb.Query(types.RequestQuery{
		Data: sugar.HandlerError(
			json.Marshal(&nsb_proto.ArgsGetStorageAt{
				Address: addr,
				Key:     key,
				Slot:    prover.NSBGetSlot(addr, nil),
			})).([]byte),
		Path: QueryKeyGetStorageAt,
	})

	assert.Equal(t, uint32(0), proofResp.Code)

	var proof merkmap.ProofJson
	sugar.HandlerError0(json.Unmarshal(
		[]byte(proofResp.Info), &proof))

	assert.Equal(t, "", proof.Log)
	assert.EqualValues(t, key, proof.Key)
	assert.EqualValues(t, value, proof.Value)

	assert.EqualValues(t, accInfo.StorageRoot, proof.Proof[0])
	provedValue, err := prover.GetMerkleProofValueWithValidateNSBMPT(
		proof.Proof[0], proof.Proof[1:],
		prover.NSBKeyByPure(
			addr, []byte{}, key))
	assert.NoError(t, err)
	assert.EqualValues(t, value, provedValue)
}

func TestNSBContractDelegate(t *testing.T) {
	nsb := createApplication(t, "./data/")
	if nsb == nil {
		return
	}

	signer := createSigner(t)
	assert.NotNil(t, signer)

	tx, ok := createNewDelegateTestPacket(t, signer)
	assert.True(t, ok)
	response := nsb.DeliverTx(tx)

	assert.Equal(t, uint32(0), response.Code)
	assert.Equal(t, fmt.Sprintf(
		"create success , this contract is deploy at %v",
		hex.EncodeToString(response.Data),
	), response.Info)

	addr := response.Data
	nsb.Commit()

	queryResp := nsb.Query(types.RequestQuery{
		Data: addr,
		Path: QueryKeyGetAccInfo,
	})

	assert.Equal(t, uint32(0), queryResp.Code)

	var accInfo common.AccountInfo
	sugar.HandlerError0(json.Unmarshal(
		[]byte(queryResp.Info), &accInfo))

	assert.EqualValues(t, "delegate", accInfo.Name)

	assertEqualProof(t, nsb, addr,
		[]byte("totalVotes"), math.NewUint256FromBytes([]byte{0}).Bytes())
	assertEqualProof(t, nsb, addr,
		[]byte("district"), []byte("chain1"))
}
