package hostfunction

import (
	"wescale-wasm-plugin-template/tools"
)

var HostInstancePtr uint64
var HostModulePtr uint64

func GetGlobalValueByKey(key string) ([]byte, error) {
	keyPtr, keySize := tools.StringToLeakedPtr(key)
	var ptr uint32
	var retSize uint32

	err := tools.StatusToError(GetGlobalValueByKeyHost(keyPtr, keySize, &ptr, &retSize))
	if err != nil {
		return nil, err
	}
	return tools.PtrToBytes(ptr, retSize), nil
}

func SetGlobalValueByKey(key string, value []byte) error {
	keyPtr, keySize := tools.StringToLeakedPtr(key)
	bytesPtr, bytesLen := tools.BytesToLeakedPtr(value)
	return tools.StatusToError(SetGlobalValueByKeyHost(keyPtr, keySize, bytesPtr, bytesLen))
}

func GetModuleValueByKey(key string) ([]byte, error) {
	keyPtr, keySize := tools.StringToLeakedPtr(key)
	var ptr uint32
	var retSize uint32

	err := tools.StatusToError(GetModuleValueByKeyHost(HostModulePtr, keyPtr, keySize, &ptr, &retSize))
	if err != nil {
		return nil, err
	}
	return tools.PtrToBytes(ptr, retSize), nil
}

func SetModuleValueByKey(key string, value []byte) error {
	keyPtr, keySize := tools.StringToLeakedPtr(key)
	bytesPtr, bytesLen := tools.BytesToLeakedPtr(value)
	return tools.StatusToError(SetModuleValueByKeyHost(HostModulePtr, keyPtr, keySize, bytesPtr, bytesLen))
}

func GetHostQuery() (string, error) {
	var ptr uint32
	var retSize uint32

	err := tools.StatusToError(GetQueryHost(HostInstancePtr, &ptr, &retSize))
	if err != nil {
		return "", err
	}
	return tools.PtrToString(ptr, retSize), nil
}

func SetHostQuery(query string) error {
	ptr, size := tools.StringToLeakedPtr(query)
	return tools.StatusToError(SetQueryHost(HostInstancePtr, ptr, size))
}

func GlobalLock() {
	GlobalLockHost()
}

func GlobalUnlock() {
	GlobalUnlockHost()
}

func ModuleLock() {
	ModuleLockHost(HostModulePtr)
}

func ModuleUnlock() {
	ModuleUnlockHost(HostModulePtr)
}
