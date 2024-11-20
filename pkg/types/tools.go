package types

import (
	"unsafe"
)

func PtrToString(ptr uint32, size uint32) string {
	return unsafe.String((*byte)(unsafe.Pointer(uintptr(ptr))), size)
}

func PtrToBytes(ptr uint32, size uint32) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(uintptr(ptr))), size)
}

////export free
//func free(ptr uint32)
//
//func PtrToStringAndFree(ptr uint32, size uint32) string {
//	temp := unsafe.String((*byte)(unsafe.Pointer(uintptr(ptr))), size)
//	result := string(temp)
//	free(ptr)
//	return result
//}
//
//func PtrToBytesAndFree(ptr uint32, size uint32) []byte {
//	temp := unsafe.Slice((*byte)(unsafe.Pointer(uintptr(ptr))), size)
//	result := make([]byte, size)
//	copy(result, temp)
//	free(ptr)
//	return result
//}

func StringToPtr(s string) (uint32, uint32) {
	ptr := unsafe.Pointer(unsafe.StringData(s))
	return uint32(uintptr(ptr)), uint32(len(s))
}

func BytesToPtr(bytes []byte) (uint32, uint32) {
	ptr := unsafe.Pointer(&bytes[0])
	return uint32(uintptr(ptr)), uint32(len(bytes))
}
