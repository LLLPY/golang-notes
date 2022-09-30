package main

import "fmt"

func sum(a int, b int) (c int) {
	c = a + b
	return
}

func main() {

	//自己声明函数的类型，然后再将函数赋值给一个变量
	type my_func func(int, int) int
	var my_sum my_func
	my_sum = sum
	res := my_sum(1, 2)
	fmt.Printf("res: %v\n", res)

	//直接通过短变量的形式将一个函数赋值给一个变量
	c := sum
	res2 := c(1, 2)
	fmt.Printf("res2: %v\n", res2)

}
