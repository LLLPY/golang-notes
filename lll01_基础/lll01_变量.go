package main

import "fmt"

func Fun() (string, int) {
	return "Tom", 18
}

func main() {

	//默认声明语法
	var price float64
	fmt.Printf("price: %v\n", price)

	//类型推断
	var age = 18
	fmt.Printf("age: %v\n", age)

	//批量声明多个变量
	var name, age2, gender = "Tom", 18, true
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age2)
	fmt.Printf("gender: %v\n", gender)

	//短变量声明
	num := 99.9
	fmt.Printf("num: %v\n", num)

	//匿名变量 匿名变量不使用也不会报错
	tom_name, _ := Fun()
	fmt.Printf("tom_name: %v\n", tom_name)

}
