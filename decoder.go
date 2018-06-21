package ngomc

import (
	"reflect"
	"unsafe"
)

func Decode(i interface{}, v []byte) interface{} {
	typ := reflect.TypeOf(i)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	eface := (*[2]uintptr)(unsafe.Pointer(&i))
	vp := *(*uintptr)(unsafe.Pointer(&v))

	offsetList, ok := data[typ.String()]
	if !ok {
		offsetList = Prepare(i)
	}
	if offsetList == nil {
		// 没有指针，高速反序列化
		return *(*interface{})(unsafe.Pointer(&[2]uintptr{eface[0], vp}))
	}
	p := eface[1]
	var index uintptr
	for _, v := range offsetList {
		if index == 0 {
			index = eface[1] + v[1]
		} else {
			index += *(*uintptr)(unsafe.Pointer(p + v[1]))
		}
		*(*uintptr)(unsafe.Pointer(p + v[0])) = index
	}
	//
	//n := typ.NumField()
	//var f reflect.StructField
	//for i := 0; i < n; i++ {
	//	f = typ.Field(i)
	//	switch f.Type.Kind() {
	//	case reflect.Slice:
	//		// TODO
	//	case reflect.String:
	//		*(*uintptr)(unsafe.Pointer(*p + f.Offset)) = *p + typ.Size()
	//	case reflect.Ptr:
	//		// TODO
	//	}
	//}
	return *(*interface{})(unsafe.Pointer(&[2]uintptr{eface[0], vp}))
}
