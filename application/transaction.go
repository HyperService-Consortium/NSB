package nsb

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"

	"github.com/HyperService-Consortium/NSB/account"
	"github.com/HyperService-Consortium/NSB/application/response"
	cmn "github.com/HyperService-Consortium/NSB/common"
	"github.com/HyperService-Consortium/NSB/crypto"
	"github.com/HyperService-Consortium/NSB/localstorage"
	"github.com/HyperService-Consortium/NSB/math"
	"github.com/tendermint/tendermint/abci/types"
	ten_cmn "github.com/tendermint/tendermint/libs/common"

	nsbrpc "github.com/HyperService-Consortium/NSB/grpc/nsbrpc"
	"github.com/gogo/protobuf/proto"
)

var (
	ContractNameNotEqual = errors.New("the name of contract to call is mismatch with providing name")
)

func (nsb *NSBApplication) parseTxHeader(txHeaderProtobuf []byte) (
	*cmn.TransactionHeader,
	*types.ResponseDeliverTx,
) {

	var txHeaderRaw nsbrpc.TransactionHeader
	err := proto.Unmarshal(txHeaderProtobuf, &txHeaderRaw)
	if err != nil {
		return nil, response.DecodeTxHeaderError(err)
	}
	if len(txHeaderRaw.Src) != 32 {
		return nil, response.DecodeTxHeaderError(errorDecodeSrcAddress)
	}
	if len(txHeaderRaw.Dst) != 32 && len(txHeaderRaw.Dst) != 0 {
		return nil, response.DecodeTxHeaderError(errorDecodeDstAddress)
	}

	var txHeader cmn.TransactionHeader
	txHeader.Value = math.NewUint256FromBytes(txHeaderRaw.Value)
	if txHeader.Value == nil {
		return nil, response.DecodeTxHeaderError(errorDecodeUint256)
	}
	txHeader.Nonce = math.NewUint256FromBytes(txHeaderRaw.Nonce)
	if txHeader.Nonce == nil {
		return nil, response.DecodeTxHeaderError(errorDecodeUint256)
	}

	byteInfo, err := nsb.txMap.TryGet(txHeaderProtobuf)
	// internal error
	if err != nil {
		return nil, response.ReTrieveTxError(err)
	}
	if byteInfo != nil {
		return nil, response.DuplicateTxError
	}
	err = nsb.txMap.TryUpdate(txHeaderProtobuf, []byte{1})
	// internal error
	if err != nil {
		return nil, response.UpdateTxTrieError(err)
	}

	txHeader.From = txHeaderRaw.Src
	txHeader.ContractAddress = txHeaderRaw.Dst
	txHeader.Data = txHeaderRaw.Data
	txHeader.Signature = txHeaderRaw.Signature

	return &txHeader, nil
}

