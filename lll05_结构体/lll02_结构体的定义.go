package main

import "fmt"

func main() {

	//一个结构体就相当于一个新的类型
	type Student struct {
		name  string
		num   int
		age   int
		email string
	}

	tom := Student{"Tom", 1, 18, "110@qq.com"}
	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("tom.name: %v\n", tom.name)
	fmt.Printf("tom.age: %v\n", tom.age)

	// 匿名结构体
	var dog struct {
		name string
		age  int
	}

	dog.name = "Jerry"
	dog.age = 3
	fmt.Printf("dog: %v\n", dog)
	fmt.Printf("dog.name: %v\n", dog.name)
}
