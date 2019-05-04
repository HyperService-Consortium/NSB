package nsb

import (
	"bytes"
	"errors"
	"encoding/json"
	"github.com/Myriad-Dreamin/NSB/math"
	"github.com/Myriad-Dreamin/NSB/account"
	"github.com/Myriad-Dreamin/NSB/localstorage"
	cmn "github.com/Myriad-Dreamin/NSB/common"
	"github.com/Myriad-Dreamin/NSB/application/response"
	"github.com/tendermint/tendermint/abci/types"
)

var (
	ContractNameNotEqual = errors.New("the name of contract to call is mismatch with providing name")
)

func (nsb *NSBApplication) parseTxHeader(txHeaderJson []byte) (
	*cmn.TransactionHeader,
	*types.ResponseDeliverTx,
) {
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
	return &txHeader, nil
}


func (nsb *NSBApplication) parseAccInfo(addr []byte) (
	*AccountInfo,
	*types.ResponseDeliverTx,
) {
	byteInfo, err := nsb.accMap.TryGet(addr)
	if err != nil {
		return nil, response.ReTrieveTxError(err)
	}

	var accInfo AccountInfo
	if byteInfo != nil {
		err = json.Unmarshal(byteInfo, &accInfo)
		if err != nil {
			return nil, response.DecodeAccountInfoError(err)
		}
	} else {
		accInfo.Balance = math.NewUint256FromBytes([]byte{0})
	}

	return &accInfo, nil
}


func (nsb *NSBApplication) parseContractInfo(
	txHeader *cmn.TransactionHeader,
	contractName []byte,
	createFlag bool,
) (
	*AccountInfo,
	*types.ResponseDeliverTx,
) {
	var contractInfo AccountInfo
	if createFlag {
		txHeader.ContractAddress = []byte(account.NewAccount([]byte{}).PublicKey)
		var byteInfo = make([]byte, 0)
		contractInfo.Balance = math.NewUint256FromBytes([]byte{0})
		contractInfo.Name = contractName
		// TODO: set CodeHash
	} else {
		byteInfo, err := nsb.accMap.TryGet(txHeader.ContractAddress)
		if err != nil {
			return nil, response.ReTrieveTxError(err)
		}
		if byteInfo == nil {
			return nil, response.MissingContract
		} else {
			err = json.Unmarshal(byteInfo, &contractInfo)
			if err != nil {
				return nil, response.DecodeAccountInfoError(err)
			}
		}
	}
	return &contractInfo, nil
}

