package main

import "fmt"

func main() {

	//切片的初始化
	//方法一
	var s1 = []int{1, 2, 3}
	fmt.Printf("s1: %v\n", s1)

	//方法二：make
	s2 := make([]int, 10)
	fmt.Printf("s2: %v\n", s2)

	//方法三：借助数组
	arr := [3]int{1, 2, 3}
	s3 := arr[:]
	fmt.Printf("s3: %v\n", s3)

	//切片/数组/字符串的切片操作： s[a:b] 左闭右开 于python不一样的是，go语言中不能修改步长，步长只能是1
	s4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("s4[1:9]: %v\n", s4[1:3])
}
