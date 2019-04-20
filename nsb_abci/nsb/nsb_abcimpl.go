package nsb

import (
	"fmt"
	"encoding/hex"
	"encoding/json"
	"bytes"
	"strings"
	"strconv"
	"github.com/Myriad-Dreamin/NetworkStatusBlockChain/crypto/trie"
	"github.com/Myriad-Dreamin/NetworkStatusBlockChain/abci/types"
	dbm "github.com/Myriad-Dreamin/NetworkStatusBlockChain/libs/db"
	"github.com/Myriad-Dreamin/NetworkStatusBlockChain/version"
	"github.com/Myriad-Dreamin/NetworkStatusBlockChain/libs/log"
	"github.com/Myriad-Dreamin/NetworkStatusBlockChain/abci/example/code"
)

var (
	stateKey = []byte("NSBStateKey")
	actionHeader = []byte("NACHD:")
)
const (
	NSBVersion version.Protocol = 0x1
)

type NSBState struct {
	db dbm.DB
	ActionRoot trie.MerkleHash `json:"action_root"`
	MerkleProofRoot trie.MerkleHash `json:"merkle_proof_root"`
	ActiveISCRoot trie.MerkleHash `json:"active_isc_root"`
	Height  int64  `json:"height"`
	AppHash []byte `json:"app_hash"`
}

func loadState(db dbm.DB) NSBState {
	stateBytes := db.Get(stateKey)
	var state NSBState
	if len(stateBytes) != 0 {
		err := json.Unmarshal(stateBytes, &state)
		if err != nil {
			panic(err)
		}
	}
	state.db = db
	return state
}

func saveState(state NSBState) {
	stateBytes, err := json.Marshal(state)
	if err != nil {
		panic(err)
	}
	state.db.Set(stateKey, stateBytes)
}


func getActionByMsgHash(db dbm.DB, msgHash []byte) Action {
	actionBytes := db.Get(msgHash)
	var action Action
	if len(actionBytes) != 0 {
		err := json.Unmarshal(actionBytes, &action)
		if err != nil {
			panic(err)
		}
	}
	return action
}

func getActionBySignature(db dbm.DB, signature []byte) Action {
	msgHash := db.Get(signature)
	var action Action
	if len(msgHash) != 0 {
		action = getActionByMsgHash(db, msgHash)
	}
	return action
}


func getMerkleProofByHash(db dbm.DB, prvHash []byte) MerkleProof {
	proofBytes := db.Get(prvHash)
	var merkleProof MerkleProof
	if len(proofBytes) != 0 {
		err := json.Unmarshal(proofBytes, &merkleProof)
		if err != nil {
			panic(err)
		}
	}
	return merkleProof
}


type NSBApplication struct {
	types.BaseApplication
	state NSBState
	
	ValUpdates []types.ValidatorUpdate
	logger log.Logger
}

func NewNSBApplication(dbDir string) (*NSBApplication, error) {
	name := "nsb"
	db, err := dbm.NewGoLevelDB(name, dbDir)
	if err != nil {
		return nil, err
	}

	state := loadState(db)

	return &NSBApplication{
		state: state,
		logger: log.NewNopLogger(),
	}, nil
}

func (nsb *NSBApplication) Info(req types.RequestInfo) types.ResponseInfo {
	return types.ResponseInfo{
		Data:       fmt.Sprintf(
			"{\"action_root\":%v, \"merkle_proof_root\":%v, \"active_isc_root\":%v, }",
			hex.EncodeToString(nsb.state.ActionRoot),
			hex.EncodeToString(nsb.state.MerkleProofRoot),
			hex.EncodeToString(nsb.state.ActiveISCRoot)),
		Version:    version.ABCIVersion,
		AppVersion: NSBVersion.Uint64(),
	}
}

