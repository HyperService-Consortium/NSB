package nsb

import (
	"fmt"
	"encoding/json"
	"github.com/Myriad-Dreamin/NSB/crypto"
	"github.com/Myriad-Dreamin/NSB/account"
	"github.com/Myriad-Dreamin/NSB/localstorage"
	"github.com/Myriad-Dreamin/NSB/application/response"
)


func (nsb *NSBApplication) parseFuncTransaction(tx []byte) types.ResponseDeliverTx {
	bytesTx := bytes.Split(tx, []byte("\x18"))
	if len(bytesTx) != 2 {
		return response.InvalidTxInputFormatWrongx18
	}
	byteInfo, err := nsb.txMap.TryGet(bytesTx[1])
	// internal error
	if err != nil {
		return response.ReTrieveTxError(err)
	}
	if byteInfo != nil {
		return response.DuplicateTxError
	}
	err = nsb.txMap.TryUpdate()
	// internal error
	if err != nil {
		return response.UpdateTxTreeError(err)
	}

	var txHeader TransactionHeader
	err = json.Unmarshal(bytesTx[1], &txHeader)
	if err != nil {
		return response.DecodeTxHeaderError(err)
	}

	// TODO: verify signature 

	byteInfo, err = nsb.accMap.TryGet(txHeader.From)
	if err != nil {
		return response.ReTrieveTxError(err)
	}

	var accInfo AccountInfo
	err = json.Unmarshal(byteInfo, &accInfo)
	if err != nil {
		return response.DecodeAccountInfoError(err)
	}

	byteInfo, err = nsb.accMap.TryGet(txHeader.ContractAddress)
	if err != nil {
		return response.ReTrieveTxError(err)
	}
	if byteInfo == nil {
		return response.MissingContract
	}

	var contractInfo AccountInfo
	err = json.Unmarshal(byteInfo, &contractInfo)
	if err != nil {
		return response.DecodeAccountInfoError(err)
	}
	// TODO: Check CodeHash

	var contractEnv = ContractEnvironment{
		From: txHeader.From,
		fromInfo: accInfo,
		ContractAddress: txHeader.ContractAddress,
		toInfo: contractInfo,
		Data: txHeader.JsonParas,
		Value: txHeader.Value,
	}
	contractEnv.Storage, err = localstorage.NewLocalStorage(
		txHeader.ContractAddress,
		contractInfo.StorageRoot,
		nsb.statedb,
	)

	// internal error
	if err != nil {
		return response.RequestStorageError(err)
	}

	return nsb.endFuncTransaction(nsb.execContractFuncs(bytesTx[0], contractEnv))
}


func (nsb *NSBApplication) parseCreateTransaction(tx []byte) types.ResponseDeliverTx {
	bytesTx := bytes.Split(tx, []byte("\x18"))
	if len(bytesTx) != 2 {
		return response.InvalidTxInputFormatWrongx18
	}
	return nsb.createContracts(bytesTx[0], bytesTx[1])
}


func (nsb *NSBApplication) endFuncTransaction(cbInfo ContractCallBackInfo) types.ResponseDeliverTx {
	return types.ResponseDeliverTx{
		Code: cbInfo.CodeResponse,
		Log: cbInfo.Log,
		// Tags:
		Info: cbInfo.Info,
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