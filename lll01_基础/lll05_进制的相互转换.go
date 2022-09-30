package main

import "fmt"

func main() {

	//十进制的数
	a := 10

	fmt.Printf("a的十进制表示a: %d\n", a)
	fmt.Printf("a的二进制表示a: %b\n", a)
	fmt.Printf("a的八进制表示a: %o\n", a)
	fmt.Printf("a的十六进制表示a: %x\n", a)

	//定义一个八进制的数 以0开头
	b := 077

	fmt.Printf("b的十进制表示a: %d\n", b)
	fmt.Printf("b的二进制表示a: %b\n", b)
	fmt.Printf("b的八进制表示a: %o\n", b)
	fmt.Printf("b的十六进制表示a: %x\n", b)

	//定义一个十六进制的数 以0x开头
	c := 0x111

	fmt.Printf("c的十进制表示a: %d\n", c)
	fmt.Printf("c的二进制表示a: %b\n", c)
	fmt.Printf("c的八进制表示a: %o\n", c)
	fmt.Printf("c的十六进制表示a: %x\n", c)

}