func (nsb *NSBApplication) parseFAPair(bytesPair []byte, createFlag bool) (
	*FAPair,
	*types.ResponseDeliverTx,
) {
	var fap FAPair
	if createFlag {
		fap.Args = bytesPair
	} else {
		err := json.Unmarshal(bytesPair, &fap)
		if err != nil {
			return nil, response.DecodeAccountInfoError(err)
		}
	}

	return &fap, nil
}
func (nsb *NSBApplication) prepareContractEnvironment(bytesTx [][]byte, createFlag bool) (
	*cmn.ContractEnvironment,
	*AccountInfo,
	*AccountInfo,
	*types.ResponseDeliverTx,
) {
	txHeader, errInfo := nsb.parseTxHeader(bytesTx[1])
	if errInfo != nil {
		return nil, nil, nil, errInfo
	}
	
	var fap *FAPair
	fap, errInfo = nsb.parseFAPair(txHeader.Data, createFlag)
	if errInfo != nil {
		return nil, nil, nil, errInfo
	}

	var accInfo, conInfo *AccountInfo
	accInfo, errInfo = nsb.parseAccInfo(txHeader.From)
	if errInfo != nil {
		return nil, nil, nil, errInfo
	}

	contractName := bytesTx[0]

	conInfo, errInfo = nsb.parseContractInfo(txHeader, contractName, createFlag)
	if errInfo != nil {
		return nil, nil, nil, errInfo
	}

	if !bytes.Equal(contractName, contractInfo.Name) {
		return nil, nil, nil, response.ReTrieveTxError(ContractNameNotEqual)
	}

	// TODO: verify signature 

	// TODO: Check CodeHash

	var contractEnv = cmn.ContractEnvironment{
		From: txHeader.From,
		ContractAddress: txHeader.ContractAddress,
		FuncName: fap.FuncName,
		Args: fap.Args,
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


func (nsb *NSBApplication) prepareSystemContractEnvironment(txHeaderJson []byte) (
	*cmn.TransactionHeader,
	*AccountInfo,
	*types.ResponseDeliverTx,
) {
	txHeader, errInfo := nsb.parseTxHeader(bytesTx)
	if errInfo != nil {
		return nil, nil, errInfo
	}
	
	var accInfo *AccountInfo
	accInfo, errInfo = nsb.parseAccInfo(txHeader.From)
	if errInfo != nil {
		return nil, nil, errInfo
	}

	return &txHeader, &accInfo, nil
}


func (nsb *NSBApplication) storeState(
	env *cmn.ContractEnvironment,
	accInfo *AccountInfo,
	conInfo *AccountInfo,
) *types.ResponseDeliverTx {
	var err error
	conInfo.StorageRoot, err = env.Storage.Commit()
	if err != nil {
		return response.CommitAccTrieError(err)
	}

	var bt []byte
	bt, err = json.Marshal(accInfo)
	if err != nil {
		return response.EncodeAccountInfoError(err)
	}

	err = nsb.accMap.TryUpdate(env.From, bt)
	if err != nil {
		return response.UpdateAccTrieError(err)
	}

	bt, err = json.Marshal(conInfo)
	if err != nil {
		return response.EncodeAccountInfoError(err)
	}

	err = nsb.accMap.TryUpdate(env.ContractAddress, bt)
	if err != nil {
		return response.UpdateAccTrieError(err)
	}

	return nil
}


func (nsb *NSBApplication) parseFuncTransaction(tx []byte) *types.ResponseDeliverTx {
	bytesTx := bytes.Split(tx, []byte("\x18"))
	if len(bytesTx) != 2 {
		return response.InvalidTxInputFormatWrongx18
	}

	env, accInfo, conInfo, errInfo := nsb.prepareContractEnvironment(bytesTx, false)
	if errInfo != nil {
		return errInfo
	}

	cb := nsb.execContractFuncs(string(bytesTx[0]), env)

	if cb.CodeResponse == uint32(response.CodeOK()) {
		// TODO: modify accInfo
		errInfo = nsb.storeState(env, accInfo, conInfo)
		if errInfo != nil {
			return errInfo
		}
	}

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

	env, accInfo, conInfo, errInfo := nsb.prepareContractEnvironment(bytesTx, true)
	if errInfo != nil {
		return errInfo
	}

	cb := nsb.createContracts(string(bytesTx[0]), env)

	if cb.CodeResponse == uint32(response.CodeOK()) {
		// TODO: modify accInfo
		errInfo = nsb.storeState(env, accInfo, conInfo)
		if errInfo != nil {
			return errInfo
		}
	}

	return &types.ResponseDeliverTx{
		Code: cb.CodeResponse,
		Log: cb.Log,
		// Tags:
		Info: cb.Info,
	}
}

func (nsb *NSBApplication) parseSystemFuncTransaction(tx []byte) *types.ResponseDeliverTx {
	bytesTx := bytes.Split(tx, []byte("\x18"))
	if len(bytesTx) != 2 {
		return response.InvalidTxInputFormatWrongx18
	}

	env, accInfo, errInfo := nsb.prepareSystemContractEnvironment(bytesTx)
	if errInfo != nil {
		return errInfo
	}
	fap := nsb.parseFAPair(env.Data)

	cb := nsb.systemCall(string(bytesTx[0]), env, accInfo, fap.FuncName, fap.Args)

	if cb.Code == uint32(response.CodeOK()) {
		bt, err := json.Marshal(accInfo)
		if err != nil {
			return response.EncodeAccountInfoError(err)
		}

		err = nsb.accMap.TryUpdate(env.From, bt)
		if err != nil {
			return response.UpdateAccTrieError(err)
		}
	}

	return cb
}
