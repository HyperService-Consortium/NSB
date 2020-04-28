package isc

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/NSB/common"
	merkle_proof "github.com/HyperService-Consortium/go-uip/const/merkle-proof-type"
	TxState "github.com/HyperService-Consortium/go-uip/const/transaction_state_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	error2 "github.com/HyperService-Consortium/go-uip/errorn"
	"github.com/HyperService-Consortium/go-uip/isc"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/op-intent/instruction"
	"github.com/HyperService-Consortium/go-uip/op-intent/parser"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var user0 = []byte{1}

func encodeInstructions(is []uip.Instruction) (bs [][]byte) {
	return sugar.HandlerError(instruction.EncodeInstructions(is)).([][]byte)
}

func closeDB() {

	if __x_ldb != nil {
		sugar.HandlerError0(__x_ldb.Close())
		__x_ldb = nil
	}
}

func TestMain(m *testing.M) {
	m.Run()
	closeDB()
}

var c2 = obj{
	"contractName": "c2",
	"domain":       2,
	"address":      "0x3723261b2a5a62b778b5c74318534d7fdf8db38c",
}

type obj map[string]interface{}

var opIntents = obj{
	"op-intents": []obj{
		{
			"name": "op1",
			"type": "Payment",
			"src": obj{
				"domain":    1,
				"user_name": "a1",
			},
			"dst": obj{
				"domain":    2,
				"user_name": "a2",
			},
			"amount": "1a",
			"unit":   "ether",
		},
		{
			"name":    "op2",
			"type":    "ContractInvocation",
			"invoker": "a2",
			"func":    "vote",
			"contract": obj{
				"domain":  2,
				"address": "0x3723261b2a5a62b778b5c74318534d7fdf8db38c",
			},
			"parameters": []obj{},
		},
		{
			"name": "if-op",
			"type": "IfStatement",
			"if": []obj{
				{
					"name":    "op3",
					"type":    "ContractInvocation",
					"invoker": "a2",
					"func":    "vote",
					"contract": obj{
						"address": "0x3723261b2a5a62b778b5c74318d34d7fdbadb38e",
					},
					"parameters": []obj{},
				},
				{
					"name": "op4",
					"type": "Payment",
					"src": obj{
						"domain":    1,
						"user_name": "a1",
					},
					"dst": obj{
						"domain":    2,
						"user_name": "a2",
					},
					"amount": "aa",
					"unit":   "ether",
				},
			},
			"else": []obj{
				{
					"name":    "op5",
					"type":    "ContractInvocation",
					"invoker": "a2",
					"func":    "vote",
					"contract": obj{
						"domain":  2,
						"address": "0x3723261b2a5a62b778b5c74318534d7fdf8db38c",
					},
					"parameters": []obj{},
				},
			},
			"condition": obj{
				"left": obj{
					"type": "uint256",
					"value": obj{
						"contract": "c2",
						"field":    "num_count",
						"pos":      "00",
					},
				},
				"right": obj{
					"type": "uint256",
					"value": obj{
						"contract": "c2",
						"field":    "totalVotes",
						"pos":      "01",
					},
				},
				"sign": "Greater",
			},
		},
		{
			"name": "loop",
			"type": "loopFunction",
			"loop": []obj{
				{
					"name":    "op6",
					"type":    "ContractInvocation",
					"invoker": "a2",
					"func":    "vote",
					"contract": obj{
						"domain":  2,
						"address": "0x3723261b2a5a62b778b5c74318534d7fdf8db38c",
					},
					"parameters": []obj{},
				},
			},
			"loopTime": "5",
		},
	},
	"dependencies": []obj{},
	"contracts": []obj{
		{
			"contractName": "c1",
			"domain":       1,
			"address":      "0xafc7d2959e72081770304f6474151293be1fbba7",
		},
		c2,
		{
			"contractName": "c3",
			"domain":       3,
			"address":      "0x3723261b2a5a62b778b5c74318d34d7fdbadb38e",
		},
	},
	"accounts": []obj{
		{
			"userName": "a1",
			"domain":   1,
			"address":  "0x7019fa779024c0a0eac1d8475733eefe10a49f3b",
		},
		{
			"userName": "a2",
			"domain":   2,
			"address":  "0x47a1cdb6594d6efed3a6b917f2fbaa2bbcf61a2e",
		},
		{
			"userName": "a3",
			"domain":   3,
			"address":  "0x47a1cdb6559d6efed3a6b917f2fbaa2bbcf61a2e",
		},
	},
}

