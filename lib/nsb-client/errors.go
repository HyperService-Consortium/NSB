package nsbcli

import "errors"

var (
	errNotJSON2d0 = errors.New("reject ret that is not jsonrpc: 2.0")
	errNilSrc     = errors.New("nil src")
	errBadJSON    = errors.New("bad format of jsonrpc")
)
