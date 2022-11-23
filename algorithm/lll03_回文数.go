package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {

	int2str := strconv.Itoa(x)
	s_len := len(int2str)
	half_len := s_len / 2
	for i := 0; i < half_len; i++ {
		if int2str[i] != int2str[s_len-i-1] {
			return false
		}
	}
	return true
}

func main() {
	x := 1
	fmt.Printf("isPalindrome(x): %v\n", isPalindrome(x))

}
