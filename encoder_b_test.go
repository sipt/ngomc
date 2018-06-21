package ngomc

import (
	"testing"
	"encoding/json"
	"encoding/gob"
	"github.com/gogo/protobuf/proto"
	"time"
	"bytes"
	"fmt"
)

type TestStruct struct {
	A  int
	B  string
	C  int
	D  string
	E  int
	F  int
	G  string
	H  string
	I  string
	J  int
	K  string
	L  string
	M  string
	N  int
	O  int
	P  string
	Q  string
	R  int
	S  int
	T  string
	U  string
	V  int
	W  string
	X  time.Duration
	Y  time.Duration
	Z  time.Duration
	A1 time.Duration
	B1 time.Duration
	C1 int
	D1 float32
	E1 float32
	F1 float32
	G1 int
	H1 int
	I1 float32
	J1 float32
	K1 float32
	L1 int
	M1 int
	N1 int
	O1 string
	P1 int
	Q1 string
	R1 string
	S1 string
	T1 string
	U1 string
	V1 string
	W1 float32
	X1 float32
	Y1 float32
	Z1 float32
	A2 float32
	B2 string
	C2 string
	D2 string
	E2 string
	F2 int
	G2 string
	H2 int
	I2 string
	J2 int
	K2 string
	L2 string
	M2 string
	N2 string
	O2 int
	P2 string
	Q2 int
	R2 int
	S2 int
	T2 string
	U2 time.Duration
	V2 int
	W2 string
	X2 time.Duration
	Y2 string
	Z2 string
	A3 string
}

type TestStruct2 struct {
	A  int64
	B  int64
	C  int64
	D  int64
	E  int64
	F  int64
	G  int64
	H  int64
	I  int64
	J  int64
	K  int64
	L  int64
	M  int64
	N  int64
	O  int64
	P  int64
	Q  int64
	R  int64
	S  int64
	T  int64
	U  int64
	V  int64
	W  int64
	X  int64
	Y  int64
	Z  int64
	A1 float64
	B1 float64
	C1 float64
	D1 float64
	E1 float64
	F1 float64
	G1 float64
	H1 float64
	I1 float64
	J1 float64
	K1 float64
	L1 float64
	M1 float64
	N1 float64
	O1 float64
	P1 float64
	Q1 float64
	R1 float64
	S1 float64
	T1 float64
	U1 float64
	V1 float64
	W1 float64
	X1 float64
	Y1 float64
	Z1 float64
	A2 int32
	B2 int32
	C2 int32
	D2 int32
	E2 int32
	F2 int32
	G2 int32
	H2 int32
	I2 int32
	J2 int32
	K2 int32
	L2 int32
	M2 int32
	N2 int32
	O2 int32
	P2 int32
	Q2 int32
	R2 int32
	S2 int32
	T2 int32
	U2 int32
	V2 int32
	W2 int32
	X2 int32
	Y2 int32
	Z2 int32
	A3 int32
}

