package main

import "fmt"

type Person struct {
	name string
}

//传递值
func (p Person) fix_name(name string) {
	p.name = name

}

//传递指针
func (p *Person) fix_name2(name string) {
	p.name = name

}

func main() {

	tom := Person{name: "Tom"}
	jerry := &Person{name: "Jerry"}
	fmt.Printf("tom.name: %v\n", tom.name)
	tom.fix_name("Tooooom")
	fmt.Printf("tom.name: %v\n", tom.name)
	print("----------------------------\n")
	fmt.Printf("jerry.name: %v\n", jerry.name)
	jerry.fix_name2("Jeeeeery")
	fmt.Printf("jerry.name: %v\n", jerry.name)

}
