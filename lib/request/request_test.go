package request

import (
	"io"
	"testing"

	mtest "github.com/Myriad-Dreamin/mydrest"
)

var s mtest.TestHelper

func TestGet(t *testing.T) {
	_, err := NewRequestClient("http://www.baidu.com").Get()
	s.AssertNoErr(t, err)
}

type mParam4 struct {
}

type service struct {
	b3 []byte
}

func (s *service) myController(r *Resp) (err error) {
	s.b3, err = r.ToBytes()
	return
}

func TestGetParamAndGroup(t *testing.T) {
	var params = &mParam4{}
	NSBApi := NewRequestClient("http://121.89.200.234:26657")
	var b, b2 = make([]byte, 1024*1024), make([]byte, 1024*1024)
	bb, err := NSBApi.Group("/net_info").GetWithStruct(params)
	s.AssertNoErr(t, err)
	_, err = bb.Read(b)
	s.AssertEqual(t, err, io.EOF)
	bb.Close()

	abciInfoApi := NSBApi.Group("/net_info")

	bb2, err := abciInfoApi.GetWithStruct(params)
	s.AssertNoErr(t, err)
	_, err = bb2.Read(b2)
	s.AssertEqual(t, err, io.EOF)
	bb2.Close()

	s.AssertEqual(t, string(b), string(b2))

	NSBApiX := NewRequestClientX("http://121.89.200.234:26657")
	bb2, err = NSBApiX.Group("/net_info").Get(params)
	s.AssertNoErr(t, err)
	_, err = bb2.Read(b2)
	s.AssertEqual(t, err, io.EOF)
	bb2.Close()

	s.AssertEqual(t, string(b), string(b2))

	bb2, err = NSBApiX.Group("/net_info").Get(&QueryParam{})
	s.AssertNoErr(t, err)
	_, err = bb2.Read(b2)
	s.AssertEqual(t, err, io.EOF)
	bb2.Close()

	s.AssertEqual(t, string(b), string(b2))
	// req.Debug = true
	yandeApi := NewRequestClientX("https://yande.re")

	_, err = yandeApi.Group("/post").Get(&QueryParam{
		"tags": "dress",
	})
	s.AssertNoErr(t, err)
	serve := new(service)
	err = yandeApi.Group("/post").Use(serve.myController).Get(&Param{
		"tags": "dress",
	})
	s.AssertNoErr(t, err)
	// s.AssertEqual(t, string(b), string(b2))
}
