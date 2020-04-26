package isc

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/HyperService-Consortium/go-uip/isc"
	"github.com/HyperService-Consortium/go-uip/op-intent/instruction"
	"github.com/HyperService-Consortium/go-uip/uip"

	cmn "github.com/HyperService-Consortium/NSB/common"
	. "github.com/HyperService-Consortium/NSB/common/contract_response"
	"github.com/HyperService-Consortium/NSB/contract/isc/transaction"
	"github.com/HyperService-Consortium/NSB/util"
)

func MustUnmarshal(data []byte, load interface{}) {
	err := json.Unmarshal(data, load)
	if err != nil {
		panic(DecodeJsonError(err))
	}
}

type ArgsCreateNewContract struct {
	IscOwners          [][]byte `json:"isc_owners"`
	Funds              []uint64 `json:"required_funds"`
	VesSig             []byte   `json:"ves_signature"`
	TransactionIntents [][]byte `json:"transaction_intents"`
}

type ArgsUpdateTxInfo struct {
	Tid               uint64                         `json:"tid"`
	TransactionIntent *transaction.TransactionIntent `json:"transaction_intent"`
}

type ArgsUpdateTxFr struct {
	Tid uint64 `json:"tid"`
	Fr  []byte `json:"from"`
}

type ArgsFreezeInfo struct {
	Tid uint64 `json:"tid"`
}

type ArgsUserAck struct {
	Address   []byte `json:"address"`
	Signature []byte `json:"signature"`
}

type ArgsInsuranceClaim struct {
	Tid uint64
	Aid uint64
}

func CreateNewContract(contractEnvironment *cmn.ContractEnvironment) *cmn.ContractCallBackInfo {
	var args ArgsCreateNewContract
	MustUnmarshal(contractEnvironment.Args, &args)
	var instance = NewISC(contractEnvironment)
	resp := instance.NewContract(args.IscOwners,
		args.Funds, decodeInstructions(args.TransactionIntents), args.TransactionIntents)
	if isc.IsOK(resp) {
		err := instance.Commit()
		if err != nil {
			panic(err)
		}
	}
	return handleResponse(resp)
}

func decodeInstructions(bs [][]byte) []uip.Instruction {
	var (
		b   = bytes.NewReader(nil)
		is  = make([]uip.Instruction, len(bs))
		err error
	)
	for i := range bs {
		b.Reset(bs[i])
		is[i], err = instruction.DecodeInstruction(b)
		if err != nil {
			panic(err)
		}
	}
	return is
}

func handleResponse(resp isc.Response) *cmn.ContractCallBackInfo {
	if isc.IsOK(resp) {
		resp := resp.(*isc.ResponseData)
		return &cmn.ContractCallBackInfo{
			CodeResponse: CodeOK(),
			Data:         resp.Data,
			Value:        resp.Value,
			OutFlag:      resp.OutFlag,
		}
	} else {
		resp := resp.(*isc.ResponseError)
		return &cmn.ContractCallBackInfo{
			CodeResponse: uint32(resp.Code),
			Log:          resp.Err,
		}
	}
}
func RegisteredMethod(env *cmn.ContractEnvironment) *cmn.ContractCallBackInfo {
	var instance = NewISC(env)
	resp := registeredMethod(instance)
	if isc.IsOK(resp) {
		err := instance.Commit()
		if err != nil {
			panic(err)
		}
	}
	return handleResponse(resp)
}

type GetPCReply struct {
	PC uint64 `json:"pc"`
}

func registeredMethod(instance *ISC) isc.Response {
	switch instance.env.FuncName {
	//case "UpdateTxInfo":
	//	var args ArgsUpdateTxInfo
	//	MustUnmarshal(instance.env.Args, &args)
	//	return iscc.UpdateTxInfo(args.Tid, args.TransactionIntent)
	//case "UpdateTxFr":
	//	var args ArgsUpdateTxFr
	//	MustUnmarshal(instance.env.Args, &args)
	//	return iscc.UpdateTxFr(args.Tid, args.Fr)
	case "FreezeInfo":
		var args ArgsFreezeInfo
		MustUnmarshal(instance.env.Args, &args)
		return instance.FreezeInfo(args.Tid)
	case "UserAck":
		var args ArgsUserAck
		MustUnmarshal(instance.env.Args, &args)
		return instance.UserAck(args.Address, args.Signature)
	case "InsuranceClaim":
		return instance.InsuranceClaim(util.BytesToUint64(instance.env.Args[0:8]), util.BytesToUint64(instance.env.Args[8:16]))
	case "SettleContract":
		if instance.env.Args != nil {
			panic(ExecContractError(errors.New("this function must have no input")))
		}
		return instance.SettleContract()
	case "GetPC":
		return isc.Reply().Param(&GetPCReply{
			PC: instance.Storage.GetPC(),
		})
	case "GetMuPC":
		return isc.Reply().Param(&GetPCReply{
			PC: instance.Storage.GetMuPC(),
		})
	default:
		panic(InvalidFunctionType(instance.env.FuncName))
	}
}
