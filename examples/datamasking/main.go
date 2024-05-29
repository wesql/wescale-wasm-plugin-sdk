package main

import (
	"bytes"
	"github.com/wesql/wescale-wasm-plugin/pkg"
	"github.com/wesql/wescale-wasm-plugin/pkg/proto/query"
)

func main() {
	pkg.SetWasmPlugin(&DataMaskingWasmPlugin{})
}

type DataMaskingWasmPlugin struct {
}

func (a *DataMaskingWasmPlugin) RunBeforeExecution() error {
	// do nothing
	return nil
}

func (a *DataMaskingWasmPlugin) RunAfterExecution(queryResult *query.QueryResult, errBefore error) (*query.QueryResult, error) {
	if queryResult == nil {
		return queryResult, errBefore
	}
	for rowIndex := range queryResult.Rows {
		newLengths := make([]int64, 0)
		newValues := bytes.Buffer{}
		var offset int64 = 0
		for colIndex, colLength := range queryResult.Rows[rowIndex].Lengths {
			if colLength == -1 {
				newLengths = append(newLengths, -1)
				continue
			}
			if isStringType(queryResult.GetFields()[colIndex].Type) {
				maskedValue := []byte("****")
				newLengths = append(newLengths, int64(4))
				newValues.Write(maskedValue)
			} else {
				newLengths = append(newLengths, queryResult.Rows[rowIndex].Lengths[colIndex])
				length := queryResult.Rows[rowIndex].Lengths[colIndex]
				val := queryResult.Rows[rowIndex].Values[offset : offset+length]
				newValues.Write(val)
			}
			offset += colLength
		}
		queryResult.Rows[rowIndex].Lengths = newLengths
		queryResult.Rows[rowIndex].Values = newValues.Bytes()
	}
	return queryResult, errBefore
}

func isStringType(t query.Type) bool {
	return t == query.Type_VARCHAR || t == query.Type_CHAR || t == query.Type_TEXT
}
