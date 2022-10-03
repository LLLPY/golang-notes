package main

import "fmt"

// 定义一个USB的读写接口
type USB interface {
	read()
	write(string)
}

// 计算机
type Computer struct {
	name string
}

//手机
type Mobile struct {
	name string
}

// 计算机实现USB接口
func (c Computer) read() {
	fmt.Printf("%v is reading...\n", c.name)
}

func (c Computer) write() {
	fmt.Printf("%v is writing...\n", c.name)
}

func main() {

	c := Computer{name: "Mac"}
	c.read()
	c.write()

}
