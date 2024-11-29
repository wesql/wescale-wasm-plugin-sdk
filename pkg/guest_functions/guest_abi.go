package guest_functions

import (
	"errors"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg/types"
)

var CurrentWasmPlugin types.WasmPlugin

//export malloc
func malloc(size uint) *byte

// Deprecated: use malloc instead
//
//export proxy_on_memory_allocate
func proxyOnMemoryAllocate(size uint) *byte {
	return malloc(size)
}

//export RunBeforeExecutionOnGuest
func RunBeforeExecutionOnGuest(hostInstanceId uint64) {
	pluginCtx := types.WasmPluginContext{Id: hostInstanceId}
	err := CurrentWasmPlugin.RunBeforeExecution(pluginCtx)
	if err != nil {
		setErrorMessage(pluginCtx, err.Error())
	}
}

//export RunAfterExecutionOnGuest
func RunAfterExecutionOnGuest(hostInstanceId uint64) {
	pluginCtx := types.WasmPluginContext{Id: hostInstanceId}
	qr, err := getQueryResult(pluginCtx)
	if err != nil && !errors.Is(err, types.StatusToError(types.StatusBadArgument)) {
		// unknown error
		setErrorMessage(pluginCtx, err.Error())
		return
	}

	errMessageBefore, err := getErrorMessage(pluginCtx)
	if err != nil {
		if !errors.Is(err, types.StatusToError(types.StatusBadArgument)) {
			// unknown error
			setErrorMessage(pluginCtx, err.Error())
			return
		} else {
			errMessageBefore = ""
		}

	}
	var errBefore error
	if errMessageBefore != "" {
		errBefore = errors.New(errMessageBefore)
	}

	finalQueryResult, finalErr := CurrentWasmPlugin.RunAfterExecution(pluginCtx, qr, errBefore)

	setQueryResult(pluginCtx, finalQueryResult)
	if finalErr != nil {
		setErrorMessage(pluginCtx, finalErr.Error())
	} else {
		setErrorMessage(pluginCtx, "")
	}
}
