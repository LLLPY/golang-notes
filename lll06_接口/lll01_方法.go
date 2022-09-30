package main

import "fmt"

type Dog struct {
	name string
	age  int
}

//为结构体Dog定义方法
func (dog Dog) say_hi() {
	fmt.Printf("我叫%v，我今年%v岁啦！\n", dog.name, dog.age)
}

func (dog Dog) woof() {
	fmt.Printf("%v正在汪汪大叫~\n", dog.name)
}

func main() {

	wc := Dog{"旺财", 2}
	wc.say_hi()
	wc.woof()

}
