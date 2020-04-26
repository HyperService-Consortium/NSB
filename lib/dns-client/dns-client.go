package dns_client

import (
	"encoding/json"
	"fmt"
	"github.com/HyperService-Consortium/NSB/lib/request"
	"github.com/HyperService-Consortium/go-uip/uip"
	"io/ioutil"
)

type DNSClient struct {
	handler    *request.Client
	Host string `json:"host"`
}

func NewDNSClient(host string) *DNSClient {
	return &DNSClient{Host:host, handler: request.NewRequestClient(host)}
}

type HealthReply struct {
	Code int `json:"code"`
}

func (dns *DNSClient) Health() (*HealthReply, error) {
	reader, err := dns.handler.Group("/health").Get()
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var reply HealthReply
	err = json.Unmarshal(b, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}


type ChainInfoReply struct {
	ChainType uip.ChainType `json:"chain_type"`
	Domain string `json:"domain"`
	MerkleProofType uip.MerkleProofType `json:"merkle_proof_type"`
}

type ChainInfoResponse struct {
	Code int `json:"code"`
	Data ChainInfoReply `json:"data"`
}

func (dns *DNSClient) GetChainInfo(chainID uip.ChainIDUnderlyingType) (
	*ChainInfoReply, error) {
	reader, err := dns.handler.Group("/chain-info").GetWithParams(
		request.QueryParam{
			"chain_id": chainID,
		})
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var reply ChainInfoResponse
	err = json.Unmarshal(b, &reply)
	if err != nil {
		return nil, err
	}
	if reply.Code != 0 {
		return nil, fmt.Errorf("code not zero: %v", reply.Code)
	}

	return &reply.Data, nil

}

