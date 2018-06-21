package ngomc

import (
	"reflect"
	"unsafe"
)

const (
	kindMask    = (1 << 5) - 1
	kindOffset  = 184 / 8
	elemOffset  = 384 / 8
	filedOffset = elemOffset + 1
)

//go:linkname memmove reflect.memmove
func memmove(to, from unsafe.Pointer, n uintptr)

func Encode(i interface{}) []byte {
	results := make([][2]uintptr, 0, 5)
	eface := (*[2]uintptr)(unsafe.Pointer(&i))
	typ := reflect.TypeOf(i)
	p := eface[0]
	if typ.Kind() == reflect.Ptr {
		p += elemOffset
		typ = typ.Elem()
	}
	p = *(*uintptr)(unsafe.Pointer(p)) // get elem rtype ptr
	p = *(*uintptr)(unsafe.Pointer(p)) // get size
	totalSize := int(p)
	results = append(results, [2]uintptr{eface[1], p})

	offsetList, ok := data[typ.String()]
	if !ok {
		offsetList = Prepare(i)
	}
	if offsetList == nil {
		// 没有指针指向，高速序列化
		s := results[0]
		return *(*[]byte)(unsafe.Pointer(&[3]uintptr{s[0], s[1], s[1]}))
	}

	var l, from uintptr
	p = eface[1]
	for _, v := range offsetList {
		l = v[3]
		if l == 0 {
			l = *(*uintptr)(unsafe.Pointer(p + v[2]))
		}
		from = *(*uintptr)(unsafe.Pointer(p + v[0]))
		results = append(results, [2]uintptr{from, l})
		totalSize += int(l)
	}
	result := make([]byte, totalSize)
	p = *(*uintptr)(unsafe.Pointer(&result))
	for i, index := 0, uintptr(0); i < len(results); i ++ {
		memmove(unsafe.Pointer(p+index), unsafe.Pointer(results[i][0]), results[i][1])
		index += results[i][1]
	}
	return result
}
