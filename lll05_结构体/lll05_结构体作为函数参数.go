package main

import "fmt"

type Student struct {
	name string
	age  int
}

//值传递
func show_student(s Student, name string, age int) {
	s.name = name
	s.age = age
	fmt.Printf("s: %v\n", s)

}

//地址传递（引用传递）
func show_student2(s *Student, name string, age int) {
	s.name = name
	s.age = age
	fmt.Printf("s: %v\n", s)
}

func main() {

	tom := Student{name: "Tom", age: 18}
	jerry := Student{name: "Jerry", age: 18}

	//tom直接传值
	fmt.Printf("tom: %v\n", tom)
	show_student(tom, "Tom2", 20)
	fmt.Printf("tom: %v\n", tom)
	fmt.Println("-------------------")
	//jerry传递地址
	fmt.Printf("jerry: %v\n", jerry)
	show_student2(&jerry, "Jerry2", 20)
	fmt.Printf("jerry: %v\n", jerry)

}
