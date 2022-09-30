package main

import "fmt"

type Person struct {
	name string
	age  int
	dog  Dog //宠物狗，另一个结构体
}

type Dog struct {
	name string
	age  int
}

func main() {

	var tom Person
	tom.name = "Tom"
	tom.age = 18

	var erha = Dog{name: "二哈", age: 3}
	tom.dog = erha

	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("tom.name: %v\n", tom.name)
	fmt.Printf("tom.dog.name: %v\n", tom.dog.name)
	tom.dog.name = "二哈哈哈"
	fmt.Printf("tom.dog: %v\n", tom.dog)
}
