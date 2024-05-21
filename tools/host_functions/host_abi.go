package hostfunction

import "wescale-wasm-plugin-template/tools"

//export GetGlobalValueByKeyOnHost
func GetGlobalValueByKeyOnHost(keyPtr, keySize uint32, returnValuePtr, returnValueSize *uint32) tools.Status

//export SetGlobalValueByKeyOnHost
func SetGlobalValueByKeyOnHost(keyPtr, keySize, valuePtr, valueSize uint32) tools.Status

//export GetModuleValueByKeyOnHost
func GetModuleValueByKeyOnHost(hostModulePtr uint64, keyPtr, keySize uint32, returnValuePtr, returnValueSize *uint32) tools.Status

//export SetModuleValueByKeyOnHost
func SetModuleValueByKeyOnHost(hostModulePtr uint64, keyPtr, keySize, valuePtr, valueSize uint32) tools.Status

//export GetQueryOnHost
func GetQueryOnHost(hostInstancePtr uint64, returnValueData *uint32, returnValueSize *uint32) tools.Status

//export SetQueryOnHost
func SetQueryOnHost(hostInstancePtr uint64, queryValuePtr uint32, queryValueSize uint32) tools.Status

//export GlobalLockOnHost
func GlobalLockOnHost()

//export GlobalUnlockOnHost
func GlobalUnlockOnHost()

//export ModuleLockOnHost
func ModuleLockOnHost(hostModulePtr uint64)

//export ModuleUnlockOnHost
func ModuleUnlockOnHost(hostModulePtr uint64)

//export GetAbiVersionOnHost
func GetAbiVersionOnHost(returnValuePtr *uint32, returnValueSize *uint32) tools.Status

//export GetRuntimeTypeOnHost
func GetRuntimeTypeOnHost(returnValuePtr *uint32, returnValueSize *uint32) tools.Status

//export InfoLogOnHost
func InfoLogOnHost(ptr uint32, size uint32) tools.Status

//export ErrorLogOnHost
func ErrorLogOnHost(ptr uint32, size uint32) tools.Status
