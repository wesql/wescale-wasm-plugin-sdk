package main

import (
	"fmt"
	"github.com/wesql/wescale-wasm-plugin/pkg"
	hostfunction "github.com/wesql/wescale-wasm-plugin/pkg/host_functions/v1alpha1"
	"github.com/wesql/wescale-wasm-plugin/pkg/proto/query"
	"github.com/xwb1989/sqlparser"
)

func main() {
	pkg.SetWasmPlugin(&ParserWasmPlugin{})
}

type ParserWasmPlugin struct {
}

func (a *ParserWasmPlugin) RunBeforeExecution() error {
	query, err := hostfunction.GetHostQuery()
	if err != nil {
		return err
	}

	stmt, err := sqlparser.Parse(query)
	if err != nil {
		hostfunction.InfoLog("parse error: " + err.Error())
		return nil
	}
	switch stmt := stmt.(type) {
	case *sqlparser.Update:
		if stmt.Where == nil {
			return fmt.Errorf("no where clause")
		}
	case *sqlparser.Delete:
		if stmt.Where == nil {
			return fmt.Errorf("no where clause")
		}
	default:
	}

	return nil
}

func (a *ParserWasmPlugin) RunAfterExecution(queryResult *query.QueryResult, errBefore error) (*query.QueryResult, error) {
	// do nothing
	return queryResult, errBefore
}
