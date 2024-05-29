package v1alpha1

import (
	"unsafe"
)

func PtrToString(ptr uint32, size uint32) string {
	return unsafe.String((*byte)(unsafe.Pointer(uintptr(ptr))), size)
}

func PtrToBytes(ptr uint32, size uint32) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(uintptr(ptr))), size)
}

func StringToPtr(s string) (uint32, uint32) {
	ptr := unsafe.Pointer(unsafe.StringData(s))
	return uint32(uintptr(ptr)), uint32(len(s))
}

func BytesToPtr(bytes []byte) (uint32, uint32) {
	ptr := unsafe.Pointer(&bytes[0])
	return uint32(uintptr(ptr)), uint32(len(bytes))
}
