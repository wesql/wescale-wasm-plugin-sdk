package hostfunction

import "wescale-wasm-plugin-template/tools"

//export GetGlobalValueByKeyHost
func GetGlobalValueByKeyHost(keyPtr, keySize uint32, returnValuePtr, returnValueSize *uint32) tools.Status

//export SetGlobalValueByKeyHost
func SetGlobalValueByKeyHost(keyPtr, keySize, valuePtr, valueSize uint32) tools.Status

//export GetModuleValueByKeyHost
func GetModuleValueByKeyHost(hostModulePtr uint64, key uint32) uint32

//export SetModuleValueByKeyHost
func SetModuleValueByKeyHost(hostModulePtr uint64, key, value uint32)

//export GetQueryHost
func GetQueryHost(hostInstancePtr uint64, returnValueData *uint32, returnValueSize *uint32) tools.Status

//export SetQueryHost
func SetQueryHost(hostInstancePtr uint64, queryValuePtr uint32, queryValueSize uint32) tools.Status
