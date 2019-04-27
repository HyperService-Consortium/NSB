package nsb

import (
	"bytes"
	"encoding/json"
	"github.com/Myriad-Dreamin/NSB/localstorage"
	"github.com/Myriad-Dreamin/NSB/application/response"
	"github.com/tendermint/tendermint/abci/types"
)


func (nsb *NSBApplication) prepareContractEnvironment(txHeaderJson []byte) (*ContractEnvironment, *types.ResponseDeliverTx) {
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

	var txHeader TransactionHeader
	err = json.Unmarshal(txHeaderJson, &txHeader)
	if err != nil {
		return nil, response.DecodeTxHeaderError(err)
	}

	// TODO: verify signature 

	byteInfo, err = nsb.accMap.TryGet(txHeader.From)
	if err != nil {
		return nil, response.ReTrieveTxError(err)
	}

	var accInfo AccountInfo
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

	var contractInfo AccountInfo
	err = json.Unmarshal(byteInfo, &contractInfo)
	if err != nil {
		return nil, response.DecodeAccountInfoError(err)
	}
	// TODO: Check CodeHash

	var contractEnv = ContractEnvironment{
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

	return nsb.endFuncTransaction(nsb.execContractFuncs(bytesTx[0], env))
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

	return nsb.endConstructTransaction(nsb.createContracts(bytesTx[0], env))
}


func (nsb *NSBApplication) endFuncTransaction(cbInfo *ContractCallBackInfo) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: cbInfo.CodeResponse,
		Log: cbInfo.Log,
		// Tags:
		Info: cbInfo.Info,
	}
}

func (nsb *NSBApplication) endConstructTransaction(cbInfo *ContractCallBackInfo) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: cbInfo.CodeResponse,
		Log: cbInfo.Log,
		// Tags:
		Info: cbInfo.Info,
	}
}
