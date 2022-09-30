package main

import "fmt"

func main() {
	a := 100
	b := 200
	//算术运算符
	res := a + b
	fmt.Printf("res: %v\n", res)

	res = a - b
	fmt.Printf("res: %v\n", res)

	res = a * b
	fmt.Printf("res: %v\n", res)

	res = a / b
	fmt.Printf("res: %v\n", res)

	res = a % b
	fmt.Printf("res: %v\n", res)

	a++

	// 关系运算符
	var res2 bool

	res2 = a == b
	fmt.Printf("res2: %v\n", res2)

	res2 = a > b
	fmt.Printf("res2: %v\n", res2)

	res2 = a >= b
	fmt.Printf("res2: %v\n", res2)

	res2 = a < b
	fmt.Printf("res2: %v\n", res2)

	res2 = a <= b
	fmt.Printf("res2: %v\n", res2)

	res2 = a != b
	fmt.Printf("res2: %v\n", res2)

	//逻辑运算
	var res3 bool
	res3 = true && true
	fmt.Printf("res3: %v\n", res3)
	res3 = true || false
	fmt.Printf("res3: %v\n", res3)
	res3 = !res3
	fmt.Printf("res3: %v\n", res3)

	//位运算
	aa := 10
	bb := 11
	fmt.Printf("aa: %b\n", aa)
	fmt.Printf("bb: %b\n", bb)

	res4 := aa & bb
	fmt.Printf("res4:%b  %v\n", res4, res4)
	res4 = aa | bb
	fmt.Printf("res4:%b  %v\n", res4, res4)
	res4 = aa ^ bb
	fmt.Printf("res4:%b  %v\n", res4, res4)

	res4 = aa << 2
	fmt.Printf("res4:%b  %v\n", res4, res4)
	res4 = aa >> 1
	fmt.Printf("res4:%b  %v\n", res4, res4)

}
