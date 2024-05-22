package hostfunction

import (
	"wescale-wasm-plugin-template/common"
)

//export GetGlobalValueByKeyOnHost
func getGlobalValueByKeyOnHost(keyPtr, keySize uint32, returnValuePtr, returnValueSize *uint32) common.Status

//export SetGlobalValueByKeyOnHost
func setGlobalValueByKeyOnHost(keyPtr, keySize, valuePtr, valueSize uint32) common.Status

//export GetModuleValueByKeyOnHost
func getModuleValueByKeyOnHost(hostModulePtr uint64, keyPtr, keySize uint32, returnValuePtr, returnValueSize *uint32) common.Status

//export SetModuleValueByKeyOnHost
func setModuleValueByKeyOnHost(hostModulePtr uint64, keyPtr, keySize, valuePtr, valueSize uint32) common.Status

//export GetQueryOnHost
func getQueryOnHost(hostInstancePtr uint64, returnValueData *uint32, returnValueSize *uint32) common.Status

//export SetQueryOnHost
func setQueryOnHost(hostInstancePtr uint64, queryValuePtr uint32, queryValueSize uint32) common.Status

//export GlobalLockOnHost
func globalLockOnHost()

//export GlobalUnlockOnHost
func globalUnlockOnHost()

//export ModuleLockOnHost
func moduleLockOnHost(hostModulePtr uint64)

//export ModuleUnlockOnHost
func moduleUnlockOnHost(hostModulePtr uint64)

//export GetAbiVersionOnHost
func getAbiVersionOnHost(returnValuePtr *uint32, returnValueSize *uint32) common.Status

//export GetRuntimeTypeOnHost
func getRuntimeTypeOnHost(returnValuePtr *uint32, returnValueSize *uint32) common.Status

//export InfoLogOnHost
func infoLogOnHost(ptr uint32, size uint32) common.Status

//export ErrorLogOnHost
func errorLogOnHost(ptr uint32, size uint32) common.Status
