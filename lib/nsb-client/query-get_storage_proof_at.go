package nsbcli

import (
	"encoding/json"
	application "github.com/HyperService-Consortium/NSB/application"
	"github.com/HyperService-Consortium/NSB/application/nsb_proto"
	system_merkle_proof "github.com/HyperService-Consortium/NSB/contract/system/merkle-proof"
	"github.com/HyperService-Consortium/NSB/merkmap"
)

//type ProofJson struct {
//	Proof [][]byte `json:"proof"`
//	Key   []byte   `json:"key"`
//	Value []byte   `json:"value"`
//	Log   string   `json:"log"`
//}

type StorageProofArgs = nsb_proto.ArgsGetStorageAt
type StorageProofResponse = merkmap.ProofJson

func (nsb *NSBClient) GetStorageProofAt(
	args StorageProofArgs) (
	*StorageProofResponse, error) {
	var data, err = json.Marshal(args)
	if err != nil {
		return nil, err
	}
	resp, err := nsb.GetProof(data, application.QueryKeyGetStorageAt)
	return resp, nil
}

func newNSBClient(host string) system_merkle_proof.NSBClient {
	return NewNSBClient(host)
}

func init() {
	system_merkle_proof.SetNewNSBClient(newNSBClient)
}
