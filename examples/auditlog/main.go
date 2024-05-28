package auditlog

import (
	"fmt"
	hostfunction "wescale-wasm-plugin-template/internal/host_functions"
	"wescale-wasm-plugin-template/proto/query"
)

func RunBeforeExecution() error {
	query, err := hostfunction.GetHostQuery()
	if err != nil {
		return err
	}
	hostfunction.InfoLog("execute SQL: " + query)
	return nil
}

func RunAfterExecution(queryResult *query.QueryResult, errBefore error) (*query.QueryResult, error) {

	if queryResult != nil {
		hostfunction.InfoLog(fmt.Sprintf("affected rows: %v", queryResult.RowsAffected))
	}
	if errBefore != nil {
		hostfunction.InfoLog("execution error: " + errBefore.Error())
	}

	return queryResult, errBefore
}
