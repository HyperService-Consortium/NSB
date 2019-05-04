package nsb

import (
	"encoding/json"
	"github.com/Myriad-Dreamin/NSB/application/response"
	cmn "github.com/Myriad-Dreamin/NSB/common"
	"github.com/Myriad-Dreamin/NSB/crypto"
	"github.com/Myriad-Dreamin/NSB/util"
	"github.com/tendermint/tendermint/abci/types"
)

/*
 * storage := actionMap
 */

func MustUnmarshal(data []byte, load interface{}) {
	err := json.Unmarshal(data, &load)
	if err != nil {
		panic(response.DecodeJsonError(err))
	}
}

func (nsb *NSBApplication) MerkleProofRigisteredMethod(
	env *cmn.TransactionHeader,
	accInfo *AccountInfo,
	funcName string,
	args []byte,
) *types.ResponseDeliverTx {
	switch funcName {
	case "validateMerkleProof":
		return nsb.validateMerkleProof(args)
	case "addMerkleProof":
		return nsb.addMerkleProof(args)
	case "getMerkleProof":
		return nsb.getMerkleProof(args)
	default:
		return response.InvalidFuncTypeError(MethodMissing)
	}
}


const (
	simpleMerkleTreeUsingSha256 uint8 = 0 + iota
	simpleMerkleTreeUsingSha512
	MerklePatriciaTrieUsingKeccak256
)

var (
	unrecognizedMerkleProofType = errors.New("unknown merkle proof type")
	evenlenSimpleMerkleProofError = errors.New(
		"MerkleProofError: simple merkle proof must have an odd number of hash nodes",
	)
	wrongSimpleMerkleTreeHash = errors.New(
		"MerkleProofError: fail to match the given hash value"
	)
	unrecognizedHashFuncType = errors.New(
		"unknown hash function type"
	)
)


type ArgsValidateMerkleProof struct {
	Type  uint8  `json:"1"`
	Proof []byte `json:"2"`
	Key   []byte `json:"3"`
	Value []byte `json:"4"`
}

func validateMerkleProofKey(typeId uint8, multiKey ...[]byte) []byte {
	return crypto.Sha512([]byte{typeId}, multiKey...)
}

func (nsb *NSBApplication) validateMerkleProof(bytesArgs []byte) *types.ResponseDeliverTx {
	var args ArgsValidateMerkleProof
	MustUnmarshal(bytesArgs, &args)
	switch args.Type {
	case simpleMerkleTreeUsingSha256, simpleMerkleTreeUsingSha512:
		return nsb.validateSimpleMerkleTree(args.Proof, args.Key, args.Type)
	case MerklePatriciaTrieUsingKeccak256:
		return nsb.validateMerklePatriciaTrie(args.Proof, args.Key, args.Value, args.Type)
	default:
		return response.ContractExecError(unrecognizedMerkleProofType)
	}
}

type SimpleMerkleProof struct {
	HashChain [][]byte `json:"h"`
}

func (nsb *NSBApplication) validateSimpleMerkleTree(
	Proof []byte,
	Key []byte,
	hfType uint8,
) *types.ResponseDeliverTx {
	var jsonProof SimpleMerkleProof
	MustUnmarshal(Proof, &jsonProof)
	if (len(jsonProof.HashChain) & 1) == 0 {
		return response.ContractExecError(evenlenSimpleMerkleProofError)
	}
	
	var hf crypto.HashFunc
	switch hfType {
	case simpleMerkleTreeUsingSha256:
		hf = crypto.Sha256
	case simpleMerkleTreeUsingSha512:
		hf = crypto.Sha512
	default:
		return response.ContractExecError(unrecognizedHashFuncType)
	}

	hashChain := append(append(jsonProof.HashChain, Key), []byte{})

	for idx := len(hashChain) - 2; idx >= 0; idx -= 2 {
		checkHash = hf(hashChain[idx], hashChain[idx + 1])
		if checkHash != hashChain[idx - 1] {
			return response.ContractExecError(wrongSimpleMerkleTreeHash)
		}
	}

	err := nsb.validMerkleProofMap.TryUpdate(
		merkleProofKey(hfType, HashChain[0]),
		Key,
	)
	if err != nil {
		return response.ContractExecError(err)
	}

	return &types.ResponseDeliverTx{
		Code: uint32(response.CodeOK()),
		Info: "nice!",
	}
}

func (nsb *NSBApplication) validateMerklePatriciaTrie(
	Proof []byte,
	Key []byte,
	Value []byte
) *types.ResponseDeliverTx {

	getEthKey

	return &types.ResponseDeliverTx{
		Code: uint32(response.CodeOK()),
		Info: "nice!",
	}
}


type ArgsAddMerkleProof struct {
	ISCAddress []byte `json:"1"`
	// hexbytes
	Tid uint64 `json:"2"`
	// hexbytes
	Aid       uint64 `json:"3"`
	Type      uint8  `json:"4"`
	Content   []byte `json:"5"`
	Signature []byte `json:"6"`
}

