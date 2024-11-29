package guest_functions

import (
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg/types"
)

//export SetErrorMessageOnHost
func _setErrorMessageOnHost(hostInstanceId uint64, errMessagePtr uint32, errMessageSize uint32) types.Status

//export GetErrorMessageOnHost
func _getErrorMessageOnHost(hostInstanceId uint64, returnErrMessagePtr *uint32, returnErrMessageSize *uint32) types.Status

//export GetQueryResultOnHost
func _getQueryResultOnHost(hostInstanceId uint64, returnQueryRstBytesPtr *uint32, returnQueryRstBytesSize *uint32) types.Status

//export SetQueryResultOnHost
func _setQueryResultOnHost(hostInstanceId uint64, returnQueryRstBytesPtr uint32, returnQueryRstBytesSize uint32) types.Status
