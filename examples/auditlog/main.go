package main

import (
	"fmt"
	"github.com/wesql/wescale-wasm-plugin/internal"
	hostfunction "github.com/wesql/wescale-wasm-plugin/internal/host_functions/v1alpha1"
	"github.com/wesql/wescale-wasm-plugin/internal/proto/query"
)

func main() {
	internal.SetWasmPlugin(&AuditLogWasmPlugin{})
}

type AuditLogWasmPlugin struct {
}

func (a *AuditLogWasmPlugin) RunBeforeExecution() error {
	query, err := hostfunction.GetHostQuery()
	if err != nil {
		return err
	}
	hostfunction.InfoLog("execute SQL: " + query)
	return nil
}

func (a *AuditLogWasmPlugin) RunAfterExecution(queryResult *query.QueryResult, errBefore error) (*query.QueryResult, error) {
	if queryResult != nil {
		hostfunction.InfoLog(fmt.Sprintf("returned rows: %v", len(queryResult.Rows)))
		hostfunction.InfoLog(fmt.Sprintf("affected rows: %v", queryResult.RowsAffected))
	}
	if errBefore != nil {
		hostfunction.InfoLog("execution error: " + errBefore.Error())
	}

	return queryResult, errBefore
}
