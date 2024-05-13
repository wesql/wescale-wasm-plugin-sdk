package wescale_wasm_plugin_template

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"wescale-wasm-plugin-template/tools"
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
	sem.Acquire(context.Background(), 1)
	//time.Sleep(3 * time.Second)
	//count++
	//if count%2 == 1 {
	//	exchange.Query = "select * from d11.t11"
	//} else {
	//	exchange.Query = "select * from d22.t22"
	//}
	exchange.Query = "select * from d1.t1"
}

func RunAfterExecution(exchange *tools.WasmPluginRunAfterExecutionExchange) {
	// TODO: Write your code here
	sem.Release(1)
	//count--
	fmt.Printf("in wasm after: %v", exchange)
	exchange.Query = "select * from d3.t3"
	fmt.Printf("in wasm after: %v", exchange)
}
