package main

import "fmt"

func main() {

	type Student struct {
		name string
		age  int
	}

	tom := Student{name: "Tom", age: 18}

	//定义一个结构体指针
	var s_p *Student
	s_p = &tom
	fmt.Printf("s_p: %v\n", s_p)
	fmt.Printf("(*s_p): %v\n", (*s_p))
	fmt.Printf("(*s_p).name: %v\n", (*s_p).name)
	fmt.Printf("s_p.name: %v\n", s_p.name) //在取成员变量的值的时候可以将*省略

	//使用new关键字创建结构体指针
	var jerry = new(Student)
	jerry.name = "Jerry"
	fmt.Printf("jerry: %v\n", jerry)
	fmt.Printf("(*jerry): %v\n", (*jerry)) //取值
	fmt.Printf("jerry.name: %v\n", jerry.name)

}
