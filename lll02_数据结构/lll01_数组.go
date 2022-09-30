package main

import "fmt"

type Student struct {
	Name string
	num  int
}

func main() {

	// 数组的定义
	var arr_int [10]int
	fmt.Printf("arr_int: %v\n", arr_int)
	var arr_str [10]Student
	fmt.Printf("arr_str: %v\n", arr_str)

	//给数组赋初始值
	var arr_int_init = [10]int{1, 2, 3}
	fmt.Printf("arr_int_init: %v\n", arr_int_init)

	//给指定位置赋初始值
	var arr_float_init = [10]float64{0: 100.0, 3: 200.0}
	fmt.Printf("arr_float_init: %v\n", arr_float_init)

	//使用...不指定数组的大小,根据初始值来判断数组的大小
	var arr_string_init = [...]string{"hello", "world"}
	fmt.Printf("arr_string_init: %v\n", arr_string_init)

}