func merkleProofKey(addr []byte, tid uint64, aid uint64) []byte {
	return crypto.Sha512(addr, util.Uint64ToBytes(tid), util.Uint64ToBytes(aid))
}

func (nsb *NSBApplication) addMerkleProof(bytesArgs []byte) *types.ResponseDeliverTx {
	var args ArgsAddMerkleProof
	MustUnmarshal(bytesArgs, &args)
	// TODO: check valid isc/tid/aid
	err := nsb.validOnchainMerkleProofMap.TryUpdate(
		merkleProofKey(args.ISCAddress, args.Tid, args.Aid),
		util.ConcatBytes([]byte{args.Type}, args.Content, args.Signature),
	)
	if err != nil {
		return response.ContractExecError(err)
	}
	return &types.ResponseDeliverTx{
		Code: uint32(response.CodeOK()),
		Info: "updateSuccess",
	}
}

type ArgsGetMerkleProof struct {
	ISCAddress []byte `json:"1"`
	// hexbytes
	Tid uint64 `json:"2"`
	// hexbytes
	Aid uint64 `json:"3"`
}

func (nsb *NSBApplication) getMerkleProof(bytesArgs []byte) *types.ResponseDeliverTx {
	var args ArgsGetMerkleProof
	MustUnmarshal(bytesArgs, &args)
	// TODO: check valid isc/tid/aid
	bt, err := nsb.validOnchainMerkleProofMap.TryGet(merkleProofKey(args.ISCAddress, args.Tid, args.Aid))
	if err != nil {
		return response.ContractExecError(err)
	}
	return &types.ResponseDeliverTx{
		Code: uint32(response.CodeOK()),
		Data: bt,
	}
}

// function addTransactionProposal(address isc_addr, uint tx_count)
// 	public
// 	returns (bool addingSuccess)
// {
// 	// InsuranceSmartContract isc = InsuranceSmartContract(isc_addr);
// 	// require(isc.isRawSender(msg.sender), "you have no access to upload ISC to NSB");
// 	// addingSuccess = false;
// 	txsStack.length++;
// 	Transactions storage txs = txsStack[txsStack.length - 1];
// 	txs.txInfo.length = tx_count;
// 	txs.contract_addr = isc_addr;
// 	txsReference[isc_addr] = txsStack[txsStack.length - 1];
// 	// for(uint idx=0; idx < txs.txInfo.length; idx++)
// 	// {
// 	//     txs.txInfo[idx].txhash = isc.getTxInfoHash(idx);
// 	// }

// 	activeISC[isc_addr] = true;
// 	addingSuccess = true;
// 	emit addISCSuccess(isc_addr, tx_count);
// }

// function addMerkleProofProposal(
// 	address isc_addr,
// 	uint txindex,
// 	string memory blockaddr,
// 	bytes32 storagehash,
// 	bytes32 key,
// 	bytes32 val
// )
// 	public
// 	returns (bytes32 keccakhash)
// {
// 	require(activeISC[isc_addr], "this isc is not active now");
// 	require(txsReference[isc_addr].txInfo.length > txindex, "index overflow");
// 	// InsuranceSmartContract isc = InsuranceSmartContract(isc_addr);
// 	// require(isc.isTransactionOwner(msg.sender, txindex), "you have no access to update the merkle proof");
// 	keccakhash = addMerkleProof(blockaddr, storagehash, key, val);
// 	proofHashCallback[keccakhash] = CallbackPair(isc_addr, txindex);
// }

// function addActionProposal(
// 	address isc_addr,
// 	uint txindex,
// 	uint actionindex,
// 	bytes32 msghash,
// 	bytes memory signature
// )
// 	public
// 	returns (bytes32 keccakhash)
// {
// 	require(activeISC[isc_addr], "this isc is not active now");
// 	// InsuranceSmartContract isc = InsuranceSmartContract(isc_addr);
// 	// assert isc.isTransactionOwner(msg.sender, txindex, actionindex)
// 	// assert actionindex < actionHash.length
// 	Transactions storage txs = txsReference[isc_addr];
// 	require(txs.txInfo.length > txindex, "index overflow");
// 	if (actionindex >= txs.txInfo[txindex].actionHash.length) {
// 		txs.txInfo[txindex].actionHash.length = actionindex + 1;
// 	}
// 	keccakhash = txs.txInfo[txindex].actionHash[actionindex] = addAction(msghash, signature);
// }

// function closeTransaction(address isc_addr)
// 	public
// 	returns (bool closeSuccess)
// {
// 	// InsuranceSmartContract isc = InsuranceSmartContract(isc_addr);
// 	closeSuccess = false;
// 	// require(isc.closed(), "ISC is active now");
// 	activeISC[isc_addr] = false;
// 	closeSuccess = true;
// }
