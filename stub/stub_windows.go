//go:build windows

package stub

import (
	"syscall"
	"unsafe"
)

const READWRITE = 0x40

var virtualProtect = syscall.NewLazyDLL("kernel32.dll").NewProc("VirtualProtect")

func setVirtualProtection(lpAddress uintptr, dwSize int, flNewProtect uint32, lpflOldProtect unsafe.Pointer) error {
	res, _, _ := virtualProtect.Call(
		lpAddress,
		uintptr(dwSize),
		uintptr(flNewProtect),
		uintptr(lpflOldProtect))

	if res == 0 {
		return syscall.GetLastError()
	}

	return nil
}

func applyStub(target uintptr, data []byte) {
	bytes := getMemory(target)

	var oldPerms uint32
	err := setVirtualProtection(target, len(data), READWRITE, unsafe.Pointer(&oldPerms))
	if err != nil {
		panic(err)
	}

	copy(bytes, data[:])

	var tmp uint32
	err = setVirtualProtection(target, len(data), oldPerms, unsafe.Pointer(&tmp))
	if err != nil {
		panic(err)
	}
}
