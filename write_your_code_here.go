package wescale_wasm_plugin_template

import (
	"golang.org/x/sync/semaphore"
	"wescale-wasm-plugin-template/tools"
	hostfunction "wescale-wasm-plugin-template/tools/host_functions"
)

const maxConcurrency = 1 // max concurrency
var sem *semaphore.Weighted

var count = 0

func init() {
	// TODO: Write your code here to init some global variables
	sem = semaphore.NewWeighted(maxConcurrency)
}

func RunBeforeExecution(exchange *tools.WasmPluginRunBeforeExecutionExchange) {
	// TODO: Write your code here
	//sem.Acquire(context.Background(), 1)
	//time.Sleep(3 * time.Second)

	globalCount := hostfunction.GetValueByKeyHost(1)
	globalCount++
	hostfunction.SetValueByKeyHost(1, globalCount)
	// todo by newborn22, move to another location

	str, _ := hostfunction.GetHostQuery()
	//if str != "" {
	//
	//}

	//if count%2 == 1 {
	//	exchange.Query = "select * from d11.t11"
	//} else {
	//	exchange.Query = "select * from d22.t22"
	//}
	//exchange.Query = fmt.Sprintf("select * from d%d.t1", globalCount)
	exchange.Query = str
}

func RunAfterExecution(exchange *tools.WasmPluginRunAfterExecutionExchange) {
	// TODO: Write your code here
	//sem.Release(1)
	//count--
	//log.Printf("in wasm after: %v", exchange)
	//exchange.Query = "select * from d3.t3"
	//log.Printf("in wasm after: %v", exchange)
}
