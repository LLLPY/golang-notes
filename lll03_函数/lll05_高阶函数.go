package main

import "fmt"

func sum(a int, b int) (c int) {
	c = a + b
	return
}

func sub(a int, b int) (c int) {
	c = a - b
	return
}

func test(a int, b int, f func(int, int) int) {
	res := f(a, b)
	fmt.Printf("%v+%v=res: %v\n", a, b, res)
}

func cal(flag string) func(int, int) int {
	switch flag {
	case "+":
		return sum
	case "-":
		return sub
	default:
		return sum
	}

}

func main() {

	//函数作为参数传给函数
	test(10, 20, sum)

	//函数作为返回值
	f_sum := cal("+")
	res := f_sum(1, 2)
	fmt.Printf("res: %v\n", res)

	f_sub := cal("-")
	res2 := f_sub(1, 2)
	fmt.Printf("res2: %v\n", res2)
}
