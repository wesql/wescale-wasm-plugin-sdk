package hostfunction

import (
	"wescale-wasm-plugin-template/internal"
)

//export GetGlobalValueByKeyOnHost
func getGlobalValueByKeyOnHost(keyPtr, keySize uint32, returnValuePtr, returnValueSize *uint32) internal.Status

//export SetGlobalValueByKeyOnHost
func setGlobalValueByKeyOnHost(keyPtr, keySize, valuePtr, valueSize uint32) internal.Status

//export GetModuleValueByKeyOnHost
func getModuleValueByKeyOnHost(hostModulePtr uint64, keyPtr, keySize uint32, returnValuePtr, returnValueSize *uint32) internal.Status

//export SetModuleValueByKeyOnHost
func setModuleValueByKeyOnHost(hostModulePtr uint64, keyPtr, keySize, valuePtr, valueSize uint32) internal.Status

//export GetQueryOnHost
func getQueryOnHost(hostInstancePtr uint64, returnValueData *uint32, returnValueSize *uint32) internal.Status

//export SetQueryOnHost
func setQueryOnHost(hostInstancePtr uint64, queryValuePtr uint32, queryValueSize uint32) internal.Status

//export GlobalLockOnHost
func globalLockOnHost()

//export GlobalUnlockOnHost
func globalUnlockOnHost()

//export ModuleLockOnHost
func moduleLockOnHost(hostModulePtr uint64)

//export ModuleUnlockOnHost
func moduleUnlockOnHost(hostModulePtr uint64)

//export GetAbiVersionOnHost
func getAbiVersionOnHost(returnValuePtr *uint32, returnValueSize *uint32) internal.Status

//export GetRuntimeTypeOnHost
func getRuntimeTypeOnHost(returnValuePtr *uint32, returnValueSize *uint32) internal.Status

//export InfoLogOnHost
func infoLogOnHost(ptr uint32, size uint32) internal.Status

//export ErrorLogOnHost
func errorLogOnHost(ptr uint32, size uint32) internal.Status

//export SetErrorMessageOnHost
func setErrorMessageOnHost(hostInstancePtr uint64, errMessagePtr uint32, errMessageSize uint32) internal.Status

//export GetErrorMessageOnHost
<<<<<<< Updated upstream:internal/host_functions/host_abi.go
func getErrorMessageOnHost(hostInstancePtr uint64, errMessagePtr *uint32, errMessageSize *uint32) internal.Status
=======
func getErrorMessageOnHost(hostInstancePtr uint64, returnErrMessagePtr *uint32, returnErrMessageSize *uint32) common.Status

//export GetQueryResultOnHost
func getQueryResultOnHost(hostInstancePtr uint64, returnQueryRstBytesPtr *uint32, returnQueryRstBytesSize *uint32) common.Status

//export SetQueryResultOnHost
func setQueryResultOnHost(hostInstancePtr uint64, returnQueryRstBytesPtr uint32, returnQueryRstBytesSize uint32) common.Status
>>>>>>> Stashed changes:common/host_functions/host_abi.go
