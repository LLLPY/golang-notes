package main

import "fmt"

func main() {

	//将匿名函数赋值给一个变量
	say_hello := func(name string, age int) {
		fmt.Printf("My name is  %v,And i',m %v old.\n", name, age)
	}
	say_hello("Tom", 18)

	//直接调用匿名函数
	func(words string) {
		fmt.Printf("words: %v\n", words)
	}("Hi,我是一个匿名函数！")

}
