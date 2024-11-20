package v1alpha2

import (
	"github.com/wesql/sqlparser/go/vt/proto/query"
)

type WasmPlugin interface {
	RunBeforeExecution() error
	RunAfterExecution(queryResult *query.QueryResult, errBefore error) (*query.QueryResult, error)
}

var wasmPlugin WasmPlugin

func SetWasmPlugin(plugin WasmPlugin) {
	wasmPlugin = plugin
}
