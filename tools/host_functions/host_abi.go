package hostfunction

//export GetValueByKeyHost
func GetValueByKeyHost(key uint32) uint32

//export SetValueByKeyHost
func SetValueByKeyHost(key, value uint32)

//export GetQueryHost
func GetQueryHost(returnValueData *uint32, returnValueSize *uint32) Status
