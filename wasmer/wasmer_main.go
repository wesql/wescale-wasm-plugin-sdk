package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	wescale_wasm_plugin_template "wescale-wasm-plugin-template"
	"wescale-wasm-plugin-template/tools"
)

func main() {}

var bytesBuffer []byte
var bufferLen uint64

//export clearBuf
func clearBuf() {
	bytesBuffer = bytesBuffer[:0]
}

//export writeBuf
func writeBuf(val, bytesLen uint64) {
	uint64Bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(uint64Bytes, uint64(val))
	bytesBuffer = append(bytesBuffer, uint64Bytes[:bytesLen]...)
}

//export readBuf
func readBuf(byteOffset uint64) (val uint64) {
	// return 8 bytes from byteOffset
	if byteOffset+8 <= bufferLen {
		return binary.LittleEndian.Uint64(bytesBuffer[byteOffset : byteOffset+8])
	}

	lastBytesPadding := make([]byte, 8)
	leftNum := bufferLen - byteOffset
	for i := uint64(0); i < 8; i++ {
		if i < leftNum {
			//fmt.Printf("guest, blen:%v, i:%v, byteOffset+i:%v\n", bufferLen, i, byteOffset+i)
			lastBytesPadding[i] = bytesBuffer[byteOffset+i]
		} else {
			lastBytesPadding[i] = 0
		}
	}
	return binary.LittleEndian.Uint64(lastBytesPadding)
}

//export getBufLen
func getBufLen() uint64 {
	bufferLen = uint64(len(bytesBuffer))
	return uint64(len(bytesBuffer))
}

//export wasmerGuestFunc
func wasmerGuestFunc() {

	// todo, how to return err?
	var err error

	e := tools.WasmPluginExchange{}
	//fmt.Printf("guest bytesbuffer len: %d buffer: %v\n", len(bytesBuffer), bytesBuffer)
	err = json.Unmarshal(bytesBuffer, &e)
	if err != nil {
		fmt.Println(err)
	}

	wescale_wasm_plugin_template.RunBeforeExecution(&e)

	bytesBuffer, err = json.Marshal(e)
	if err != nil {
		fmt.Println(err)
	}
}
