package main

const (
	CodeOK int = iota
	CodeInternalError
	CodeIOError
	CodeLogicError
	CodeConflictError
)

type errObj struct {
	err error
}

func (errobj *errObj) Error() string {
	return errobj.err.Error()
}

type InternalErrorObj struct {
	errObj
}

func (errobj *InternalErrorObj) ExitCode() int {
	return CodeInternalError
}

func InternalError(err error) *InternalErrorObj {
	return &InternalErrorObj{errObj: errObj{err}}
}

type IOErrorObj struct {
	errObj
}

func (errobj *IOErrorObj) ExitCode() int {
	return CodeIOError
}

func IOError(err error) *IOErrorObj {
	return &IOErrorObj{errObj: errObj{err}}
}

type LogicErrorObj struct {
	errObj
}

func (errobj *LogicErrorObj) ExitCode() int {
	return CodeLogicError
}

func LogicError(err error) *LogicErrorObj {
	return &LogicErrorObj{errObj: errObj{err}}
}

type ConflictErrorObj struct {
	errObj
}

func (errobj *ConflictErrorObj) ExitCode() int {
	return CodeConflictError
}

func ConflictError(err error) *ConflictErrorObj {
	return &ConflictErrorObj{errObj: errObj{err}}
}