func setupOpIntent(t *testing.T) (env *common.ContractEnvironment, instance *ISC) {
	t.Helper()
	var intents parser.TxIntents

	ier, err := opintent.NewInitializer(uip.BlockChainGetterNilImpl{}, mAccountProvider{})
	if err != nil {
		t.Error(err)
		return
	}

	p := packet{
		content: sugar.HandlerError(json.Marshal(opIntents)).([]byte),
	}

	intents, err = ier.ParseR(p)
	if err != nil {
		t.Error(err)
		pe := err.(*error2.ParseError)
		fmt.Println(string(sugar.HandlerError(pe.Serialize()).([]byte)))
		return
	}
	var txIntents = intents.GetTxIntents()
	var instructions []uip.Instruction
	for i := range txIntents {
		fmt.Println(i, txIntents[i].GetName(), txIntents[i].GetInstruction().GetType())
		instructions = append(instructions, txIntents[i].GetInstruction())
	}

	env = createRoot(t, user0, []byte{2})

	instance = NewISC(env)

	var newContractReply isc.NewContractReply

	unpack(instance.NewContract([][]byte{env.From}, []uint64{0}, instructions, encodeInstructions(instructions)), &newContractReply)
	commit(t, instance)
	fmt.Println(newContractReply)
	assert.EqualValues(t, isc.StateInitializing, instance.Storage.GetISCState())

	for i := range instructions {
		assert.EqualValues(t, isc.OK, instance.FreezeInfo(uint64(i)))
		commit(t, instance)
	}
	assert.EqualValues(t, isc.StateInitialized, instance.Storage.GetISCState())

	assert.EqualValues(t, isc.OK, instance.UserAck(user0, []byte("todo")))
	assert.EqualValues(t, isc.StateOpening, instance.Storage.GetISCState())
	return
}

func TestIfScenario_IfYes(t *testing.T) {
	env, instance := setupOpIntent(t)
	//0 op1.cna 0
	doTransaction(t, instance, uint64(0))
	//1 op1.cnb 0
	doTransaction(t, instance, uint64(1))
	//2 op2 1

	BranchIfTest0(env, true)
	doTransaction(t, instance, uint64(2))
	//3 if-op.goto.if 3
	//4 op5 1
	//doTransaction(t, instance, uint64(4))
	//5 if-op.goto.else 2
	//6 op3 1
	doTransaction(t, instance, uint64(6))
	////7 op4.cna 0
	doTransaction(t, instance, uint64(7))
	////8 op4.cnb 0
	doTransaction(t, instance, uint64(8))
	//9 loop.loopBegin 3
	//10 op6 1
	for i := 0; i < 5; i++ {
		doTransaction(t, instance, uint64(10))
	}
	//11 loop.addLoopVar 4
	//12 loop.loopEnd 2
	//13 loop.resetLoopVar 4

	assert.EqualValues(t, isc.StateSettling, instance.Storage.GetISCState())

	assert.EqualValues(t, isc.OK, instance.SettleContract())
	commit(t, instance)

	assert.EqualValues(t, isc.StateClosed, instance.Storage.GetISCState())
}

func TestIfScenario_IfNo(t *testing.T) {
	env, instance := setupOpIntent(t)
	//0 op1.cna 0
	doTransaction(t, instance, uint64(0))
	//1 op1.cnb 0
	doTransaction(t, instance, uint64(1))
	//2 op2 1

	BranchIfTest0(env, false)
	doTransaction(t, instance, uint64(2))
	//3 if-op.goto.if 3
	//4 op5 1
	doTransaction(t, instance, uint64(4))
	//5 if-op.goto.else 2
	//6 op3 1
	//7 op4.cna 0
	//8 op4.cnb 0
	//9 loop.loopBegin 3
	//10 op6 1
	for i := 0; i < 5; i++ {
		doTransaction(t, instance, uint64(10))
	}
	//11 loop.addLoopVar 4
	//12 loop.loopEnd 2
	//13 loop.resetLoopVar 4

	assert.EqualValues(t, isc.StateSettling, instance.Storage.GetISCState())

	assert.EqualValues(t, isc.OK, instance.SettleContract())
	commit(t, instance)

	assert.EqualValues(t, isc.StateClosed, instance.Storage.GetISCState())
}

func doTransaction(t *testing.T, instance *ISC, pc uint64) {
	t.Helper()
	assert.EqualValues(t, pc, instance.GetPC())
	assert.EqualValues(t, isc.OK, instance.InsuranceClaim(pc, TxState.Instantiating))
	commit(t, instance)

	assert.EqualValues(t, pc, instance.GetPC())
	assert.EqualValues(t, isc.OK, instance.InsuranceClaim(pc, TxState.Open))
	commit(t, instance)

	assert.EqualValues(t, pc, instance.GetPC())
	assert.EqualValues(t, isc.OK, instance.InsuranceClaim(pc, TxState.Opened))
	commit(t, instance)

	assert.EqualValues(t, pc, instance.GetPC())
	assert.EqualValues(t, isc.OK, instance.InsuranceClaim(pc, TxState.Closed))
	commit(t, instance)
}

