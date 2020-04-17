package isc

import (
	"encoding/json"
	cmn "github.com/HyperService-Consortium/NSB/common"
	. "github.com/HyperService-Consortium/NSB/common/contract_response"
	"github.com/HyperService-Consortium/NSB/merkmap"
	"github.com/HyperService-Consortium/go-uip/isc"
	"github.com/HyperService-Consortium/go-uip/storage"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
)

type ISC struct {
	isc.ISC
	env *cmn.ContractEnvironment
}

func NewISC(env *cmn.ContractEnvironment) (instance *ISC) {
	instance = new(ISC)
	instance.env = env
	ctx := context{env}
	instance.ISC = *isc.NewISC(ctx, storage.NewVM(ctx))
	return
}

type context struct {
	env *cmn.ContractEnvironment
}

type merkMapI struct {
	merk *merkmap.MerkMap
}

func (m merkMapI) Update(key []byte, value []byte) error {
	return m.merk.TryUpdate(key, value)
}

func (m merkMapI) Get(key []byte) ([]byte, error) {
	return m.merk.TryGet(key)
}

func (m merkMapI) Delete(key []byte) error {
	return m.merk.TryDelete(key)
}

func (m merkMapI) Prove(key []byte) ([][]byte, error) {
	return m.merk.TryProve(key)
}

func (m merkMapI) MakeProof(key []byte) string {
	return m.merk.MakeProof(key)
}

func (m merkMapI) MakeErrorProof(err error) string {
	return m.merk.MakeErrorProof(err)
}

func (m merkMapI) MakeErrorProofFromString(str string) string {
	return m.merk.MakeErrorProofFromString(str)
}

func (c context) ArrangeSlot(newSlot string) storage.MerkMap {
	return merkMapI{c.env.Storage.MakeStorageSlot(newSlot)}
}

func (c context) Commit() error {
	_, err := c.env.Storage.Commit()
	return err
}

func (c context) Sender() []byte {
	return c.env.From
}

func (c context) Address() []byte {
	return c.env.ContractAddress
}

type impl struct {
	v uip.Variable
}

func (i impl) GetGVMType() gvm.RefType {
	return gvm.RefType(i.v.GetType())
}

func (i impl) Unwrap() interface{} {
	return i.v.GetValue()
}

func (i impl) Encode() ([]byte, error) {
	panic("implement me")
}

func (c context) GetExternalStorageAt(chainID uip.ChainID, typeID uip.TypeID,
	contractAddress uip.ContractAddress, pos []byte, description []byte) (gvm.Ref, error) {
	v, err := c.env.BN.GetStorageAt(chainID, typeID, contractAddress, pos, description)
	if err != nil {
		return nil, err
	}
	return impl{v}, nil
}

func (iscc *ISC) GetOwners() *cmn.ContractCallBackInfo {
	owners := iscc.env.Storage.NewBytesArray("owners")
	var ret [][]byte
	for idx := uint64(0); idx < owners.Length(); idx++ {
		ret = append(ret, owners.Get(idx))
	}

	b, err := json.Marshal(ret)
	if err != nil {
		return ExecContractError(err)
	}

	return &cmn.ContractCallBackInfo{
		CodeResponse: CodeOK(),
		Data:         b,
	}
}

func (iscc *ISC) IsOwner(address []byte) *cmn.ContractCallBackInfo {
	isOwner := iscc.env.Storage.NewBoolMap("iscOwner")
	b, err := json.Marshal(isOwner.Get(address))
	if err != nil {
		return ExecContractError(err)
	}

	return &cmn.ContractCallBackInfo{
		CodeResponse: CodeOK(),
		Data:         b,
	}
}

func (iscc *ISC) GetState() *cmn.ContractCallBackInfo {
	return &cmn.ContractCallBackInfo{
		CodeResponse: CodeOK(),
		Data:         []byte{iscc.env.Storage.GetUint8("iscState")},
	}
}
