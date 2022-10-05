package main

import (
	"fmt"
	"time"
)

// 将值传入通道
func send(c *chan string, val string) {
	*c <- val
}

// 将值从通道中取出
func get(c *chan string) *string {
	val := <-*c
	fmt.Printf("val: %v\n", val)
	return &val
}

func main() {

	c := make(chan string) //无缓冲通道
	defer close(c)

	go send(&c, "hello")
	go get(&c)
	go send(&c, "world")
	go get(&c)

	time.Sleep(time.Millisecond * 1000)
	fmt.Printf("main end...")
}
