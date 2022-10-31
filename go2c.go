package syscallutils

import "C"

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var Go2C Go2CBase

type Go2CBase struct{}

func (g2c *Go2CBase) AnyPtr(d interface{}) uintptr {
	return uintptr(unsafe.Pointer(&d))
}

func (g2c *Go2CBase) Uintptr(d interface{}) uintptr {
	return d.(uintptr)
}

func (g2c *Go2CBase) Uint8Slice(d []uint8) uintptr {
	if d == nil {
		return 0
	}
	ret := unsafe.Pointer(&d[0])
	return uintptr(unsafe.Pointer(ret))
}

func (g2c *Go2CBase) UintPtrSlice(d []uintptr) uintptr {
	if d == nil {
		return 0
	}
	ret := unsafe.Pointer(&d[0])
	return uintptr(unsafe.Pointer(ret))
}

func (g2c *Go2CBase) Float64Slice(d []float64) uintptr {
	if d == nil {
		return 0
	}
	ret := unsafe.Pointer(&d[0])
	return uintptr(ret)
}

func (g2c *Go2CBase) Float32Slice(d []float32) uintptr {
	if d == nil {
		return 0
	}
	ret := unsafe.Pointer(&d[0])
	return uintptr(ret)
}

func (g2c *Go2CBase) Int(d int) uintptr {
	return uintptr(d)
}

func (g2c *Go2CBase) Int32(d int32) uintptr {
	return uintptr(d)
}

func (g2c *Go2CBase) Int64(d int64) uintptr {
	return uintptr(d)
}

func (g2c *Go2CBase) Float32(d float32) uintptr {
	return uintptr(d)
}

func (g2c *Go2CBase) Float64(d float64) uintptr {
	// return uintptr(math.Float64bits(d)) //?
	return uintptr(d)
}

func (g2c *Go2CBase) Byte(d byte) uintptr {
	return uintptr(d)
}

func (g2c *Go2CBase) Uint8(d uint) uintptr {
	return uintptr(d)
}

func (g2c *Go2CBase) Chars(d string) uintptr {
	b := []byte(d)
	var newByte []byte
	for i := range b {
		newByte = append(newByte, b[i])
	}
	newByte = append(newByte, 0)
	// log.Println("Chars", b)
	return uintptr(unsafe.Pointer(&newByte[0]))
}

func wcharPtrFromString(s string) (*C.wchar_t, error) {
	p, err := windows.UTF16PtrFromString(s)
	return (*C.wchar_t)(p), err
}

func (g2c *Go2CBase) WChars(d string) uintptr {
	sp, _ := wcharPtrFromString(d)
	return uintptr(unsafe.Pointer(&sp))
}
