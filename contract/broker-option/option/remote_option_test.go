package option_test

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	application "github.com/HyperService-Consortium/NSB/application"
	"github.com/HyperService-Consortium/NSB/application/nsb_proto"
	"github.com/HyperService-Consortium/NSB/common"
	"github.com/HyperService-Consortium/NSB/contract/broker-option/option"
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

	assert.EqualValues(t, "option", accInfo.Name)

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
	rawReply, err := nsb.CreateOptionContract(
		signer, math.NewUint256FromBytes([]byte{11}), signer.GetPublicKey(), math.NewUint256FromBytes([]byte{2}))

	assert.NoError(t, err)

	var reply option.NewContractReply
	assert.NoError(t, json.Unmarshal(rawReply, &reply))
	addr := reply.Address
	fmt.Println(hex.EncodeToString(addr))

	queryResp, err := nsb.GetQuery(addr, application.QueryKeyGetAccInfo)
	assert.NoError(t, err)

	assert.Equal(t, uint32(0), queryResp.Response.Code)

	var accInfo common.AccountInfo
	sugar.HandlerError0(json.Unmarshal(
		[]byte(queryResp.Response.Info), &accInfo))

	assert.EqualValues(t, "option", accInfo.Name)
	//
	//assert.EqualValues(t, user0, option.GetOwner())
	//assert.EqualValues(t, math.NewUint256FromBytes([]byte{10}), option.GetMinStake())
	//assert.EqualValues(t, _2, option.GetStrikePrice())
	//assert.EqualValues(t, _11, option.GetRemainingFund())
	assertEqualProof(t, nsb, addr,
		[]byte("owner"), signer.GetPublicKey())
	assertEqualProof(t, nsb, addr,
		[]byte("minStake"), math.NewUint256FromBytes([]byte{10}).Bytes())
	assertEqualProof(t, nsb, addr,
		[]byte("strikePrice"), math.NewUint256FromBytes([]byte{2}).Bytes())
	assertEqualProof(t, nsb, addr,
		[]byte("remainingFund"), math.NewUint256FromBytes([]byte{11}).Bytes())
}
