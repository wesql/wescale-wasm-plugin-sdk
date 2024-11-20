package guest_functions

import (
	"errors"
	"github.com/wesql/sqlparser/go/vt/proto/query"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg/types"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg/v1alpha2/host_functions"
)

type WasmPlugin interface {
	RunBeforeExecution() error
	RunAfterExecution(queryResult *query.QueryResult, errBefore error) (*query.QueryResult, error)
}

//export proxy_on_memory_allocate
func proxyOnMemoryAllocate(size uint) *byte {
	buf := make([]byte, size)
	return &buf[0]
}

//export RunBeforeExecutionOnGuest
func RunBeforeExecutionOnGuest(hostInstancePtr, hostModulePtr uint64) {
	types.CurrentWasmPluginContext.HostInstancePtr = hostInstancePtr
	types.CurrentWasmPluginContext.HostModulePtr = hostModulePtr

	err := types.CurrentWasmPlugin.RunBeforeExecution()
	if err != nil {
		host_functions.SetErrorMessage(err.Error())
	}
}

//export RunAfterExecutionOnGuest
func RunAfterExecutionOnGuest() {
	qr, err := host_functions.GetQueryResult()
	if err != nil && !errors.Is(err, types.StatusToError(types.StatusBadArgument)) {
		// unknown error
		host_functions.SetErrorMessage(err.Error())
		return
	}

	errMessageBefore, err := host_functions.GetErrorMessage()
	if err != nil {
		if !errors.Is(err, types.StatusToError(types.StatusBadArgument)) {
			// unknown error
			host_functions.SetErrorMessage(err.Error())
			return
		} else {
			errMessageBefore = ""
		}

	}
	var errBefore error
	if errMessageBefore != "" {
		errBefore = errors.New(errMessageBefore)
	}

	finalQueryResult, finalErr := types.CurrentWasmPlugin.RunAfterExecution(qr, errBefore)

	host_functions.SetQueryResult(finalQueryResult)
	if finalErr != nil {
		host_functions.SetErrorMessage(finalErr.Error())
	}
}
