package request

import (
	"net/http"
	"time"

	"github.com/imroc/req"
)

var (
	requestTimeout = 50 * time.Second
)

type Client struct {
	BaseURL string
	Header  Header
}

func NewRequestClient(url string) *Client {
	return &Client{BaseURL: url}
}

func (jc *Client) SetHeader(strMap map[string]string) *Client {
	jc.Header = strMap
	return jc
}

func (jc *Client) SetHeaderWithReqHeader(header Header) *Client {
	jc.Header = header
	return jc
}

func (jc *Client) Group(sub string) *Client {
	return &Client{BaseURL: jc.BaseURL + sub, Header: jc.Header}
}

type ClientX struct {
	BaseURL string
	Header  Header
	path    string
}

func NewRequestClientX(url string) *ClientX {
	return &ClientX{BaseURL: url}
}

func (jc *ClientX) SetHeader(i interface{}) *ClientX {
	switch s := i.(type) {
	case map[string]string:
		jc.Header = s
	case Header:
		jc.Header = s
	default:
	}
	return jc
}

func (jc *ClientX) Group(sub string) *ClientX {
	return &ClientX{BaseURL: jc.BaseURL + sub, Header: jc.Header}
}

func (jc *ClientX) Path(path string) *ClientX {
	jc.path = path
	return jc
}

func (jc *ClientX) Use(handler func(*Resp) error) *Context {
	return &Context{BaseURL: jc.BaseURL + jc.path, Header: jc.Header, handler: handler}
}

type Context struct {
	BaseURL string
	Header  Header
	handler func(*Resp) error
}

func (jc *Context) Path(path string) *Context {
	jc.BaseURL += path
	return jc
}

func SetConnPool() {
	client := &http.Client{}
	client.Transport = &http.Transport{
		MaxIdleConnsPerHost: 500,
	}

	req.SetClient(client)
	req.SetTimeout(requestTimeout)
}

func init() {
	SetConnPool()
}
