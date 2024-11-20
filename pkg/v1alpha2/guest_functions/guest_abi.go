package guest_functions

import (
	"errors"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg/v1alpha2/host_functions"
)

//export proxy_on_memory_allocate
func proxyOnMemoryAllocate(size uint) *byte {
	buf := make([]byte, size)
	return &buf[0]
}

//export RunBeforeExecutionOnGuest
func RunBeforeExecutionOnGuest(hostInstancePtr, hostModulePtr uint64) {
	pkg.CurrentWasmPluginContext.HostInstancePtr = hostInstancePtr
	pkg.CurrentWasmPluginContext.HostModulePtr = hostModulePtr

	err := pkg.CurrentWasmPlugin.RunBeforeExecution()
	if err != nil {
		host_functions.SetErrorMessage(err.Error())
	}
}

//export RunAfterExecutionOnGuest
func RunAfterExecutionOnGuest() {
	qr, err := host_functions.GetQueryResult()
	if err != nil && !errors.Is(err, pkg.StatusToError(pkg.StatusBadArgument)) {
		// unknown error
		host_functions.SetErrorMessage(err.Error())
		return
	}

	errMessageBefore, err := host_functions.GetErrorMessage()
	if err != nil {
		if !errors.Is(err, pkg.StatusToError(pkg.StatusBadArgument)) {
			// unknown error
			host_functions.SetErrorMessage(err.Error())
			return
		} else {
			errMessageBefore = ""
		}

	}
	var errBefore error
	if errMessageBefore != "" {
		errBefore = errors.New(errMessageBefore)
	}

	finalQueryResult, finalErr := pkg.CurrentWasmPlugin.RunAfterExecution(qr, errBefore)

	host_functions.SetQueryResult(finalQueryResult)
	if finalErr != nil {
		host_functions.SetErrorMessage(finalErr.Error())
	}
}