// Save the validators in the merkle tree
func (nsb *NSBApplication) InitChain(req types.RequestInitChain) types.ResponseInitChain {
	for _, v := range req.Validators {
		r := nsb.updateValidator(v)
		if r.IsErr() {
			nsb.logger.Error("Error updating validators", "r", r)
		}
	}
	return types.ResponseInitChain{}
}

// Track the block hash and header information
func (nsb *NSBApplication) BeginBlock(req types.RequestBeginBlock) types.ResponseBeginBlock {
	// reset valset changes
	nsb.ValUpdates = make([]types.ValidatorUpdate, 0)
	return types.ResponseBeginBlock{}
}

// Update the validator set
func (nsb *NSBApplication) EndBlock(req types.RequestEndBlock) types.ResponseEndBlock {
	return types.ResponseEndBlock{ValidatorUpdates: nsb.ValUpdates}
}

func (nsb *NSBApplication) CheckTx(tx []byte) types.ResponseCheckTx {
	return types.ResponseCheckTx{Code: code.CodeTypeOK, GasWanted: 1}
}

func (nsb *NSBApplication) SetLogger(l log.Logger) {
	nsb.logger = l
}

func (nsb *NSBApplication) deliverTx(tx []byte) types.ResponseDeliverTx {
	return types.ResponseDeliverTx{Code: code.CodeTypeOK}
}

func (nsb *NSBApplication) DeliverTx(tx []byte) types.ResponseDeliverTx {
	// if it starts with "val:", update the validator set
	// format is "val:pubkey/power"
	if isValidatorTx(tx) {
		// update validators in the merkle tree
		// and in app.ValUpdates
		return nsb.execValidatorTx(tx)
	}

	// otherwise, update the key-value store
	return nsb.deliverTx(tx)
}

func (nsb *NSBApplication) Validators() (validators []types.ValidatorUpdate) {
	itr := nsb.state.db.Iterator(nil, nil)
	for ; itr.Valid(); itr.Next() {
		if isValidatorTx(itr.Key()) {
			validator := new(types.ValidatorUpdate)
			err := types.ReadMessage(bytes.NewBuffer(itr.Value()), validator)
			if err != nil {
				panic(err)
			}
			validators = append(validators, *validator)
		}
	}
	return
}

func MakeValSetChangeTx(pubkey types.PubKey, power int64) []byte {
	return []byte(fmt.Sprintf("val:%X/%d", pubkey.Data, power))
}

func isValidatorTx(tx []byte) bool {
	return true// strings.HasPrefix(string(tx), ValidatorSetChangePrefix)
}

func (nsb *NSBApplication) execValidatorTx(tx []byte) types.ResponseDeliverTx {
	// tx = tx[len(ValidatorSetChangePrefix):]

	//get the pubkey and power
	pubKeyAndPower := strings.Split(string(tx), "/")
	if len(pubKeyAndPower) != 2 {
		return types.ResponseDeliverTx{
			Code: code.CodeTypeEncodingError,
			Log:  fmt.Sprintf("Expected 'pubkey/power'. Got %v", pubKeyAndPower)}
	}
	pubkeyS, powerS := pubKeyAndPower[0], pubKeyAndPower[1]

	// decode the pubkey
	pubkey, err := hex.DecodeString(pubkeyS)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: code.CodeTypeEncodingError,
			Log:  fmt.Sprintf("Pubkey (%s) is invalid hex", pubkeyS)}
	}

	// decode the power
	power, err := strconv.ParseInt(powerS, 10, 64)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: code.CodeTypeEncodingError,
			Log:  fmt.Sprintf("Power (%s) is not an int", powerS)}
	}

	// update
	return nsb.updateValidator(types.Ed25519ValidatorUpdate(pubkey, int64(power)))
}

func (nsb *NSBApplication) updateValidator(v types.ValidatorUpdate) types.ResponseDeliverTx {
	return types.ResponseDeliverTx{Code: code.CodeTypeOK}
}


// func (nsb *NSBApplication) Query(req types.RequestQuery) types.ResponseQuery {

// }

