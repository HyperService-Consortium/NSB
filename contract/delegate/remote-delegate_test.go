package delegate_test

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	application "github.com/HyperService-Consortium/NSB/application"
	"github.com/HyperService-Consortium/NSB/application/nsb_proto"
	"github.com/HyperService-Consortium/NSB/common"
	nsbcli "github.com/HyperService-Consortium/NSB/lib/nsb-client"
	"github.com/HyperService-Consortium/NSB/lib/prover"
	"github.com/HyperService-Consortium/NSB/math"
	"github.com/HyperService-Consortium/NSB/merkmap"
	"github.com/HyperService-Consortium/go-uip/signaturer"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func assertEqualProof(t *testing.T, nsb *nsbcli.NSBClient,
	addr []byte, key []byte, value []byte) {
	queryResp, err := nsb.GetQuery(addr, application.QueryKeyGetAccInfo)
	assert.NoError(t, err)

	assert.Equal(t, uint32(0), queryResp.Response.Code)

	var accInfo common.AccountInfo
	sugar.HandlerError0(json.Unmarshal(
		[]byte(queryResp.Response.Info), &accInfo))

	assert.EqualValues(t, "delegate", accInfo.Name)

	proofResp, err := nsb.GetQuery(sugar.HandlerError(
		json.Marshal(&nsb_proto.ArgsGetStorageAt{
			Address: addr,
			Key:     key,
			Slot:    prover.NSBGetSlot(addr, nil),
		})).([]byte), application.QueryKeyGetStorageAt)
	assert.NoError(t, err)

	assert.Equal(t, uint32(0), proofResp.Response.Code)

	var proof merkmap.ProofJson
	sugar.HandlerError0(json.Unmarshal(
		[]byte(proofResp.Response.Info), &proof))

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

func TestRemoteCreate(t *testing.T) {
	nsb := nsbcli.NewNSBClient("http://121.89.200.234:26657")
	signer, err := signaturer.NewTendermintNSBSigner(
		sugar.HandlerError(hex.DecodeString(
			"2333bbffffffffffffff2333bbffffffffffffff2333bbffffffffffffffffff2333bbffffffffffffff2333bbffffffffffffff2333bbffffffffffffffffff")).([]byte))
	assert.NoError(t, err)
	addr, err := nsb.CreateDelegate(signer, [][]byte{
		signer.GetPublicKey(),
	}, "chain1", math.NewUint256FromBytes([]byte{1}))

	assert.NoError(t, err)
	fmt.Println(hex.EncodeToString(addr))
	queryResp, err := nsb.GetQuery(addr, application.QueryKeyGetAccInfo)
	assert.NoError(t, err)

	assert.Equal(t, uint32(0), queryResp.Response.Code)

	var accInfo common.AccountInfo
	sugar.HandlerError0(json.Unmarshal(
		[]byte(queryResp.Response.Info), &accInfo))

	assert.EqualValues(t, "delegate", accInfo.Name)

	assertEqualProof(t, nsb, addr,
		[]byte("totalVotes"), math.NewUint256FromBytes([]byte{1}).Bytes())
	assertEqualProof(t, nsb, addr,
		[]byte("district"), []byte("chain1"))
}

func TestRemoteVote(t *testing.T) {
	nsb := nsbcli.NewNSBClient("http://121.89.200.234:26657")
	signer, err := signaturer.NewTendermintNSBSigner(
		sugar.HandlerError(hex.DecodeString(
			"2333bbffffffffffffff2333bbffffffffffffff2333bbffffffffffffffffff2333bbffffffffffffff2333bbffffffffffffff2333bbffffffffffffffffff")).([]byte))
	assert.NoError(t, err)
	addr, err := nsb.CreateDelegate(signer, [][]byte{
		signer.GetPublicKey(),
	}, "chain1", math.NewUint256FromBytes([]byte{1}))
	assert.NoError(t, err)

	res, err := nsb.DelegateVote(signer, addr)
	assert.NoError(t, err)

	assert.EqualValues(t, math.NewUint256FromBytes([]byte{2}),
		math.NewUint256FromBytes(res.DeliverTx.Data))

	queryResp, err := nsb.GetQuery(addr, application.QueryKeyGetAccInfo)
	assert.NoError(t, err)

	assert.Equal(t, uint32(0), queryResp.Response.Code)

	var accInfo common.AccountInfo
	sugar.HandlerError0(json.Unmarshal(
		[]byte(queryResp.Response.Info), &accInfo))

	assert.EqualValues(t, "delegate", accInfo.Name)

	assertEqualProof(t, nsb, addr,
		[]byte("totalVotes"), res.DeliverTx.Data)
	assertEqualProof(t, nsb, addr,
		[]byte("district"), []byte("chain1"))
}

func TestRemoteStorage(t *testing.T) {
	//	c3953ad10a4a01e197f1e2dd7424c7c9126d8d8df334e867789f1c83b9c3ca46
}
