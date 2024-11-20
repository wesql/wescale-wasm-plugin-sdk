package main

import (
	"github.com/wesql/sqlparser/go/vt/proto/query"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg"
	"testing"
)

type TestWasmPlugin struct {
}

func (a *TestWasmPlugin) RunBeforeExecution() error {
	return nil
}

func (a *TestWasmPlugin) RunAfterExecution(queryResult *query.QueryResult, errBefore error) (*query.QueryResult, error) {
	return queryResult, errBefore
}

func TestSdkCompileSuccess(t *testing.T) {
	// This test is to ensure that the SDK compiles successfully
	pkg.SetWasmPlugin(&TestWasmPlugin{})
}
