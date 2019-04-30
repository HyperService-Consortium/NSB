package nsb

import (
	"fmt"
	"bytes"
	"encoding/json"
	"github.com/Myriad-Dreamin/NSB/math"
	"github.com/Myriad-Dreamin/NSB/account"
	"github.com/Myriad-Dreamin/NSB/localstorage"
	cmn "github.com/Myriad-Dreamin/NSB/common"
	"github.com/Myriad-Dreamin/NSB/application/response"
	"github.com/tendermint/tendermint/abci/types"
)


func (nsb *NSBApplication) prepareContractEnvironment(txHeaderJson []byte, createFlag bool) (
	*cmn.ContractEnvironment,
	*AccountInfo,
	*AccountInfo,
	*types.ResponseDeliverTx,
) {
	// fmt.Println("prepare to create", txHeaderJson, createFlag)

	byteInfo, err := nsb.txMap.TryGet(txHeaderJson)
	// internal error
	if err != nil {
		return nil, nil, nil, response.ReTrieveTxError(err)
	}
	if byteInfo != nil {
		return nil, nil, nil, response.DuplicateTxError
	}
	err = nsb.txMap.TryUpdate(txHeaderJson, []byte{1})
	// internal error
	if err != nil {
		return nil, nil, nil, response.UpdateTxTrieError(err)
	}

	// fmt.Println("Check TxTrie OK")

	var txHeader cmn.TransactionHeader
	err = json.Unmarshal(txHeaderJson, &txHeader)
	if err != nil {
		return nil, nil, nil, response.DecodeTxHeaderError(err)
	}

	// fmt.Println("decoded TxHeader", txHeader)

	// TODO: verify signature 

	byteInfo, err = nsb.accMap.TryGet(txHeader.From)
	if err != nil {
		return nil, nil, nil, response.ReTrieveTxError(err)
	}

	var accInfo AccountInfo
	if byteInfo != nil {
		err = json.Unmarshal(byteInfo, &accInfo)
		if err != nil {
			return nil, nil, nil, response.DecodeAccountInfoError(err)
		}
	} else {
		accInfo.Balance = math.NewUint256FromBytes([]byte{0})
	}
	
	// fmt.Println("decoded accInfo", accInfo)

	var contractInfo AccountInfo
	if createFlag {
		txHeader.ContractAddress = []byte(account.NewAccount([]byte{}).PublicKey)
		byteInfo = make([]byte, 0)
		contractInfo.Balance = math.NewUint256FromBytes([]byte{0})
		// TODO: set CodeHash
	} else {
		byteInfo, err = nsb.accMap.TryGet(txHeader.ContractAddress)
		if err != nil {
			return nil, nil, nil, response.ReTrieveTxError(err)
		}
		if byteInfo == nil {
			return nil, nil, nil, response.MissingContract
		} else {
			err = json.Unmarshal(byteInfo, &contractInfo)
			if err != nil {
				return nil, nil, nil, response.DecodeAccountInfoError(err)
			}
		}
	}

	// TODO: Check CodeHash
	
	// fmt.Println("decoded contractInfo", contractInfo)

	var contractEnv = cmn.ContractEnvironment{
		From: txHeader.From,
		ContractAddress: txHeader.ContractAddress,
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
		return nil, nil, nil, response.RequestStorageError(err)
	}

	return &contractEnv, &accInfo, &contractInfo, nil
}


func (nsb *NSBApplication) parseFuncTransaction(tx []byte) *types.ResponseDeliverTx {
	bytesTx := bytes.Split(tx, []byte("\x18"))
	if len(bytesTx) != 2 {
		return response.InvalidTxInputFormatWrongx18
	}

	env, accInfo, conInfo, err := nsb.prepareContractEnvironment(bytesTx[1], false)
	if err != nil {
		return err
	}

	cb := nsb.execContractFuncs(string(bytesTx[0]), env)
	fmt.Println(accInfo, conInfo)

	return &types.ResponseDeliverTx{
		Code: cb.CodeResponse,
		Log: cb.Log,
		// Tags:
		Info: cb.Info,
	}
}


func (nsb *NSBApplication) parseCreateTransaction(tx []byte) *types.ResponseDeliverTx {
	bytesTx := bytes.Split(tx, []byte("\x18"))
	if len(bytesTx) != 2 {
		return response.InvalidTxInputFormatWrongx18
	}

	env, accInfo, conInfo, err := nsb.prepareContractEnvironment(bytesTx[1], true)
	if err != nil {
		return err
	}

	cb := nsb.createContracts(string(bytesTx[0]), env)
	fmt.Println(accInfo, conInfo)

	return &types.ResponseDeliverTx{
		Code: cb.CodeResponse,
		Log: cb.Log,
		// Tags:
		Info: cb.Info,
	}
}

func (nsb *NSBApplication) transact(tx []byte) *types.ResponseDeliverTx {
	return &types.ResponseDeliverTx{
		Code: uint32(response.CodeTODO),
	}
}
