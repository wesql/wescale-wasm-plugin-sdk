package types

import (
	"unsafe"
)

//export free
func free(ptr uint32)

// Deprecated: use PtrToStringWithFree instead
func PtrToString(ptr uint32, size uint32) string {
	return unsafe.String((*byte)(unsafe.Pointer(uintptr(ptr))), size)
}

// Deprecated: use PtrToBytesWithFree instead
func PtrToBytes(ptr uint32, size uint32) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(uintptr(ptr))), size)
}

func PtrToStringWithFree(ptr uint32, size uint32) string {
	bytes := PtrToBytesWithFree(ptr, size)
	return string(bytes)
}

func PtrToBytesWithFree(ptr uint32, size uint32) []byte {
	temp := unsafe.Slice((*byte)(unsafe.Pointer(uintptr(ptr))), size)
	result := make([]byte, size)
	copy(result, temp)
	free(ptr)
	return result
}

func StringToPtr(s string) (uint32, uint32) {
	ptr := unsafe.Pointer(unsafe.StringData(s))
	return uint32(uintptr(ptr)), uint32(len(s))
}

func BytesToPtr(bytes []byte) (uint32, uint32) {
	ptr := unsafe.Pointer(&bytes[0])
	return uint32(uintptr(ptr)), uint32(len(bytes))
}
