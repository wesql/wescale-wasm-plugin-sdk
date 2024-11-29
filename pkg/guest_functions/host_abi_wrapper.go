package guest_functions

import (
	"github.com/wesql/sqlparser/go/vt/proto/query"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg/types"
)

func setErrorMessage(pluginCtx types.WasmPluginContext, errMessage string) {
	if len(errMessage) == 0 {
		return
	}
	ptr, size := types.StringToPtr(errMessage)
	_setErrorMessageOnHost(pluginCtx.Id, ptr, size)
}

func getErrorMessage(pluginCtx types.WasmPluginContext) (string, error) {
	var ptr uint32
	var retSize uint32

	err := types.StatusToError(_getErrorMessageOnHost(pluginCtx.Id, &ptr, &retSize))
	if err != nil {
		return "", err
	}
	return types.PtrToStringWithFree(ptr, retSize), nil
}

func getQueryResult(pluginCtx types.WasmPluginContext) (*query.QueryResult, error) {
	var ptr uint32
	var retSize uint32

	err := types.StatusToError(_getQueryResultOnHost(pluginCtx.Id, &ptr, &retSize))
	if err != nil {
		return nil, err
	}
	bytes := types.PtrToBytesWithFree(ptr, retSize)
	queryResult := &query.QueryResult{}
	err = queryResult.UnmarshalVT(bytes)
	if err != nil {
		return nil, err
	}
	return queryResult, nil
}

func setQueryResult(pluginCtx types.WasmPluginContext, queryResult *query.QueryResult) error {
	if queryResult == nil {
		return nil
	}
	bytes, err := queryResult.MarshalVT()
	if err != nil {
		return nil
	}
	ptr, size := types.BytesToPtr(bytes)
	return types.StatusToError(_setQueryResultOnHost(pluginCtx.Id, ptr, size))
}
