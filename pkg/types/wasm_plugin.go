package types

import (
	"github.com/wesql/sqlparser/go/vt/proto/query"
)

type WasmPlugin interface {
	RunBeforeExecution(ctx WasmPluginContext) error
	RunAfterExecution(ctx WasmPluginContext, queryResult *query.QueryResult, errBefore error) (*query.QueryResult, error)
}
