package wescale_wasm_plugin_template

import (
	"wescale-wasm-plugin-template/tools"
	hostfunction "wescale-wasm-plugin-template/tools/host_functions"
)

func RunBeforeExecution() {
	// TODO: Write your code here

	globalCount := hostfunction.GetValueByKeyHost(1)
	globalCount++
	hostfunction.SetValueByKeyHost(1, globalCount)

	if globalCount%2 == 0 {
		hostfunction.SetHostQuery("select * from guest.setquerytest;")
	} else {
		str, _ := hostfunction.GetHostQuery()
		hostfunction.SetHostQuery(str)
	}
}

func RunAfterExecution(exchange *tools.WasmPluginRunAfterExecutionExchange) {
	// TODO: Write your code here
}
