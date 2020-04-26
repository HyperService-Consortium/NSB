package request

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"

	"github.com/google/go-querystring/query"
	"github.com/imroc/req"
)

type QueryParam = req.QueryParam
type Param = req.Param
type Resp = req.Resp
type Header = req.Header

func get(s string, v ...interface{}) (io.ReadCloser, error) {
	//fmt.Println("url", s, "params", v)
	resp, err := req.Get(s, v...)
	if err != nil {
		return nil, err
	}
	var sc = resp.Response().StatusCode
	if sc != 200 {
		return nil, fmt.Errorf("error code: %v", sc)
	}
	return resp.Response().Body, nil
}

func getx(s string, v ...interface{}) (*Resp, error) {
	resp, err := req.Get(s, v...)
	if err != nil {
		return nil, err
	}
	var sc = resp.Response().StatusCode
	if sc != 200 {
		return nil, fmt.Errorf("error code: %v", sc)
	}
	return resp, err
}

func getc(s string, handler func(*Resp) error, v ...interface{}) error {
	resp, err := req.Get(s, v...)
	if err != nil {
		return err
	}
	var sc = resp.Response().StatusCode
	if sc != 200 {
		return fmt.Errorf("error code: %v", sc)
	}

	err = handler(resp)
	return err
}

func (jc *Client) Get() (io.ReadCloser, error) {
	return get(jc.BaseURL, jc.Header)
}

func (jc *Client) GetWithKVMap(request map[string]interface{}) (io.ReadCloser, error) {
	return get(jc.BaseURL, jc.Header, request)
}

func (jc *Client) GetWithParams(params ...interface{}) (io.ReadCloser, error) {
	return get(jc.BaseURL, append(params, jc.Header)...)
}

func (jc *Client) GetWithStruct(request interface{}) (io.ReadCloser, error) {
	v, err := query.Values(request)
	if err != nil {
		return nil, err
	}
	s := bytes.NewBufferString(jc.BaseURL)
	err = s.WriteByte('?')
	if err != nil {
		return nil, err
	}
	_, err = s.WriteString(v.Encode())
	if err != nil {
		return nil, err
	}
	return get(s.String(), jc.Header)
}

func (jc *ClientX) Get(params ...interface{}) (io.ReadCloser, error) {
	fin, r := false, jc.BaseURL
	for idx, param := range params {
		if reflect.TypeOf(param).Kind() == reflect.Struct {
			v, err := query.Values(param)
			if err != nil {
				return nil, err
			}
			s := bytes.NewBufferString(r)
			err = s.WriteByte('?')
			if err != nil {
				return nil, err
			}
			_, err = s.WriteString(v.Encode())
			if err != nil {
				return nil, err
			}
			fin = true
			r = s.String()
			params = append(params[:idx], params[idx+1:]...)
			continue
		}
		if reflect.TypeOf(param).Kind() == reflect.Ptr && reflect.ValueOf(param).Elem().Kind() == reflect.Struct {
			v, err := query.Values(param)
			if err != nil {
				return nil, err
			}
			s := bytes.NewBufferString(r)
			err = s.WriteByte('?')
			if err != nil {
				return nil, err
			}
			_, err = s.WriteString(v.Encode())
			if err != nil {
				return nil, err
			}
			fin = true
			r = s.String()
			params = append(params[:idx], params[idx+1:]...)
			continue
		}

		switch i := param.(type) {
		case *req.QueryParam:
			params[idx] = *i
			if fin {
				return nil, errors.New("struct mapping and map[string]interface{} cannot be used at the same time")
			}
		case *req.Param:
			params[idx] = *i
			if fin {
				return nil, errors.New("struct mapping and map[string]interface{} cannot be used at the same time")
			}
		case req.Param, req.QueryParam:
			if fin {
				return nil, errors.New("struct mapping and map[string]interface{} cannot be used at the same time")
			}
		default:
		}
	}
	params = append(params, jc.Header)
	return get(r, params...)
}

func (jc *Context) Get(params ...interface{}) error {
	fin, r := false, jc.BaseURL
	for idx, param := range params {
		if reflect.TypeOf(param).Kind() == reflect.Struct {
			v, err := query.Values(param)
			if err != nil {
				return err
			}
			s := bytes.NewBufferString(r)
			err = s.WriteByte('?')
			if err != nil {
				return err
			}
			_, err = s.WriteString(v.Encode())
			if err != nil {
				return err
			}
			fin = true
			r = s.String()
			params = append(params[:idx], params[idx+1:]...)
			continue
		}
		if reflect.TypeOf(param).Kind() == reflect.Ptr && reflect.ValueOf(param).Elem().Kind() == reflect.Struct {
			v, err := query.Values(param)
			if err != nil {
				return err
			}
			s := bytes.NewBufferString(r)
			err = s.WriteByte('?')
			if err != nil {
				return err
			}
			_, err = s.WriteString(v.Encode())
			if err != nil {
				return err
			}
			fin = true
			r = s.String()
			params = append(params[:idx], params[idx+1:]...)
			continue
		}

		switch i := param.(type) {
		case *req.QueryParam:
			params[idx] = *i
			if fin {
				return errors.New("struct mapping and map[string]interface{} cannot be used at the same time")
			}
		case *req.Param:
			params[idx] = *i
			if fin {
				return errors.New("struct mapping and map[string]interface{} cannot be used at the same time")
			}
		case req.Param, req.QueryParam:
			if fin {
				return errors.New("struct mapping and map[string]interface{} cannot be used at the same time")
			}
		default:
		}
	}
	params = append(params, jc.Header)
	return getc(r, jc.handler, params...)
}
