package wescale_wasm_plugin_template

import (
	"errors"
	"fmt"
	"strconv"
	"wescale-wasm-plugin-template/tools"
	hostfunction "wescale-wasm-plugin-template/tools/host_functions"
)

func RunBeforeExecution() {
	// TODO: Write your code here
	hostfunction.GlobalLock()
	var moduleCount int
	countBytes, err := hostfunction.GetModuleValueByKey("moduleCount")
	if errors.Is(err, tools.ErrorStatusNotFound) {
		moduleCount = 0
		hostfunction.SetModuleValueByKey("moduleCount", []byte(strconv.Itoa(moduleCount)))
	}

	countBytes, _ = hostfunction.GetModuleValueByKey("moduleCount")
	moduleCount, _ = strconv.Atoi(string(countBytes))
	moduleCount++

	if moduleCount%2 == 0 {
		hostfunction.SetHostQuery("select * from guest.setquerytest;")
	} else {
		str, _ := hostfunction.GetHostQuery()
		hostfunction.SetHostQuery(str)
	}

	runtimeType, _ := hostfunction.GetRuntimeType()
	version, _ := hostfunction.GetAbiVersion()
	if runtimeType != "" {
		if version != "" {

		}
	}
	hostfunction.InfoLog(fmt.Sprintf("wasm guest runtime type:%v version:%v", runtimeType, version))

	hostfunction.SetModuleValueByKey("moduleCount", []byte(strconv.Itoa(moduleCount)))
	hostfunction.GlobalUnlock()

}

func RunAfterExecution(exchange *tools.WasmPluginRunAfterExecutionExchange) {
	// TODO: Write your code here
}
