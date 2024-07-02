package v1alpha1

import (
	"errors"
	"github.com/wesql/sqlparser/go/vt/proto/query"
)

var HostInstancePtr uint64
var HostModulePtr uint64

type SharedScope uint32

const (
	SharedScope_MODULE SharedScope = 0
	SharedScope_TABLET SharedScope = 1
)

func GetValueByKey(scope SharedScope, key string) ([]byte, error) {
	if len(key) == 0 {
		return nil, errors.New("key is empty")
	}
	keyPtr, keySize := StringToPtr(key)
	var ptr uint32
	var retSize uint32

	err := StatusToError(getValueByKeyOnHost(uint32(scope), HostModulePtr, keyPtr, keySize, &ptr, &retSize))
	if err != nil {
		return nil, err
	}
	return PtrToBytes(ptr, retSize), nil
}

func SetValueByKey(scope SharedScope, key string, value []byte) error {
	if len(key) == 0 {
		return errors.New("key is empty")
	}
	if len(value) == 0 {
		return errors.New("value is empty")
	}
	keyPtr, keySize := StringToPtr(key)
	bytesPtr, bytesLen := BytesToPtr(value)
	return StatusToError(setValueByKeyOnHost(uint32(scope), HostModulePtr, keyPtr, keySize, bytesPtr, bytesLen))
}

func Lock(scope SharedScope) {
	lockOnHost(uint32(scope), HostModulePtr)
}

func Unlock(scope SharedScope) {
	unlockOnHost(uint32(scope), HostModulePtr)
}

func GetHostQuery() (string, error) {
	var ptr uint32
	var retSize uint32

	err := StatusToError(getQueryOnHost(HostInstancePtr, &ptr, &retSize))
	if err != nil {
		return "", err
	}
	return PtrToString(ptr, retSize), nil
}

func SetHostQuery(query string) error {
	if len(query) == 0 {
		return errors.New("query is empty")
	}
	ptr, size := StringToPtr(query)
	return StatusToError(setQueryOnHost(HostInstancePtr, ptr, size))
}

func GetAbiVersion() (string, error) {
	var ptr uint32
	var retSize uint32

	err := StatusToError(getAbiVersionOnHost(&ptr, &retSize))
	if err != nil {
		return "", err
	}
	return PtrToString(ptr, retSize), nil
}

func GetRuntimeType() (string, error) {
	var ptr uint32
	var retSize uint32

	err := StatusToError(getRuntimeTypeOnHost(&ptr, &retSize))
	if err != nil {
		return "", err
	}
	return PtrToString(ptr, retSize), nil
}

func InfoLog(message string) {
	if len(message) == 0 {
		return
	}
	ptr, size := StringToPtr(message)
	infoLogOnHost(ptr, size)
}

func ErrorLog(message string) {
	if len(message) == 0 {
		return
	}
	ptr, size := StringToPtr(message)
	errorLogOnHost(ptr, size)
}

func SetErrorMessage(errMessage string) {
	if len(errMessage) == 0 {
		return
	}
	ptr, size := StringToPtr(errMessage)
	setErrorMessageOnHost(HostInstancePtr, ptr, size)
}

func GetErrorMessage() (string, error) {
	var ptr uint32
	var retSize uint32

	err := StatusToError(getErrorMessageOnHost(HostInstancePtr, &ptr, &retSize))
	if err != nil {
		return "", err
	}
	return PtrToString(ptr, retSize), nil
}

func GetQueryResult() (*query.QueryResult, error) {
	var ptr uint32
	var retSize uint32

	err := StatusToError(getQueryResultOnHost(HostInstancePtr, &ptr, &retSize))
	if err != nil {
		return nil, err
	}
	bytes := PtrToBytes(ptr, retSize)
	queryResult := &query.QueryResult{}
	err = queryResult.UnmarshalVT(bytes)
	if err != nil {
		return nil, err
	}
	return queryResult, nil
}

func SetQueryResult(queryResult *query.QueryResult) error {
	if queryResult == nil {
		return nil
	}
	bytes, err := queryResult.MarshalVT()
	if err != nil {
		return nil
	}
	ptr, size := BytesToPtr(bytes)
	return StatusToError(setQueryResultOnHost(HostInstancePtr, ptr, size))
}
