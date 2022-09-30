package main

import "fmt"

func main() {

	//切片的定义
	var s1 []int
	s1 = append(s1, 1)
	fmt.Printf("s1: %v\n", s1)

	//make定义切片的同时会将其初始化
	s2 := make([]string, 10)
	fmt.Printf("s2: %v\n", s2)

	//访问切片中的元素
	fmt.Printf("s1[0]: %v\n", s1[0])

	//在切片的末尾添加元素
	s1 = append(s1, 2)
	fmt.Printf("s1: %v\n", s1)
	//修改切片中的元素
	s1[0] = 100
	fmt.Printf("s1: %v\n", s1)

	//获取切片的大小
	fmt.Printf("len(s1): %v\n", len(s1))

	fmt.Printf("cap(s1): %v\n", cap(s1))

}
