package host_functions

import (
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg/types"
)

//export GetQueryOnHost
func getQueryOnHost(hostInstancePtr uint64, returnValueData *uint32, returnValueSize *uint32) types.Status

//export SetQueryOnHost
func setQueryOnHost(hostInstancePtr uint64, queryValuePtr uint32, queryValueSize uint32) types.Status

//export GetAbiVersionOnHost
func getAbiVersionOnHost(returnValuePtr *uint32, returnValueSize *uint32) types.Status

//export GetRuntimeTypeOnHost
func getRuntimeTypeOnHost(returnValuePtr *uint32, returnValueSize *uint32) types.Status

//export InfoLogOnHost
func infoLogOnHost(ptr uint32, size uint32) types.Status

//export ErrorLogOnHost
func errorLogOnHost(ptr uint32, size uint32) types.Status
