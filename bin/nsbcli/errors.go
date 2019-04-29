package main

const (
	CodeOK int = iota
	CodeDecodeError
)

type DecodeErrorObj struct {
	err error
}

func (errobj *DecodeErrorObj) Error() string {
	return errobj.err.Error()
}

func (errobj *DecodeErrorObj) ExitCode() int {
	return CodeDecodeError
}

func DecodeError(err error) *DecodeErrorObj {
	return &DecodeErrorObj{err}
}