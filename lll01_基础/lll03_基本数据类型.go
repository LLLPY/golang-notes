package main

import "fmt"

func foo() {}

func main() {

	//布尔类型
	var a bool = false
	fmt.Printf("%T \n", a) //%T可以用来打印输出变量的类型

	//数字类型
	b := 10
	c := 10.00
	var d float32 = 1.00
	fmt.Printf("%T \n", b) //int
	fmt.Printf("%T \n", c) //float64
	fmt.Printf("%T \n", d) //float32

	//字符类型
	e := "Tom"
	fmt.Printf("%T \n", e) //string

	//数组类型
	arr := [10]int{1, 2, 3, 4}
	fmt.Printf("%T \n", arr) //[10]int

	//切片类型
	slice := []int{1, 2, 3, 4}
	fmt.Printf("%T \n", slice) //[]int

	//指针类型
	f := 10000
	g := &f
	fmt.Printf("%T \n", g) //*int

	//函数类型
	fmt.Printf("%T \n", foo) //func()

	//结构体类型
	type Student struct {
		Name string
		Age  int
	}

	tom := Student{"Tom", 18}
	fmt.Printf("%T \n", tom) //main.Student
	aa := 2 ^ 8
	fmt.Printf("aa: %v\n", aa)

}
