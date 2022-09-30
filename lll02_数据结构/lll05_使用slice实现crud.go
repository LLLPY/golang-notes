package main

import (
	"fmt"
)

func main() {

	s := []int{1, 2, 3, 4, 5}

	//add
	s = append(s, 1)
	fmt.Printf("s: %v\n", s)

	//delete：删除索引为index的元素
	index := 2
	s = append(s[:index], s[index+1:]...)
	fmt.Printf("s: %v\n", s)

	//update
	s[4] = 6
	fmt.Printf("s: %v\n", s)

	//query
	target := 5
	for i, v := range s {
		if v == target {
			println("找到了！", i)
			break
		}

	}

}
