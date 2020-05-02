package nsb_message

import (
	"time"
)

/******************************* abci_info ************************************/

// AbciInfo is struct description of abci information in json
type AbciInfo struct {
	Response *AbciInfoResponse `json:"response"`
}

// AbciInfoResponse is struct description of abci information response in json
type AbciInfoResponse struct {
	Data             string `json:"data"`
	Version          string `json:"version"`
	AppVersion       string `json:"app_version"`
	LastBlockHeight  string `json:"last_block_height"`
	LastBlockAppHash []byte `json:"last_block_app_hash"`
}

/******************************* block_info ***********************************/

// BlockInfo is struct description of block information in json
type BlockInfo struct {
	BlockMeta *BlockMeta `json:"block_meta"`
	Block     *Block     `json:"block"`
}

// BlockMeta is struct description of block meta in json
type BlockMeta struct {
	BlockID *BlockID `json:"block_id"`
	Header  Header   `json:"header"`
}

// Block is struct description of block in json
type Block struct {
	Header     *Header     `json:"header"`
	Data       *Data       `json:"data"`
	Evidence   *Evidence   `json:"evidence"`
	LastCommit *LastCommit `json:"last_commit"`
}

// BlockID is struct description of block id in json
type BlockID struct {
	Hash  string `json:"hash"`
	Parts *Parts `json:"parts"`
}

// Parts is struct description of parts in json
type Parts struct {
	Total string `json:"total"`
	Hash  string `json:"hash"`
}

// Header is struct description of header in json
type Header struct {
	Version            *Version     `json:"version"`
	ChainID            string       `json:"chain_id"`
	Height             string       `json:"height"`
	Time               time.Time    `json:"time"`
	NumTxs             string       `json:"num_txs"`
	TotalTxs           string       `json:"total_txs"`
	LastBlockID        *LastBlockID `json:"last_block_id"`
	LastCommitHash     string       `json:"last_commit_hash"`
	DataHash           string       `json:"data_hash"`
	ValidatorsHash     string       `json:"validators_hash"`
	NextValidatorsHash string       `json:"next_validators_hash"`
	ConsensusHash      string       `json:"consensus_hash"`
	AppHash            string       `json:"app_hash"`
	LastResultsHash    string       `json:"last_results_hash"`
	EvidenceHash       string       `json:"evidence_hash"`
	ProposerAddress    string       `json:"proposer_address"`
}

// Version is struct description of version in json
type Version struct {
	Block string `json:"block"`
	App   string `json:"app"`
}

// LastBlockID is struct description of last block id in json
type LastBlockID struct {
	Hash  string `json:"hash"`
	Parts *Parts `json:"parts"`
}

// Data is struct description of data in json
type Data struct {
	Txs []string `json:"txs"`
}

// Evidence temporarily unknown
type Evidence struct {
	MaxAge string `json:"max_age"`
}

// LastCommit is struct description of last commit in json
type LastCommit struct {
	BlockID    *BlockID      `json:"block_id"`
	Precommits []*Precommits `json:"precommits"`
}

// Precommits is struct description of precommits in json
type Precommits struct {
	Type             int       `json:"type"`
	Height           string    `json:"height"`
	Round            string    `json:"round"`
	BlockID          *BlockID  `json:"block_id"`
	Timestamp        time.Time `json:"timestamp"`
	ValidatorAddress string    `json:"validator_address"`
	ValidatorIndex   string    `json:"validator_index"`
	Signature        string    `json:"signature"`
}

/**************************** block_results_info ******************************/

// BlockResultsInfo is struct description of block results information in json
type BlockResultsInfo struct {
	Height  string   `json:"height"`
	Results *Results `json:"results"`
}

// Results is struct description of results in json
type Results struct {
	DeliverTxInfo  []*DeliverTx    `json:"DeliverTx"`
	EndBlockInfo   *EndBlockInfo   `json:"EndBlock"`
	BeginBlockInfo *BeginBlockInfo `json:"BeginBlock"`
}

