package pkg

import (
	"errors"
	hostfunction "github.com/wesql/wescale-wasm-plugin-sdk/pkg/host_functions/v1alpha1"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg/proto/query"
)

type WasmPlugin interface {
	RunBeforeExecution() error
	RunAfterExecution(queryResult *query.QueryResult, errBefore error) (*query.QueryResult, error)
}

var wasmPlugin WasmPlugin

func SetWasmPlugin(plugin WasmPlugin) {
	wasmPlugin = plugin
}

//export RunBeforeExecutionOnGuest
func RunBeforeExecutionOnGuest(hostInstancePtr, hostModulePtr uint64) {
	hostfunction.HostInstancePtr = hostInstancePtr
	hostfunction.HostModulePtr = hostModulePtr

	err := wasmPlugin.RunBeforeExecution()
	if err != nil {
		hostfunction.SetErrorMessage(err.Error())
	}
}

//export RunAfterExecutionOnGuest
func RunAfterExecutionOnGuest() {
	qr, err := hostfunction.GetQueryResult()
	if err != nil && !errors.Is(err, hostfunction.StatusToError(hostfunction.StatusBadArgument)) {
		// unknown error
		hostfunction.SetErrorMessage(err.Error())
		return
	}

	errMessageBefore, err := hostfunction.GetErrorMessage()
	if err != nil {
		if !errors.Is(err, hostfunction.StatusToError(hostfunction.StatusBadArgument)) {
			// unknown error
			hostfunction.SetErrorMessage(err.Error())
			return
		} else {
			errMessageBefore = ""
		}

	}
	var errBefore error
	if errMessageBefore != "" {
		errBefore = errors.New(errMessageBefore)
	}

	finalQueryResult, finalErr := wasmPlugin.RunAfterExecution(qr, errBefore)

	hostfunction.SetQueryResult(finalQueryResult)
	if finalErr != nil {
		hostfunction.SetErrorMessage(finalErr.Error())
	}
}
