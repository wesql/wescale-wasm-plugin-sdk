package types

type WasmPluginContext struct {
	HostInstancePtr uint64
	HostModulePtr   uint64
}

var CurrentWasmPluginContext = WasmPluginContext{}
