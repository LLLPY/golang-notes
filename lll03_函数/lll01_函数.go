package main

import "fmt"

//定义一个两数之和
func sum(a int, b int) (res int) {
	res = a + b
	return res
}
func main() {

	a := 10
	b := 100
	res := sum(a, b)
	fmt.Printf("res: %v\n", res)

}
