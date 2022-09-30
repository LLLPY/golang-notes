package main

import "fmt"

//没有返回值
func foo1() {
	print("没有参数，也没有返回值！\n")
}

//有参数，没有返回值
func foo2(name string) {
	print("有一个参数：")
	fmt.Printf("name: %v\n", name)
}

//有参数，有返回值 且给返回值命名了
func foo3(name string, age int) (name2 string, age2 int) {
	name2 = name
	age2 = age
	return name2, age2
	// return    如果return不指定返回值，则默认返回上面定义的返回值
}

// 有参数，有返回值，但没有给返回值命名 这个时候就必须需要return来指定返回值
func foo4(name string, age int) (string, int) {
	return name, age

}

func main() {

	foo1()
	foo2("Tom")
	name2, age2 := foo3("Tom", 18)
	fmt.Printf("name2: %v\n", name2)
	fmt.Printf("age2: %v\n", age2)
	name, _ := foo4("Tom", 18) //丢弃age
	fmt.Printf("name: %v\n", name)

}
