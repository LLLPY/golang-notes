package main

import (
	"fmt"
	"sync"
)

var mg sync.WaitGroup

func show_msg(msg int) {
	fmt.Printf("msg: %d\n", msg)
	defer mg.Add(-1) //任务完成就减一
}

func main() {

	for i := 0; i < 10; i++ {
		mg.Add(1) //添加任务就加一
		go show_msg(i)
	}

	mg.Wait() //如果WaitGroup的计数为0就停止等待
	print("end main...")
}
