package nsb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/HyperServiceOne/NSB/account"
	"github.com/HyperServiceOne/NSB/application/response"
	cmn "github.com/HyperServiceOne/NSB/common"
	"github.com/HyperServiceOne/NSB/crypto"
	"github.com/HyperServiceOne/NSB/localstorage"
	"github.com/HyperServiceOne/NSB/math"
	"github.com/tendermint/tendermint/abci/types"
	ten_cmn "github.com/tendermint/tendermint/libs/common"
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
		fmt.Println("creating", contractName)
		txHeader.ContractAddress = []byte(account.NewAccount(txHeader.Signature, txHeader.Nonce.Bytes(), nsb.state.StateRoot).PublicKey)

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

	// fmt.Println("cmp", conInfo.Name, contractName, "cmp")

	// if !bytes.Equal(contractName, conInfo.Name) {
	// 	return nil, nil, nil, response.ReTrieveTxError(ContractNameNotEqual)
	// }

	// TODO: verify signature

	// TODO: Check CodeHash
	var err error
	var contractEnv = cmn.ContractEnvironment{
		From:            txHeader.From,
		ContractAddress: txHeader.ContractAddress,
		FuncName:        fap.FuncName,
		Args:            fap.Args,
		Value:           txHeader.Value,
	}
	contractEnv.Storage, err = localstorage.NewLocalStorage(
		txHeader.ContractAddress,
		conInfo.StorageRoot,
		nsb.statedb,
	)

	// internal error
	if err != nil {
		return nil, nil, nil, response.RequestStorageError(err)
	}

	return &contractEnv, accInfo, conInfo, nil
}

func (nsb *NSBApplication) prepareSystemContractEnvironment(txHeaderJson []byte) (
	*cmn.TransactionHeader,
	*AccountInfo,
	*AccountInfo,
	*types.ResponseDeliverTx,
) {
	txHeader, errInfo := nsb.parseTxHeader(txHeaderJson)
	if errInfo != nil {
		return nil, nil, nil, errInfo
	}

	var frInfo, toInfo *AccountInfo
	frInfo, errInfo = nsb.parseAccInfo(txHeader.From)
	if errInfo != nil {
		return nil, nil, nil, errInfo
	}
	if txHeader.ContractAddress != nil {
		if bytes.Equal(txHeader.ContractAddress, txHeader.From) {
			toInfo = frInfo
		} else {
			toInfo, errInfo = nsb.parseAccInfo(txHeader.ContractAddress)
			if errInfo != nil {
				return nil, nil, nil, errInfo
			}
		}

	}

	return txHeader, frInfo, toInfo, nil
}

