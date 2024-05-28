package internal

import "github.com/wesql/wescale-wasm-plugin-template/proto/query"

type Plugin interface {
	RunBeforeExecution() error
	RunAfterExecution(queryResult *query.QueryResult, errBefore error) (*query.QueryResult, error)
}
