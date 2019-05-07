package merkleproof

import (
	"testing"
	"fmt"
	"encoding/json"
	"github.com/Myriad-Dreamin/NSB/contract/isc/merkleproof/MerkleProofType"
	"github.com/Myriad-Dreamin/NSB/contract/isc/merkleproof/MerkleProofError"
)


func TestAddEthMerkleProof(t *testing.T) {
	// ethereum merkle proof test
	var merkleProof = &MerkleProof{
		Mtype: MerkleProofType.EthereumMerkleProof,
		ChainId: "BuptChainA",
		StorageHash: []byte("\x12\x34\x56\x78\x12\x34\x56\x78\x12\x34\x56\x78\x12\x34\x56\x78\x12\x34\x56\x78\x12\x34\x56\x78\x12\x34\x56\x78\x12\x34\x56\x78"),
		Key: []byte("\x12\x34\x56\x78\x12\x34\x56\x78\x12\x34\x56\x78\x12\x34\x56\x78\x12\x34\x56\x78\x12\x34\x56\x78\x12\x34\x56\x78\x12\x34\x56\x78"),
		Value: []byte("\x12\x34\x56\x78\x12\x34\x56\x78\x12\x34\x56\x78\x12\x34\x56\x78\x12\x34\x56\x78\x12\x34\x56\x78\x12\x34\x56\x78\x12\x34\x56\x78")}	
	
	var jsonMerkleProof, err = json.Marshal(merkleProof)
	if err != nil {
		t.Error(err)
		return
	}
	var proof *MerkleProof
	proof, err = AddMerkleProof(jsonMerkleProof)
	fmt.Println(proof)
	if err != nil {
		t.Error(err)
	}
}

func TestInvalidMerkleProof(t *testing.T) {
	// invalid merkle proof test
	var merkleProof = &MerkleProof{
		Mtype: 5,
		ChainId: "BuptChainA",
		StorageHash: []byte(""),
		Key: []byte(""),
		Value: []byte("")}	
	
	var jsonMerkleProof, err = json.Marshal(merkleProof)
	if err != nil {
		t.Error(err)
		return
	}
	
	var proof *MerkleProof
	proof, err = AddMerkleProof(jsonMerkleProof)
	fmt.Println(proof)
	if err != nil && err != MerkleProofError.UnrecognizedType {
		t.Error(err)
	}
}