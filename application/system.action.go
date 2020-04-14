package nsb

import (
	"encoding/json"
	"errors"
	TxState "github.com/HyperService-Consortium/go-uip/const/transaction_state_type"

	autil "github.com/HyperService-Consortium/NSB/action"
	"github.com/HyperService-Consortium/NSB/application/response"
	cmn "github.com/HyperService-Consortium/NSB/common"
	"github.com/HyperService-Consortium/NSB/crypto"
	"github.com/HyperService-Consortium/NSB/localstorage"
	"github.com/HyperService-Consortium/NSB/util"
	"github.com/tendermint/tendermint/abci/types"

	transaction "github.com/HyperService-Consortium/NSB/contract/isc/transaction"
)

/*
 * storage := actionMap
 */

type ArgsAddAction struct {
	ISCAddress []byte `json:"1"`
	Tid        uint64 `json:"2"`
	Aid        uint64 `json:"3"`

	// uip.Signature
	Type      uint32 `json:"4"`
	Signature []byte `json:"6"`

	// the content to be signed
	Content []byte `json:"5"`
}

type ArgsAddActions struct {
	Args []ArgsAddAction `json:"1"`
}

func (nsb *NSBApplication) ActionRigisteredMethod(
	env *cmn.TransactionHeader,
	frInfo *cmn.AccountInfo,
	toInfo *cmn.AccountInfo,
	funcName string,
	args []byte,
) *types.ResponseDeliverTx {
	switch funcName {
	case "addAction":
		var argsAddAction ArgsAddAction
		util.MustUnmarshal(args, &argsAddAction)
		return nsb.addAction(&argsAddAction)
	case "getAction":
		return nsb.getAction(args)
	case "addActions":
		var argsAddActions ArgsAddActions
		util.MustUnmarshal(args, &argsAddActions)
		return nsb.addActions(&argsAddActions)
	default:
		return response.InvalidFuncTypeError(response.MethodMissing)
	}
}

func actionKey(addr []byte, tid uint64, aid uint64) []byte {
	return crypto.Sha512(addr, util.Uint64ToBytes(tid), util.Uint64ToBytes(aid))
}

func (nsb *NSBApplication) _addAction(args *ArgsAddAction) *types.ResponseDeliverTx {
	// TODO: check valid isc/tid/aid
	if args.Aid > TxState.Closed {
		return response.ExecContractError(errors.New("action index overflow"))
	}

	action, err := autil.NewAction(args.Type, args.Signature, args.Content)
	if err != nil {
		response.ExecContractError(err)
		return response.ExecContractError(err)
	}

	conInfo, errInfo := nsb.extractAddress(args.ISCAddress)
	if errInfo != nil {
		return errInfo
	}
	storage, err := localstorage.NewLocalStorage(
		args.ISCAddress,
		conInfo.StorageRoot,
		nsb.statedb,
	)
	//todo: encapsulate storage operation

	txArr := storage.NewBytesArray("transactions")
	if txArr.Length() <= args.Tid {
		return response.ExecContractError(errors.New("transaction index overflow"))
	}

	txb := txArr.Get(args.Tid)
	var tx transaction.TransactionIntent
	err = json.Unmarshal(txb, &tx)
	if err != nil {
		return response.ExecContractError(err)
	}

	// if is src
	if ((TxState.Instantiating ^ args.Aid) & 1) == 0 {
		if action.Verify(tx.Fr) == false {
			return response.ExecContractError(errors.New("validate failed"))
		}
	} else {
		if action.Verify(tx.To) == false {
			return response.ExecContractError(errors.New("validate failed"))
		}
	}

	err = nsb.actionMap.TryUpdate(
		actionKey(args.ISCAddress, args.Tid, args.Aid),
		action.Concat(),
	)
	if err != nil {
		return response.ExecContractError(err)
	}
	return nil
}

func (nsb *NSBApplication) addAction(args *ArgsAddAction) *types.ResponseDeliverTx {

	if err := nsb._addAction(args); err != nil {
		return err
	}

	return &types.ResponseDeliverTx{
		Code: uint32(response.CodeOK()),
		Info: "updateSuccess",
	}
}

func (nsb *NSBApplication) addActions(args *ArgsAddActions) *types.ResponseDeliverTx {
	for _, batchArgs := range args.Args {
		if err := nsb._addAction(&batchArgs); err != nil {
			return err
		}
	}
	return &types.ResponseDeliverTx{
		Code: uint32(response.CodeOK()),
		Info: "updateSuccess",
	}
}

type ArgsGetAction struct {
	ISCAddress []byte `json:"1"`
	// hexbytes
	Tid uint64 `json:"2"`
	// hexbytes
	Aid uint64 `json:"3"`
}

func (nsb *NSBApplication) getAction(bytesArgs []byte) *types.ResponseDeliverTx {
	var args ArgsGetAction
	util.MustUnmarshal(bytesArgs, &args)
	// TODO: check valid isc/tid/aid
	bt, err := nsb.actionMap.TryGet(actionKey(args.ISCAddress, args.Tid, args.Aid))
	if err != nil {
		return response.ExecContractError(err)
	}
	return &types.ResponseDeliverTx{
		Code: uint32(response.CodeOK()),
		Data: bt,
	}
}
