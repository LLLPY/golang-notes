package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i == 10 {
			break
		}
	}

	//配合标签使用
label:
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i == 5 {
			break label
		}

	}
	fmt.Println("END...")

}
