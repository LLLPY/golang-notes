package main

import "fmt"

//"父类"
type Pet struct {
	name string
	age  int
}

func (p Pet) eat() {
	fmt.Printf("%v is eating\n", p.name)
}

func (p Pet) sleep() {
	fmt.Printf("%v is sleeping\n", p.name)
}

// “子类”
type Dog struct {
	Pet
}

func main() {

	erha := Dog{
		Pet{name: "二哈", age: 2},
	}

	// 子类可以直接调用父类中的方法
	erha.eat()
	erha.sleep()

}
