package hostfunction

import (
	"errors"
	"github.com/wesql/wescale-wasm-plugin-template/internal"
	"github.com/wesql/wescale-wasm-plugin-template/internal/proto/query"
)

var HostInstancePtr uint64
var HostModulePtr uint64

func GetGlobalValueByKey(key string) ([]byte, error) {
	if len(key) == 0 {
		return nil, errors.New("key is empty")
	}
	keyPtr, keySize := internal.StringToPtr(key)
	var ptr uint32
	var retSize uint32

	err := internal.StatusToError(getGlobalValueByKeyOnHost(keyPtr, keySize, &ptr, &retSize))
	if err != nil {
		return nil, err
	}
	return internal.PtrToBytes(ptr, retSize), nil
}

func SetGlobalValueByKey(key string, value []byte) error {
	if len(key) == 0 {
		return errors.New("key is empty")
	}
	if len(value) == 0 {
		return errors.New("value is empty")
	}
	keyPtr, keySize := internal.StringToPtr(key)
	bytesPtr, bytesLen := internal.BytesToPtr(value)
	return internal.StatusToError(setGlobalValueByKeyOnHost(keyPtr, keySize, bytesPtr, bytesLen))
}

func GetModuleValueByKey(key string) ([]byte, error) {
	if len(key) == 0 {
		return nil, errors.New("key is empty")
	}
	keyPtr, keySize := internal.StringToPtr(key)
	var ptr uint32
	var retSize uint32

	err := internal.StatusToError(getModuleValueByKeyOnHost(HostModulePtr, keyPtr, keySize, &ptr, &retSize))
	if err != nil {
		return nil, err
	}
	return internal.PtrToBytes(ptr, retSize), nil
}

func SetModuleValueByKey(key string, value []byte) error {
	if len(key) == 0 {
		return errors.New("key is empty")
	}
	if len(value) == 0 {
		return errors.New("value is empty")
	}
	keyPtr, keySize := internal.StringToPtr(key)
	bytesPtr, bytesLen := internal.BytesToPtr(value)
	return internal.StatusToError(setModuleValueByKeyOnHost(HostModulePtr, keyPtr, keySize, bytesPtr, bytesLen))
}

func GetHostQuery() (string, error) {
	var ptr uint32
	var retSize uint32

	err := internal.StatusToError(getQueryOnHost(HostInstancePtr, &ptr, &retSize))
	if err != nil {
		return "", err
	}
	return internal.PtrToString(ptr, retSize), nil
}

func SetHostQuery(query string) error {
	if len(query) == 0 {
		return errors.New("query is empty")
	}
	ptr, size := internal.StringToPtr(query)
	return internal.StatusToError(setQueryOnHost(HostInstancePtr, ptr, size))
}

func GlobalLock() {
	globalLockOnHost()
}

func GlobalUnlock() {
	globalUnlockOnHost()
}

func ModuleLock() {
	moduleLockOnHost(HostModulePtr)
}

func ModuleUnlock() {
	moduleUnlockOnHost(HostModulePtr)
}

func GetAbiVersion() (string, error) {
	var ptr uint32
	var retSize uint32

	err := internal.StatusToError(getAbiVersionOnHost(&ptr, &retSize))
	if err != nil {
		return "", err
	}
	return internal.PtrToString(ptr, retSize), nil
}

func GetRuntimeType() (string, error) {
	var ptr uint32
	var retSize uint32

	err := internal.StatusToError(getRuntimeTypeOnHost(&ptr, &retSize))
	if err != nil {
		return "", err
	}
	return internal.PtrToString(ptr, retSize), nil
}

func InfoLog(message string) {
	if len(message) == 0 {
		return
	}
	ptr, size := internal.StringToPtr(message)
	infoLogOnHost(ptr, size)
}

func ErrorLog(message string) {
	if len(message) == 0 {
		return
	}
	ptr, size := internal.StringToPtr(message)
	errorLogOnHost(ptr, size)
}

func SetErrorMessage(errMessage string) {
	if len(errMessage) == 0 {
		return
	}
	ptr, size := internal.StringToPtr(errMessage)
	setErrorMessageOnHost(HostInstancePtr, ptr, size)
}

func GetErrorMessage() (string, error) {
	var ptr uint32
	var retSize uint32

	err := internal.StatusToError(getErrorMessageOnHost(HostInstancePtr, &ptr, &retSize))
	if err != nil {
		return "", err
	}
	return internal.PtrToString(ptr, retSize), nil
}

func GetQueryResult() (*query.QueryResult, error) {
	var ptr uint32
	var retSize uint32

	err := internal.StatusToError(getQueryResultOnHost(HostInstancePtr, &ptr, &retSize))
	if err != nil {
		return nil, err
	}
	bytes := internal.PtrToBytes(ptr, retSize)
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
	ptr, size := internal.BytesToPtr(bytes)
	return internal.StatusToError(setQueryResultOnHost(HostInstancePtr, ptr, size))
}
