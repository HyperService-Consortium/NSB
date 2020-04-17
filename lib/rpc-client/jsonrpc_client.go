package jsonrpcclient

import (
	"errors"
	"io"
	"strings"

	gjson "github.com/tidwall/gjson"

	bytespool "github.com/HyperService-Consortium/NSB/lib/bytes-pool"
	request "github.com/HyperService-Consortium/NSB/lib/request"
)

func decorateHost(host string) string {
	if strings.HasPrefix(host, httpPrefix) || strings.HasPrefix(host, httpsPrefix) {
		return host
	}
	return "http://" + host
}

type JsonRpcClient struct {
	handler    *request.Client
	bufferPool *bytespool.BytesPool
}

func NewJsonRpcClient(host string) *JsonRpcClient {
	return &JsonRpcClient{
		handler:    request.NewRequestClient(decorateHost(host)).SetHeader(fixedJsonHeader),
		bufferPool: bytespool.NewBytesPool(maxBytesSize),
	}
}

// todo: test invalid json
func (nc *JsonRpcClient) preloadJsonResponse(bb io.ReadCloser) ([]byte, error) {

	var b = nc.bufferPool.Get()
	defer nc.bufferPool.Put(b)

	_, err := bb.Read(b)
	if err != nil && err != io.EOF {
		return nil, err
	}
	bb.Close()
	var jm = gjson.ParseBytes(b)
	if s := jm.Get("jsonrpc"); !s.Exists() || s.String() != "2.0" {
		return nil, errors.New("reject ret that is not jsonrpc: 2.0")
	}
	if s := jm.Get("error"); s.Exists() {
		return nil, FromGJsonResultError(s)
	}
	if s := jm.Get("result"); s.Exists() {
		if s.Index > 0 {
			return []byte(s.Raw), nil
		}
	}
	return nil, errors.New("bad format of jsonrpc")
}

func (nc *JsonRpcClient) GetRequestParams(params map[string]interface{}) ([]byte, error) {
	b, err := nc.handler.GetWithParams(params)
	if err != nil {
		return nil, err
	}
	return nc.preloadJsonResponse(b)
}

func (nc *JsonRpcClient) PostRequestWithJsonObj(jsonObj interface{}) ([]byte, error) {
	b, err := nc.handler.PostWithJsonObj(jsonObj)
	if err != nil {
		return nil, err
	}
	return nc.preloadJsonResponse(b)
}

func (nc *JsonRpcClient) PostWithBody(jsonBody []byte) ([]byte, error) {
	b, err := nc.handler.PostWithBody(jsonBody)
	if err != nil {
		return nil, err
	}
	return staticJsonRpcClient.preloadJsonResponse(b)
}

func Do(url string, jsonBody []byte) ([]byte, error) {
	b, err := request.PostWithBody(url, jsonBody)
	if err != nil {
		return nil, err
	}
	return staticJsonRpcClient.preloadJsonResponse(b)
}
