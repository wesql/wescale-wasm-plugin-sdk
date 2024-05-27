package main

// #include <stdlib.h>
import "C"

import (
	"errors"
	wescale_wasm_plugin_template "wescale-wasm-plugin-template"
	"wescale-wasm-plugin-template/internal"
	"wescale-wasm-plugin-template/internal/host_functions"
)

// main is required for TinyGo to compile to Wasm.
func main() {}

//export WazeroGuestFuncBeforeExecution
func WazeroGuestFuncBeforeExecution(hostInstancePtr, hostModulePtr uint64) {
	hostfunction.HostInstancePtr = hostInstancePtr
	hostfunction.HostModulePtr = hostModulePtr

	err := wescale_wasm_plugin_template.RunBeforeExecution()
	if err != nil {
		hostfunction.SetErrorMessage(err.Error())
	}
}

//export WazeroGuestFuncAfterExecution
func WazeroGuestFuncAfterExecution() {
	qr, err := hostfunction.GetQueryResult()
	if err != nil && !errors.Is(err, internal.StatusToError(internal.StatusBadArgument)) {
		// unknown error
		hostfunction.SetErrorMessage(err.Error())
		return
	}

	errMessageBefore, err := hostfunction.GetErrorMessage()
	if err != nil {
		if !errors.Is(err, internal.StatusToError(internal.StatusBadArgument)) {
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

	finalQueryResult, finalErr := wescale_wasm_plugin_template.RunAfterExecution(qr, errBefore)

	hostfunction.SetQueryResult(finalQueryResult)
	if finalErr != nil {
		hostfunction.SetErrorMessage(finalErr.Error())
	}
}

//export proxy_on_memory_allocate
func proxyOnMemoryAllocate(size uint) *byte {
	buf := make([]byte, size)
	return &buf[0]
}
