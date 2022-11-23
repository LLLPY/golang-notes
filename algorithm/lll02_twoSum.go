package main

import (
	"fmt"
)

/**
 *
 * @param numbers int整型一维数组
 * @param target int整型
 * @return int整型一维数组
 */
func twoSum(numbers []int, target int) []int {
	// write code here

	tmp_map := make(map[int]int)

	for i, v := range numbers {
		dif := target - v
		_, ok := tmp_map[dif]
		if ok {
			if i < tmp_map[dif] {
				return []int{i, tmp_map[dif]}
			} else {
				return []int{tmp_map[dif], i}
			}

		}
		tmp_map[v] = i
		fmt.Printf("tmp_map: %v\n", tmp_map)
	}

	return []int{1, 1}

}

func main() {
	numbers := []int{1, 2, 3, 4}
	target := 4
	res := twoSum(numbers, target)
	fmt.Printf("res: %v\n", res)
}
