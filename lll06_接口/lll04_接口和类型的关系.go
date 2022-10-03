package main

import "fmt"

//一个类型实现多个接口
type Music interface {
	PlayMusic()
	a(aa string, bb string) (cc string, dd string)
}

type Vide interface {
	PlayVideo()
}

type Mobile struct {
	name string
}

func (m Mobile) PlayMusic() {
	fmt.Printf("%v播放音乐...\n", m.name)
}

func (m Mobile) PlayVideo() {
	fmt.Printf("%v播放视频...\n", m.name)

}

// 多个类型实现一个接口
type Eat interface {
	eat()
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func main() {

	m := Mobile{name: "Apple"}
	m.PlayMusic()
	m.PlayVideo()

	erha := Dog{name: "二哈"}
	fmt.Printf("%v is eating...\n", erha.name)
	tom := Cat{name: "Tom"}
	fmt.Printf("%v is eating...\n", tom.name)

}
