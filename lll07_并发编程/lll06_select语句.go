package main

import (
	"fmt"
	"time"
)

var chanName = make(chan string, 0)
var chanAge = make(chan int, 0)

func main() {

	go func() {
		for i := 0; i < 5; i++ {
			name := "Tom"
			chanName <- name
			chanAge <- i
		}
	}()

	for {
		select {
		case name := <-chanName: //读chanName
			fmt.Printf("name: %v\n", name)
		case age := <-chanAge:
			fmt.Printf("age: %v\n", age)

		default:
			print("default...\n")

		}
		time.Sleep(time.Second)
	}

}
