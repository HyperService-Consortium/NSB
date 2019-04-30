package merkleproof

import (
	// dbm "github.com/tendermint/tendermint/libs/db"
	"encoding/json"
	"fmt"
	"github.com/Myriad-Dreamin/NSB/contract/isc/merkleproof/MerkleProofType"
	"github.com/Myriad-Dreamin/NSB/contract/isc/merkleproof/MerkleProofError"
)

type MerkleProof struct {
	Mtype       MerkleProofType.Type        `json:"merkle_proof_type"`
	ChainId     string                      `json:"chain_id"`
	StorageHash []byte                      `json:"storage_hash"`
	Key         []byte                      `json:"key"`
	Value       []byte                      `json:"value"`
}


// func getMerkleProofByHash(db dbm.DB, prvHash []byte) MerkleProof {
// 	proofBytes := db.Get(prvHash)
// 	var merkleProof MerkleProof
// 	if len(proofBytes) != 0 {
// 		err := json.Unmarshal(proofBytes, &merkleProof)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// 	return merkleProof
// }

// func validMerkleProoforNot(db.dbm.DB, prvHash []byte) bool {

// }

func checkEthereumMerkleProof(proof *MerkleProof) (retproof *MerkleProof, err error) {
	fmt.Println("adding EthereumMerkleProof", proof)
	return nil, nil
}

func checkNebulasMerkleProof(proof *MerkleProof) (retproof *MerkleProof, err error) {
	fmt.Println("adding NebulasMerkleProof", proof)
	return nil, nil
}

func checkTendermintMerkleProof(proof *MerkleProof) (retproof *MerkleProof, err error) {
	fmt.Println("adding TendermintMerkleProof", proof)
	return nil, nil
}

func AddMerkleProof(byteJson []byte) (*MerkleProof, error) {
	var proof = new(MerkleProof)
	
	err := json.Unmarshal(byteJson, proof)
	if err != nil {
		return nil, err
	}

	switch proof.Mtype {

	case MerkleProofType.EthereumMerkleProof:
		return checkEthereumMerkleProof(proof)

	case MerkleProofType.NebulasMerkleProof:
		return checkNebulasMerkleProof(proof)

	case MerkleProofType.TendermintMerkleProof:
		return checkTendermintMerkleProof(proof)

	default:
		return nil, MerkleProofError.UnrecognizedType
	}
}
