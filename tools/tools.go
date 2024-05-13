package tools

// todo, it's a silly way
// todo, add a field to set log here?
type WasmPluginRunBeforeExecutionExchange struct {
	Query string
}

// todo, which fields? query result?
type WasmPluginRunAfterExecutionExchange struct {
	Query string
}
