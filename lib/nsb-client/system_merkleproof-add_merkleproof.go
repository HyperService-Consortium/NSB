package nsbcli

import (
	"encoding/json"
	transactiontype "github.com/HyperService-Consortium/NSB/application/transaction-type"
	system_merkle_proof "github.com/HyperService-Consortium/NSB/contract/system/merkle-proof"
	"github.com/HyperService-Consortium/NSB/lib/nsb-client/nsb-message"

	"github.com/HyperService-Consortium/NSB/grpc/nsbrpc"
	uip "github.com/HyperService-Consortium/go-uip/uip"
)

func (nsb *NSBClient) AddMerkleProof(
	user uip.Signer, toAddress []byte,
	merkletype uint16, rootHash, proof, key, value []byte,
) (*nsb_message.ResultInfo, error) {
	// fmt.Println(string(buf.Bytes()))
	fap, err := nsb.addMerkleProof(merkletype, rootHash, proof, key, value)
	if err != nil {
		return nil, err
	}
	txHeader, err := nsb.CreateContractPacket(user, toAddress, []byte{0}, fap)
	if err != nil {
		return nil, err
	}
	ret, err := nsb.sendContractTx(transactiontype.SystemCall, txHeader)
	if err != nil {
		return nil, err
	}
	// fmt.Println(PretiStruct(ret), err)
	return ret, nil
}

func (nsb *NSBClient) addMerkleProof(
	merkletype uint16, rootHash []byte, proof []byte, key []byte, value []byte,
) (*nsbrpc.FAPair, error) {
	var args system_merkle_proof.ArgsValidateMerkleProof

	args.Type = merkletype
	args.RootHash = rootHash

	args.Proof = proof
	args.Key = key
	args.Value = value
	b, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	var fap = new(nsbrpc.FAPair)
	fap.FuncName = "system.merkleproof@validateMerkleProof"
	fap.Args = b
	// fmt.Println(PretiStruct(args), b)
	return fap, err
}
