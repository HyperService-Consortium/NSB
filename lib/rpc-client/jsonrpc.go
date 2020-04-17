package jsonrpcclient

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
)

type jsonMap = map[string]interface{}

type JsonError struct {
	errorx string
}

func (je JsonError) Error() string {
	return je.errorx
}

func FromJsonMapError(jm jsonMap) *JsonError {
	return &JsonError{
		errorx: fmt.Sprintf("jsonrpc error: %v(%v), %v", jm["message"], jm["code"], jm["data"]),
	}
}

func FromBytesError(b []byte) *JsonError {
	var jm jsonMap
	err := json.Unmarshal(b, &jm)
	if err != nil {
		return &JsonError{
			errorx: fmt.Sprintf("bad format of json error: %v", err),
		}
	}
	return FromJsonMapError(jm)
}

func FromGJsonResultError(b gjson.Result) *JsonError {
	return &JsonError{
		errorx: fmt.Sprintf("gjsonrpc error: %v(%v), %v", b.Get("message"), b.Get("code"), b.Get("data")),
	}
}
