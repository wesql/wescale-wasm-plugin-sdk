package tools

// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

// todo, it's a silly way
// todo, add a field to set log here?
type WasmPluginRunBeforeExecutionExchange struct {
	Query string
}

// todo, which fields? query result?
type WasmPluginRunAfterExecutionExchange struct {
	Query string
}

// ptrToString returns a string from WebAssembly compatible numeric types
// representing its pointer and length.
func PtrToString(ptr uint32, size uint32) string {
	return unsafe.String((*byte)(unsafe.Pointer(uintptr(ptr))), size)
}

func PtrToBytes(ptr uint32, size uint32) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(uintptr(ptr))), size)
}

// stringToLeakedPtr returns a pointer and size pair for the given string in a way
// compatible with WebAssembly numeric types.
// The pointer is not automatically managed by TinyGo hence it must be freed by the host.
func StringToLeakedPtr(s string) (uint32, uint32) {
	size := C.ulong(len(s))
	ptr := unsafe.Pointer(C.malloc(size))
	copy(unsafe.Slice((*byte)(ptr), size), s)
	return uint32(uintptr(ptr)), uint32(size)
}

func BytesToLeakedPtr(b []byte) (uint32, uint32) {
	size := C.ulong(len(b))
	ptr := unsafe.Pointer(C.malloc(size))
	copy(unsafe.Slice((*byte)(ptr), size), b)
	return uint32(uintptr(ptr)), uint32(size)
}
