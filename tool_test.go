package ngomc

import (
	"testing"
	"fmt"
)

type OffsetTest struct {
	A int64
	B string
	C string
	D string
	E int64
}

func TestPrepare_string(t *testing.T) {
	answer := OffsetType{{8, 64, 16, 0}, {24, 16, 32, 0}, {40, 32, 48, 0}}
	reply := Prepare(&OffsetTest{})
	if fmt.Sprint(answer) != fmt.Sprint(reply) {
		t.Errorf("reply failed:%v", reply)
	}
}
