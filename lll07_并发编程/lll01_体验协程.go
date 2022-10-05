package main

import (
	"fmt"
	"time"
)

func show(name string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("name: %v\n", name)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {

	go show("Python") //开启一个go协程
	go show("Golang") //开启另外一个go协程
	time.Sleep(time.Millisecond * 2000)
	print("hello world!\n") //默认情况下：主线程结束后其他线程也会结束
}
