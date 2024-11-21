package guest_functions

import (
	"errors"
	"github.com/wesql/sqlparser/go/vt/proto/query"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg/types"
)

type WasmPlugin interface {
	RunBeforeExecution() error
	RunAfterExecution(queryResult *query.QueryResult, errBefore error) (*query.QueryResult, error)
}

//export malloc
func malloc(size uint) *byte

// Deprecated: use malloc instead
//
//export proxy_on_memory_allocate
func proxyOnMemoryAllocate(size uint) *byte {
	return malloc(size)
}

//export RunBeforeExecutionOnGuest
func RunBeforeExecutionOnGuest(hostInstancePtr, hostModulePtr uint64) {
	types.CurrentWasmPluginContext.HostInstancePtr = hostInstancePtr
	types.CurrentWasmPluginContext.HostModulePtr = hostModulePtr

	err := types.CurrentWasmPlugin.RunBeforeExecution()
	if err != nil {
		setErrorMessage(err.Error())
	}
}

//export RunAfterExecutionOnGuest
func RunAfterExecutionOnGuest() {
	qr, err := getQueryResult()
	if err != nil && !errors.Is(err, types.StatusToError(types.StatusBadArgument)) {
		// unknown error
		setErrorMessage(err.Error())
		return
	}

	errMessageBefore, err := getErrorMessage()
	if err != nil {
		if !errors.Is(err, types.StatusToError(types.StatusBadArgument)) {
			// unknown error
			setErrorMessage(err.Error())
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

	setQueryResult(finalQueryResult)
	if finalErr != nil {
		setErrorMessage(finalErr.Error())
	} else {
		setErrorMessage("")
	}
}
