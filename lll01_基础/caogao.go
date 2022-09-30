package main

import "fmt"

/**
 *
 * @param numbers int整型一维数组
 * @param target int整型
 * @return int整型一维数组
 */
func twoSum(numbers []int, target int) []int {
	// write code here
	var res []int
	var m = make(map[int]int, 0)
	for i, v := range numbers {
		diff := target - v
		ii, ok := m[diff]
		if ok {
			fmt.Println(v, ii)
			if i > ii {
				res = append(res, ii+1)
				res = append(res, i+1)
			} else {
				res = append(res, i+1)
				res = append(res, ii+1)
			}

			goto end
		}

		m[v] = i

	}
end:
	return res
}

func main() {
	numbers := []int{3, 2, 4}
	target := 6
	res := twoSum(numbers, target)
	fmt.Printf("res: %v\n", res)
}
