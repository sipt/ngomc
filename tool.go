package ngomc

import "reflect"

type OffsetType [][4]uintptr // [[offset, value offset, len offset, len value],[]]

var data map[string]OffsetType

func init() {
	data = make(map[string]OffsetType)
}

func Prepare(i interface{}) OffsetType {
	typ := reflect.TypeOf(i)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	r := MakeOffset(typ)
	data[typ.String()] = r
	return r
}

func MakeOffset(typ reflect.Type) OffsetType {
	offsets := make([][4]uintptr, 0, 5)
	var f reflect.StructField
	var offset = typ.Size()
	for i, l := 0, typ.NumField(); i < l; i++ {
		f = typ.Field(i)
		switch f.Type.Kind() {
		case reflect.String:
			offsets = append(offsets, [4]uintptr{f.Offset, offset, f.Offset + 8})
			offset = f.Offset + 8
		}
	}
	if len(offsets) == 0 {
		return nil
	}
	return offsets
}
