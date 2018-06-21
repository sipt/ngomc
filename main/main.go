package main

import (
	"github.com/sipt/ngomc"
	"fmt"
)

type A struct {
	A int64
	B int32
	C int8
	D string
	E float64
	F string
	G string
	H bool
}

func main() {
	a := &A{0, 1, 2, "abc", 1.1, "123", "ABC", false}
	fmt.Println(ngomc.Prepare(a))
	bytes := ngomc.Encode(a)
	fmt.Println(bytes)
	b := &A{}
	b = ngomc.Decode(b, bytes).(*A)
	fmt.Println(b)
	bytes[2] = 1
	fmt.Println(b)
	//a = nil
	//fmt.Println(*(*uintptr)(unsafe.Pointer(&a)))
}
