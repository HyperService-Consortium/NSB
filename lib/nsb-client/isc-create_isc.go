package nsbcli

import (
	"encoding/json"
	transactiontype "github.com/HyperService-Consortium/NSB/application/transaction-type"
	ISC "github.com/HyperService-Consortium/NSB/contract/isc"
	"github.com/HyperService-Consortium/NSB/grpc/nsbrpc"
	"github.com/HyperService-Consortium/go-uip/isc"
	uip "github.com/HyperService-Consortium/go-uip/uip"
)

func (nsb *NSBClient) CreateISC(
	user uip.Signer,
	funds []uint64, iscOwners [][]byte,
	bytesTransactionIntents [][]byte,
	vesSig []byte,
) ([]byte, error) {
	fap, err := nsb.createISC(funds, iscOwners, bytesTransactionIntents, vesSig)
	if err != nil {
		return nil, err
	}
	txHeader, err := nsb.CreateContractPacket(user, nil, []byte{0}, fap)
	if err != nil {
		return nil, err
	}
	ret, err := nsb.sendContractTx(transactiontype.CreateContract, txHeader)

	if err != nil {
		return nil, err
	}

	var reply isc.NewContractReply
	err = json.Unmarshal(ret.DeliverTx.Data, &reply)
	if err != nil {
		return nil, err
	}
	return reply.Address, nil
}

func (nsb *NSBClient) createISC(
	funds []uint64, iscOwners [][]byte,
	bytesTransactionIntents [][]byte,
	vesSig []byte,
) (*nsbrpc.FAPair, error) {
	//for idx, txb := range bytesTransactionIntents {
	//	err := json.Unmarshal(txb, &txm)
	//	if err != nil {
	//		return nil, err
	//	}
	//	var txi = new(iscTransactionIntent.TransactionIntent)
	//	if txm["src"] == nil && txm["from"] == nil {
	//		return nil, errNilSrc
	//	}
	//	if txm["src"] != nil {
	//		txi.Fr, err = base64.StdEncoding.DecodeString(txm["src"].(string))
	//		if err != nil {
	//			return nil, err
	//		}
	//	} else {
	//		txi.Fr, err = base64.StdEncoding.DecodeString(txm["from"].(string))
	//		if err != nil {
	//			return nil, err
	//		}
	//	}
	//	if txm["dst"] != nil {
	//		txi.To, err = base64.StdEncoding.DecodeString(txm["dst"].(string))
	//		if err != nil {
	//			return nil, err
	//		}
	//	} else if txm["from"] != nil {
	//		txi.To, err = base64.StdEncoding.DecodeString(txm["from"].(string))
	//		if err != nil {
	//			return nil, err
	//		}
	//	}
	//	if txm["meta"] != nil {
	//		txi.Meta, err = base64.StdEncoding.DecodeString(txm["meta"].(string))
	//		if err != nil {
	//			return nil, err
	//		}
	//	}
	//	txi.Seq = nsbmath.NewUint256FromBigInt(big.NewInt(int64(idx)))
	//	if txm["amt"] != nil {
	//		b, _ := hex.DecodeString(txm["amt"].(string))
	//		txi.Amt = nsbmath.NewUint256FromBytes(b)
	//	} else {
	//		txi.Amt = nsbmath.NewUint256FromBytes([]byte{0})
	//	}
	//	transactionIntents = append(transactionIntents, txi)
	//	// fmt.Println("encoding", txm)
	//}

	var args ISC.ArgsCreateNewContract
	args.IscOwners = iscOwners
	args.Funds = funds
	args.TransactionIntents = bytesTransactionIntents
	args.VesSig = vesSig
	b, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	var fap = new(nsbrpc.FAPair)
	fap.FuncName = "isc"
	fap.Args = b
	// fmt.Println(PretiStruct(args), b)
	return fap, err
}
