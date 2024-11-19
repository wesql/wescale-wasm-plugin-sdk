package pkg

import (
	"errors"
	hostfunction "github.com/wesql/wescale-wasm-plugin-sdk/pkg/host_functions/v1alpha1"
)

//export proxy_on_memory_allocate
func proxyOnMemoryAllocate(size uint) *byte {
	buf := make([]byte, size)
	return &buf[0]
}

//export RunBeforeExecutionOnGuest
func RunBeforeExecutionOnGuest(hostInstancePtr, hostModulePtr uint64) {
	hostfunction.HostInstancePtr = hostInstancePtr
	hostfunction.HostModulePtr = hostModulePtr

	err := wasmPlugin.RunBeforeExecution()
	if err != nil {
		hostfunction.SetErrorMessage(err.Error())
	}
}

//export RunAfterExecutionOnGuest
func RunAfterExecutionOnGuest() {
	qr, err := hostfunction.GetQueryResult()
	if err != nil && !errors.Is(err, hostfunction.StatusToError(hostfunction.StatusBadArgument)) {
		// unknown error
		hostfunction.SetErrorMessage(err.Error())
		return
	}

	errMessageBefore, err := hostfunction.GetErrorMessage()
	if err != nil {
		if !errors.Is(err, hostfunction.StatusToError(hostfunction.StatusBadArgument)) {
			// unknown error
			hostfunction.SetErrorMessage(err.Error())
			return
		} else {
			errMessageBefore = ""
		}

	}
	var errBefore error
	if errMessageBefore != "" {
		errBefore = errors.New(errMessageBefore)
	}

	finalQueryResult, finalErr := wasmPlugin.RunAfterExecution(qr, errBefore)

	hostfunction.SetQueryResult(finalQueryResult)
	if finalErr != nil {
		hostfunction.SetErrorMessage(finalErr.Error())
	}
}
