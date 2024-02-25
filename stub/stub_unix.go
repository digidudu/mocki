//go:build !windows

package stub

import (
	"reflect"
	"syscall"
	"unsafe"
)

func rawMemoryAccess(p uintptr, length int) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: p,
		Len:  length,
		Cap:  length,
	}))
}

func mProtect(addr uintptr, length int, prot int) {
	pageSize := syscall.Getpagesize()
	for p := pageStart(addr); p < addr+uintptr(length); p += uintptr(pageSize) {
		page := rawMemoryAccess(p, pageSize)
		err := syscall.Mprotect(page, prot)
		if err != nil {
			panic(err)
		}
	}
}

func applyStub(location uintptr, data []byte) {
	bytes := getMemory(location)

	mProtect(location, len(data), syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC)
	copy(bytes, data[:])
	mProtect(location, len(data), syscall.PROT_READ|syscall.PROT_EXEC)
}
