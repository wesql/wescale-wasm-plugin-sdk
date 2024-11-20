package host_functions

import (
	"errors"
	"github.com/wesql/sqlparser/go/vt/proto/query"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg/types"
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
	keyPtr, keySize := types.StringToPtr(key)
	var ptr uint32
	var retSize uint32

	err := types.StatusToError(getValueByKeyOnHost(uint32(scope), types.CurrentWasmPluginContext.HostModulePtr, keyPtr, keySize, &ptr, &retSize))
	if err != nil {
		return nil, err
	}
	return types.PtrToBytes(ptr, retSize), nil
}

func SetValueByKey(scope SharedScope, key string, value []byte) error {
	if len(key) == 0 {
		return errors.New("key is empty")
	}
	if len(value) == 0 {
		return errors.New("value is empty")
	}
	keyPtr, keySize := types.StringToPtr(key)
	bytesPtr, bytesLen := types.BytesToPtr(value)
	return types.StatusToError(setValueByKeyOnHost(uint32(scope), types.CurrentWasmPluginContext.HostModulePtr, keyPtr, keySize, bytesPtr, bytesLen))
}

func Lock(scope SharedScope) {
	lockOnHost(uint32(scope), types.CurrentWasmPluginContext.HostModulePtr)
}

func Unlock(scope SharedScope) {
	unlockOnHost(uint32(scope), types.CurrentWasmPluginContext.HostModulePtr)
}

func GetHostQuery() (string, error) {
	var ptr uint32
	var retSize uint32

	err := types.StatusToError(getQueryOnHost(types.CurrentWasmPluginContext.HostInstancePtr, &ptr, &retSize))
	if err != nil {
		return "", err
	}
	return types.PtrToString(ptr, retSize), nil
}

func SetHostQuery(query string) error {
	if len(query) == 0 {
		return errors.New("query is empty")
	}
	ptr, size := types.StringToPtr(query)
	return types.StatusToError(setQueryOnHost(types.CurrentWasmPluginContext.HostInstancePtr, ptr, size))
}

func GetAbiVersion() (string, error) {
	var ptr uint32
	var retSize uint32

	err := types.StatusToError(getAbiVersionOnHost(&ptr, &retSize))
	if err != nil {
		return "", err
	}
	return types.PtrToString(ptr, retSize), nil
}

func GetRuntimeType() (string, error) {
	var ptr uint32
	var retSize uint32

	err := types.StatusToError(getRuntimeTypeOnHost(&ptr, &retSize))
	if err != nil {
		return "", err
	}
	return types.PtrToString(ptr, retSize), nil
}

func InfoLog(message string) {
	if len(message) == 0 {
		return
	}
	ptr, size := types.StringToPtr(message)
	infoLogOnHost(ptr, size)
}

func ErrorLog(message string) {
	if len(message) == 0 {
		return
	}
	ptr, size := types.StringToPtr(message)
	errorLogOnHost(ptr, size)
}

func SetErrorMessage(errMessage string) {
	if len(errMessage) == 0 {
		return
	}
	ptr, size := types.StringToPtr(errMessage)
	setErrorMessageOnHost(types.CurrentWasmPluginContext.HostInstancePtr, ptr, size)
}

func GetErrorMessage() (string, error) {
	var ptr uint32
	var retSize uint32

	err := types.StatusToError(getErrorMessageOnHost(types.CurrentWasmPluginContext.HostInstancePtr, &ptr, &retSize))
	if err != nil {
		return "", err
	}
	return types.PtrToString(ptr, retSize), nil
}

func GetQueryResult() (*query.QueryResult, error) {
	var ptr uint32
	var retSize uint32

	err := types.StatusToError(getQueryResultOnHost(types.CurrentWasmPluginContext.HostInstancePtr, &ptr, &retSize))
	if err != nil {
		return nil, err
	}
	bytes := types.PtrToBytes(ptr, retSize)
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
	ptr, size := types.BytesToPtr(bytes)
	return types.StatusToError(setQueryResultOnHost(types.CurrentWasmPluginContext.HostInstancePtr, ptr, size))
}
