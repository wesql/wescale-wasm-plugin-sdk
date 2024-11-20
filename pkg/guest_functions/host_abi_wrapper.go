package guest_functions

import (
	"github.com/wesql/sqlparser/go/vt/proto/query"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg/types"
)

func setErrorMessage(errMessage string) {
	if len(errMessage) == 0 {
		return
	}
	ptr, size := types.StringToPtr(errMessage)
	_setErrorMessageOnHost(types.CurrentWasmPluginContext.HostInstancePtr, ptr, size)
}

func getErrorMessage() (string, error) {
	var ptr uint32
	var retSize uint32

	err := types.StatusToError(_getErrorMessageOnHost(types.CurrentWasmPluginContext.HostInstancePtr, &ptr, &retSize))
	if err != nil {
		return "", err
	}
	return types.PtrToStringWithFree(ptr, retSize), nil
}

func getQueryResult() (*query.QueryResult, error) {
	var ptr uint32
	var retSize uint32

	err := types.StatusToError(_getQueryResultOnHost(types.CurrentWasmPluginContext.HostInstancePtr, &ptr, &retSize))
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

func setQueryResult(queryResult *query.QueryResult) error {
	if queryResult == nil {
		return nil
	}
	bytes, err := queryResult.MarshalVT()
	if err != nil {
		return nil
	}
	ptr, size := types.BytesToPtr(bytes)
	return types.StatusToError(_setQueryResultOnHost(types.CurrentWasmPluginContext.HostInstancePtr, ptr, size))
}
