package nsb

import (
	"bytes"
	"encoding/json"
	"github.com/Myriad-Dreamin/NSB/localstorage"
	cmn "github.com/Myriad-Dreamin/NSB/common"
	"github.com/Myriad-Dreamin/NSB/application/response"
	"github.com/tendermint/tendermint/abci/types"
)


func (nsb *NSBApplication) prepareContractEnvironment(txHeaderJson []byte) (*cmn.ContractEnvironment, *types.ResponseDeliverTx) {
	byteInfo, err := nsb.txMap.TryGet(txHeaderJson)
	// internal error
	if err != nil {
		return nil, response.ReTrieveTxError(err)
	}
	if byteInfo != nil {
		return nil, response.DuplicateTxError
	}
	err = nsb.txMap.TryUpdate(txHeaderJson, []byte{1})
	// internal error
	if err != nil {
		return nil, response.UpdateTxTrieError(err)
	}

	var txHeader cmn.TransactionHeader
	err = json.Unmarshal(txHeaderJson, &txHeader)
	if err != nil {
		return nil, response.DecodeTxHeaderError(err)
	}

	// TODO: verify signature 

	byteInfo, err = nsb.accMap.TryGet(txHeader.From)
	if err != nil {
		return nil, response.ReTrieveTxError(err)
	}

	var accInfo cmn.AccountInfo
	err = json.Unmarshal(byteInfo, &accInfo)
	if err != nil {
		return nil, response.DecodeAccountInfoError(err)
	}

	byteInfo, err = nsb.accMap.TryGet(txHeader.ContractAddress)
	if err != nil {
		return nil, response.ReTrieveTxError(err)
	}
	if byteInfo == nil {
		return nil, response.MissingContract
	}

	var contractInfo cmn.AccountInfo
	err = json.Unmarshal(byteInfo, &contractInfo)
	if err != nil {
		return nil, response.DecodeAccountInfoError(err)
	}
	// TODO: Check CodeHash

	var contractEnv = cmn.ContractEnvironment{
		From: txHeader.From,
		fromInfo: &accInfo,
		ContractAddress: txHeader.ContractAddress,
		toInfo: &contractInfo,
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
		return nil, response.RequestStorageError(err)
	}

	return &contractEnv, nil
}


func (nsb *NSBApplication) parseFuncTransaction(tx []byte) *types.ResponseDeliverTx {
	bytesTx := bytes.Split(tx, []byte("\x18"))
	if len(bytesTx) != 2 {
		return response.InvalidTxInputFormatWrongx18
	}

	env, err := nsb.prepareContractEnvironment(bytesTx[1])
	if err != nil {
		return err
	}

	return nsb.endFuncTransaction(nsb.execContractFuncs(string(bytesTx[0]), env))
}


func (nsb *NSBApplication) parseCreateTransaction(tx []byte) *types.ResponseDeliverTx {
	bytesTx := bytes.Split(tx, []byte("\x18"))
	if len(bytesTx) != 2 {
		return response.InvalidTxInputFormatWrongx18
	}

	env, err := nsb.prepareContractEnvironment(bytesTx[1])
	if err.Code != 0 {
		return err
	}

	return nsb.endConstructTransaction(nsb.createContracts(string(bytesTx[0]), env))
}


func (nsb *NSBApplication) endFuncTransaction(cbInfo *cmn.ContractCallBackInfo) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: cbInfo.CodeResponse,
		Log: cbInfo.Log,
		// Tags:
		Info: cbInfo.Info,
	}
}

func (nsb *NSBApplication) endConstructTransaction(cbInfo *cmn.ContractCallBackInfo) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: cbInfo.CodeResponse,
		Log: cbInfo.Log,
		// Tags:
		Info: cbInfo.Info,
	}
}
