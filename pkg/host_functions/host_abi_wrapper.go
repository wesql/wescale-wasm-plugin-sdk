package host_functions

import (
	"errors"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg/types"
)

func GetHostQuery(ctx types.WasmPluginContext) (string, error) {
	var ptr uint32
	var retSize uint32

	err := types.StatusToError(getQueryOnHost(ctx.Id, &ptr, &retSize))
	if err != nil {
		return "", err
	}
	return types.PtrToStringWithFree(ptr, retSize), nil
}

func SetHostQuery(ctx types.WasmPluginContext, query string) error {
	if len(query) == 0 {
		return errors.New("query is empty")
	}
	ptr, size := types.StringToPtr(query)
	return types.StatusToError(setQueryOnHost(ctx.Id, ptr, size))
}

func GetAbiVersion(ctx types.WasmPluginContext) (string, error) {
	var ptr uint32
	var retSize uint32

	err := types.StatusToError(getAbiVersionOnHost(&ptr, &retSize))
	if err != nil {
		return "", err
	}
	return types.PtrToStringWithFree(ptr, retSize), nil
}

func GetRuntimeType(ctx types.WasmPluginContext) (string, error) {
	var ptr uint32
	var retSize uint32

	err := types.StatusToError(getRuntimeTypeOnHost(&ptr, &retSize))
	if err != nil {
		return "", err
	}
	return types.PtrToStringWithFree(ptr, retSize), nil
}

func InfoLog(ctx types.WasmPluginContext, message string) {
	if len(message) == 0 {
		return
	}
	ptr, size := types.StringToPtr(message)
	infoLogOnHost(ptr, size)
}

func ErrorLog(ctx types.WasmPluginContext, message string) {
	if len(message) == 0 {
		return
	}
	ptr, size := types.StringToPtr(message)
	errorLogOnHost(ptr, size)
}
