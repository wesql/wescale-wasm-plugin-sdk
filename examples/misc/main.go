package main

import (
	"errors"
	"fmt"
	"github.com/wesql/wescale-wasm-plugin/pkg"
	hostfunction "github.com/wesql/wescale-wasm-plugin/pkg/host_functions/v1alpha1"
	"github.com/wesql/wescale-wasm-plugin/pkg/proto/query"
	"strconv"
)

func main() {
	pkg.SetWasmPlugin(&MiscWasmPlugin{})
}

type MiscWasmPlugin struct {
}

func (a *MiscWasmPlugin) RunBeforeExecution() error {
	// TODO: Write your code here
	hostfunction.GlobalLock()
	var moduleCount int
	countBytes, err := hostfunction.GetModuleValueByKey("moduleCount")
	if errors.Is(err, hostfunction.ErrorStatusNotFound) {
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
	return nil
}

func (a *MiscWasmPlugin) RunAfterExecution(queryResult *query.QueryResult, errBefore error) (*query.QueryResult, error) {
	// TODO: Write your code here
	// you should know that the queryResult can be nil

	if queryResult == nil {
		return nil, fmt.Errorf("new error in after: %v", errBefore)
	}
	queryResult.Rows = append(queryResult.Rows, queryResult.Rows...)
	return queryResult, nil
}
