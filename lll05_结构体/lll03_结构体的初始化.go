package main

import "fmt"

func main() {
	type Student struct {
		name  string
		age   int
		email string
	}

	tom := Student{}
	fmt.Printf("tom: %v\n", tom) //未初始化的结构体每个成员属性都是零值

	//“k-v”式初始化 可以对部分值进行初始化，也可以对全部值初始化
	jerry := Student{name: "Jerry", age: 18}
	fmt.Printf("jerry: %v\n", jerry)

	//“列表”式初始化 必须将所有值初始化
	autumn := Student{"Autumn", 18, "110@qq.com"}
	fmt.Printf("autumn: %v\n", autumn)
}