// DeliverTx is struct description of deliver tx in json
type DeliverTx struct {
	Info   string   `json:"info"`
	Log    string   `json:"log"`
	Data   []byte   `json:"data"`
	Events []Events `json:"events"`
}

type Attributes struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Tags is struct description of tags in json
type Events struct {
	Type       string       `json:"type"`
	Attributes []Attributes `json:"attributes"`
}

// EndBlockInfo temporarily unknown
type EndBlockInfo struct {
	ValidatorUpdates interface{} `json:"validator_updates"`
}

// BeginBlockInfo is struct description of begin block information in json
type BeginBlockInfo struct {
}

/****************************** blocks_info ***********************************/

// BlocksInfo is struct description of blocks information in json
type BlocksInfo struct {
	LastHeight string       `json:"last_height"`
	BlockMetas []*BlockMeta `json:"block_metas"`
}

/****************************** commit_info ***********************************/

// CommitInfo is struct description of commit information in json
type CommitInfo struct {
	SignedHeader *SignedHeader `json:"signed_header"`
	Canonical    bool          `json:"canonical"`
}

// SignedHeader is struct description of signed header in json
type SignedHeader struct {
	Header *Header `json:"header"`
	Commit *Commit `json:"commit"`
}

// Commit is struct description of commit in json
type Commit struct {
	BlockID    *BlockID      `json:"block_id"`
	Precommits []*Precommits `json:"precommits"`
}

/************************* consensus_params_info ******************************/

// ConsensusParamsInfo is struct description of consensus params information in json
type ConsensusParamsInfo struct {
	BlockHeight     string           `json:"block_height"`
	ConsensusParams *ConsensusParams `json:"consensus_params"`
}

// ConsensusParams is struct description of consensus params in json
type ConsensusParams struct {
	BlockConsensus *BlockConsensus `json:"block"`
	Evidence       *Evidence       `json:"evidence"`
	Validator      *Validator      `json:"validator"`
}

// BlockConsensus is struct description of block consensus in json
type BlockConsensus struct {
	MaxBytes   string `json:"max_bytes"`
	MaxGas     string `json:"max_gas"`
	TimeIotaMs string `json:"time_iota_ms"`
}

// Validator is struct description of validator in json
type Validator struct {
	PubKeyTypes []string `json:"pub_key_types"`
}

/************************* consensus_state_info ******************************/

// ConsensusStateInfo is struct description of consensus state information in json
type ConsensusStateInfo struct {
	RdState *RoundState `json:"round_state"`
}

// RoundState is struct description of round state in json
type RoundState struct {
	HeightRoundStep   string        `json:"height/round/step"`
	StartTime         string        `json:"start_time"`
	ProposalBlockHash string        `json:"proposal_block_hash"`
	LockedBlockHash   string        `json:"locked_block_hash"`
	ValidBlockHash    string        `json:"valid_block_hash"`
	HeightVoteSet     []*HeightVote `json:"height_vote_set"`
}

// HeightVote is struct description of height vote in json
type HeightVote struct {
	Round              string   `json:"round"`
	Prevotes           []string `json:"prevotes"`
	PrevotesBitArray   string   `json:"prevotes_bit_array"`
	Precommits         []string `json:"precommits"`
	PrecommitsBitArray string   `json:"precommits_bit_array"`
}

/************************* genesis_info ******************************/

// GenesisInfo is struct description of genesis information in json
type GenesisInfo struct {
	Genesis *Genesis `json:"genesis"`
}

// Genesis is struct description of genesis block in json
type Genesis struct {
	GenesisTime     string          `json:"genesis_time"`
	ChainID         string          `json:"chain_id"`
	ConsensusParams ConsensusParams `json:"consensus_params"`
	Validators      []*Validator    `json:"validators"`
	AppHash         string          `json:"app_hash"`
}

/************************* net_info ******************************/

