package hostfunction

import (
	"errors"
	"wescale-wasm-plugin-template/tools"
)

var HostInstancePtr uint64
var HostModulePtr uint64

func GetGlobalValueByKey(key string) ([]byte, error) {
	if len(key) == 0 {
		return nil, errors.New("key is empty")
	}
	keyPtr, keySize := tools.StringToPtr(key)
	var ptr uint32
	var retSize uint32

	err := tools.StatusToError(GetGlobalValueByKeyOnHost(keyPtr, keySize, &ptr, &retSize))
	if err != nil {
		return nil, err
	}
	return tools.PtrToBytes(ptr, retSize), nil
}

func SetGlobalValueByKey(key string, value []byte) error {
	if len(key) == 0 {
		return errors.New("key is empty")
	}
	if len(value) == 0 {
		return errors.New("value is empty")
	}
	keyPtr, keySize := tools.StringToPtr(key)
	bytesPtr, bytesLen := tools.BytesToPtr(value)
	return tools.StatusToError(SetGlobalValueByKeyOnHost(keyPtr, keySize, bytesPtr, bytesLen))
}

func GetModuleValueByKey(key string) ([]byte, error) {
	if len(key) == 0 {
		return nil, errors.New("key is empty")
	}
	keyPtr, keySize := tools.StringToPtr(key)
	var ptr uint32
	var retSize uint32

	err := tools.StatusToError(GetModuleValueByKeyOnHost(HostModulePtr, keyPtr, keySize, &ptr, &retSize))
	if err != nil {
		return nil, err
	}
	return tools.PtrToBytes(ptr, retSize), nil
}

func SetModuleValueByKey(key string, value []byte) error {
	if len(key) == 0 {
		return errors.New("key is empty")
	}
	if len(value) == 0 {
		return errors.New("value is empty")
	}
	keyPtr, keySize := tools.StringToPtr(key)
	bytesPtr, bytesLen := tools.BytesToPtr(value)
	return tools.StatusToError(SetModuleValueByKeyOnHost(HostModulePtr, keyPtr, keySize, bytesPtr, bytesLen))
}

func GetHostQuery() (string, error) {
	var ptr uint32
	var retSize uint32

	err := tools.StatusToError(GetQueryOnHost(HostInstancePtr, &ptr, &retSize))
	if err != nil {
		return "", err
	}
	return tools.PtrToString(ptr, retSize), nil
}

func SetHostQuery(query string) error {
	if len(query) == 0 {
		return errors.New("query is empty")
	}
	ptr, size := tools.StringToPtr(query)
	return tools.StatusToError(SetQueryOnHost(HostInstancePtr, ptr, size))
}

func GlobalLock() {
	GlobalLockOnHost()
}

func GlobalUnlock() {
	GlobalUnlockOnHost()
}

func ModuleLock() {
	ModuleLockOnHost(HostModulePtr)
}

func ModuleUnlock() {
	ModuleUnlockOnHost(HostModulePtr)
}

func GetAbiVersion() (string, error) {
	var ptr uint32
	var retSize uint32

	err := tools.StatusToError(GetAbiVersionOnHost(&ptr, &retSize))
	if err != nil {
		return "", err
	}
	return tools.PtrToString(ptr, retSize), nil
}

func GetRuntimeType() (string, error) {
	var ptr uint32
	var retSize uint32

	err := tools.StatusToError(GetRuntimeTypeOnHost(&ptr, &retSize))
	if err != nil {
		return "", err
	}
	return tools.PtrToString(ptr, retSize), nil
}

func InfoLog(message string) {
	if len(message) == 0 {
		return
	}
	ptr, size := tools.StringToPtr(message)
	InfoLogOnHost(ptr, size)
}

func ErrorLog(message string) {
	if len(message) == 0 {
		return
	}
	ptr, size := tools.StringToPtr(message)
	ErrorLogOnHost(ptr, size)
}
