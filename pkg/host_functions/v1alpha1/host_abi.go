package v1alpha1

//export GetValueByKeyOnHost
func getValueByKeyOnHost(scope uint32, hostModulePtr uint64, keyPtr, keySize uint32, returnValuePtr, returnValueSize *uint32) Status

//export SetValueByKeyOnHost
func setValueByKeyOnHost(scope uint32, hostModulePtr uint64, keyPtr, keySize, valuePtr, valueSize uint32) Status

//export LockOnHost
func lockOnHost(scope uint32, hostModulePtr uint64)

//export UnlockOnHost
func unlockOnHost(scope uint32, hostModulePtr uint64)

//export GetQueryOnHost
func getQueryOnHost(hostInstancePtr uint64, returnValueData *uint32, returnValueSize *uint32) Status

//export SetQueryOnHost
func setQueryOnHost(hostInstancePtr uint64, queryValuePtr uint32, queryValueSize uint32) Status

//export GetAbiVersionOnHost
func getAbiVersionOnHost(returnValuePtr *uint32, returnValueSize *uint32) Status

//export GetRuntimeTypeOnHost
func getRuntimeTypeOnHost(returnValuePtr *uint32, returnValueSize *uint32) Status

//export InfoLogOnHost
func infoLogOnHost(ptr uint32, size uint32) Status

//export ErrorLogOnHost
func errorLogOnHost(ptr uint32, size uint32) Status

//export SetErrorMessageOnHost
func setErrorMessageOnHost(hostInstancePtr uint64, errMessagePtr uint32, errMessageSize uint32) Status

//export GetErrorMessageOnHost
func getErrorMessageOnHost(hostInstancePtr uint64, returnErrMessagePtr *uint32, returnErrMessageSize *uint32) Status

//export GetQueryResultOnHost
func getQueryResultOnHost(hostInstancePtr uint64, returnQueryRstBytesPtr *uint32, returnQueryRstBytesSize *uint32) Status

//export SetQueryResultOnHost
func setQueryResultOnHost(hostInstancePtr uint64, returnQueryRstBytesPtr uint32, returnQueryRstBytesSize uint32) Status
