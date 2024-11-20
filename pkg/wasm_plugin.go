package pkg

import (
	"github.com/wesql/sqlparser/go/vt/proto/query"
)

type WasmPluginContext struct {
	HostInstancePtr uint64
	HostModulePtr   uint64
}

type WasmPlugin interface {
	RunBeforeExecution() error
	RunAfterExecution(queryResult *query.QueryResult, errBefore error) (*query.QueryResult, error)
}

var CurrentWasmPluginContext = WasmPluginContext{}

var CurrentWasmPlugin WasmPlugin

func SetWasmPlugin(plugin WasmPlugin) {
	CurrentWasmPlugin = plugin
}
