package main

import "fmt"

func main() {

	//类型定义
	type my_string string

	var name my_string = "Tom"
	fmt.Printf("name: %T %v\n", name, name) //类型为自己定义的新类型

	//类型别名
	type my_string2 = string
	var name2 my_string2 = "Tom"
	fmt.Printf("name2: %T %v\n", name2, name2) //类型仍旧是string

}
