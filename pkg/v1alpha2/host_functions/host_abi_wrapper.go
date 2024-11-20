package host_functions

import (
	"errors"
	"github.com/wesql/sqlparser/go/vt/proto/query"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg"
)

type SharedScope uint32

const (
	SharedScope_MODULE SharedScope = 0
	SharedScope_TABLET SharedScope = 1
)

func GetValueByKey(scope SharedScope, key string) ([]byte, error) {
	if len(key) == 0 {
		return nil, errors.New("key is empty")
	}
	keyPtr, keySize := pkg.StringToPtr(key)
	var ptr uint32
	var retSize uint32

	err := pkg.StatusToError(getValueByKeyOnHost(uint32(scope), pkg.CurrentWasmPluginContext.HostModulePtr, keyPtr, keySize, &ptr, &retSize))
	if err != nil {
		return nil, err
	}
	return pkg.PtrToBytes(ptr, retSize), nil
}

func SetValueByKey(scope SharedScope, key string, value []byte) error {
	if len(key) == 0 {
		return errors.New("key is empty")
	}
	if len(value) == 0 {
		return errors.New("value is empty")
	}
	keyPtr, keySize := pkg.StringToPtr(key)
	bytesPtr, bytesLen := pkg.BytesToPtr(value)
	return pkg.StatusToError(setValueByKeyOnHost(uint32(scope), pkg.CurrentWasmPluginContext.HostModulePtr, keyPtr, keySize, bytesPtr, bytesLen))
}

func Lock(scope SharedScope) {
	lockOnHost(uint32(scope), pkg.CurrentWasmPluginContext.HostModulePtr)
}

func Unlock(scope SharedScope) {
	unlockOnHost(uint32(scope), pkg.CurrentWasmPluginContext.HostModulePtr)
}

func GetHostQuery() (string, error) {
	var ptr uint32
	var retSize uint32

	err := pkg.StatusToError(getQueryOnHost(pkg.CurrentWasmPluginContext.HostInstancePtr, &ptr, &retSize))
	if err != nil {
		return "", err
	}
	return pkg.PtrToString(ptr, retSize), nil
}

func SetHostQuery(query string) error {
	if len(query) == 0 {
		return errors.New("query is empty")
	}
	ptr, size := pkg.StringToPtr(query)
	return pkg.StatusToError(setQueryOnHost(pkg.CurrentWasmPluginContext.HostInstancePtr, ptr, size))
}

func GetAbiVersion() (string, error) {
	var ptr uint32
	var retSize uint32

	err := pkg.StatusToError(getAbiVersionOnHost(&ptr, &retSize))
	if err != nil {
		return "", err
	}
	return pkg.PtrToString(ptr, retSize), nil
}

func GetRuntimeType() (string, error) {
	var ptr uint32
	var retSize uint32

	err := pkg.StatusToError(getRuntimeTypeOnHost(&ptr, &retSize))
	if err != nil {
		return "", err
	}
	return pkg.PtrToString(ptr, retSize), nil
}

func InfoLog(message string) {
	if len(message) == 0 {
		return
	}
	ptr, size := pkg.StringToPtr(message)
	infoLogOnHost(ptr, size)
}

func ErrorLog(message string) {
	if len(message) == 0 {
		return
	}
	ptr, size := pkg.StringToPtr(message)
	errorLogOnHost(ptr, size)
}

func SetErrorMessage(errMessage string) {
	if len(errMessage) == 0 {
		return
	}
	ptr, size := pkg.StringToPtr(errMessage)
	setErrorMessageOnHost(pkg.CurrentWasmPluginContext.HostInstancePtr, ptr, size)
}

func GetErrorMessage() (string, error) {
	var ptr uint32
	var retSize uint32

	err := pkg.StatusToError(getErrorMessageOnHost(pkg.CurrentWasmPluginContext.HostInstancePtr, &ptr, &retSize))
	if err != nil {
		return "", err
	}
	return pkg.PtrToString(ptr, retSize), nil
}

func GetQueryResult() (*query.QueryResult, error) {
	var ptr uint32
	var retSize uint32

	err := pkg.StatusToError(getQueryResultOnHost(pkg.CurrentWasmPluginContext.HostInstancePtr, &ptr, &retSize))
	if err != nil {
		return nil, err
	}
	bytes := pkg.PtrToBytes(ptr, retSize)
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
	ptr, size := pkg.BytesToPtr(bytes)
	return pkg.StatusToError(setQueryResultOnHost(pkg.CurrentWasmPluginContext.HostInstancePtr, ptr, size))
}
