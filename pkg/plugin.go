package pkg

import (
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg/guest_functions"
	_ "github.com/wesql/wescale-wasm-plugin-sdk/pkg/guest_functions"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg/types"
)

func InitWasmPlugin(plugin types.WasmPlugin) {
	guest_functions.CurrentWasmPlugin = plugin
}
