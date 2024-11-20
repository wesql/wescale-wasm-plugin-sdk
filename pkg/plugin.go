package pkg

import (
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg/types"
	_ "github.com/wesql/wescale-wasm-plugin-sdk/pkg/v1alpha2/guest_functions"
)

func SetWasmPlugin(plugin types.WasmPlugin) {
	types.CurrentWasmPlugin = plugin
}