func BenchmarkEncodeDecode(b *testing.B) {
	x := &TestStruct{1, "2", 3, "4", 5, 6, "7", "8", "9", 10, "11", "12", "13", 14, 15, "16", "17", 18, 19, "20", "21", 22, "23", 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, "41", 42, "43", "44", "45", "46", "47", "48", 49, 50, 51, 52, 53, "54", "55", "56", "57", 58, "59", 60, "61", 62, "63", "64", "65", "66", 67, "68", 69, 70, 71, "72", 73, 74, "75", 76, "77", "78", "79"}
	var bs []byte
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	b.Run("ngomc Encode", func(b *testing.B) {
		Prepare(x)
		for i := 0; i < b.N; i++ {
			bs = Encode(x)
		}
	})
	fmt.Println("ngomc encode len([]byte):", len(bs))
	b.Run("ngomc Decode", func(b *testing.B) {
		y := &TestStruct{}
		Prepare(x)
		for i := 0; i < b.N; i++ {
			_ = Decode(y, bs).(*TestStruct)
		}
	})

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	b.Run("Json Encode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bs, _ = json.Marshal(x)
		}
	})
	fmt.Println("Json encode len([]byte):", len(bs))
	b.Run("Json Decode", func(b *testing.B) {
		y := &TestStruct{}
		for i := 0; i < b.N; i++ {
			json.Unmarshal(bs, y)
		}
	})

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	b.Run("GOB Encode", func(b *testing.B) {
		var buffer bytes.Buffer
		enc := gob.NewEncoder(&buffer)
		bs = nil
		for i := 0; i < b.N; i++ {
			enc.Encode(x)
			if len(bs) == 0 {
				bs = buffer.Bytes()
			}
			buffer.Reset()
		}
	})
	fmt.Println("GOB encode len([]byte):", len(bs))
	b.Run("GOB Decode", func(b *testing.B) {
		var buffer bytes.Buffer
		dec := gob.NewDecoder(&buffer)
		y := &TestStruct{}
		for i := 0; i < b.N; i++ {
			buffer.Write(bs)
			dec.Decode(y)
		}
	})

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	b.Run("Proto Encode", func(b *testing.B) {
		x := &Test{1, "2", 3, "4", 5, 6, "7", "8", "9", 10, "11", "12", "13", 14, 15, "16", "17", 18, 19, "20", "21", 22, "23", 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, "41", 42, "43", "44", "45", "46", "47", "48", 49, 50, 51, 52, 53, "54", "55", "56", "57", 58, "59", 60, "61", 62, "63", "64", "65", "66", 67, "68", 69, 70, 71, "72", 73, 74, "75", 76, "77", "78", "79"}
		for i := 0; i < b.N; i++ {
			bs, _ = proto.Marshal(x)
		}
	})
	fmt.Println("Proto encode len([]byte):", len(bs))
	b.Run("Proto Decode", func(b *testing.B) {
		y := &Test{}
		for i := 0; i < b.N; i++ {
			proto.Unmarshal(bs, y)
		}
	})
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
}

func BenchmarkEncodeDecode2(b *testing.B) {
	x := &TestStruct2{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79}
	var bs []byte
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	b.Run("ngomc Encode", func(b *testing.B) {
		Prepare(x)
		for i := 0; i < b.N; i++ {
			bs = Encode(x)
		}
	})
	fmt.Println("ngomc encode len([]byte):", len(bs))
	b.Run("ngomc Decode", func(b *testing.B) {
		y := &TestStruct2{}
		Prepare(x)
		for i := 0; i < b.N; i++ {
			_ = Decode(y, bs).(*TestStruct2)
		}
	})

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	b.Run("Json Encode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bs, _ = json.Marshal(x)
		}
	})
	fmt.Println("Json encode len([]byte):", len(bs))
	b.Run("Json Decode", func(b *testing.B) {
		y := &TestStruct2{}
		for i := 0; i < b.N; i++ {
			json.Unmarshal(bs, y)
		}
	})

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	b.Run("GOB Encode", func(b *testing.B) {
		var buffer bytes.Buffer
		enc := gob.NewEncoder(&buffer)
		bs = nil
		for i := 0; i < b.N; i++ {
			enc.Encode(x)
			if len(bs) == 0 {
				bs = buffer.Bytes()
			}
			buffer.Reset()
		}
	})
	fmt.Println("GOB encode len([]byte):", len(bs))
	b.Run("GOB Decode", func(b *testing.B) {
		var buffer bytes.Buffer
		dec := gob.NewDecoder(&buffer)
		y := &TestStruct2{}
		for i := 0; i < b.N; i++ {
			buffer.Write(bs)
			dec.Decode(y)
		}
	})

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	b.Run("Proto Encode", func(b *testing.B) {
		x := &Test2{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79}
		for i := 0; i < b.N; i++ {
			bs, _ = proto.Marshal(x)
		}
	})
	fmt.Println("Proto encode len([]byte):", len(bs))
	b.Run("Proto Decode", func(b *testing.B) {
		y := &Test{}
		for i := 0; i < b.N; i++ {
			proto.Unmarshal(bs, y)
		}
	})
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
}
