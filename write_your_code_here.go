package wescale_wasm_plugin_template

import (
	"errors"
	"strconv"
	"wescale-wasm-plugin-template/tools"
	hostfunction "wescale-wasm-plugin-template/tools/host_functions"
)

func RunBeforeExecution() {
	// TODO: Write your code here

	var globalCount int
	countBytes, err := hostfunction.GetGlobalValueByKey("globalCount")
	if errors.Is(err, tools.ErrorStatusNotFound) {
		globalCount = 0
		hostfunction.SetGlobalValueByKey("globalCount", []byte(strconv.Itoa(globalCount)))
	}

	countBytes, _ = hostfunction.GetGlobalValueByKey("globalCount")
	globalCount, _ = strconv.Atoi(string(countBytes))
	globalCount++

	if globalCount%2 == 0 {
		hostfunction.SetHostQuery("select * from guest.setquerytest;")
	} else {
		str, _ := hostfunction.GetHostQuery()
		hostfunction.SetHostQuery(str)
	}

	hostfunction.SetGlobalValueByKey("globalCount", []byte(strconv.Itoa(globalCount)))
}

func RunAfterExecution(exchange *tools.WasmPluginRunAfterExecutionExchange) {
	// TODO: Write your code here
}
