package main

import (
	"fmt"
)

func main() {

	//常规语法
	const PI float32 = 3.14
	const H = 100

	//定义多个
	const (
		NAME = "Tom"
		AGE  = 18
	)

	//注意：声明多个变量时，如果省略了赋值，则表示和上面的一行值相同
	const (
		a = 1
		b //b=a=1
		c = 100
		d //d=c=100
	)

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)

	//同时定义
	const X, Y, Z = 2, 3, 4
	fmt.Printf("X: %v\n", X)
	fmt.Printf("Y: %v\n", Y)

	//iota
	const (
		A = iota //0
		B = iota //1
		C = iota //2

	)
	fmt.Printf("A: %v\n", A)
	fmt.Printf("B: %v\n", B)
	fmt.Printf("C: %v\n", C)

	//iota中间截断
	const (
		aa = iota //0
		_
		cc = iota //2
	)

	fmt.Printf("aa: %v\n", aa)
	fmt.Printf("cc: %v\n", cc)

	const (
		aaa = iota //0
		bbb = 1000 //1000
		ccc = iota //2
	)

	fmt.Printf("aaa: %v\n", aaa)
	fmt.Printf("ccc: %v\n", ccc)

}
