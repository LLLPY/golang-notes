package main

import (
	"fmt"
)

func main() {

	name := "你好，China!"
	fmt.Println(len(name))
	fmt.Println(len([]rune(name)))
	fmt.Println(len([]int32(name)))
	fmt.Println(len([]byte(name)))

	rune_name := []rune(name)
	byte_name := []byte(name)
	int32_name := []int32(name)
	fmt.Printf("name: %v\n", name)
	fmt.Printf("int32_name: %v\n", string(int32_name))
	fmt.Printf("rune_name: %v\n", string(rune_name))
	fmt.Printf("byte_name: %v\n", string(byte_name))

	rune_name = append(rune_name[:len(rune_name)-1])
	fmt.Printf("rune_name: %v\n", string(rune_name))
	// for i := 0; i < len(name); i++ {
	// 	fmt.Printf("name[i]: %c\n", name[i])
	// }

	// for i := 0; i < len(rune_name); i++ {
	// 	fmt.Printf("rune_name[i]: %c\n", rune_name[i])
	// }
	// for i := 0; i < len(byte_name); i++ {
	// 	fmt.Printf("byte_name[i]: %c\n", byte_name[i])
	// }

}
