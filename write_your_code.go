package wescale_wasm_plugin_template

import "wescale-wasm-plugin-template/tools"

func WasmPlugin(exchange *tools.WasmPluginExchange) {
	// TODO: Write your code here
	exchange.Query = "select * from d1.t1;"
}
