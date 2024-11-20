package pkg

import (
	_ "github.com/wesql/wescale-wasm-plugin-sdk/pkg/guest_functions"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg/types"
)

func InitWasmPlugin(plugin types.WasmPlugin) {
	types.CurrentWasmPlugin = plugin
}
