package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {

	var a uint8
	var b uint16
	var c uint32
	var d uint64

	fmt.Printf("a: %v %T %dB  %v~%v\n", a, a, unsafe.Sizeof(a), 0, math.MaxUint8)
	fmt.Printf("b: %v %T %dB  %v~%v\n", b, b, unsafe.Sizeof(b), 0, math.MaxUint16)
	fmt.Printf("c: %v %T %dB  %v~%v\n", c, c, unsafe.Sizeof(c), 0, math.MaxUint32)
	fmt.Printf("d: %v %T %dB  %v~%v\n", d, d, unsafe.Sizeof(d), 0, math.MaxInt64)

	var e int8
	var f int16
	var g int32
	var h int64
	fmt.Printf("e: %v %T %dB  %v~%v\n", e, e, unsafe.Sizeof(e), math.MinInt8, math.MaxInt8)
	fmt.Printf("f: %v %T %dB  %v~%v\n", f, f, unsafe.Sizeof(f), math.MinInt16, math.MaxInt16)
	fmt.Printf("g: %v %T %dB  %v~%v\n", g, g, unsafe.Sizeof(g), math.MinInt32, math.MaxInt32)
	fmt.Printf("h: %v %T %dB  %v~%v\n", h, h, unsafe.Sizeof(h), math.MinInt64, math.MaxInt64)

	var i float32
	var j float64
	fmt.Printf("i: %v %T %dB  %v~%v\n", i, i, unsafe.Sizeof(i), -math.MaxFloat32, math.MaxFloat32)
	fmt.Printf("j: %v %T %dB  %v~%v\n", j, j, unsafe.Sizeof(j), -math.MaxFloat64, math.MaxFloat64)

	var k complex64
	var l complex128
	fmt.Printf("k: %v %T %dB  \n", k, k, unsafe.Sizeof(k))
	fmt.Printf("l: %v %T %dB  \n", l, l, unsafe.Sizeof(l))

}
