package hostfunction

import "wescale-wasm-plugin-template/tools"

//export GetValueByKeyHost
func GetValueByKeyHost(key uint32) uint32

//export SetValueByKeyHost
func SetValueByKeyHost(key, value uint32)

//export GetQueryHost
func GetQueryHost(hostInstancePtr uint64, returnValueData *uint32, returnValueSize *uint32) tools.Status

//export SetQueryHost
func SetQueryHost(hostInstancePtr uint64, queryValuePtr uint32, queryValueSize uint32) tools.Status