func (nsb *NSBApplication) parseAccInfo(addr []byte) (
	*cmn.AccountInfo,
	*types.ResponseDeliverTx,
) {
	byteInfo, err := nsb.accMap.TryGet(addr)
	if err != nil {
		return nil, response.ReTrieveTxError(err)
	}

	var accInfo cmn.AccountInfo
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

func (nsb *NSBApplication) extractAddress(contractAddress []byte) (
	*cmn.AccountInfo,
	*types.ResponseDeliverTx,
) {
	byteInfo, err := nsb.accMap.TryGet(contractAddress)
	if err != nil {
		return nil, response.ReTrieveTxError(err)
	}
	if byteInfo == nil {
		return nil, response.MissingContract
	} else {
		var contractInfo cmn.AccountInfo
		err = json.Unmarshal(byteInfo, &contractInfo)
		if err != nil {
			return nil, response.DecodeAccountInfoError(err)
		}
		return &contractInfo, nil
	}
}

func (nsb *NSBApplication) createContractAccount(
	txHeader *cmn.TransactionHeader,
	contractName string,
) (
	*cmn.AccountInfo,
	*types.ResponseDeliverTx,
) {
	nsb.logger.Info("creating", "contractName", contractName)
	txHeader.ContractAddress = []byte(account.NewAccount(txHeader.Signature, txHeader.Nonce.Bytes(), nsb.state.StateRoot).PublicKey)

	// TODO: merk: TryGet -> TestExistence
	byteInfo, err := nsb.accMap.TryGet(txHeader.ContractAddress)
	if err != nil {
		return nil, response.ExecContractError(err)
	}
	if byteInfo != nil {
		return nil, response.ConflictAddress
	}
	var contractInfo cmn.AccountInfo
	contractInfo.Balance = math.NewUint256FromBytes([]byte{0})
	contractInfo.Name = []byte(contractName)
	// TODO: set CodeHash
	return &contractInfo, nil
}

func (nsb *NSBApplication) parseFAPair(bytesPair []byte) (
	*FAPair,
	*types.ResponseDeliverTx,
) {
	var fap FAPair
	err := proto.Unmarshal(bytesPair, &fap)
	if err != nil {
		return nil, response.DecodeFAPairError(err)
	}
	return &fap, nil
}
func (nsb *NSBApplication) prepareContractEnvironment(
	bytesTx []byte, createFlag bool,
) (
	*cmn.ContractEnvironment,
	*cmn.AccountInfo,
	*cmn.AccountInfo,
	*types.ResponseDeliverTx,
) {
	txHeader, errInfo := nsb.parseTxHeader(bytesTx)
	if errInfo != nil {
		return nil, nil, nil, errInfo
	}

	var fap *FAPair
	fap, errInfo = nsb.parseFAPair(txHeader.Data)
	if errInfo != nil {
		return nil, nil, nil, errInfo
	}

	var accInfo, conInfo *cmn.AccountInfo
	accInfo, errInfo = nsb.parseAccInfo(txHeader.From)
	if errInfo != nil {
		return nil, nil, nil, errInfo
	}

	if createFlag {
		conInfo, errInfo = nsb.createContractAccount(txHeader, fap.FuncName)
	} else {
		conInfo, errInfo = nsb.extractAddress(txHeader.ContractAddress)
	}

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
		BN:              nsb.system.merkleProof,
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

func (nsb *NSBApplication) prepareSystemContractEnvironment(txHeaderProtobuf []byte) (
	*cmn.TransactionHeader,
	*cmn.AccountInfo,
	*cmn.AccountInfo,
	*types.ResponseDeliverTx,
) {
	txHeader, errInfo := nsb.parseTxHeader(txHeaderProtobuf)
	if errInfo != nil {
		return nil, nil, nil, errInfo
	}

	var frInfo, toInfo *cmn.AccountInfo
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
	accInfo *cmn.AccountInfo,
	conInfo *cmn.AccountInfo,
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
	accInfo *cmn.AccountInfo,
	conInfo *cmn.AccountInfo,
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

	env, accInfo, conInfo, errInfo := nsb.prepareContractEnvironment(tx, false)
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
				Value: crypto.Keccak256([]byte("tx:"), tx),
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
				Value: crypto.Keccak256([]byte("tx:"), tx),
			})
		}

	}

	return &types.ResponseDeliverTx{
		Code: cb.CodeResponse,
		Log:  cb.Log,
		// Tags:
		Info: cb.Info,
		Data: cb.Data,
		Events: []types.Event{
			types.Event{
				Type:       "sendTransaction",
				Attributes: Tags,
			},
		},
	}
}

func (nsb *NSBApplication) parseCreateTransaction(tx []byte) *types.ResponseDeliverTx {

	env, accInfo, conInfo, errInfo := nsb.prepareContractEnvironment(tx, true)
	if errInfo != nil {
		return errInfo
	}

	nsb.logger.Info("parsed", "src", accInfo.Name, "src_balance", accInfo.Balance,
		"dst", conInfo.Name, "dst_balance", conInfo.Balance,
	)

	cb := nsb.createContracts(env)

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
				Value: crypto.Keccak256([]byte("tx:"), tx),
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
				Value: crypto.Keccak256([]byte("tx:"), tx),
			})
		}
	}

	return &types.ResponseDeliverTx{
		Code: cb.CodeResponse,
		Log:  cb.Log,
		// Tags:
		Info: cb.Info,
		Data: cb.Data,
		Events: []types.Event{
			types.Event{
				Type:       "createContract",
				Attributes: Tags,
			},
		},
	}
}

func (nsb *NSBApplication) parseSystemFuncTransaction(tx []byte) *types.ResponseDeliverTx {

	env, frInfo, toInfo, errInfo := nsb.prepareSystemContractEnvironment(tx)
	if errInfo != nil {
		return errInfo
	}

	var fap *FAPair
	fap, errInfo = nsb.parseFAPair(env.Data)
	if errInfo != nil {
		return errInfo
	}
	env.Data = nil
	names := strings.Split(fap.FuncName, "@")
	if len(names) != 2 {
		return response.InvalidTxInputFormatWrongFunctionName
	}

	cb := nsb.systemCall(names[0], env, frInfo, toInfo, names[1], fap.Args)

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

		if cb.Events == nil {
			cb.Events = []types.Event{
				types.Event{
					Type: "systemCall",
					Attributes: []ten_cmn.KVPair{ten_cmn.KVPair{
						Key:   []byte("TransactionHash"),
						Value: crypto.Keccak256([]byte("tx:"), tx),
					}},
				},
			}
		} else {
			cb.Events = append(cb.Events, types.Event{
				Type: "systemCall",
				Attributes: []ten_cmn.KVPair{ten_cmn.KVPair{
					Key:   []byte("TransactionHash"),
					Value: crypto.Keccak256([]byte("tx:"), tx),
				}},
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