type StorageKey struct {
	chainID         uip.ChainID
	typeID          uip.TypeID
	contractAddress string
	pos             string
	description     string
}

type storageImpl struct {
	externalStorage map[StorageKey]uip.Variable
}

func (c *storageImpl) GetTransactionProof(chainID uip.ChainID, blockID uip.BlockID, color []byte) (uip.MerkleProof, error) {
	panic("implement me")
}

func (c *storageImpl) GetStorageAt(chainID uip.ChainID, typeID uip.TypeID,
	contractAddress uip.ContractAddress, pos []byte, description []byte) (uip.Variable, error) {
	if c.externalStorage == nil {
		c.externalStorage = make(map[StorageKey]uip.Variable)
	}
	if x, ok := c.externalStorage[StorageKey{
		chainID:         chainID,
		typeID:          typeID,
		contractAddress: string(contractAddress),
		pos:             string(pos),
		description:     string(description),
	}]; ok {
		return x, nil
	} else {
		return nil, errors.New("no found")
	}
}

func (c *storageImpl) ProvideExternalStorageAt(chainID uip.ChainID, typeID uip.TypeID,
	contractAddress uip.ContractAddress, pos []byte, description []byte, ref uip.Variable) {
	if c.externalStorage == nil {
		c.externalStorage = make(map[StorageKey]uip.Variable)
	}
	c.externalStorage[StorageKey{
		chainID:         chainID,
		typeID:          typeID,
		contractAddress: string(contractAddress),
		pos:             string(pos),
		description:     string(description),
	}] = ref
}

func BranchIfTest0(env *common.ContractEnvironment, ifOrNot bool) {
	//"left": obj{
	//	"type": "uint256",
	//	"value": obj{
	//		"contract": "c2",
	//		"field":    "num_count",
	//		"pos":      "00",
	//	},
	//},
	var l0 uip.Variable
	//
	if ifOrNot {
		l0 = uip.VariableImpl{Type: value_type.Uint256, Value: big.NewInt(2)}
	} else {
		l0 = uip.VariableImpl{Type: value_type.Uint256, Value: big.NewInt(1)}
	}
	env.BN.(*storageImpl).ProvideExternalStorageAt(
		uip.ChainIDUnderlyingType(c2["domain"].(int)), value_type.Uint256,
		sugar.HandlerError(hex.DecodeString(c2["address"].(string)[2:])).([]byte), []byte{0}, []byte("num_count"),
		l0)
	//	"right": obj{
	//	"type": "uint256",
	//	"value": obj{
	//		"contract": "c2",
	//		"field":    "totalVotes",
	//		"pos":      "01",
	//	},
	//},
	env.BN.(*storageImpl).ProvideExternalStorageAt(
		uip.ChainIDUnderlyingType(c2["domain"].(int)), value_type.Uint256,
		sugar.HandlerError(hex.DecodeString(c2["address"].(string)[2:])).([]byte), []byte{1}, []byte("totalVotes"),
		uip.VariableImpl{Type: value_type.Uint256, Value: big.NewInt(1)})
	//	"sign": "Greater",
}

func commit(t *testing.T, isc *ISC) {
	t.Helper()
	sugar.HandlerError0(isc.Commit())
	sugar.HandlerError(isc.env.Storage.Commit())
}

func unpack(response isc.Response, n *isc.NewContractReply) {
	var data = response.(*isc.ResponseData).Data
	//fmt.Println(string(data))
	sugar.HandlerError0(json.Unmarshal(data, n))
}

type packet struct {
	content []byte
}

func (p packet) GetContent() (content []byte) {
	return p.content
}

type mAccountProvider struct {
}

func (a mAccountProvider) AccountBase() uip.AccountBase {
	return a
}

func (mAccountProvider) Get(_ string, chainId uint64) (uip.Account, error) {
	return &uip.AccountImpl{
		ChainId: chainId,
		Address: []byte("121313212313133123333333333333333313"),
	}, nil
}

func (mAccountProvider) GetRelay(domain uint64) (uip.Account, error) {
	return &uip.AccountImpl{
		ChainId: domain,
		Address: []byte("99999"),
	}, nil
}

func (mAccountProvider) GetTransactionProofType(_ uint64) (uip.MerkleProofType, error) {
	return merkle_proof.MerklePatriciaTrieUsingKeccak256, nil
}
