package hostfunction

import (
	"errors"
	"wescale-wasm-plugin-template/common"
)

var HostInstancePtr uint64
var HostModulePtr uint64

func GetGlobalValueByKey(key string) ([]byte, error) {
	if len(key) == 0 {
		return nil, errors.New("key is empty")
	}
	keyPtr, keySize := common.StringToPtr(key)
	var ptr uint32
	var retSize uint32

	err := common.StatusToError(getGlobalValueByKeyOnHost(keyPtr, keySize, &ptr, &retSize))
	if err != nil {
		return nil, err
	}
	return common.PtrToBytes(ptr, retSize), nil
}

func SetGlobalValueByKey(key string, value []byte) error {
	if len(key) == 0 {
		return errors.New("key is empty")
	}
	if len(value) == 0 {
		return errors.New("value is empty")
	}
	keyPtr, keySize := common.StringToPtr(key)
	bytesPtr, bytesLen := common.BytesToPtr(value)
	return common.StatusToError(setGlobalValueByKeyOnHost(keyPtr, keySize, bytesPtr, bytesLen))
}

func GetModuleValueByKey(key string) ([]byte, error) {
	if len(key) == 0 {
		return nil, errors.New("key is empty")
	}
	keyPtr, keySize := common.StringToPtr(key)
	var ptr uint32
	var retSize uint32

	err := common.StatusToError(getModuleValueByKeyOnHost(HostModulePtr, keyPtr, keySize, &ptr, &retSize))
	if err != nil {
		return nil, err
	}
	return common.PtrToBytes(ptr, retSize), nil
}

func SetModuleValueByKey(key string, value []byte) error {
	if len(key) == 0 {
		return errors.New("key is empty")
	}
	if len(value) == 0 {
		return errors.New("value is empty")
	}
	keyPtr, keySize := common.StringToPtr(key)
	bytesPtr, bytesLen := common.BytesToPtr(value)
	return common.StatusToError(setModuleValueByKeyOnHost(HostModulePtr, keyPtr, keySize, bytesPtr, bytesLen))
}

func GetHostQuery() (string, error) {
	var ptr uint32
	var retSize uint32

	err := common.StatusToError(getQueryOnHost(HostInstancePtr, &ptr, &retSize))
	if err != nil {
		return "", err
	}
	return common.PtrToString(ptr, retSize), nil
}

func SetHostQuery(query string) error {
	if len(query) == 0 {
		return errors.New("query is empty")
	}
	ptr, size := common.StringToPtr(query)
	return common.StatusToError(setQueryOnHost(HostInstancePtr, ptr, size))
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

	err := common.StatusToError(getAbiVersionOnHost(&ptr, &retSize))
	if err != nil {
		return "", err
	}
	return common.PtrToString(ptr, retSize), nil
}

func GetRuntimeType() (string, error) {
	var ptr uint32
	var retSize uint32

	err := common.StatusToError(getRuntimeTypeOnHost(&ptr, &retSize))
	if err != nil {
		return "", err
	}
	return common.PtrToString(ptr, retSize), nil
}

func InfoLog(message string) {
	if len(message) == 0 {
		return
	}
	ptr, size := common.StringToPtr(message)
	infoLogOnHost(ptr, size)
}

func ErrorLog(message string) {
	if len(message) == 0 {
		return
	}
	ptr, size := common.StringToPtr(message)
	errorLogOnHost(ptr, size)
}

func SetErrorMessage(errMessage string) {
	if len(errMessage) == 0 {
		return
	}
	ptr, size := common.StringToPtr(errMessage)
	setErrorMessageOnHost(HostInstancePtr, ptr, size)
}
