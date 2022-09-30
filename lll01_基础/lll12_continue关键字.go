package main

import "fmt"

func main() {
label:
	fmt.Println("i=1,j=1时跳出了循环。")

	//打印10以内的偶数
	for i := 0; i < 11; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Printf("i: %v\n", i)
	}

	for i := 0; i < 3; i++ {

		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				continue label //跳到label所在的地方
			}
			fmt.Printf("i: %v  j: %v\n", i, j)
		}

	}

}
