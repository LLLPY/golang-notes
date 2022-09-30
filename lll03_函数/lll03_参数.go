package main

import "fmt"

func test(a int) {
	a = 200
	fmt.Printf("里面的a: %v\n", a)
}

//两数之和
func sum(a int, b int) (c int) {
	c = a + b
	return
}

//...接收无数个参数
func foo(name string, age int, args ...string) {
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}
func main() {

	a := 100
	test(a)
	fmt.Printf("外面的a=%v\n", a)

	res := sum(1, 2)
	fmt.Printf("res: %v\n", res)
	foo("Tom", 18, "Hello", "My", "name", "is", "Tom")

}
