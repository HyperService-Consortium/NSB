package nsb

import (
	dbm "github.com/tendermint/tendermint/libs/db"
	"encoding/json"
	"fmt"
)


type MerkleProofType uint8;
const (
	EthereumMerkleProof MerkleProofType = 0 + iota
	NebulasMerkleProof
	TendermintMerkleProof
)


type MerkleProof struct {
	Mtype       MerkleProofType     `json:"merkle_proof_type"`
	ChainId     string              `json:"chain_id"`
	StorageHash []byte              `json:"storage_hash"`
	key         []byte              `json:"key"`
	value       []byte              `json:"value"`
}


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


func checkEthMerkleProof(proof *MerkleProof) bool {

}

func addMerkleProof(byteJson []byte) (addSuccess bool ,err error) {
	addSuccess = false
	
	var proof MerkleProof
	err = json.Unmarshal(byteJson, proof)
	if err != nil {
		return
	}

	fmt.Println()

}
// function addMerkleProof(string memory blockaddr, bytes32 storagehash, bytes32 key, bytes32 val)
// 	public
// 	//ownerExists(msg.sender)
// 	validMerkleProof(storagehash, key)
// 	returns (bytes32 keccakhash)
// {
// 	MerkleProof memory toAdd = MerkleProof(blockaddr, storagehash, key, val);
// 	keccakhash = keccak256(abi.encodePacked(blockaddr, storagehash, key, val));

// 	require(MerkleProofTree[keccakhash].storagehash == 0, "already in MerkleProofTree");

// 	proofPointer[keccakhash] = uint32(waitingVerifyProof.length);
// 	waitingVerifyProof.push(keccakhash);
// 	MerkleProofTree[keccakhash] = toAdd;
// 	validCount.length ++;
// 	votedCount.length ++;

// 	emit addingMerkleProof(blockaddr, storagehash, key, val);
// }

// function voteProofByHash(bytes32 keccakhash, bool validProof)
// 	public
// 	ownerExists(msg.sender)
// 	validVoteByHash(msg.sender, keccakhash)
// {
// 	uint32 curPointer = proofPointer[keccakhash];

// 	//update counts
// 	if(validProof) {
// 		validCount[curPointer] ++;
// 	}
// 	votedCount[curPointer] ++;

// 	//judge if there is enough owners voted
// 	if (votedCount[curPointer] == requiredOwnerCount) {
// 		if (validCount[curPointer] >= requiredValidVotesCount) {
// 			verifiedMerkleProof[keccakhash] = true;
// 			CallbackPair storage cb = proofHashCallback[keccakhash];
// 			if(cb.isc_addr != address(0)) {
// 				txsReference[cb.isc_addr].txInfo[cb.tx_index].proofHash.push(keccakhash);
// 			}
// 		} else {
// 			delete MerkleProofTree[keccakhash];
// 		}
// 		delete waitingVerifyProof[curPointer];
// 		delete validCount[curPointer];
// 		delete votedCount[curPointer];
// 	}

// 	//update the pointer
// 	while(votedPointer < waitingVerifyProof.length &&
// 		waitingVerifyProof[votedPointer] == 0)votedPointer ++;
// }

// function voteProofByPointer(uint32 curPointer, bool validProof)
// 	public
// 	ownerExists(msg.sender)
// 	validVoteByPointer(msg.sender, curPointer)
// {
// 	//update counts
// 	if(validProof) {
// 		validCount[curPointer] ++;
// 	}
// 	votedCount[curPointer] ++;

// 	//judge if there is enough owners voted
// 	if (votedCount[curPointer] == requiredOwnerCount) {
// 		if (validCount[curPointer] >= requiredValidVotesCount) {
// 			verifiedMerkleProof[waitingVerifyProof[curPointer]] = true;
// 			CallbackPair storage cb = proofHashCallback[waitingVerifyProof[curPointer]];
// 			if(cb.isc_addr != address(0)) {
// 				txsReference[cb.isc_addr].txInfo[cb.tx_index].proofHash.push(waitingVerifyProof[curPointer]);
// 			}
// 		} else {
// 			delete MerkleProofTree[waitingVerifyProof[curPointer]];
// 		}
// 		delete waitingVerifyProof[curPointer];
// 		delete validCount[curPointer];
// 		delete votedCount[curPointer];
		
// 	}

// 	//update the pointer
// 	while(votedPointer < waitingVerifyProof.length &&
// 		waitingVerifyProof[votedPointer] == 0)votedPointer ++;
// }

// function getMerkleProofByHash(bytes32 keccakhash)
// 	public
// 	view
// 	ownerExists(msg.sender)
// 	returns (string memory a, bytes32 s, bytes32 k, bytes32 v)
// {
// 	MerkleProof storage toGet = MerkleProofTree[keccakhash];
// 	a = toGet.blockaddr;
// 	s = toGet.storagehash;
// 	k = toGet.key;
// 	v = toGet.value;
// }

// function getMerkleProofByPointer(uint32 curPointer)
// 	public
// 	view
// 	ownerExists(msg.sender)
// 	returns (string memory a, bytes32 s, bytes32 k, bytes32 v)
// {
// 	MerkleProof storage toGet = MerkleProofTree[waitingVerifyProof[curPointer]];
// 	a = toGet.blockaddr;
// 	s = toGet.storagehash;
// 	k = toGet.key;
// 	v = toGet.value;
// }

// function validMerkleProoforNot(bytes32 keccakhash)
//          external
//         view
//         returns (bool)
//     {
//         return verifiedMerkleProof[keccakhash];
//     }

//     function getVaildMerkleProof(bytes32 keccakhash)
//         public
//         view
//         returns (string memory a, bytes32 s, bytes32 k, bytes32 v)
//     {
//         require(verifiedMerkleProof[keccakhash] == true, "invalid");
//         MerkleProof storage toGet = MerkleProofTree[keccakhash];
//         a = toGet.blockaddr;
//         s = toGet.storagehash;
//         k = toGet.key;
//         v = toGet.value;
//     }

// function getVaildMerkleProof(bytes32 keccakhash)
//         public
//         view
//         returns (string memory a, bytes32 s, bytes32 k, bytes32 v)
//     {
//         require(verifiedMerkleProof[keccakhash] == true, "invalid");
//         MerkleProof storage toGet = MerkleProofTree[keccakhash];
//         a = toGet.blockaddr;
//         s = toGet.storagehash;
//         k = toGet.key;
//         v = toGet.value;
//     }