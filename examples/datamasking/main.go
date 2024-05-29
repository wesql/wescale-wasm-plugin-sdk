package main

import (
	"github.com/wesql/wescale-wasm-plugin-template/internal"
	"github.com/wesql/wescale-wasm-plugin-template/internal/proto/query"
)

func main() {
	internal.SetWasmPlugin(&DataMaskingWasmPlugin{})
}

type DataMaskingWasmPlugin struct {
}

func (a *DataMaskingWasmPlugin) RunBeforeExecution() error {
	// do nothing
	return nil
}

func (a *DataMaskingWasmPlugin) RunAfterExecution(queryResult *query.QueryResult, errBefore error) (*query.QueryResult, error) {
	if queryResult != nil {
		for i := range queryResult.Rows {
			for j := range queryResult.Rows[i].Values {
				if isStringType(queryResult.GetFields()[j].Type) {
					//TODO: mask the string
				}
			}
		}
	}
	return queryResult, errBefore
}

func isStringType(t query.Type) bool {
	return t == query.Type_VARCHAR || t == query.Type_CHAR || t == query.Type_TEXT
}
