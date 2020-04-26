package nsb

import (
	_ "encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/HyperService-Consortium/NSB/application/nsb_proto"
	"github.com/HyperService-Consortium/NSB/common"
	"github.com/HyperService-Consortium/NSB/lib/prover"
	"github.com/HyperService-Consortium/NSB/merkmap"
	"github.com/HyperService-Consortium/go-uip/isc"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/abci/types"
	"testing"
)

func TestNSBApplication_Query(t *testing.T) {
	nsb := createApplication(t, "./data/")
	if nsb == nil {
		return
	}

	signer := createSigner(t)
	assert.NotNil(t, signer)
	uu := createU(t)
	assert.NotNil(t, uu)
	vv := createV(t)
	assert.NotNil(t, vv)

	tx, ok := createISCTestPacket(t, signer, uu, vv)
	assert.True(t, ok)
	response := nsb.DeliverTx(tx)

	assert.Equal(t, "", response.Log)
	assert.Equal(t, "", response.Info)
	var contractReply isc.NewContractReply
	sugar.HandlerError0(json.Unmarshal(response.Data, &contractReply))
	addr := contractReply.Address
	nsb.Commit()

	queryResp := nsb.Query(types.RequestQuery{
		Data: addr,
		Path: QueryKeyGetAccInfo,
	})

	assert.Equal(t, uint32(0), queryResp.Code)

	var accInfo common.AccountInfo
	sugar.HandlerError0(json.Unmarshal(
		[]byte(queryResp.Info), &accInfo))

	assert.EqualValues(t, "isc", accInfo.Name)

	proofResp := nsb.Query(types.RequestQuery{
		Data: sugar.HandlerError(
			json.Marshal(&nsb_proto.ArgsGetStorageAt{
				Address: addr,
				Key:     signer.GetPublicKey(),
				Slot:    prover.NSBGetSlot(addr, []byte("IsOwner")),
			})).([]byte),
		Path: QueryKeyGetStorageAt,
	})

	assert.Equal(t, uint32(0), proofResp.Code)

	var proof merkmap.ProofJson
	sugar.HandlerError0(json.Unmarshal(
		[]byte(proofResp.Info), &proof))

	assert.Equal(t, "", proof.Log)

	assert.EqualValues(t, signer.GetPublicKey(), proof.Key)
	fmt.Println(proof.Value)
	assert.EqualValues(t, accInfo.StorageRoot, proof.Proof[0])
	value, err := prover.GetMerkleProofValueWithValidateNSBMPT(
		proof.Proof[0], proof.Proof[1:],
		prover.NSBKeyByPure(
			addr, []byte("IsOwner"), signer.GetPublicKey()))
	assert.NoError(t, err)
	assert.EqualValues(t, proof.Value, value)
}
