package nsb

import (
	dbm "github.com/tendermint/tendermint/libs/db"
	"encoding/json"
)

func getMerkleProofByHash(db dbm.DB, prvHash []byte) MerkleProof {
	proofBytes := db.Get(prvHash)
	var merkleProof MerkleProof
	if len(proofBytes) != 0 {
		err := json.Unmarshal(proofBytes, &merkleProof)
		if err != nil {
			panic(err)
		}
	}
	return merkleProof
}