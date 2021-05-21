package syscallutils

import "unsafe"

var C2Go C2GoBase

type C2GoBase struct {
}

func (c2g *C2GoBase) Int(p uintptr) int {
	return int(p)
}

func (c2g *C2GoBase) Int32(p uintptr) int32 {
	return int32(p)
}

func (c2g *C2GoBase) Int64(p uintptr) int64 {
	return int64(p)
}

func (c2g *C2GoBase) Float64Slice(p uintptr, len int) (ret []float64) {
	for i := 0; i < len; i++ {
		ret = append(ret, *((*float64)(unsafe.Pointer(p))))
		p += unsafe.Sizeof(float64(0))
	}
	return ret
}

func (c2g *C2GoBase) Float32Slice(p uintptr, len int) (ret []float32) {
	for i := 0; i < len; i++ {
		ret = append(ret, *((*float32)(unsafe.Pointer(p))))
		p += unsafe.Sizeof(float32(0))
	}
	return ret
}

func (c2g *C2GoBase) Bytes(p uintptr) (d []byte) {
	ret := (*byte)(unsafe.Pointer(p))
	// 遍历C返回的char指针，直到 '\0' 为止
	for *ret != 0 {
		d = append(d, *ret)
		p += unsafe.Sizeof(byte(0))      // 移动指针，指向下一个char
		ret = (*byte)(unsafe.Pointer(p)) // 获取指针的值，此时指针已经指向下一个char
	}
	return
}

func (c2g *C2GoBase) Bytes4Void(p uintptr, len int) (d []byte) {
	for i := 0; i < len; i++ {
		d = append(d, *(*byte)(unsafe.Pointer(p)))
		p += unsafe.Sizeof(byte(0)) // 移动指针，指向下一个char
	}
	return
}

func (c2g *C2GoBase) Uint8Slice(p uintptr, len int) (d []uint8) {
	for i := 0; i < len; i++ {
		d = append(d, *(*uint8)(unsafe.Pointer(p)))
		p += unsafe.Sizeof(uint8(0)) // 移动指针，指向下一个char
	}
	return
}
func (c2g *C2GoBase) Uint16Slice(p uintptr, len int) (d []uint16) {
	for i := 0; i < len; i++ {
		d = append(d, *(*uint16)(unsafe.Pointer(p)))
		p += unsafe.Sizeof(uint16(0)) // 移动指针，指向下一个char
	}
	return
}
func (c2g *C2GoBase) Uint32Slice(p uintptr, len int) (d []uint32) {
	for i := 0; i < len; i++ {
		d = append(d, *(*uint32)(unsafe.Pointer(p)))
		p += unsafe.Sizeof(uint32(0)) // 移动指针，指向下一个char
	}
	return
}

func (c2g *C2GoBase) IntSlice(p uintptr, len int) (d []int) {
	for i := 0; i < len; i++ {
		d = append(d, *(*int)(unsafe.Pointer(p)))
		p += unsafe.Sizeof(int(0)) // 移动指针，指向下一个char
	}
	return
}

func (c2g *C2GoBase) String4CharPtr(p uintptr) string {
	return string(c2g.Bytes(p))
}
