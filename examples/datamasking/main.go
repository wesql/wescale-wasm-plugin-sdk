package main

import (
	"fmt"
	"github.com/wesql/wescale-wasm-plugin-template/internal"
	hostfunction "github.com/wesql/wescale-wasm-plugin-template/internal/host_functions/v1alpha1"
	"github.com/wesql/wescale-wasm-plugin-template/internal/proto/query"
)

func main() {
	internal.SetWasmPlugin(&DataMaskingWasmPlugin{})
}

type DataMaskingWasmPlugin struct {
}

func (a *DataMaskingWasmPlugin) RunBeforeExecution() error {
	query, err := hostfunction.GetHostQuery()
	if err != nil {
		return err
	}
	hostfunction.InfoLog("execute SQL: " + query)
	return nil
}

func (a *DataMaskingWasmPlugin) RunAfterExecution(queryResult *query.QueryResult, errBefore error) (*query.QueryResult, error) {
	if queryResult != nil {
		hostfunction.InfoLog(fmt.Sprintf("returned rows: %v", len(queryResult.Rows)))
		hostfunction.InfoLog(fmt.Sprintf("affected rows: %v", queryResult.RowsAffected))
	}
	if errBefore != nil {
		hostfunction.InfoLog("execution error: " + errBefore.Error())
	}

	return queryResult, errBefore
}
