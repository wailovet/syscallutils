package syscallutils

import "unsafe"

func PtrPtr(i uintptr, cb func(pp uintptr)) uintptr {
	pp := uintptr(unsafe.Pointer(&i))
	cb(pp)
	return (uintptr)(unsafe.Pointer((*uintptr)(unsafe.Pointer(ptr))))
}
