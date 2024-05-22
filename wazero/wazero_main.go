package main

// #include <stdlib.h>
import "C"

import (
	wescale_wasm_plugin_template "wescale-wasm-plugin-template"
	"wescale-wasm-plugin-template/common/host_functions"
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
	err := wescale_wasm_plugin_template.RunAfterExecution()
	if err != nil {
		hostfunction.SetErrorMessage(err.Error())
	}
}

//export proxy_on_memory_allocate
func proxyOnMemoryAllocate(size uint) *byte {
	buf := make([]byte, size)
	return &buf[0]
}
