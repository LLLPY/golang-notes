package main

import "fmt"

func main() {

	//append函数 在切片的末尾添加元素
	i := append([]int{1, 2, 3}, 4)
	fmt.Printf("i: %v\n", i)

	//len函数 求切片/string等的长度
	fmt.Printf("len(i): %v\n", len(i))

	//panic 抛出异常
	// panic("抛出异常!!!")

	s := new(string)
	fmt.Printf("s: %T\n", s) //*string
	fmt.Printf("s: %v\n", *s)

	i2 := new([]int)
	fmt.Printf("i2: %T\n", i2) //*[]int
	fmt.Printf("i2: %v\n", *i2)

	i3 := make([]int, 10, 100) //初始化容量为100，长度为10
	fmt.Printf("i3: %T\n", i3)
	fmt.Printf("i3: %v\n", i3)

}
