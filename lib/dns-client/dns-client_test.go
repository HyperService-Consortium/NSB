package dns_client

import (
	ChainType "github.com/HyperService-Consortium/go-uip/const/chain_type"
	merkle_proof "github.com/HyperService-Consortium/go-uip/const/merkle-proof-type"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDNSClient_Health(t *testing.T) {

	reply := sugar.HandlerError(NewDNSClient("http://localhost:26668").Health()).(*HealthReply)
	assert.EqualValues(t, 0, reply.Code)
}

func TestDNSClient_GetChainInfo(t *testing.T) {

	reply := sugar.HandlerError(
		NewDNSClient("http://localhost:26668").GetChainInfo(
			1)).(*ChainInfoReply)
	assert.EqualValues(t, ChainType.Ethereum, reply.ChainType)
	assert.EqualValues(t, merkle_proof.SecureMerklePatriciaTrieUsingKeccak256,
		reply.MerkleProofType)
}