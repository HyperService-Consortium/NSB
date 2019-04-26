package nsb

import (
	"fmt"
	"encoding/json"
	"github.com/Myriad-Dreamin/NSB/crypto"
	"github.com/Myriad-Dreamin/NSB/account"
	"github.com/Myriad-Dreamin/NSB/localstorage"
)

type TransactionHeader struct {
	From []byte  `json:"from"`
	ContractAddress []byte  `json:"to"`
	JsonParas []byte `json:"data"`
	Value []byte `json:"value"`
	Nonce []byte `json:"nonce"`
	Signature []byte `json:"signature"`
}


type ContractEnvironment struct {
	Storage *localstorage.LocalStorage
	From []byte
	Value []byte
}


func (nsb *NSBApplication) parseFuncTransaction(tx []byte) types.ResponseDeliverTx {
	bytesTx := bytes.Split(tx, []byte("\x18"))
	if len(bytesTx) != 2 {
		return invalidTxInputFormatWrongx18
	}
	
	var txHeader TransactionHeader
	err := json.Unmarshal(bytesTx[1], &txHeader)
	if err != nil {
		return DecodeTxHeaderError(err)
	}



	return nsb.foundContracts(bytesTx[0], txHeader.JsonParas)
}


func (nsb *NSBApplication) parseCreateTransaction(tx []byte) types.ResponseDeliverTx {
	bytesTx := bytes.Split(tx, []byte("\x18"))
	if len(bytesTx) != 2 {
		return invalidTxInputFormatWrongx18
	}
	return nsb.createContracts(bytesTx[0], bytesTx[1])
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