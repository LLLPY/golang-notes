package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i == 5 {
			goto label
		}

	}

label:
	fmt.Println("循环结束！")

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				fmt.Printf("i=%v j=%v k=%v\n", i, j, k)
				if i == 1 && j == 1 && k == 1 {
					goto mylabel
				}

			}

		}

	}
mylabel:
	fmt.Println("循环结束！")
}
