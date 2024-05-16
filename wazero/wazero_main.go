package main

// #include <stdlib.h>
import "C"

import (
	"encoding/json"
	wescale_wasm_plugin_template "wescale-wasm-plugin-template"
	"wescale-wasm-plugin-template/tools"
)

// main is required for TinyGo to compile to Wasm.
func main() {}

//export wazeroGuestFunc
func wazeroGuestFunc(ptr, size uint32) (ptrSize uint64) {
	dataFromHost := tools.PtrToString(ptr, size)

	w := tools.WasmPluginRunBeforeExecutionExchange{}
	// todo, how to handle errors?
	json.Unmarshal([]byte(dataFromHost), &w)

	wescale_wasm_plugin_template.RunBeforeExecution(&w)

	// todo, how to handle errors?
	dataToHost, _ := json.Marshal(&w)
	dataToHostString := string(dataToHost)

	ptr, size = tools.StringToLeakedPtr(dataToHostString)
	return (uint64(ptr) << uint64(32)) | uint64(size)
}

//export wazeroGuestFuncAfterExecution
func wazeroGuestFuncAfterExecution(ptr, size uint32) (ptrSize uint64) {
	dataFromHost := tools.PtrToString(ptr, size)

	w := tools.WasmPluginRunAfterExecutionExchange{}
	// todo, how to handle errors?
	json.Unmarshal([]byte(dataFromHost), &w)

	wescale_wasm_plugin_template.RunAfterExecution(&w)

	// todo, how to handle errors?
	dataToHost, _ := json.Marshal(&w)
	dataToHostString := string(dataToHost)

	ptr, size = tools.StringToLeakedPtr(dataToHostString)
	return (uint64(ptr) << uint64(32)) | uint64(size)
}
