package hostfunction

import (
	"wescale-wasm-plugin-template/tools"
)

func GetHostQuery() (string, error) {
	var ptr uint32
	var retSize uint32

	err := tools.StatusToError(GetQueryHost(&ptr, &retSize))
	if err != nil {
		return "", err
	}
	return tools.PtrToString(ptr, retSize), nil
}

func SetHostQuery() {
}
