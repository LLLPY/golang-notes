package main

import "fmt"

func main() {

	var arr = [5]int{1, 2, 3, 4, 5}
	var arr_p [5]*int //数组类型的指针
	fmt.Printf("arr: %v\n", arr)
	for i := 0; i < len(arr); i++ {
		arr_p[i] = &arr[i]
	}

	for i := 0; i < len(arr_p); i++ {
		fmt.Printf("arr_p[i]=%v *arr_p[i]=%v arr[i]=%v\n", arr_p[i], *arr_p[i], arr[i])

	}

}
