package syscallutils

import (
	"unsafe"
)

var Go2C Go2CBase

type Go2CBase struct{}

func (g2c *Go2CBase) AnyPtr(d interface{}) uintptr {
	return uintptr(unsafe.Pointer(&d))
}

func (g2c *Go2CBase) Float64Slice(d []float64) uintptr {
	ret := unsafe.Pointer(&d[0])
	return uintptr(ret)
}

func (g2c *Go2CBase) Float32Slice(d []float64) uintptr {
	ret := unsafe.Pointer(&d[0])
	return uintptr(ret)
}

func (g2c *Go2CBase) Int(d int) uintptr {
	return uintptr(d)
}

func (g2c *Go2CBase) Int32(d int) uintptr {
	return uintptr(d)
}

func (g2c *Go2CBase) Int64(d int) uintptr {
	return uintptr(d)
}

func (g2c *Go2CBase) Chars(d string) uintptr {
	b := append([]byte(d), 0)
	return uintptr(unsafe.Pointer(&b[0]))
}
