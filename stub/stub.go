package stub

import (
	"reflect"
	"sync"
	"syscall"
	"unsafe"
)

type pointerValue struct {
	_   uintptr
	ptr unsafe.Pointer
}

var lock = &sync.Mutex{}

type stub struct {
	original []byte
	new      *reflect.Value
}

var originalFunctions = make(map[uintptr]stub)

// Replaces values of reflected functions
func StubFunc(target, replacement reflect.Value) {
	lock.Lock()

	defer lock.Unlock()

	if replacement.Kind() != reflect.Func && target.Type() != replacement.Type() {
		panic("target and replacementneed need to be a Func!")
	}

	// if already exists revert stub first
	UnStubFunc(target.Pointer())

	bytes := replaceFunction(target.Pointer(), (uintptr)(getPointer(replacement)))
	originalFunctions[target.Pointer()] = stub{bytes, &replacement}
}

// remove stub if exists for given pointer
func UnStubFunc(target uintptr) {
	stub, ok := originalFunctions[target]

	if ok {
		applyStub(target, stub.original)
	}
}

func replaceFunction(from, to uintptr) (original []byte) {
	val := funcValue(to)
	mem := getMemory(from)
	original = make([]byte, len(mem))
	copy(original, mem)

	applyStub(from, val)
	return
}

// Read bytes of provided pointer
func getMemory(p uintptr) []byte {
	return *(*[]byte)(unsafe.Pointer(&p))
}

func getPointer(v reflect.Value) unsafe.Pointer {
	return (*pointerValue)(unsafe.Pointer(&v)).ptr
}

func pageStart(ptr uintptr) uintptr {
	return ptr & ^(uintptr(syscall.Getpagesize() - 1))
}
