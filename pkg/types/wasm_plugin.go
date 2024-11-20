package types

import (
	"github.com/wesql/sqlparser/go/vt/proto/query"
)

type WasmPlugin interface {
	RunBeforeExecution() error
	RunAfterExecution(queryResult *query.QueryResult, errBefore error) (*query.QueryResult, error)
}

var CurrentWasmPlugin WasmPlugin
