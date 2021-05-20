package syscallutils

import "unsafe"

func PtrPtr(ret *uintptr) (b uintptr) {
	return uintptr(unsafe.Pointer(ret))
}