// NetInfo is struct description of net information in json
type NetInfo struct {
	Listening bool     `json:"listening"`
	Listeners []string `json:"listeners"` // /NOT CONFIRMED
	NPeers    string   `json:"n_peers"`
	Peers     []string `json:"peers"` // NOT CONFIRMED
}

/************************* num_unconfirmed_txs_info ******************************/

// NumUnconfirmedTxsInfo is struct description of number unconfirmed txs information in
// json
type NumUnconfirmedTxsInfo struct {
	NTxs       string   `json:"n_txs"`
	Total      string   `json:"total"`
	TotalBytes string   `json:"total_bytes"`
	Txs        []string `json:"txs"` // NOT CONFIRMED
}

/************************* status_info ******************************/

// StatusInfo is struct description of status information in json
type StatusInfo struct {
	NodeInfo      *NodeInfo      `json:"node_info"`
	SyncInfo      *SyncInfo      `json:"sync_info"`
	ValidatorInfo *ValidatorInfo `json:"validator_info"`
}

// NodeInfo is struct description of node information in json
type NodeInfo struct {
	ProtocolVersion *ProtocolVersion `json:"protocol_version"`
	ID              string           `json:"id"`
	ListenAddr      string           `json:"listen_addr"`
	NetWork         string           `json:"network"`
	Version         string           `json:"version"`
	Channels        string           `json:"channels"`
	Moniker         string           `json:"moniker"`
	Other           *Other           `json:"other"`
}

// ProtocolVersion is struct description of protocol version in json
type ProtocolVersion struct {
	P2P   string `json:"p2p"`
	Block string `json:"block"`
	App   string `json:"app"`
}

// Other is struct description of other in json
type Other struct {
	TxIndex    string `json:"tx_index"`
	RPCAddress string `json:"rpc_address"`
}

// SyncInfo is struct description of sync information in json
type SyncInfo struct {
	LatestBlockHash   string `json:"latest_block_hash"`
	LatestAppHash     string `json:"latest_app_hash"`
	LatestBlockHeight string `json:"latest_block_height"`
	LatestBlockTime   string `json:"latest_block_time"`
	CatchingUp        bool   `json:"catching_up"`
}

// ValidatorInfo is struct description of validator information in json
type ValidatorInfo struct {
	Address     string  `json:"address"`
	PubKey      *PubKey `json:"pub_key"`
	VotingPower string  `json:"voting_power"`
}

// PubKey is struct description of public key in json
type PubKey struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

/************************* validators_info ******************************/

// ValidatorsInfo is struct description of validators information in json
type ValidatorsInfo struct {
	BlockHeight string               `json:"block_height"`
	Validators  []*FullValidatorInfo `json:"validators"`
}

// FullValidatorInfo is struct description of full validator information in json
type FullValidatorInfo struct {
	Address          string  `json:"address"`
	PubKey           *PubKey `json:"pub_key"`
	VotingPower      string  `json:"voting_power"`
	ProposerPriority string  `json:"proposer_priority"`
}

/*************************** result_info ********************************/

// ResultInfo is struct description of result information in json
type ResultInfo struct {
	CheckTx   interface{} `json:"check_tx"`
	DeliverTx DeliverTx   `json:"deliver_tx"`
	Hash      string      `json:"hash"`
	Height    string      `json:"height"`
}

// CheckTx temporarily unknown
type CheckTx struct {
	Info   string   `json:"info"`
	Log    string   `json:"log"`
	Data   []byte   `json:"data"`
	Events []Events `json:"events"`
}

/**************************** proof_info ********************************/

type ProofResponse struct {
	Value  []byte `json:"value"`
	Key    []byte `json:"key"`
	Height string `json:"height"`
	Info   string `json:"info"`
	Log    string `json:"log"`
	Code   uint32 `json:"code"`
}

type ProofInfo struct {
	Response ProofResponse `json:"response"`
}

/****************************** receipt *********************************/

type TransactionReceipt struct {
	Code int    `json:"code"`
	Data string `json:"data"`
	Log  string `json:"log"`
	Hash string `json:"hash"`
}
