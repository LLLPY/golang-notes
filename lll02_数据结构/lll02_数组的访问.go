package main

import "fmt"

func main() {

	var arr_int = [10]int{1, 2, 3, 4, 5, 6}

	//访问第一个元素
	fmt.Printf("arr_int[0]: %v\n", arr_int[0])

	//访问第3个元素
	fmt.Printf("arr_int[2]: %v\n", arr_int[2])

	//访问最后一个元素
	fmt.Printf("arr_int[len(arr_int)-1]: %v\n", arr_int[len(arr_int)-1])

	//for遍历数组
	for i := 0; i < len(arr_int); i++ {

		fmt.Printf("a[%v]=%v \n", i, arr_int[i])

	}
	print("#########################################\n")
	//for range遍历数组
	for i, v := range arr_int {
		fmt.Printf("a[%v]=%v \n", i, v)

	}
}
