package main

//定义一个宠物接口
type Pet interface {
	eat()
	sleep()
}

//定义两个接口
type Dog struct{}
type Cat struct{}

func (d Dog) eat() {
	print("dog is eating...\n")
}

func (d Dog) sleep() {
	print("dog is sleeping...\n")
}

func (c Cat) eat() {
	print("cat is eating...\n")
}

func (cat Cat) sleep() {
	print("cat is sleeping...\n")
}

type Person struct {
	name string
}

func (p Person) care(pet Pet) {

	pet.eat()
	pet.sleep()

}

func main() {

	d := Dog{}
	c := Cat{}
	p := Person{name: "Tom"}
	p.care(d)
	p.care(c)

}
