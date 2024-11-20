package host_functions

import (
	"errors"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg/types"
)

func GetHostQuery() (string, error) {
	var ptr uint32
	var retSize uint32

	err := types.StatusToError(getQueryOnHost(types.CurrentWasmPluginContext.HostInstancePtr, &ptr, &retSize))
	if err != nil {
		return "", err
	}
	return types.PtrToString(ptr, retSize), nil
}

func SetHostQuery(query string) error {
	if len(query) == 0 {
		return errors.New("query is empty")
	}
	ptr, size := types.StringToPtr(query)
	return types.StatusToError(setQueryOnHost(types.CurrentWasmPluginContext.HostInstancePtr, ptr, size))
}

func GetAbiVersion() (string, error) {
	var ptr uint32
	var retSize uint32

	err := types.StatusToError(getAbiVersionOnHost(&ptr, &retSize))
	if err != nil {
		return "", err
	}
	return types.PtrToString(ptr, retSize), nil
}

func GetRuntimeType() (string, error) {
	var ptr uint32
	var retSize uint32

	err := types.StatusToError(getRuntimeTypeOnHost(&ptr, &retSize))
	if err != nil {
		return "", err
	}
	return types.PtrToString(ptr, retSize), nil
}

func InfoLog(message string) {
	if len(message) == 0 {
		return
	}
	ptr, size := types.StringToPtr(message)
	infoLogOnHost(ptr, size)
}

func ErrorLog(message string) {
	if len(message) == 0 {
		return
	}
	ptr, size := types.StringToPtr(message)
	errorLogOnHost(ptr, size)
}