func (nsb *NSBApplication) modifyState(
	cb *cmn.ContractCallBackInfo,
	env *cmn.ContractEnvironment,
	accInfo *AccountInfo,
	conInfo *AccountInfo,
) *types.ResponseDeliverTx {
	if cb.Value == nil {
		return nil
	}

	if cb.Value.BitLen() != 0 {
		if cb.OutFlag {
			checkErr := conInfo.Balance.Sub(cb.Value)
			if checkErr {
				return response.InsufficientBalanceToTransfer("contract")
			}
			checkErr = accInfo.Balance.Add(cb.Value)
			if checkErr {
				return response.BalanceOverflow("user")
			}
		} else {
			checkErr := accInfo.Balance.Sub(cb.Value)
			if checkErr {
				return response.InsufficientBalanceToTransfer("user")
			}
			checkErr = conInfo.Balance.Add(cb.Value)
			if checkErr {
				return response.BalanceOverflow("contract")
			}
		}
	}
	return nil
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

	cb := nsb.execContractFuncs(string(conInfo.Name), env)

	var Tags []ten_cmn.KVPair
	if cb.CodeResponse == uint32(response.CodeOK()) {
		// TODO: modify accInfo
		errInfo = nsb.modifyState(cb, env, accInfo, conInfo)
		if errInfo != nil {
			return errInfo
		}
		errInfo = nsb.storeState(env, accInfo, conInfo)
		if errInfo != nil {
			return errInfo
		}
		if cb.Tags == nil {
			Tags = []ten_cmn.KVPair{ten_cmn.KVPair{
				Key:   []byte("TransactionHash"),
				Value: crypto.Keccak256([]byte("tx:"), bytesTx[1]),
			}}
		} else {
			Tags = make([]ten_cmn.KVPair, 0, len(cb.Tags)+1)
			for _, tag := range cb.Tags {
				Tags = append(Tags, ten_cmn.KVPair{
					Key:   tag.Key(),
					Value: tag.Value(),
				})
			}
			Tags = append(Tags, ten_cmn.KVPair{
				Key:   []byte("TransactionHash"),
				Value: crypto.Keccak256([]byte("tx:"), bytesTx[1]),
			})
		}

	}

	return &types.ResponseDeliverTx{
		Code: cb.CodeResponse,
		Log:  cb.Log,
		// Tags:
		Info: cb.Info,
		Data: cb.Data,
		Tags: Tags,
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

	fmt.Println(accInfo, conInfo)

	cb := nsb.createContracts(string(bytesTx[0]), env)

	var Tags []ten_cmn.KVPair
	if cb.CodeResponse == uint32(response.CodeOK()) {
		// TODO: modify accInfo
		errInfo = nsb.modifyState(cb, env, accInfo, conInfo)
		if errInfo != nil {
			return errInfo
		}
		errInfo = nsb.storeState(env, accInfo, conInfo)
		if errInfo != nil {
			return errInfo
		}

		if cb.Tags == nil {
			Tags = []ten_cmn.KVPair{ten_cmn.KVPair{
				Key:   []byte("TransactionHash"),
				Value: crypto.Keccak256([]byte("tx:"), bytesTx[1]),
			}}
		} else {
			Tags = make([]ten_cmn.KVPair, 0, len(cb.Tags)+1)
			for _, tag := range cb.Tags {
				Tags = append(Tags, ten_cmn.KVPair{
					Key:   tag.Key(),
					Value: tag.Value(),
				})
			}
			Tags = append(Tags, ten_cmn.KVPair{
				Key:   []byte("TransactionHash"),
				Value: crypto.Keccak256([]byte("tx:"), bytesTx[1]),
			})
		}
	}

	return &types.ResponseDeliverTx{
		Code: cb.CodeResponse,
		Log:  cb.Log,
		// Tags:
		Info: cb.Info,
		Data: cb.Data,
		Tags: Tags,
	}
}

func (nsb *NSBApplication) parseSystemFuncTransaction(tx []byte) *types.ResponseDeliverTx {
	bytesTx := bytes.Split(tx, []byte("\x18"))
	if len(bytesTx) != 2 {
		return response.InvalidTxInputFormatWrongx18
	}

	env, frInfo, toInfo, errInfo := nsb.prepareSystemContractEnvironment(bytesTx[1])
	if errInfo != nil {
		return errInfo
	}

	var fap *FAPair
	fap, errInfo = nsb.parseFAPair(env.Data, false)
	if errInfo != nil {
		return errInfo
	}

	cb := nsb.systemCall(string(bytesTx[0]), env, frInfo, toInfo, fap.FuncName, fap.Args)

	if cb.Code == uint32(response.CodeOK()) {
		bt, err := json.Marshal(frInfo)
		if err != nil {
			return response.EncodeAccountInfoError(err)
		}

		err = nsb.accMap.TryUpdate(env.From, bt)
		if err != nil {
			return response.UpdateAccTrieError(err)
		}

		if toInfo != nil {
			bt, err = json.Marshal(toInfo)
			if err != nil {
				return response.EncodeAccountInfoError(err)
			}

			err = nsb.accMap.TryUpdate(env.ContractAddress, bt)
			if err != nil {
				return response.UpdateAccTrieError(err)
			}
		}

		if cb.Tags == nil {
			cb.Tags = []ten_cmn.KVPair{ten_cmn.KVPair{
				Key:   []byte("TransactionHash"),
				Value: crypto.Keccak256([]byte("tx:"), bytesTx[1]),
			}}
		} else {
			cb.Tags = append(cb.Tags, ten_cmn.KVPair{
				Key:   []byte("TransactionHash"),
				Value: crypto.Keccak256([]byte("tx:"), bytesTx[1]),
			})
		}

	}
	// else if cb.Code == uint32(response.CodeUndateBalanceIn()) {
	// 	value := math.NewUint256FromBytes(cb.Data)

	// 	if value == nil {
	// 		return response.DecodeBalanceError()
	// 	}

	// 	checkErr := accInfo.Balance.Sub(value)
	// 	if checkErr {
	// 		return response.InsufficientBalanceToTransfer("user")
	// 	}

	// 	bt, err := json.Marshal(accInfo)
	// 	if err != nil {
	// 		return response.EncodeAccountInfoError(err)
	// 	}

	// 	err = nsb.accMap.TryUpdate(env.From, bt)
	// 	if err != nil {
	// 		return response.UpdateAccTrieError(err)
	// 	}

	// 	cb.Code = uint32(response.CodeOK())
	// } else if cb.Code == uint32(response.CodeUndateBalanceOut()) {
	// 	value := math.NewUint256FromBytes(cb.Data)

	// 	if value == nil {
	// 		return response.DecodeBalanceError()
	// 	}

	// 	checkErr := accInfo.Balance.Add(value)
	// 	if checkErr {
	// 		return response.BalanceOverflow("user")
	// 	}

	// 	bt, err := json.Marshal(accInfo)
	// 	if err != nil {
	// 		return response.EncodeAccountInfoError(err)
	// 	}

	// 	err = nsb.accMap.TryUpdate(env.From, bt)
	// 	if err != nil {
	// 		return response.UpdateAccTrieError(err)
	// 	}

	// 	cb.Code = uint32(response.CodeOK())
	// }

	return cb
}
