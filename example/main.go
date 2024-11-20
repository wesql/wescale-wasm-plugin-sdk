package main

import (
	"fmt"
	"github.com/wesql/sqlparser"
	"github.com/wesql/sqlparser/go/vt/proto/query"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg"
	hostfunction "github.com/wesql/wescale-wasm-plugin-sdk/pkg/v1alpha2/host_functions"
)

func main() {
	pkg.InitWasmPlugin(&ParserWasmPlugin{})
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
